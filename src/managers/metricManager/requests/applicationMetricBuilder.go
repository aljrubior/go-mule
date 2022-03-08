package requests

import (
	"github.com/aljrubior/go-mule/application"
	"math/rand"
	"time"
)

func NewApplicationMetricBuilder(application *application.Application) *ApplicationMetricBuilder {
	return &ApplicationMetricBuilder{
		application: application,
	}
}

type ApplicationMetricBuilder struct {
	application *application.Application
}

func (t *ApplicationMetricBuilder) Build() ApplicationMetricRequest {

	rand.Seed(time.Now().Unix())
	min := 1
	max := 30

	messageCount := 0
	responseTime := 0
	errorCount := 0

	flows := make(map[string]*FlowMetric)

	now := time.Now()
	time := now.Format(time.RFC3339)

	for _, v := range t.application.Flows {
		randomMessageCount := rand.Intn(max-min) + min
		randomResponseTime := rand.Intn(max-min) + min
		randomErrorCount := rand.Intn(max-min) + min

		flows[v.Name] = NewFlowMetricBuilder(time, randomMessageCount, randomResponseTime, randomErrorCount).Build()

		messageCount += randomMessageCount
		responseTime += randomResponseTime
		errorCount += randomErrorCount
	}

	return ApplicationMetricRequest{
		MessageCount: []Metric{NewMetric(time, messageCount)},
		ResponseTime: []Metric{NewMetric(time, responseTime)},
		ErrorCount:   []Metric{NewMetric(time, errorCount)},
		Flows:        flows,
	}
}
