package redis

// Error messages for Redis Service operations.
const (
	ErrPingRedis = "failed to ping Redis: %w"
	ErrGet       = "failed to get value for key: %w"
	ErrSet       = "failed to set value for key: %w"
	ErrDelete    = "failed to delete key: %w"
	ErrExists    = "failed to check key existence: %w"
	ErrExpire    = "failed to set expiration on key: %w"
	ErrTTL       = "failed to get TTL for key: %w"
	ErrIncr      = "failed to increment key: %w"
	ErrIncrBy    = "failed to increment key: %w"
)

// Error messages for Redis Stream operations.
const (
	ErrAddToStream          = "failed to add entry to stream: %w"
	ErrReadFromStream       = "failed to read from stream: %w"
	ErrReadGroupFromStream  = "failed to read from consumer group: %w"
	ErrAcknowledgeMessage   = "failed to acknowledge message: %w"
	ErrCreateConsumerGroup  = "failed to create consumer group: %w"
	ErrClaimPendingMessages = "failed to claim pending messages: %w"
)
