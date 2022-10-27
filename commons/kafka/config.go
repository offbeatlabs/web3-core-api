package kafka

// Config kafka config
type Config struct {
	Brokers    []string `mapstructure:"brokers"`
	GroupID    string   `mapstructure:"group_id"`
	InitTopics bool     `mapstructure:"init_topics"`
}

// TopicConfig kafka topic config
type TopicConfig struct {
	TopicName         string `mapstructure:"topic_name"`
	Partitions        int    `mapstructure:"partitions"`
	ReplicationFactor int    `mapstructure:"replication_factor"`
}
