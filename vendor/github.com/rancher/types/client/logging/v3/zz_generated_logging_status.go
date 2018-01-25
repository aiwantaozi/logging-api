package client

const (
	LoggingStatusType               = "loggingStatus"
	LoggingStatusFieldCurrentTarget = "currentTarget"
)

type LoggingStatus struct {
	CurrentTarget string `json:"currentTarget,omitempty"`
}
