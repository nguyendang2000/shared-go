package redis

import "crypto/tls"

// Config holds the configuration options required to connect to a Redis instance.
// It includes YAML tags for easy configuration using a YAML file.
type Config struct {
	Address      string      `yaml:"address"`        // Redis server address in the format "host:port"
	Password     string      `yaml:"password"`       // Redis password (optional)
	DB           int         `yaml:"db"`             // Redis database number (default 0)
	TLSConfig    *tls.Config `yaml:"tls_config"`     // TLS configuration for secure connections (optional)
	PoolSize     int         `yaml:"pool_size"`      // Maximum number of connections in the connection pool
	MinIdleConns int         `yaml:"min_idle_conns"` // Minimum number of idle connections in the pool
	Timeout      int64       `yaml:"timeout"`        // Timeout for connection operations, in seconds
}
