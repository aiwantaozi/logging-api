package v4

import (
	"github.com/rancher/norman/condition"
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LoggingConditionType string

var (
	LoggingConditionUpdateting condition.Cond = "Updateing"
	LoggingConditionReloading  condition.Cond = "Reloading"
	LoggingConditionRunning    condition.Cond = "Running"
)

type Logging struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec LoggingSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status        LoggingStatus `json:"status"`
	DisplayName   string        `json:"displayName,omitempty"`
	CurrentTarget string        `json:"currentTarget"`
	ClusterName   string        `json:"clusterName" norman:"type=reference[cluster]"`
}

type ProjectLogging struct {
	types.Namespaced
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec LoggingSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status        LoggingStatus `json:"status"`
	DisplayName   string        `json:"displayName,omitempty"`
	CurrentTarget string        `json:"currentTarget"`
	ProjectName   string        `json:"projectName,omitempty" norman:"type=reference[project]"`
}

type LoggingSpec struct {
	ElasticsearchConfig *ElasticsearchConfig `json:"elasticsearchConfig,omitempty"`
	SplunkConfig        *SplunkConfig        `json:"splunkConfig,omitempty"`
	EmbeddedConfig      *EmbeddedConfig      `json:"embeddedConfig,omitempty"`
	KafkaConfig         *KafkaConfig         `json:"kafkaConfig,omitempty"`
}

type ProjectLoggingSpec struct {
	ElasticsearchConfig *ElasticsearchConfig `json:"elasticsearchConfig,omitempty"`
	SplunkConfig        *SplunkConfig        `json:"splunkConfig,omitempty"`
	KafkaConfig         *KafkaConfig         `json:"kafkaConfig,omitempty"`
}
type LoggingStatus struct {
	CurrentTarget   string `json:"currentTarget"`
	ComponentStatus string `json:"componentStatus"`
}

type ElasticsearchConfig struct {
	Host               string `json:"host"`
	Port               int    `json:"port"`
	LogstashPrefix     string `json:"logstashPrefix"`
	LogstashDateformat string `json:"logstashDateformat"`
	AuthUser           string `json:"authUser"`     //secret
	AuthPassword       string `json:"authPassword"` //secret
}

type SplunkConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Protocol   string `json:"protocol"`
	Source     string `json:"source"`
	TimeFormat string `json:"timeFormat"`
	Token      string `json:"token"` //secret
}

type EmbeddedConfig struct {
	RequestCPU    string `json:"requestCPU"`
	RequestMemory string `json:"requestMemory"`
}

type KafkaConfig struct {
	BrokerType     string   `json:"brokerType"`
	Brokers        []string `json:"brokers"`
	ZookeeperHost  string   `json:"zookeeperHost"`
	ZookeeperPort  int      `json:"zookeeperPort"`
	Topic          string   `json:"topic"`
	OutputDataType string   `json:"outputDataType"`
	MaxSendRetries int      `json:"maxSendRetries"`
}
