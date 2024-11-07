package redis

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
