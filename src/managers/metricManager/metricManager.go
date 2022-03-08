package metricManager

import "github.com/aljrubior/standalone-runtime/managers/metricManager/requests"

type MetricManager interface {
	PostApplicationMetrics(applicationName string, metrics requests.ApplicationMetricRequest) error
	PostServerMetrics(metrics requests.ServerMetricRequest) error
}
