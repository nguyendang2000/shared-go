package redis

// Error messages for Redis Service operations.
// These constants define error messages for general Redis operations,
// formatted with placeholders to allow dynamic values.
const (
	// ErrPingRedis is returned when a connection ping to Redis fails.
	ErrPingRedis = "failed to ping Redis: %w"

	// ErrGet is returned when a GET operation for a key fails.
	ErrGet = "failed to get key %s: %w"

	// ErrSet is returned when a SET operation for a key fails.
	ErrSet = "failed to set key %s: %w"

	// ErrDelete is returned when a DELETE operation for one or more keys fails.
	ErrDelete = "failed to delete keys %+v: %w"

	// ErrExists is returned when checking the existence of one or more keys fails.
	ErrExists = "failed to check existence of keys %+v: %w"

	// ErrExpire is returned when setting an expiration time for a key fails.
	ErrExpire = "failed to set expiration for key %s: %w"

	// ErrTTL is returned when retrieving the TTL of a key fails.
	ErrTTL = "failed to get TTL of key %s: %w"

	// ErrIncr is returned when incrementing a key by 1 fails.
	ErrIncr = "failed to increment key %s: %w"

	// ErrIncrBy is returned when incrementing a key by a specified value fails.
	ErrIncrBy = "failed to increment key %s by %d: %w"
)

// Error messages for Redis Hash operations.
// These constants define error messages for operations involving Redis hash data types.
const (
	// ErrHGet is returned when retrieving a field from a hash fails.
	ErrHGet = "failed to get field %s in key %s: %w"

	// ErrHGetAll is returned when retrieving all fields and values from a hash fails.
	ErrHGetAll = "failed to get all fields in key %s: %w"

	// ErrHSet is returned when setting fields and values in a hash fails.
	ErrHSet = "failed to set fields and values for key %s: %w"

	// ErrHDel is returned when deleting fields from a hash fails.
	ErrHDel = "failed to delete fields in key %s: %w"

	// ErrHExists is returned when checking if a field exists in a hash fails.
	ErrHExists = "failed to check if field %s exists in key %s: %w"

	// ErrHExpire is returned when setting expiration for hash fields fails.
	ErrHExpire = "failed to set expiration for key %s: %w"

	// ErrHTTL is returned when retrieving TTL for hash fields fails.
	ErrHTTL = "failed to get TTL for fields in key %s: %w"

	// ErrHIncrBy is returned when incrementing a hash field by a specified value fails.
	ErrHIncrBy = "failed to increment field %s by %d in key %s: %w"

	// ErrHKeys is returned when retrieving all field names from a hash fails.
	ErrHKeys = "failed to get fields in key %s: %w"

	// ErrHVals is returned when retrieving all values from a hash fails.
	ErrHVals = "failed to get values in key %s: %w"

	// ErrHLen is returned when retrieving the length (number of fields) of a hash fails.
	ErrHLen = "failed to get length of key %s: %w"
)

// Error messages for Redis Stream operations.
// These constants define error messages for operations involving Redis streams.
const (
	// ErrAddToStream is returned when adding an entry to a Redis stream fails.
	ErrAddToStream = "failed to add entry to stream: %w"

	// ErrReadFromStream is returned when reading from a Redis stream fails.
	ErrReadFromStream = "failed to read from stream: %w"

	// ErrReadGroupFromStream is returned when reading from a Redis consumer group fails.
	ErrReadGroupFromStream = "failed to read from consumer group: %w"

	// ErrAcknowledgeMessage is returned when acknowledging a message in a Redis stream fails.
	ErrAcknowledgeMessage = "failed to acknowledge message: %w"

	// ErrCreateConsumerGroup is returned when creating a Redis consumer group fails.
	ErrCreateConsumerGroup = "failed to create consumer group: %w"

	// ErrClaimPendingMessages is returned when claiming pending messages in a Redis stream fails.
	ErrClaimPendingMessages = "failed to claim pending messages: %w"
)
