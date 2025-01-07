package kafka

// Error messages for Kafka Service operations.
// These constants define error messages for general Kafka operations,
// formatted with placeholders to allow dynamic values.
const (
	// ErrKafkaClientSetup is returned when setting up the Kafka client fails.
	ErrKafkaClientSetup = "failed to set up Kafka client: %w"

	// ErrKafkaPing is returned when a ping to Kafka brokers fails.
	ErrKafkaPing = "failed to ping Kafka brokers: %w"

	// ErrMarshalData is returned when marshaling message data into JSON fails.
	ErrMarshalData = "failed to marshal data for message %+v: %w"

	// ErrMarshalKey is returned when marshaling the message key into JSON fails.
	ErrMarshalKey = "failed to marshal key %+v: %w"

	// ErrAsyncProduce is returned when an asynchronous message production fails.
	ErrAsyncProduce = "failed to async produce message %+v: %+v"

	// ErrSyncProduce is returned when a synchronous message production fails.
	ErrSyncProduce = "failed to produce message(s): %w"

	// ErrFetchMessages is returned when fetching messages from Kafka topics fails.
	ErrFetchMessages = "failed to fetch messages from topic %s: %+v"
)
