package client

const (
	EmbeddedConfigType               = "embeddedConfig"
	EmbeddedConfigFieldRequestCPU    = "requestCPU"
	EmbeddedConfigFieldRequestMemory = "requestMemory"
)

type EmbeddedConfig struct {
	RequestCPU    string `json:"requestCPU,omitempty"`
	RequestMemory string `json:"requestMemory,omitempty"`
}
