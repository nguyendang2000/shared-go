package elastic

// Config represents the configuration settings for connecting to an Elasticsearch cluster.
// The structure includes fields for connection addresses, authentication, certificates, and timeout settings.
type Config struct {
	// Addresses is a list of Elasticsearch server URLs.
	// This field is required for establishing a connection to the cluster.
	Addresses []string `yaml:"addresses"`

	// Username is used for basic authentication with the Elasticsearch cluster.
	// This field is optional and can be left empty if authentication is not required.
	Username string `yaml:"username"`

	// Password is the password for basic authentication with the Elasticsearch cluster.
	// This field is optional and should be used alongside the Username field for secure access.
	Password string `yaml:"password"`

	// CertificateFingerprint is an optional field used to specify the certificate fingerprint.
	// This helps ensure secure connection verification.
	CertificateFingerprint string `yaml:"certificate_fingerprint"`

	// CACert represents the optional file path to the Certificate Authority (CA) certificate.
	// This can be used to establish a secure connection with self-signed certificates.
	CACert string `yaml:"ca_cert"`

	// Timeout specifies the maximum time (in milliseconds) to wait for a connection.
	// This field is optional, and if not set, the default timeout is used.
	Timeout int64 `yaml:"timeout"`
}
