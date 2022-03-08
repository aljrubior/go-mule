package services

import (
	"encoding/json"
	"github.com/aljrubior/standalone-runtime/clients/metricClient"
	"github.com/aljrubior/standalone-runtime/managers/metricManager/requests"
)

func NewDefaultMetricService(metricClient metricClient.MetricClient) DefaultMetricService {
	return DefaultMetricService{
		metricClient: metricClient,
	}
}

type DefaultMetricService struct {
	metricClient metricClient.MetricClient
}

func (t *DefaultMetricService) PostApplicationMetrics(applicationName string, request requests.ApplicationMetricRequest) error {

	body, err := json.Marshal(request)

	if err != nil {
		return err
	}

	return t.metricClient.PostApplicationMetrics(applicationName, body)
}

func (t *DefaultMetricService) PostServerMetrics(request requests.ServerMetricRequest) error {

	body, err := json.Marshal(request)

	if err != nil {
		return err
	}

	return t.metricClient.PostServerMetrics(body)
}
