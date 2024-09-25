package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// AddToStream adds an entry to a Redis stream with the given values and returns the message ID.
// By default, it uses an auto-generated ID unless a custom ID is provided.
func (inst *Service) AddToStream(stream string, values map[string]interface{}, id ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
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
// Supports blocking and uses XRead. The `lastID` is optional and defaults to DefaultLastID.
// Returns an empty slice if no messages are found.
func (inst *Service) ReadFromStream(stream string, count int64, block time.Duration, lastID string) ([]redis.XMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), block+time.Duration(inst.timeout)*time.Second)
	defer cancel()

	if lastID == "" {
		lastID = DefaultLastID // Default for XRead
	}

	result, err := inst.client.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, lastID},
		Count:   count,
		Block:   block,
	}).Result()

	if err != nil {
		return nil, fmt.Errorf(ErrReadFromStream, err)
	}

	if len(result) > 0 {
		return result[0].Messages, nil
	}

	return []redis.XMessage{}, nil
}

// ReadGroupFromStream reads entries from a Redis stream within a consumer group, starting from a specific message ID.
// Supports blocking and uses XReadGroup. The `lastID` is optional and defaults to DefaultGroupLastID.
// Optionally, it can auto-acknowledge messages upon reading.
// Returns an empty slice if no messages are found.
func (inst *Service) ReadGroupFromStream(stream, group, consumer string, count int64, block time.Duration, lastID string, autoAck bool) ([]redis.XMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), block+time.Duration(inst.timeout)*time.Second)
	defer cancel()

	if lastID == "" {
		lastID = DefaultGroupLastID // Default for XReadGroup
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

	messages := result[0].Messages
	if autoAck {
		for _, msg := range messages {
			_, ackErr := inst.AcknowledgeMessage(stream, group, msg.ID)
			if ackErr != nil {
				return nil, fmt.Errorf(ErrAcknowledgeMessage, ackErr)
			}
		}
	}

	if len(messages) > 0 {
		return messages, nil
	}

	return []redis.XMessage{}, nil
}

// AcknowledgeMessage acknowledges a message in a consumer group by its ID.
// It returns the number of messages acknowledged or an error if the operation fails.
func (inst *Service) AcknowledgeMessage(stream, group, id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	result, err := inst.client.XAck(ctx, stream, group, id).Result()
	if err != nil {
		return 0, fmt.Errorf(ErrAcknowledgeMessage, err)
	}

	return result, nil
}

// CreateConsumerGroup creates a consumer group for a Redis stream with the given group name.
// The starting ID defaults to DefaultStartID but can be customized.
// It returns an error if the group cannot be created.
func (inst *Service) CreateConsumerGroup(stream, group, startID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	if startID == "" {
		startID = DefaultStartID // Default start ID
	}

	err := inst.client.XGroupCreateMkStream(ctx, stream, group, startID).Err()
	if err != nil {
		return fmt.Errorf(ErrCreateConsumerGroup, err)
	}

	return nil
}

// ClaimPendingMessages claims pending messages in a consumer group that have exceeded the minimum idle time.
// If `count` is less than or equal to 0, it defaults to DefaultClaimCount. It leverages XAutoClaim to claim messages automatically.
// Optionally, it can auto-acknowledge the claimed messages after processing.
func (inst *Service) ClaimPendingMessages(stream, group, consumer string, minIdleTime time.Duration, startID string, count int64, autoAck bool) ([]redis.XMessage, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	if count <= 0 {
		count = DefaultClaimCount // Default count
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
