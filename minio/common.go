package minio

// DefaultTimeout defines the default request timeout in seconds
const DefaultTimeout int64 = 30 // 30 seconds

// Config represents the configuration for MinIO connection
type Config struct {
	Address   string `yaml:"address"`    // The MinIO server address (e.g., play.min.io)
	AccessKey string `yaml:"access_key"` // The access key for authentication
	SecretKey string `yaml:"secret_key"` // The secret key for authentication
	UseSSL    bool   `yaml:"use_ssl"`    // Indicates whether to use SSL (true for https, false for http)
	Timeout   int64  `yaml:"timeout"`    // The number of seconds before a request times out (optional)
}
