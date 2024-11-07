package mongo

// Config represents the configuration for MongoDB connection
type Config struct {
	Address  string `yaml:"address"`  // The address of the MongoDB server
	Username string `yaml:"username"` // The username for MongoDB authentication (optional)
	Password string `yaml:"password"` // The password for MongoDB authentication (optional)
	AuthDB   string `yaml:"auth_db"`  // The name of the authentication database
	Timeout  int64  `yaml:"timeout"`  // The number of seconds before a request times out (optional)
}