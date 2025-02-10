package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nguyendang2000/shared-go/logger"
	"github.com/twmb/franz-go/pkg/kgo"
)

// Service provides an abstraction for Kafka producers and consumers.
// It manages Kafka clients, topics, and logging within a specific context.
type Service struct {
	client         *kgo.Client     // Kafka client instance
	topicsConsumer []string        // Topics used for consuming messages
	context        context.Context // Context for managing service lifecycle
	logger         logger.Logger   // Logger for structured logging
}

// NewService initializes a new Kafka Service.
// It configures the Kafka client using the provided configuration and logger.
// If the client fails to initialize or the Kafka brokers are unreachable, it returns an error.
func NewService(ctx context.Context, conf Config, lg logger.Logger) (*Service, error) {
	options := []kgo.Opt{kgo.SeedBrokers(conf.Brokers...)}

	if len(conf.ConsumerTopics) > 0 {
		options = append(options, kgo.ConsumeTopics(conf.ConsumerTopics...))
	}

	if conf.ConsumerGroup != "" {
		options = append(options, kgo.ConsumerGroup(conf.ConsumerGroup))
	}

	if conf.AutoTopicCreation {
		options = append(options, kgo.AllowAutoTopicCreation())
	}

	client, err := kgo.NewClient(options...)
	if err != nil {
		return nil, fmt.Errorf(ErrKafkaClientSetup, err)
	}

	service := &Service{
		client:         client,
		topicsConsumer: conf.ConsumerTopics,
		context:        ctx,
		logger:         lg,
	}

	// Flush buffers and close the service when the context is canceled.
	go func() {
		<-ctx.Done()
		service.Flush()
		service.Close()
	}()

	if err := service.Ping(); err != nil {
		return nil, fmt.Errorf(ErrKafkaPing, err)
	}

	return service, nil
}

// Client returns the underlying Kafka client instance.
func (inst *Service) Client() *kgo.Client {
	return inst.client
}

// Ping tests the connection to Kafka brokers.
// It returns an error if the ping fails.
func (inst *Service) Ping() error {
	if err := inst.client.Ping(inst.context); err != nil {
		return err
	}
	return nil
}

// Flush ensures that all pending messages in the Kafka client's internal buffer
// are sent to Kafka before returning. This is useful for ensuring message delivery
// during shutdown or low throughput scenarios.
func (inst *Service) Flush() error {
	if err := inst.client.Flush(context.Background()); err != nil {
		return err
	}
	return nil
}

// Close shuts down the Kafka client.
func (inst *Service) Close() {
	inst.client.Close()
}

// prepareRecord creates a Kafka record with the specified topic, key, and data.
// It serializes the key and data into JSON format.
func prepareRecord(topic string, key interface{}, data interface{}) (*kgo.Record, error) {
	btsData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf(ErrMarshalData, data, err)
	}

	record := &kgo.Record{
		Topic: topic,
		Value: btsData,
	}

	if key != nil {
		btsKey, err := json.Marshal(key)
		if err != nil {
			return nil, fmt.Errorf(ErrMarshalKey, key, err)
		}
		record.Key = btsKey
	}

	return record, nil
}

// Produce asynchronously sends messages to the specified Kafka topic.
// It logs errors encountered during the production process.
func (inst *Service) Produce(topic string, key interface{}, data ...interface{}) error {
	for _, d := range data {
		record, err := prepareRecord(topic, key, d)
		if err != nil {
			return err
		}

		inst.client.Produce(inst.context, record, func(_ *kgo.Record, err error) {
			if err != nil {
				inst.logger.Errorf(ErrAsyncProduce, d, err)
			}
		})
	}

	return nil
}

// ProduceSync sends messages to the specified Kafka topic synchronously.
// It returns an error if any message fails to produce.
func (inst *Service) ProduceSync(topic string, key interface{}, data ...interface{}) error {
	records := make([]*kgo.Record, len(data))

	for i, d := range data {
		record, err := prepareRecord(topic, key, d)
		if err != nil {
			return err
		}
		records[i] = record
	}

	err := inst.client.ProduceSync(inst.context, records...).FirstErr()
	if err != nil {
		return fmt.Errorf(ErrSyncProduce, err)
	}

	return nil
}

// Consume retrieves messages from Kafka topics as a channel of Kafka records.
// The channel closes when the context is canceled or the client stops.
func (inst *Service) Consume() chan *kgo.Record {
	recordCh := make(chan *kgo.Record)

	go func() {
		defer close(recordCh)
		for {
			fetches := inst.client.PollFetches(inst.context)

			if inst.context.Err() != nil || fetches.IsClientClosed() {
				return
			}

			if errs := fetches.Errors(); len(errs) > 0 {
				for _, fetchErr := range errs {
					inst.logger.Panicf(ErrFetchMessages, fetchErr.Topic, fetchErr.Err)
				}
				continue
			}

			fetches.EachRecord(func(record *kgo.Record) {
				select {
				case recordCh <- record:
				case <-inst.context.Done():
					return
				}
			})
		}
	}()

	return recordCh
}
