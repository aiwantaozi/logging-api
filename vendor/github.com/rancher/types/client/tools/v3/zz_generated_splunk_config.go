package client

const (
	SplunkConfigType            = "splunkConfig"
	SplunkConfigFieldHost       = "host"
	SplunkConfigFieldPort       = "port"
	SplunkConfigFieldProtocol   = "protocol"
	SplunkConfigFieldSource     = "source"
	SplunkConfigFieldTimeFormat = "timeFormat"
	SplunkConfigFieldToken      = "token"
)

type SplunkConfig struct {
	Host       string `json:"host,omitempty"`
	Port       *int64 `json:"port,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
	Source     string `json:"source,omitempty"`
	TimeFormat string `json:"timeFormat,omitempty"`
	Token      string `json:"token,omitempty"`
}
