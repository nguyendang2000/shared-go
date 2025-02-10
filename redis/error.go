package redis

import "github.com/redis/go-redis/v9"

// ErrNil represents a Redis Nil response, typically indicating a missing key.
var ErrNil = error(redis.Nil)

// Error messages for Redis Service operations.
// These constants define error messages for general Redis key-value operations.
const (
	// ErrPingRedis occurs when a connection ping to Redis fails.
	ErrPingRedis = "failed to ping Redis: %w"

	// ErrGet occurs when retrieving the value of a key fails.
	ErrGet = "failed to get key %s: %w"

	// ErrSet occurs when setting the value of a key fails.
	ErrSet = "failed to set key %s: %w"

	// ErrDelete occurs when attempting to delete one or more keys fails.
	ErrDelete = "failed to delete keys %+v: %w"

	// ErrExists occurs when checking the existence of one or more keys fails.
	ErrExists = "failed to check existence of keys %+v: %w"

	// ErrExpire occurs when setting an expiration time for a key fails.
	ErrExpire = "failed to set expiration for key %s: %w"

	// ErrTTL occurs when retrieving the time-to-live (TTL) of a key fails.
	ErrTTL = "failed to get TTL of key %s: %w"

	// ErrIncr occurs when incrementing the value of a key by 1 fails.
	ErrIncr = "failed to increment key %s: %w"

	// ErrIncrBy occurs when incrementing the value of a key by a specified amount fails.
	ErrIncrBy = "failed to increment key %s by %d: %w"
)

// Error messages for Redis Hash operations.
// These constants define error messages for Redis hash-related operations.
const (
	// ErrHGet occurs when retrieving a specific field from a hash fails.
	ErrHGet = "failed to get field %s in key %s: %w"

	// ErrHGetAll occurs when retrieving all fields and values from a hash fails.
	ErrHGetAll = "failed to get all fields in key %s: %w"

	// ErrHSet occurs when setting fields and values in a hash fails.
	ErrHSet = "failed to set fields and values for key %s: %w"

	// ErrHDel occurs when deleting specific fields from a hash fails.
	ErrHDel = "failed to delete fields in key %s: %w"

	// ErrHExists occurs when checking if a specific field exists in a hash fails.
	ErrHExists = "failed to check if field %s exists in key %s: %w"

	// ErrHExpire occurs when setting an expiration time for a hash fails.
	ErrHExpire = "failed to set expiration for key %s: %w"

	// ErrHTTL occurs when retrieving the TTL for a hash fails.
	ErrHTTL = "failed to get TTL for fields in key %s: %w"

	// ErrHIncrBy occurs when incrementing a field in a hash by a specified value fails.
	ErrHIncrBy = "failed to increment field %s by %d in key %s: %w"

	// ErrHKeys occurs when retrieving all field names from a hash fails.
	ErrHKeys = "failed to get fields in key %s: %w"

	// ErrHVals occurs when retrieving all values from a hash fails.
	ErrHVals = "failed to get values in key %s: %w"

	// ErrHLen occurs when retrieving the number of fields in a hash fails.
	ErrHLen = "failed to get length of key %s: %w"
)

// Error messages for Redis Stream operations.
// These constants define error messages for Redis stream-related operations.
const (
	// ErrAddToStream occurs when adding an entry to a Redis stream fails.
	ErrAddToStream = "failed to add entry to stream: %w"

	// ErrReadFromStream occurs when reading from a Redis stream fails.
	ErrReadFromStream = "failed to read from stream: %w"

	// ErrReadGroupFromStream occurs when reading messages from a Redis consumer group fails.
	ErrReadGroupFromStream = "failed to read from consumer group: %w"

	// ErrAcknowledgeMessage occurs when acknowledging a message in a Redis stream fails.
	ErrAcknowledgeMessage = "failed to acknowledge message: %w"

	// ErrCreateConsumerGroup occurs when creating a Redis consumer group fails.
	ErrCreateConsumerGroup = "failed to create consumer group: %w"

	// ErrClaimPendingMessages occurs when attempting to claim pending messages in a Redis stream fails.
	ErrClaimPendingMessages = "failed to claim pending messages: %w"
)

// Error messages for Redis Sorted Set (ZSET) operations.
// These constants define error messages for Redis sorted set-related operations.
const (
	// ErrZAdd occurs when adding a member to a sorted set fails.
	ErrZAdd = "failed to add to sorted set: %w"

	// ErrZAddArgs occurs when adding multiple members to a sorted set with arguments fails.
	ErrZAddArgs = "failed to add to sorted set with args: %w"

	// ErrZCard occurs when retrieving the cardinality (size) of a sorted set fails.
	ErrZCard = "failed to get sorted set cardinality: %w"

	// ErrZCount occurs when counting the number of members in a sorted set within a score range fails.
	ErrZCount = "failed to count sorted set members: %w"

	// ErrZIncrBy occurs when incrementing the score of a member in a sorted set fails.
	ErrZIncrBy = "failed to increment score of sorted set member: %w"

	// ErrZRange occurs when retrieving a range of members from a sorted set by rank fails.
	ErrZRange = "failed to get sorted set range: %w"

	// ErrZRangeByLex occurs when retrieving a range of members from a sorted set by lexicographic order fails.
	ErrZRangeByLex = "failed to get sorted set range by lex: %w"

	// ErrZRangeByScore occurs when retrieving a range of members from a sorted set by score fails.
	ErrZRangeByScore = "failed to get sorted set range by score: %w"

	// ErrZRank occurs when retrieving the rank (position) of a member in a sorted set fails.
	ErrZRank = "failed to get member rank in sorted set: %w"

	// ErrZRem occurs when removing one or more members from a sorted set fails.
	ErrZRem = "failed to remove sorted set members: %w"

	// ErrZRemRangeByLex occurs when removing members from a sorted set within a lexicographic range fails.
	ErrZRemRangeByLex = "failed to remove sorted set members by lex: %w"

	// ErrZRemRangeByRank occurs when removing members from a sorted set within a rank range fails.
	ErrZRemRangeByRank = "failed to remove sorted set members by rank: %w"

	// ErrZRemRangeByScore occurs when removing members from a sorted set within a score range fails.
	ErrZRemRangeByScore = "failed to remove sorted set members by score: %w"

	// ErrZScore occurs when retrieving the score of a member in a sorted set fails.
	ErrZScore = "failed to get member score in sorted set: %w"
)
