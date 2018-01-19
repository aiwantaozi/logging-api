package client

const (
	LoggingSpecType                     = "loggingSpec"
	LoggingSpecFieldElasticsearchConfig = "elasticsearchConfig"
	LoggingSpecFieldEmbeddedConfig      = "embeddedConfig"
	LoggingSpecFieldKafkaConfig         = "kafkaConfig"
	LoggingSpecFieldSplunkConfig        = "splunkConfig"
)

type LoggingSpec struct {
	ElasticsearchConfig *ElasticsearchConfig `json:"elasticsearchConfig,omitempty"`
	EmbeddedConfig      *EmbeddedConfig      `json:"embeddedConfig,omitempty"`
	KafkaConfig         *KafkaConfig         `json:"kafkaConfig,omitempty"`
	SplunkConfig        *SplunkConfig        `json:"splunkConfig,omitempty"`
}
