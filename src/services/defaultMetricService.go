package services

import (
	"encoding/json"
	"github.com/aljrubior/go-mule/clients/metricClient"
	"github.com/aljrubior/go-mule/managers/metricManager/requests"
)

func NewDefaultMetricService(metricClient metricClient.MetricClient) DefaultMetricService {
	return DefaultMetricService{
		metricClient: metricClient,
	}
}

type DefaultMetricService struct {
	metricClient metricClient.MetricClient
}

func (t DefaultMetricService) PostApplicationMetrics(applicationName string, request requests.ApplicationMetricRequest) error {

	data, err := json.Marshal(request)

	if err != nil {
		return err
	}

	return t.metricClient.PostApplicationMetrics(applicationName, data)
}

func (t DefaultMetricService) PostServerMetrics(request requests.ServerMetricRequest) error {

	data, err := json.Marshal(request)

	if err != nil {
		return err
	}

	return t.metricClient.PostServerMetrics(data)
}
