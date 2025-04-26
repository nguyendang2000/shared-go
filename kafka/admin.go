package kafka

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kadm"
)

// CreatePartitions adds partitions to existing Kafka topics.
// It returns an error if adding partitions fails.
func (inst *Service) CreatePartitions(add int, topics ...string) error {
	adminClient := kadm.NewClient(inst.client)

	if _, err := adminClient.CreatePartitions(context.Background(), add, topics...); err != nil {
		return fmt.Errorf(ErrCreatePartitions, add, topics, err)
	}

	return nil
}

// CreateTopic creates a new Kafka topic with the specified configuration.
// It returns an error if the topic creation fails.
func (inst *Service) CreateTopic(partitions int32, replicationFactor int16, configs TopicConfig, topic string) error {
	adminClient := kadm.NewClient(inst.client)

	// Attempt to create the topic with the provided configuration
	_, err := adminClient.CreateTopic(context.Background(), partitions, replicationFactor, configs.Parse(), topic)
	if err != nil {
		return fmt.Errorf(ErrCreateTopic, topic, err)
	}

	return nil
}

// CreateTopics creates multiple Kafka topics with identical configurations.
// It returns an error if any topic creation fails.
func (inst *Service) CreateTopics(partitions int32, replicationFactor int16, configs TopicConfig, topics ...string) error {
	adminClient := kadm.NewClient(inst.client)

	// Attempt to create the topics with the provided configuration
	_, err := adminClient.CreateTopics(context.Background(), partitions, replicationFactor, configs.Parse(), topics...)
	if err != nil {
		return fmt.Errorf(ErrCreateTopics, topics, err)
	}

	return nil
}

// DeleteGroup deletes a consumer group from Kafka.
// It returns an error if the deletion fails.
func (inst *Service) DeleteGroup(group string) error {
	adminClient := kadm.NewClient(inst.client)

	_, err := adminClient.DeleteGroup(context.Background(), group)
	if err != nil {
		return fmt.Errorf(ErrDeleteGroup, group, err)
	}

	return nil
}

// DeleteGroups deletes multiple consumer groups from Kafka.
// It returns an error if any deletion fails.
func (inst *Service) DeleteGroups(groups ...string) error {
	adminClient := kadm.NewClient(inst.client)

	_, err := adminClient.DeleteGroups(context.Background(), groups...)
	if err != nil {
		return fmt.Errorf(ErrDeleteGroups, groups, err)
	}

	return nil
}

// DeleteTopic deletes a Kafka topic.
// It returns an error if the deletion fails.
func (inst *Service) DeleteTopic(topic string) error {
	adminClient := kadm.NewClient(inst.client)

	_, err := adminClient.DeleteTopic(context.Background(), topic)
	if err != nil {
		return fmt.Errorf(ErrDeleteTopic, topic, err)
	}

	return nil
}

// DeleteTopics deletes multiple Kafka topics.
// It returns an error if any deletion fails.
func (inst *Service) DeleteTopics(topics ...string) error {
	adminClient := kadm.NewClient(inst.client)

	_, err := adminClient.DeleteTopics(context.Background(), topics...)
	if err != nil {
		return fmt.Errorf(ErrDeleteTopics, topics, err)
	}

	return nil
}

// UpdatePartitions sets the number of partitions for existing Kafka topics.
// It returns an error if updating the partitions fails.
func (inst *Service) UpdatePartitions(set int, topics ...string) error {
	adminClient := kadm.NewClient(inst.client)

	_, err := adminClient.UpdatePartitions(context.Background(), set, topics...)
	if err != nil {
		return fmt.Errorf(ErrUpdatePartitions, set, topics, err)
	}

	return nil
}
