package kafka

// Config defines the configuration for the Kafka Service.
// It includes broker addresses, consumer group, consumer topics, and additional options.
type Config struct {
	// Brokers is a list of Kafka broker addresses used for connecting to the cluster.
	Brokers []string `yaml:"brokers"`

	// ConsumerGroup specifies the Kafka consumer group ID for consuming messages.
	// This allows multiple consumers within the same group to share message processing.
	ConsumerGroup string `yaml:"consumer_group"`

	// ConsumerTopics defines the list of Kafka topics from which messages will be consumed.
	// Messages from these topics will be fetched and processed by the consumer.
	ConsumerTopics []string `yaml:"consumer_topics"`

	// AutoTopicCreation determines whether Kafka should automatically create topics
	// that do not already exist when a producer or consumer interacts with them.
	AutoTopicCreation bool `yaml:"auto_topic_creation"`
}
