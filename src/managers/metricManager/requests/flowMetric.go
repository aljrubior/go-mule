package requests

type FlowMetric struct {
	MessageCount []Metric `json:"message-count"`
	ResponseTime []Metric `json:"response-time"`
	ErrorCount   []Metric `json:"error-count"`
}
