package kafka

// Config defines the configuration for the Kafka Service.
// It includes broker addresses, producer topics, consumer group, and consumer topics.
type Config struct {
	// Brokers is a list of Kafka broker addresses used for connecting to the cluster.
	Brokers []string `yaml:"brokers"`

	// ConsumerGroup specifies the Kafka consumer group ID for consuming messages.
	ConsumerGroup string `yaml:"consumer_group"`

	// ConsumerTopics defines the list of Kafka topics from which messages will be consumed.
	ConsumerTopics []string `yaml:"consumer_topics"`
}
