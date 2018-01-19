package client

const (
	KafkaConfigType                = "kafkaConfig"
	KafkaConfigFieldBrokerType     = "brokerType"
	KafkaConfigFieldBrokers        = "brokers"
	KafkaConfigFieldMaxSendRetries = "maxSendRetries"
	KafkaConfigFieldOutputDataType = "outputDataType"
	KafkaConfigFieldTopic          = "topic"
	KafkaConfigFieldZookeeperHost  = "zookeeperHost"
	KafkaConfigFieldZookeeperPort  = "zookeeperPort"
)

type KafkaConfig struct {
	BrokerType     string   `json:"brokerType,omitempty"`
	Brokers        []string `json:"brokers,omitempty"`
	MaxSendRetries *int64   `json:"maxSendRetries,omitempty"`
	OutputDataType string   `json:"outputDataType,omitempty"`
	Topic          string   `json:"topic,omitempty"`
	ZookeeperHost  string   `json:"zookeeperHost,omitempty"`
	ZookeeperPort  *int64   `json:"zookeeperPort,omitempty"`
}
