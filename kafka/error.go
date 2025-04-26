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

const (
	// ErrCreatePartitions is returned when adding partitions to topics fails.
	ErrCreatePartitions = "failed to add %d partitions to topics %v: %w"

	// ErrCreateTopic is returned when creating a Kafka topic fails.
	ErrCreateTopic = "failed to create topic %s: %w"

	// ErrCreateTopics is returned when creating multiple Kafka topics fails.
	ErrCreateTopics = "failed to create topics %v: %w"

	// ErrDeleteGroup is returned when deleting a consumer group fails.
	ErrDeleteGroup = "failed to delete consumer group %s: %w"

	// ErrDeleteGroups is returned when deleting multiple consumer groups fails.
	ErrDeleteGroups = "failed to delete consumer groups %v: %w"

	// ErrDeleteTopic is returned when deleting a Kafka topic fails.
	ErrDeleteTopic = "failed to delete topic %s: %w"

	// ErrDeleteTopics is returned when deleting multiple Kafka topics fails.
	ErrDeleteTopics = "failed to delete topics %v: %w"

	// ErrUpdatePartitions is returned when setting the number of partitions for topics fails.
	ErrUpdatePartitions = "failed to set %d partitions for topics %v: %w"
)
