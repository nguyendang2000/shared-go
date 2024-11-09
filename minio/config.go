package minio

// Config represents the configuration settings required for connecting to a MinIO server.
type Config struct {
	// Address specifies the MinIO server address (e.g., play.min.io).
	Address string `yaml:"address"`

	// AccessKey is the access key used for authentication with the MinIO server.
	AccessKey string `yaml:"access_key"`

	// SecretKey is the secret key used for authentication with the MinIO server.
	SecretKey string `yaml:"secret_key"`

	// UseSSL indicates whether to use SSL for the connection.
	// If true, the connection uses HTTPS; if false, it uses HTTP.
	UseSSL bool `yaml:"use_ssl"`

	// Timeout defines the number of seconds before a request to the MinIO server times out.
	// This field is optional.
	Timeout int64 `yaml:"timeout"`
}
