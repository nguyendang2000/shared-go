package mongo

// Config represents the configuration settings required for connecting to a MongoDB server.
type Config struct {
	// Address specifies the address of the MongoDB server.
	Address string `yaml:"address"`

	// Username is the optional username for MongoDB authentication.
	Username string `yaml:"username"`

	// Password is the optional password for MongoDB authentication.
	Password string `yaml:"password"`

	// AuthDB defines the name of the authentication database.
	AuthDB string `yaml:"auth_db"`

	// Timeout specifies the number of seconds before a request to MongoDB times out.
	// This field is optional.
	Timeout int64 `yaml:"timeout"`
}
