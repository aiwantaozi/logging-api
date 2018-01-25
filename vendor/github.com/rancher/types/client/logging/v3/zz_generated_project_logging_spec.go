package client

const (
	ProjectLoggingSpecType                     = "projectLoggingSpec"
	ProjectLoggingSpecFieldCurrentTarget       = "currentTarget"
	ProjectLoggingSpecFieldElasticsearchConfig = "elasticsearchConfig"
	ProjectLoggingSpecFieldKafkaConfig         = "kafkaConfig"
	ProjectLoggingSpecFieldOutputFlushInterval = "outputFlushInterval"
	ProjectLoggingSpecFieldOutputTags          = "outputTags"
	ProjectLoggingSpecFieldSplunkConfig        = "splunkConfig"
	ProjectLoggingSpecFieldSyslogConfig        = "syslogConfig"
)

type ProjectLoggingSpec struct {
	CurrentTarget       string               `json:"currentTarget,omitempty"`
	ElasticsearchConfig *ElasticsearchConfig `json:"elasticsearchConfig,omitempty"`
	KafkaConfig         *KafkaConfig         `json:"kafkaConfig,omitempty"`
	OutputFlushInterval *int64               `json:"outputFlushInterval,omitempty"`
	OutputTags          map[string]string    `json:"outputTags,omitempty"`
	SplunkConfig        *SplunkConfig        `json:"splunkConfig,omitempty"`
	SyslogConfig        *SyslogConfig        `json:"syslogConfig,omitempty"`
}
