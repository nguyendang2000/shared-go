package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Service represents a wrapper around a Redis client connection.
type Service struct {
	client  *redis.Client // Redis client connection instance
	timeout int64         // Timeout for Redis operations, in seconds
}

// NewService initializes a Redis connection using the provided configuration and a user-defined context.
// It returns a Service struct that contains the Redis client or an error if the connection fails.
// The service will automatically close the Redis connection when the passed context is canceled.
func NewService(ctx context.Context, conf Config) (*Service, error) {
	// Set timeout, use default if not provided.
	timeout := conf.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	// Set pool size, use default if not provided.
	poolSize := conf.PoolSize
	if poolSize == 0 {
		poolSize = DefaultPoolSize
	}

	// Set minimum idle connections, use default if not provided.
	minIdleConns := conf.MinIdleConns
	if minIdleConns == 0 {
		minIdleConns = DefaultMinIdleConns
	}

	// Configure Redis client options.
	options := &redis.Options{
		Addr:         conf.Address,                         // Address in the format "host:port"
		Password:     conf.Password,                        // Password for Redis authentication
		DB:           conf.DB,                              // Redis database number
		TLSConfig:    conf.TLSConfig,                       // TLS configuration for secure connections (optional)
		PoolSize:     poolSize,                             // Maximum number of connections in the pool
		MinIdleConns: minIdleConns,                         // Minimum number of idle connections
		DialTimeout:  time.Duration(timeout) * time.Second, // Timeout for establishing new connections
		ReadTimeout:  time.Duration(timeout) * time.Second, // Timeout for reading from Redis
		WriteTimeout: time.Duration(timeout) * time.Second, // Timeout for writing to Redis
	}

	// Create a new Redis client.
	client := redis.NewClient(options)

	// Create the Service instance and store the timeout.
	service := &Service{
		client:  client,
		timeout: timeout,
	}

	// Goroutine to close the Redis client when the context is canceled.
	go func() {
		<-ctx.Done()    // Wait for the context to be canceled.
		service.Close() // Close the Redis connection.
	}()

	// Ping the Redis server to verify the connection.
	if err := service.Ping(); err != nil {
		return nil, fmt.Errorf(ErrPingRedis, err)
	}

	return service, nil
}

// Client returns the underlying Redis client instance.
func (inst *Service) Client() *redis.Client {
	return inst.client
}

// Ping tests the connection to the Redis server by sending a ping command.
// It uses the stored timeout in the Service struct and returns an error if the ping fails.
func (inst *Service) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	err := inst.client.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf(ErrPingRedis, err)
	}

	return nil
}

// Close gracefully closes the Redis connection.
func (inst *Service) Close() error {
	return inst.client.Close()
}

// Get retrieves the value associated with the given key from Redis.
// It returns an error if the operation fails.
func (inst *Service) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	result, err := inst.client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf(ErrGet, err)
	}

	return result, nil
}

// Set stores a key-value pair in Redis with an optional expiration time.
// It returns an error if the operation fails.
func (inst *Service) Set(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	err := inst.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf(ErrSet, err)
	}

	return nil
}

// Del deletes one or more keys from Redis.
// It returns the number of keys deleted and an error if the operation fails.
func (inst *Service) Del(keys ...string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	result, err := inst.client.Del(ctx, keys...).Result()
	if err != nil {
		return 0, fmt.Errorf(ErrDelete, err)
	}

	return result, nil
}

// Exists checks if one or more keys exist in Redis.
// It returns the number of existing keys and an error if the operation fails.
func (inst *Service) Exists(keys ...string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	result, err := inst.client.Exists(ctx, keys...).Result()
	if err != nil {
		return 0, fmt.Errorf(ErrExists, err)
	}

	return result, nil
}

// Expire sets a timeout on a specific key, after which the key will expire.
// It returns an error if the operation fails.
func (inst *Service) Expire(key string, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	err := inst.client.Expire(ctx, key, expiration).Err()
	if err != nil {
		return fmt.Errorf(ErrExpire, err)
	}

	return nil
}

// TTL retrieves the time-to-live (TTL) remaining for a specific key.
// It returns the TTL duration and an error if the operation fails.
func (inst *Service) TTL(key string) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	ttl, err := inst.client.TTL(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf(ErrTTL, err)
	}

	return ttl, nil
}

// Incr increments the integer value of a key by one.
// It returns the new value and an error if the operation fails.
func (inst *Service) Incr(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	result, err := inst.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf(ErrIncr, err)
	}

	return result, nil
}
