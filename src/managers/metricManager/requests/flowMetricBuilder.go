package requests

func NewFlowMetricBuilder(
	time string,
	messageCount,
	responseTime,
	errorCount int) *FlowMetricBuilder {

	return &FlowMetricBuilder{
		time:         time,
		messageCount: messageCount,
		responseTime: responseTime,
		errorCount:   errorCount,
	}
}

type FlowMetricBuilder struct {
	time string
	messageCount,
	responseTime,
	errorCount int
}

func (t *FlowMetricBuilder) Build() *FlowMetric {
	return &FlowMetric{
		MessageCount: []Metric{NewMetric(t.time, t.messageCount)},
		ResponseTime: []Metric{NewMetric(t.time, t.responseTime)},
		ErrorCount:   []Metric{NewMetric(t.time, t.errorCount)},
	}
}
