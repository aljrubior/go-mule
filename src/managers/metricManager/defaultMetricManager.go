package metricManager

import (
	"github.com/aljrubior/standalone-runtime/managers/metricManager/requests"
	"github.com/aljrubior/standalone-runtime/services"
)

func NewDefaultMetricManager(metricService services.MetricService) DefaultMetricManager {
	return DefaultMetricManager{
		metricService: metricService,
	}
}

type DefaultMetricManager struct {
	metricService services.MetricService
}

func (t DefaultMetricManager) PostApplicationMetrics(applicationName string, metrics requests.ApplicationMetricRequest) error {
	return t.metricService.PostApplicationMetrics(applicationName, metrics)
}

func (t DefaultMetricManager) PostServerMetrics(metrics requests.ServerMetricRequest) error {
	return t.metricService.PostServerMetrics(metrics)
}
