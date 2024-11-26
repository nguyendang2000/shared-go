package redis

import (
	"fmt"
	"time"
)

// HGet retrieves the value of a specific field in a Redis hash.
// It uses the stored timeout in the Service struct and returns the value or an error if the operation fails.
func (inst *Service) HGet(key, field string) (string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	result, err := inst.client.HGet(ctx, key, field).Result()
	if err != nil {
		return "", fmt.Errorf(ErrHGet, field, key, err)
	}

	return result, nil
}

// HGetAll retrieves all fields and their values from a Redis hash.
// It uses the stored timeout in the Service struct and returns a map of field-value pairs or an error if the operation fails.
func (inst *Service) HGetAll(key string) (map[string]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	result, err := inst.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrHGetAll, key, err)
	}

	return result, nil
}

// HSet sets multiple fields and their values in a Redis hash.
// It uses the stored timeout in the Service struct and returns an error if the operation fails.
func (inst *Service) HSet(key string, fieldValues map[string]interface{}) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	err := inst.client.HSet(ctx, key, fieldValues).Err()
	if err != nil {
		return fmt.Errorf(ErrHSet, key, err)
	}

	return nil
}

// HDel deletes specific fields from a Redis hash.
// It uses the stored timeout in the Service struct and returns an error if the operation fails.
func (inst *Service) HDel(key string, fields ...string) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	err := inst.client.HDel(ctx, key, fields...).Err()
	if err != nil {
		return fmt.Errorf(ErrHDel, key, err)
	}

	return nil
}

// HExists checks if a specific field exists in a Redis hash.
// It uses the stored timeout in the Service struct and returns true if the field exists, or false with an error if the operation fails.
func (inst *Service) HExists(key, field string) (bool, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	exists, err := inst.client.HExists(ctx, key, field).Result()
	if err != nil {
		return false, fmt.Errorf(ErrHExists, field, key, err)
	}

	return exists, nil
}

// HExpire sets a timeout for fields in a Redis hash.
// It uses the stored timeout in the Service struct and returns an error if the operation fails.
func (inst *Service) HExpire(key string, expiration time.Duration, fields ...string) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	err := inst.client.HExpire(ctx, key, expiration, fields...).Err()
	if err != nil {
		return fmt.Errorf(ErrHExpire, key, err)
	}

	return nil
}

// HTTL retrieves the time-to-live (TTL) for fields in a Redis hash.
// It uses the stored timeout in the Service struct and returns a slice of TTL durations or an error if the operation fails.
func (inst *Service) HTTL(key string, fields ...string) ([]int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	result, err := inst.client.HTTL(ctx, key, fields...).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrHTTL, key, err)
	}

	return result, nil
}

// HIncrBy increments the value of a specific field in a Redis hash by the given amount.
// It uses the stored timeout in the Service struct and returns the new value or an error if the operation fails.
func (inst *Service) HIncrBy(key, field string, increment int64) (int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	result, err := inst.client.HIncrBy(ctx, key, field, increment).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrHIncrBy, field, increment, key, err)
	}

	return result, nil
}

// HKeys retrieves all field names from a Redis hash.
// It uses the stored timeout in the Service struct and returns a slice of field names or an error if the operation fails.
func (inst *Service) HKeys(key string) ([]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	result, err := inst.client.HKeys(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrHKeys, key, err)
	}

	return result, nil
}

// HVals retrieves all values from a Redis hash.
// It uses the stored timeout in the Service struct and returns a slice of values or an error if the operation fails.
func (inst *Service) HVals(key string) ([]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	result, err := inst.client.HVals(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrHVals, key, err)
	}

	return result, nil
}

// HLen retrieves the number of fields in a Redis hash.
// It uses the stored timeout in the Service struct and returns the field count or an error if the operation fails.
func (inst *Service) HLen(key string) (int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	result, err := inst.client.HLen(ctx, key).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrHLen, key, err)
	}

	return result, nil
}
