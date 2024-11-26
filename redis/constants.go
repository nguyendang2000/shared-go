package redis

// Default constants for Redis configurations.
// These constants provide sensible defaults for various Redis operations
// and are used when no specific values are provided by the user.
const (
	// DefaultTimeout specifies the default timeout in seconds for Redis operations.
	DefaultTimeout = 5

	// DefaultPoolSize defines the default number of connections in the Redis connection pool.
	DefaultPoolSize = 10

	// DefaultMinIdleConns specifies the default minimum number of idle connections
	// maintained in the connection pool to improve response times.
	DefaultMinIdleConns = 2

	// DefaultStreamID is the default ID for auto-generating stream entries.
	// This is typically used for appending new entries to a stream.
	DefaultStreamID = "*"

	// DefaultLastID represents the default ID for the XRead command,
	// which starts reading messages from the beginning of the stream.
	DefaultLastID = "0"

	// DefaultGroupLastID is the default ID for the XReadGroup command.
	// It indicates that only new messages should be read by the consumer group.
	DefaultGroupLastID = ">"

	// DefaultStartID is the default ID used when creating consumer groups.
	// It specifies that the group should start processing from new messages.
	DefaultStartID = "$"

	// DefaultClaimCount sets the default number of pending messages to claim
	// when using the XAutoClaim command.
	DefaultClaimCount = 100
)
