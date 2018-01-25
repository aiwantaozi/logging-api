package client

import (
	"github.com/rancher/norman/types"
)

const (
	LoggingType                      = "logging"
	LoggingFieldAnnotations          = "annotations"
	LoggingFieldClusterId            = "clusterId"
	LoggingFieldCreated              = "created"
	LoggingFieldCreatorID            = "creatorId"
	LoggingFieldElasticsearchConfig  = "elasticsearchConfig"
	LoggingFieldEmbeddedConfig       = "embeddedConfig"
	LoggingFieldKafkaConfig          = "kafkaConfig"
	LoggingFieldLabels               = "labels"
	LoggingFieldName                 = "name"
	LoggingFieldOutputFlushInterval  = "outputFlushInterval"
	LoggingFieldOutputTags           = "outputTags"
	LoggingFieldOwnerReferences      = "ownerReferences"
	LoggingFieldRemoved              = "removed"
	LoggingFieldSplunkConfig         = "splunkConfig"
	LoggingFieldState                = "state"
	LoggingFieldStatus               = "status"
	LoggingFieldSyslogConfig         = "syslogConfig"
	LoggingFieldTransitioning        = "transitioning"
	LoggingFieldTransitioningMessage = "transitioningMessage"
	LoggingFieldUuid                 = "uuid"
)

type Logging struct {
	types.Resource
	Annotations          map[string]string    `json:"annotations,omitempty"`
	ClusterId            string               `json:"clusterId,omitempty"`
	Created              string               `json:"created,omitempty"`
	CreatorID            string               `json:"creatorId,omitempty"`
	ElasticsearchConfig  *ElasticsearchConfig `json:"elasticsearchConfig,omitempty"`
	EmbeddedConfig       *EmbeddedConfig      `json:"embeddedConfig,omitempty"`
	KafkaConfig          *KafkaConfig         `json:"kafkaConfig,omitempty"`
	Labels               map[string]string    `json:"labels,omitempty"`
	Name                 string               `json:"name,omitempty"`
	OutputFlushInterval  *int64               `json:"outputFlushInterval,omitempty"`
	OutputTags           map[string]string    `json:"outputTags,omitempty"`
	OwnerReferences      []OwnerReference     `json:"ownerReferences,omitempty"`
	Removed              string               `json:"removed,omitempty"`
	SplunkConfig         *SplunkConfig        `json:"splunkConfig,omitempty"`
	State                string               `json:"state,omitempty"`
	Status               *LoggingStatus       `json:"status,omitempty"`
	SyslogConfig         *SyslogConfig        `json:"syslogConfig,omitempty"`
	Transitioning        string               `json:"transitioning,omitempty"`
	TransitioningMessage string               `json:"transitioningMessage,omitempty"`
	Uuid                 string               `json:"uuid,omitempty"`
}
type LoggingCollection struct {
	types.Collection
	Data   []Logging `json:"data,omitempty"`
	client *LoggingClient
}

type LoggingClient struct {
	apiClient *Client
}

type LoggingOperations interface {
	List(opts *types.ListOpts) (*LoggingCollection, error)
	Create(opts *Logging) (*Logging, error)
	Update(existing *Logging, updates interface{}) (*Logging, error)
	ByID(id string) (*Logging, error)
	Delete(container *Logging) error
}

func newLoggingClient(apiClient *Client) *LoggingClient {
	return &LoggingClient{
		apiClient: apiClient,
	}
}

func (c *LoggingClient) Create(container *Logging) (*Logging, error) {
	resp := &Logging{}
	err := c.apiClient.Ops.DoCreate(LoggingType, container, resp)
	return resp, err
}

func (c *LoggingClient) Update(existing *Logging, updates interface{}) (*Logging, error) {
	resp := &Logging{}
	err := c.apiClient.Ops.DoUpdate(LoggingType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LoggingClient) List(opts *types.ListOpts) (*LoggingCollection, error) {
	resp := &LoggingCollection{}
	err := c.apiClient.Ops.DoList(LoggingType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LoggingCollection) Next() (*LoggingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LoggingCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LoggingClient) ByID(id string) (*Logging, error) {
	resp := &Logging{}
	err := c.apiClient.Ops.DoByID(LoggingType, id, resp)
	return resp, err
}

func (c *LoggingClient) Delete(container *Logging) error {
	return c.apiClient.Ops.DoResourceDelete(LoggingType, &container.Resource)
}
