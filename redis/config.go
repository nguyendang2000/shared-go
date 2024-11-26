package redis

import "crypto/tls"

// Config represents the configuration settings for connecting to a Redis instance.
// This struct supports YAML-based configuration for seamless integration with external config files.
type Config struct {
	// Address specifies the Redis server address in the format "host:port".
	Address string `yaml:"address"`

	// Password provides the optional password for authenticating with the Redis server.
	Password string `yaml:"password"`

	// DB indicates the Redis database number to use. The default database is 0.
	DB int `yaml:"db"`

	// TLSConfig contains the TLS settings for establishing secure connections.
	// This field is optional and should be set only when a secure connection is required.
	TLSConfig *tls.Config `yaml:"tls_config"`

	// PoolSize defines the maximum number of connections allowed in the Redis connection pool.
	// A larger pool size can handle more concurrent requests but consumes more resources.
	PoolSize int `yaml:"pool_size"`

	// MinIdleConns specifies the minimum number of idle connections to maintain in the pool.
	// Keeping idle connections pre-warmed improves response times for new requests.
	MinIdleConns int `yaml:"min_idle_conns"`

	// Timeout sets the maximum time, in seconds, for connection operations before they fail.
	// This includes connection attempts and read/write operations.
	Timeout int64 `yaml:"timeout"`
}
