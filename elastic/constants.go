package elastic

// DefaultTimeout specifies the default timeout duration for Elasticsearch requests, in milliseconds.
// This timeout is applied when no custom timeout is provided in the service configuration.
const DefaultTimeout int64 = 3000 // Timeout in milliseconds (3 seconds)
