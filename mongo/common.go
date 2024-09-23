package mongo

// DefaultTimeout is the default number of seconds before a request times out
const DefaultTimeout int64 = 30 // 30 seconds

// Config represents the configuration for MongoDB connection
type Config struct {
	Address  string `yaml:"address"`  // The address of the MongoDB server
	Username string `yaml:"username"` // The username for MongoDB authentication (optional)
	Password string `yaml:"password"` // The password for MongoDB authentication (optional)
	AuthDB   string `yaml:"auth_db"`  // The name of the authentication database
	Timeout  int64  `yaml:"timeout"`  // The number of seconds before a request times out (optional)
}

// Query struct is a wrapper around map[string]interface{} for MongoDB query filters
type Query struct {
	Filter map[string]interface{} // The filter used to query documents
}

// NewQuery initializes a new Query
func NewQuery() *Query {
	return &Query{
		Filter: make(map[string]interface{}),
	}
}

// Add adds a key-value pair to the Query filter
func (q *Query) Add(key string, value interface{}) *Query {
	q.Filter[key] = value
	return q
}
