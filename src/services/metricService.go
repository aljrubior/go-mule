package services

import "github.com/aljrubior/go-mule/managers/metricManager/requests"

type MetricService interface {
	PostApplicationMetrics(applicationName string, request requests.ApplicationMetricRequest) error
	PostServerMetrics(request requests.ServerMetricRequest) error
}
