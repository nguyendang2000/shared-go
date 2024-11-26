package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// AddToStream adds an entry to a Redis stream with the given values.
// By default, an auto-generated ID is used unless a custom ID is provided.
// It returns the message ID of the added entry or an error if the operation fails.
func (inst *Service) AddToStream(stream string, values map[string]interface{}, id ...string) (string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	streamID := DefaultStreamID
	if len(id) > 0 {
		streamID = id[0]
	}

	result, err := inst.client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		ID:     streamID,
		Values: values,
	}).Result()

	if err != nil {
		return "", fmt.Errorf(ErrAddToStream, err)
	}

	return result, nil
}

// ReadFromStream reads entries from a Redis stream starting from a specific message ID.
// It uses XRead and supports blocking. The `lastID` defaults to DefaultLastID if not provided.
// Returns the read messages or an error if the operation fails.
func (inst *Service) ReadFromStream(stream string, count int64, block time.Duration, lastID string) ([]redis.XMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), block+time.Duration(inst.timeout)*time.Second)
	defer cancel()

	if lastID == "" {
		lastID = DefaultLastID
	}

	result, err := inst.client.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, lastID},
		Count:   count,
		Block:   block,
	}).Result()

	if err != nil {
		return nil, fmt.Errorf(ErrReadFromStream, err)
	}

	var messages []redis.XMessage
	if len(result) > 0 {
		messages = result[0].Messages
	}

	return messages, nil
}

// ReadGroupFromStream reads entries from a Redis stream within a consumer group.
// It uses XReadGroup and supports blocking. The `lastID` defaults to DefaultGroupLastID if not provided.
// Optionally, messages can be auto-acknowledged after reading. Returns the read messages or an error if the operation fails.
func (inst *Service) ReadGroupFromStream(stream, group, consumer string, count int64, block time.Duration, lastID string, autoAck bool) ([]redis.XMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), block+time.Duration(inst.timeout)*time.Second)
	defer cancel()

	if lastID == "" {
		lastID = DefaultGroupLastID
	}

	result, err := inst.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumer,
		Streams:  []string{stream, lastID},
		Count:    count,
		Block:    block,
	}).Result()

	if err != nil {
		return nil, fmt.Errorf(ErrReadGroupFromStream, err)
	}

	var messages []redis.XMessage
	if len(result) > 0 {
		messages = result[0].Messages
	}

	if autoAck {
		for _, msg := range messages {
			_, ackErr := inst.AcknowledgeMessage(stream, group, msg.ID)
			if ackErr != nil {
				return nil, fmt.Errorf(ErrAcknowledgeMessage, ackErr)
			}
		}
	}

	return messages, nil
}

// AcknowledgeMessage acknowledges a message in a consumer group by its ID.
// It returns the number of acknowledged messages or an error if the operation fails.
func (inst *Service) AcknowledgeMessage(stream, group, id string) (int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	result, err := inst.client.XAck(ctx, stream, group, id).Result()
	if err != nil {
		return 0, fmt.Errorf(ErrAcknowledgeMessage, err)
	}

	return result, nil
}

// CreateConsumerGroup creates a new consumer group for a Redis stream.
// The starting ID defaults to DefaultStartID if not provided. Returns an error if the operation fails.
func (inst *Service) CreateConsumerGroup(stream, group, startID string) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	if startID == "" {
		startID = DefaultStartID
	}

	err := inst.client.XGroupCreateMkStream(ctx, stream, group, startID).Err()
	if err != nil {
		return fmt.Errorf(ErrCreateConsumerGroup, err)
	}

	return nil
}

// ClaimPendingMessages claims pending messages in a consumer group that have exceeded the minimum idle time.
// If `count` is less than or equal to 0, DefaultClaimCount is used. Uses XAutoClaim for claiming messages.
// Optionally, messages can be auto-acknowledged after claiming. Returns the claimed messages, the new start ID, or an error.
func (inst *Service) ClaimPendingMessages(stream, group, consumer string, minIdleTime time.Duration, startID string, count int64, autoAck bool) ([]redis.XMessage, string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	if count <= 0 {
		count = DefaultClaimCount
	}

	result, start, err := inst.client.XAutoClaim(ctx, &redis.XAutoClaimArgs{
		Stream:   stream,
		Group:    group,
		Consumer: consumer,
		MinIdle:  minIdleTime,
		Start:    startID,
		Count:    count,
	}).Result()

	if err != nil {
		return nil, "", fmt.Errorf(ErrClaimPendingMessages, err)
	}

	if autoAck {
		for _, msg := range result {
			_, ackErr := inst.AcknowledgeMessage(stream, group, msg.ID)
			if ackErr != nil {
				return nil, "", fmt.Errorf(ErrAcknowledgeMessage, ackErr)
			}
		}
	}

	return result, start, nil
}
