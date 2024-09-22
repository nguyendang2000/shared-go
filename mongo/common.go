package mongo

// Config represents the configuration for MongoDB connection
type Config struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	AuthDB   string `yaml:"auth_db"`
}

// Query struct is a wrapper around map[string]interface{} for MongoDB query filters
type Query struct {
	Filter map[string]interface{}
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
