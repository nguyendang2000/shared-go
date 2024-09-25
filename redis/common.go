package redis

import "crypto/tls"

// Default constants for Redis configurations
const (
	DefaultTimeout      = 5   // Default timeout in seconds for Redis operations
	DefaultPoolSize     = 10  // Default pool size if not provided
	DefaultMinIdleConns = 2   // Default minimum idle connections if not provided
	DefaultStreamID     = "*" // Default ID for auto-generating stream entries
	DefaultLastID       = "0" // Default ID for XRead (reading from the start)
	DefaultGroupLastID  = ">" // Default ID for XReadGroup (only new messages)
	DefaultStartID      = "$" // Default ID for creating consumer groups (start from new messages)
	DefaultClaimCount   = 100 // Default count for claiming pending messages
)

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
