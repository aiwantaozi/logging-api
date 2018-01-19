package client

const (
	ElasticsearchConfigType                    = "elasticsearchConfig"
	ElasticsearchConfigFieldAuthPassword       = "authPassword"
	ElasticsearchConfigFieldAuthUser           = "authUser"
	ElasticsearchConfigFieldHost               = "host"
	ElasticsearchConfigFieldLogstashDateformat = "logstashDateformat"
	ElasticsearchConfigFieldLogstashPrefix     = "logstashPrefix"
	ElasticsearchConfigFieldPort               = "port"
)

type ElasticsearchConfig struct {
	AuthPassword       string `json:"authPassword,omitempty"`
	AuthUser           string `json:"authUser,omitempty"`
	Host               string `json:"host,omitempty"`
	LogstashDateformat string `json:"logstashDateformat,omitempty"`
	LogstashPrefix     string `json:"logstashPrefix,omitempty"`
	Port               *int64 `json:"port,omitempty"`
}
