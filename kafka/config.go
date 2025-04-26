package kafka

import (
	"fmt"
	"strconv"
)

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

type (
	// TopicConfig represents the configuration options for a Kafka topic.
	// It contains various settings that control topic behavior, retention,
	// compression, and performance characteristics.
	// All fields are pointers to allow distinguishing between zero values and unset values.
	TopicConfig struct {
		// CleanupPolicy determines how Kafka handles old log segments.
		// Options are "delete" (remove old segments) or "compact" (retain latest key values).
		CleanupPolicy *CleanupPolicy

		// Compression contains settings for topic-level message compression.
		Compression *TopicCompressionConfig

		// DeleteRetentionMs specifies how long deleted records are retained (in milliseconds).
		// Only applicable when cleanup.policy=compact.
		DeleteRetentionMs *int

		// FileDeleteDelayMs specifies the time to wait before deleting a file from the filesystem (in milliseconds).
		FileDeleteDelayMs *int

		// FlushMessages specifies the number of messages accumulated before messages are flushed to disk.
		FlushMessages *int64

		// FlushMs specifies the maximum time in milliseconds before messages are flushed to disk.
		FlushMs *int64

		// FollowerReplicationThrottledReplicas lists the replicas to throttle during follower replication.
		FollowerReplicationThrottledReplicas *string

		// IndexIntervalBytes specifies the interval at which Kafka adds an entry to the offset index.
		IndexIntervalBytes *int

		// LeaderReplicationThrottledReplicas lists the replicas to throttle during leader replication.
		LeaderReplicationThrottledReplicas *string

		// LocalRetentionBytes specifies the maximum size of the local log before deleting messages.
		LocalRetentionBytes *int64

		// LocalRetentionMs specifies the time to retain local logs before deleting them (in milliseconds).
		LocalRetentionMs *int64

		// MaxCompactionLagMs specifies the maximum time a message will remain uncompacted (in milliseconds).
		MaxCompactionLagMs *int64

		// MaxMessageBytes specifies the largest record batch size allowed by Kafka.
		MaxMessageBytes *int

		// MessageDownconversionEnable enables/disables automatic message format downconversion.
		MessageDownconversionEnable *bool

		// MessageTimestampAfterMaxMs specifies the maximum timestamp difference allowed for messages (in milliseconds).
		MessageTimestampAfterMaxMs *int64

		// MessageTimestampBeforeMaxMs specifies how far in the past messages can be to be accepted (in milliseconds).
		MessageTimestampBeforeMaxMs *int64

		// MessageTimestampType specifies whether the timestamp in the message is the creation time or log append time.
		MessageTimestampType *MessageTimestampType

		// MinCleanableDirtyRatio specifies the minimum ratio of dirty/total log for a log to be eligible for compaction.
		MinCleanableDirtyRatio *float64

		// MinCompactionLagMs specifies the minimum time a message will remain uncompacted (in milliseconds).
		MinCompactionLagMs *int64

		// MinInsyncReplicas specifies the minimum number of replicas that must acknowledge a write.
		MinInsyncReplicas *int

		// Preallocate determines whether to preallocate the file when creating a new log segment.
		Preallocate *bool

		// RemoteLogCopyDisable enables/disables copying of log segments to remote storage.
		RemoteLogCopyDisable *bool

		// RemoteLogDeleteOnDisable determines whether to delete remote logs when remote storage is disabled.
		RemoteLogDeleteOnDisable *bool

		// RemoteStorageEnable enables/disables the use of remote storage for log segments.
		RemoteStorageEnable *bool

		// RetentionBytes specifies the maximum size of the log before deleting messages.
		RetentionBytes *int64

		// RetentionMs specifies the time to retain logs before deleting them (in milliseconds).
		RetentionMs *int64

		// SegmentBytes specifies the maximum size of a single log segment file.
		SegmentBytes *int

		// SegmentIndexBytes specifies the maximum size of the offset index.
		SegmentIndexBytes *int

		// SegmentJitterMs specifies the maximum random jitter subtracted from segmentMs (in milliseconds).
		SegmentJitterMs *int64

		// SegmentMs specifies the time after which Kafka will force the log to roll (in milliseconds).
		SegmentMs *int64

		// UncleanLeaderElectionEnable enables/disables unclean leader election (allowing out-of-sync replicas to become leaders).
		UncleanLeaderElectionEnable *bool
	}

	// TopicCompressionConfig represents compression-related settings for a Kafka topic.
	// It contains configuration for different compression algorithms and their levels.
	TopicCompressionConfig struct {
		// GzipLevel specifies the compression level for gzip compression (higher = more compression but more CPU).
		GzipLevel *int

		// Lz4Level specifies the compression level for LZ4 compression.
		Lz4Level *int

		// Type specifies the compression algorithm to use for the topic.
		Type *CompressionType

		// ZstdLevel specifies the compression level for Zstandard compression.
		ZstdLevel *int
	}
)

// CleanupPolicy defines how Kafka handles old log segments for a topic.
type CleanupPolicy string

var (
	// CleanupPolicyCompact indicates that Kafka should compact the topic's logs,
	// retaining at least the last known value for each message key.
	CleanupPolicyCompact CleanupPolicy = "compact"

	// CleanupPolicyDelete indicates that Kafka should delete old log segments
	// when their retention time or size limit is reached.
	CleanupPolicyDelete CleanupPolicy = "delete"
)

// CompressionType defines the algorithm used to compress messages in a topic.
type CompressionType string

var (
	// CompressionTypeUncompressed indicates that messages should not be compressed.
	CompressionTypeUncompressed CompressionType = "uncompressed"

	// CompressionTypeZstd indicates that messages should be compressed using the Zstandard algorithm.
	CompressionTypeZstd CompressionType = "zstd"

	// CompressionTypeLz4 indicates that messages should be compressed using the LZ4 algorithm.
	CompressionTypeLz4 CompressionType = "lz4"

	// CompressionTypeSnappy indicates that messages should be compressed using the Snappy algorithm.
	CompressionTypeSnappy CompressionType = "snappy"

	// CompressionTypeGzip indicates that messages should be compressed using the Gzip algorithm.
	CompressionTypeGzip CompressionType = "gzip"

	// CompressionTypeProducer indicates that the producer's compression settings should be used.
	CompressionTypeProducer CompressionType = "producer"
)

// MessageTimestampType defines the type of timestamp used for messages in a topic.
type MessageTimestampType string

var (
	// MessageTimestampTypeCreateTime indicates that the message timestamp represents
	// the time the message was created by the producer.
	MessageTimestampTypeCreateTime MessageTimestampType = "CreateTime"

	// MessageTimestampTypeLogAppendTime indicates that the message timestamp represents
	// the time the message was appended to the broker log.
	MessageTimestampTypeLogAppendTime MessageTimestampType = "LogAppendTime"
)

// Parse converts the TopicConfig struct fields to a map of string key-value pairs.
// It only includes fields that have non-nil values, and converts the field names
// from CamelCase to lowercase dot-separated format (e.g., CleanupPolicy -> cleanup.policy).
// The returned map contains the configuration keys and their string values.
func (inst *TopicConfig) Parse() map[string]*string {
	result := make(map[string]*string)

	// Helper function to convert a value to string pointer
	toString := func(value any) *string {
		var str string
		switch v := value.(type) {
		case string:
			str = v
		case bool:
			str = strconv.FormatBool(v)
		case int:
			str = strconv.Itoa(v)
		case int64:
			str = strconv.FormatInt(v, 10)
		case float64:
			str = strconv.FormatFloat(v, 'f', -1, 64)
		default:
			str = fmt.Sprintf("%v", v)
		}
		return &str
	}

	// CleanupPolicy -> cleanup.policy
	if inst.CleanupPolicy != nil {
		result["cleanup.policy"] = toString(string(*inst.CleanupPolicy))
	}

	// Handle nested Compression struct
	if inst.Compression != nil {
		// Compression.Type -> compression.type
		if inst.Compression.Type != nil {
			result["compression.type"] = toString(string(*inst.Compression.Type))
		}
		// Compression.GzipLevel -> compression.gzip.level
		if inst.Compression.GzipLevel != nil {
			result["compression.gzip.level"] = toString(*inst.Compression.GzipLevel)
		}
		// Compression.Lz4Level -> compression.lz4.level
		if inst.Compression.Lz4Level != nil {
			result["compression.lz4.level"] = toString(*inst.Compression.Lz4Level)
		}
		// Compression.ZstdLevel -> compression.zstd.level
		if inst.Compression.ZstdLevel != nil {
			result["compression.zstd.level"] = toString(*inst.Compression.ZstdLevel)
		}
	}

	// DeleteRetentionMs -> delete.retention.ms
	if inst.DeleteRetentionMs != nil {
		result["delete.retention.ms"] = toString(*inst.DeleteRetentionMs)
	}

	// FileDeleteDelayMs -> file.delete.delay.ms
	if inst.FileDeleteDelayMs != nil {
		result["file.delete.delay.ms"] = toString(*inst.FileDeleteDelayMs)
	}

	// FlushMessages -> flush.messages
	if inst.FlushMessages != nil {
		result["flush.messages"] = toString(*inst.FlushMessages)
	}

	// FlushMs -> flush.ms
	if inst.FlushMs != nil {
		result["flush.ms"] = toString(*inst.FlushMs)
	}

	// FollowerReplicationThrottledReplicas -> follower.replication.throttled.replicas
	if inst.FollowerReplicationThrottledReplicas != nil {
		result["follower.replication.throttled.replicas"] = inst.FollowerReplicationThrottledReplicas
	}

	// IndexIntervalBytes -> index.interval.bytes
	if inst.IndexIntervalBytes != nil {
		result["index.interval.bytes"] = toString(*inst.IndexIntervalBytes)
	}

	// LeaderReplicationThrottledReplicas -> leader.replication.throttled.replicas
	if inst.LeaderReplicationThrottledReplicas != nil {
		result["leader.replication.throttled.replicas"] = inst.LeaderReplicationThrottledReplicas
	}

	// LocalRetentionBytes -> local.retention.bytes
	if inst.LocalRetentionBytes != nil {
		result["local.retention.bytes"] = toString(*inst.LocalRetentionBytes)
	}

	// LocalRetentionMs -> local.retention.ms
	if inst.LocalRetentionMs != nil {
		result["local.retention.ms"] = toString(*inst.LocalRetentionMs)
	}

	// MaxCompactionLagMs -> max.compaction.lag.ms
	if inst.MaxCompactionLagMs != nil {
		result["max.compaction.lag.ms"] = toString(*inst.MaxCompactionLagMs)
	}

	// MaxMessageBytes -> max.message.bytes
	if inst.MaxMessageBytes != nil {
		result["max.message.bytes"] = toString(*inst.MaxMessageBytes)
	}

	// MessageDownconversionEnable -> message.downconversion.enable
	if inst.MessageDownconversionEnable != nil {
		result["message.downconversion.enable"] = toString(*inst.MessageDownconversionEnable)
	}

	// MessageTimestampAfterMaxMs -> message.timestamp.after.max.ms
	if inst.MessageTimestampAfterMaxMs != nil {
		result["message.timestamp.after.max.ms"] = toString(*inst.MessageTimestampAfterMaxMs)
	}

	// MessageTimestampBeforeMaxMs -> message.timestamp.before.max.ms
	if inst.MessageTimestampBeforeMaxMs != nil {
		result["message.timestamp.before.max.ms"] = toString(*inst.MessageTimestampBeforeMaxMs)
	}

	// MessageTimestampType -> message.timestamp.type
	if inst.MessageTimestampType != nil {
		result["message.timestamp.type"] = toString(string(*inst.MessageTimestampType))
	}

	// MinCleanableDirtyRatio -> min.cleanable.dirty.ratio
	if inst.MinCleanableDirtyRatio != nil {
		result["min.cleanable.dirty.ratio"] = toString(*inst.MinCleanableDirtyRatio)
	}

	// MinCompactionLagMs -> min.compaction.lag.ms
	if inst.MinCompactionLagMs != nil {
		result["min.compaction.lag.ms"] = toString(*inst.MinCompactionLagMs)
	}

	// MinInsyncReplicas -> min.insync.replicas
	if inst.MinInsyncReplicas != nil {
		result["min.insync.replicas"] = toString(*inst.MinInsyncReplicas)
	}

	// Preallocate -> preallocate
	if inst.Preallocate != nil {
		result["preallocate"] = toString(*inst.Preallocate)
	}

	// RemoteLogCopyDisable -> remote.log.copy.disable
	if inst.RemoteLogCopyDisable != nil {
		result["remote.log.copy.disable"] = toString(*inst.RemoteLogCopyDisable)
	}

	// RemoteLogDeleteOnDisable -> remote.log.delete.on.disable
	if inst.RemoteLogDeleteOnDisable != nil {
		result["remote.log.delete.on.disable"] = toString(*inst.RemoteLogDeleteOnDisable)
	}

	// RemoteStorageEnable -> remote.storage.enable
	if inst.RemoteStorageEnable != nil {
		result["remote.storage.enable"] = toString(*inst.RemoteStorageEnable)
	}

	// RetentionBytes -> retention.bytes
	if inst.RetentionBytes != nil {
		result["retention.bytes"] = toString(*inst.RetentionBytes)
	}

	// RetentionMs -> retention.ms
	if inst.RetentionMs != nil {
		result["retention.ms"] = toString(*inst.RetentionMs)
	}

	// SegmentBytes -> segment.bytes
	if inst.SegmentBytes != nil {
		result["segment.bytes"] = toString(*inst.SegmentBytes)
	}

	// SegmentIndexBytes -> segment.index.bytes
	if inst.SegmentIndexBytes != nil {
		result["segment.index.bytes"] = toString(*inst.SegmentIndexBytes)
	}

	// SegmentJitterMs -> segment.jitter.ms
	if inst.SegmentJitterMs != nil {
		result["segment.jitter.ms"] = toString(*inst.SegmentJitterMs)
	}

	// SegmentMs -> segment.ms
	if inst.SegmentMs != nil {
		result["segment.ms"] = toString(*inst.SegmentMs)
	}

	// UncleanLeaderElectionEnable -> unclean.leader.election.enable
	if inst.UncleanLeaderElectionEnable != nil {
		result["unclean.leader.election.enable"] = toString(*inst.UncleanLeaderElectionEnable)
	}

	return result
}
