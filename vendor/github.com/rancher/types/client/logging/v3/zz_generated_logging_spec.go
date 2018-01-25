package client

const (
	LoggingSpecType                     = "loggingSpec"
	LoggingSpecFieldElasticsearchConfig = "elasticsearchConfig"
	LoggingSpecFieldEmbeddedConfig      = "embeddedConfig"
	LoggingSpecFieldKafkaConfig         = "kafkaConfig"
	LoggingSpecFieldOutputFlushInterval = "outputFlushInterval"
	LoggingSpecFieldOutputTags          = "outputTags"
	LoggingSpecFieldSplunkConfig        = "splunkConfig"
	LoggingSpecFieldSyslogConfig        = "syslogConfig"
)

type LoggingSpec struct {
	ElasticsearchConfig *ElasticsearchConfig `json:"elasticsearchConfig,omitempty"`
	EmbeddedConfig      *EmbeddedConfig      `json:"embeddedConfig,omitempty"`
	KafkaConfig         *KafkaConfig         `json:"kafkaConfig,omitempty"`
	OutputFlushInterval *int64               `json:"outputFlushInterval,omitempty"`
	OutputTags          map[string]string    `json:"outputTags,omitempty"`
	SplunkConfig        *SplunkConfig        `json:"splunkConfig,omitempty"`
	SyslogConfig        *SyslogConfig        `json:"syslogConfig,omitempty"`
}
