package requests

type ApplicationMetricRequest struct {
	MessageCount []Metric               `json:"message-count"`
	ResponseTime []Metric               `json:"response-time"`
	ErrorCount   []Metric               `json:"error-count"`
	Flows        map[string]*FlowMetric `json:"flows"`
}
