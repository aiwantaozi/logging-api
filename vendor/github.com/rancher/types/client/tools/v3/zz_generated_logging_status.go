package client

const (
	LoggingStatusType                 = "loggingStatus"
	LoggingStatusFieldComponentStatus = "componentStatus"
	LoggingStatusFieldCurrentTarget   = "currentTarget"
)

type LoggingStatus struct {
	ComponentStatus string `json:"componentStatus,omitempty"`
	CurrentTarget   string `json:"currentTarget,omitempty"`
}
