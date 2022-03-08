package metricManager

import "github.com/aljrubior/go-mule/managers/metricManager/requests"

type MetricManager interface {
	PostApplicationMetrics(applicationName string, metrics requests.ApplicationMetricRequest) error
	PostServerMetrics(metrics requests.ServerMetricRequest) error
}
