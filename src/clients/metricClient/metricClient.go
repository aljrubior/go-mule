package metricClient

type MetricClient interface {
	PostApplicationMetrics(applicationName string, body []byte) error
	PostServerMetrics(body []byte) error
}
