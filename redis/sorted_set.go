package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Z represents a member of a sorted set in Redis.
type Z redis.Z

// ZAddArgs defines the arguments for adding multiple members to a sorted set.
type ZAddArgs struct {
	NX      bool // Only add elements that do not already exist.
	XX      bool // Only update elements that already exist.
	LT      bool // Only update elements if the new score is less than the current score.
	GT      bool // Only update elements if the new score is greater than the current score.
	Ch      bool // Return the number of elements added or updated.
	Members []Z  // Members to be added or updated.
}

// RangeArgs represents the range parameters for Redis sorted set operations.
type RangeArgs struct {
	Gt  *float64 // Greater than (exclusive).
	Gte *float64 // Greater than or equal to (inclusive).
	Lt  *float64 // Less than (exclusive).
	Lte *float64 // Less than or equal to (inclusive).
}

// Parse converts the range arguments into Redis-compatible range strings.
// If `lex` is false, the function returns numeric range boundaries for score-based queries.
// If `lex` is true, the function returns lexicographic boundaries for string-based queries.
func (inst *RangeArgs) Parse(lex bool) (string, string) {
	var min, max string

	// Determine the lower bound
	if inst.Gt != nil {
		min = fmt.Sprintf("(%g", *inst.Gt) // Exclusive lower bound
	} else if inst.Gte != nil {
		if !lex {
			min = fmt.Sprintf("%g", *inst.Gte) // Inclusive numeric lower bound
		} else {
			min = fmt.Sprintf("[%g", *inst.Gte) // Inclusive lexicographic lower bound
		}
	} else {
		if !lex {
			min = "-inf" // No lower limit for numeric range
		} else {
			min = "-" // No lower limit for lexicographic range
		}
	}

	// Determine the upper bound
	if inst.Lt != nil {
		max = fmt.Sprintf("(%g", *inst.Lt) // Exclusive upper bound
	} else if inst.Lte != nil {
		if !lex {
			max = fmt.Sprintf("%g", *inst.Lte) // Inclusive numeric upper bound
		} else {
			max = fmt.Sprintf("[%g", *inst.Lte) // Inclusive lexicographic upper bound
		}
	} else {
		if !lex {
			max = "+inf" // No upper limit for numeric range
		} else {
			max = "+" // No upper limit for lexicographic range
		}
	}

	return min, max
}

// ZAdd adds a single member to a sorted set with the specified score.
// If the member already exists, its score is updated.
func (inst *Service) ZAdd(key string, member interface{}, score float64) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	if err := inst.client.ZAdd(ctx, key, redis.Z{Member: member, Score: score}).Err(); err != nil {
		return fmt.Errorf(ErrZAdd, err)
	}

	return nil
}

// ZAddArgs adds multiple members to a sorted set with configurable options.
// - NX: Only add elements that do not already exist.
// - XX: Only update elements that already exist.
// - LT: Only update elements if the new score is lower than the current score.
// - GT: Only update elements if the new score is higher than the current score.
// - Ch: Returns the number of elements that were changed (added or updated).
func (inst *Service) ZAddArgs(key string, args ZAddArgs) (int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	// Prepare arguments for the ZADD command
	addArgs := redis.ZAddArgs{
		XX:      args.XX,
		Ch:      args.Ch,
		Members: make([]redis.Z, len(args.Members)),
	}

	// Convert custom Z members into redis.Z format
	for i := range args.Members {
		addArgs.Members[i] = redis.Z(args.Members[i])
	}

	// Apply only one of NX, LT, or GT options
	switch {
	case args.NX:
		addArgs.NX = args.NX
	case args.LT:
		addArgs.LT = args.LT
	case args.GT:
		addArgs.GT = args.GT
	}

	// Execute ZADD with the provided arguments
	count, err := inst.client.ZAddArgs(ctx, key, addArgs).Result()
	if err != nil {
		return count, fmt.Errorf(ErrZAddArgs, err)
	}

	return count, nil
}

// ZCard returns the number of members in the sorted set stored at the given key.
// If the key does not exist, it is treated as an empty sorted set and returns 0.
func (inst *Service) ZCard(key string) (int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	card, err := inst.client.ZCard(ctx, key).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrZCard, err)
	}

	return card, nil
}

// ZCount returns the number of elements in the sorted set within the specified score range.
// The range boundaries are determined using the RangeArgs struct, which supports both inclusive and exclusive limits.
func (inst *Service) ZCount(key string, rangeArgs RangeArgs) (int64, error) {
	min, max := rangeArgs.Parse(false) // Convert RangeArgs into Redis-compatible range strings.
	ctx, cancel := inst.getTimeout()
	defer cancel()

	count, err := inst.client.ZCount(ctx, key, min, max).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrZCount, err)
	}

	return count, nil
}

// ZIncrBy increments the score of a member in the sorted set by the given increment.
// If the member does not exist, it is added to the sorted set with the specified score.
func (inst *Service) ZIncrBy(key string, increment float64, member string) (float64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	score, err := inst.client.ZIncrBy(ctx, key, increment, member).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrZIncrBy, err)
	}

	return score, nil
}

// ZRange retrieves a range of members from a sorted set based on their rank (position).
// - `start` and `stop` define the range of ranks to retrieve, where `0` is the first member.
// - If `limit` is `0`, it is set to `-1` to return all matching elements.
// - `offset` specifies how many elements to skip before returning results.
func (inst *Service) ZRange(key string, start, stop, limit, offset int64) ([]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	if limit == 0 {
		limit = -1
	}

	members, err := inst.client.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:    key,
		Start:  start,
		Stop:   stop,
		Offset: offset,
		Count:  limit,
	}).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrZRange, err)
	}

	return members, nil
}

// ZRangeByLex retrieves a range of members from a sorted set in lexicographic order.
// - `rangeArgs` defines the lexicographic range using inclusive/exclusive boundaries.
// - If `limit` is `0`, it is set to `-1` to return all matching elements.
// - `offset` specifies how many elements to skip before returning results.
func (inst *Service) ZRangeByLex(key string, rangeArgs RangeArgs, limit, offset int64) ([]string, error) {
	min, max := rangeArgs.Parse(true) // Convert range arguments to lexicographic format.

	if limit == 0 {
		limit = -1
	}

	ctx, cancel := inst.getTimeout()
	defer cancel()

	members, err := inst.client.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:    key,
		Start:  min,
		Stop:   max,
		ByLex:  true, // Enable lexicographic ordering.
		Offset: offset,
		Count:  limit,
	}).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrZRangeByLex, err)
	}

	return members, nil
}

// ZRangeByScore retrieves a range of members from a sorted set based on their score.
// - `rangeArgs` defines the score range using inclusive/exclusive boundaries.
// - If `limit` is `0`, it is set to `-1` to return all matching elements.
// - `offset` specifies how many elements to skip before returning results.
func (inst *Service) ZRangeByScore(key string, rangeArgs RangeArgs, limit, offset int64) ([]string, error) {
	min, max := rangeArgs.Parse(false) // Convert range arguments to numeric score format.

	if limit == 0 {
		limit = -1
	}

	ctx, cancel := inst.getTimeout()
	defer cancel()

	members, err := inst.client.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:     key,
		Start:   min,
		Stop:    max,
		ByScore: true, // Enable score-based ordering.
		Offset:  offset,
		Count:   limit,
	}).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrZRangeByScore, err)
	}

	return members, nil
}

// ZRank returns the rank (zero-based index) of a member in a sorted set.
func (inst *Service) ZRank(key, member string) (int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	rank, err := inst.client.ZRank(ctx, key, member).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrZRank, err)
	}

	return rank, nil
}

// ZRem removes one or more members from a sorted set.
func (inst *Service) ZRem(key string, members ...interface{}) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	if err := inst.client.ZRem(ctx, key, members...).Err(); err != nil {
		return fmt.Errorf(ErrZRem, err)
	}

	return nil
}

// ZRemRangeByLex removes members in a sorted set within a lexicographic range.
func (inst *Service) ZRemRangeByLex(key string, rangeArgs RangeArgs) error {
	min, max := rangeArgs.Parse(true)
	ctx, cancel := inst.getTimeout()
	defer cancel()

	if err := inst.client.ZRemRangeByLex(ctx, key, min, max).Err(); err != nil {
		return fmt.Errorf(ErrZRemRangeByLex, err)
	}

	return nil
}

// ZRemRangeByRank removes members in a sorted set within a given rank range.
func (inst *Service) ZRemRangeByRank(key string, start, stop int64) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	if err := inst.client.ZRemRangeByRank(ctx, key, start, stop).Err(); err != nil {
		return fmt.Errorf(ErrZRemRangeByRank, err)
	}

	return nil
}

// ZRemRangeByScore removes members in a sorted set within a given score range.
func (inst *Service) ZRemRangeByScore(key string, rangeArgs RangeArgs) error {
	min, max := rangeArgs.Parse(false)
	ctx, cancel := inst.getTimeout()
	defer cancel()

	if err := inst.client.ZRemRangeByScore(ctx, key, min, max).Err(); err != nil {
		return fmt.Errorf(ErrZRemRangeByScore, err)
	}

	return nil
}

// ZScore retrieves the score of a member in a sorted set.
func (inst *Service) ZScore(key, member string) (float64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	score, err := inst.client.ZScore(ctx, key, member).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrZScore, err)
	}

	return score, nil
}
