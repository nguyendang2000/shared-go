package elastic

type Config struct {
	Addresses              []string `yaml:"addresses"`               // Required
	Username               string   `yaml:"username"`                // Optional
	Password               string   `yaml:"password"`                // Optional
	CertificateFingerprint string   `yaml:"certificate_fingerprint"` // Optional
	CACert                 string   `yaml:"ca_cert"`                 // Optional file path to CA cert
	Timeout                int64    `yaml:"timeout"`                 // Optional timeout in milliseconds
}
