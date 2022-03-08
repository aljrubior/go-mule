package services

import "github.com/aljrubior/standalone-runtime/managers/metricManager/requests"

type MetricService interface {
	PostApplicationMetrics(applicationName string, request requests.ApplicationMetricRequest) error
	PostServerMetrics(request requests.ServerMetricRequest) error
}
