package application

import (
	"fmt"
	"math/rand"
	"time"
)

func NewApplicationBuilder(
	applicationName string,
	totalFixedSchedulers,
	totalCronSchedulers int) *ApplicationBuilder {

	return &ApplicationBuilder{
		applicationName:      applicationName,
		totalFixedSchedulers: totalFixedSchedulers,
		totalCronSchedulers:  totalCronSchedulers,
	}
}

type ApplicationBuilder struct {
	applicationName      string
	totalFixedSchedulers int
	totalCronSchedulers  int
}

func (builder ApplicationBuilder) Build() *Application {

	fixedSchedulers := builder.buildFixedSchedulers()

	cronScheduler := builder.buildCronSchedulers()

	flows := builder.buildFlows(fixedSchedulers, cronScheduler)

	return NewApplication(builder.applicationName, flows, fixedSchedulers, cronScheduler)

}

func (builder ApplicationBuilder) buildFlows(fixedSchedulers []FixedScheduler, cronSchedulers []CronScheduler) []Flow {

	var flows []Flow

	for _, v := range fixedSchedulers {
		flows = append(flows, NewFlow(v.FlowName))
	}

	for _, v := range cronSchedulers {
		flows = append(flows, NewFlow(v.FlowName))
	}

	return flows

}

func (builder ApplicationBuilder) buildFixedSchedulers() []FixedScheduler {

	var schedulers []FixedScheduler

	for i := 0; i < builder.totalFixedSchedulers; i++ {
		flowName := fmt.Sprintf("fixedScheduler%v", i)
		schedulers = append(schedulers, FixedScheduler{
			Name:       fmt.Sprintf("polling://%s/", flowName),
			FlowName:   flowName,
			TimeUnit:   builder.randomTimeUnit(),
			Frequency:  builder.random(1, 60),
			StartDelay: builder.random(1, 30),
			Enabled:    true,
		})
	}

	return schedulers
}

func (builder ApplicationBuilder) buildCronSchedulers() []CronScheduler {

	var schedulers []CronScheduler

	for i := 0; i < builder.totalCronSchedulers; i++ {
		flowName := fmt.Sprintf("cronScheduler%v", i)
		schedulers = append(schedulers, CronScheduler{
			Name:       fmt.Sprintf("polling://%s/", flowName),
			FlowName:   flowName,
			Expression: "0 0/1 * 1/1 * ? *",
			TimeZone:   builder.randomTimeZone(),
			Enabled:    true,
		})
	}

	return schedulers
}

func (builder ApplicationBuilder) random(min, max int) int {

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func (builder ApplicationBuilder) randomTimeUnit() string {

	timeUnit := TimeUnit(builder.random(0, 4))

	return timeUnit.String()
}

func (builder ApplicationBuilder) randomTimeZone() string {

	timeZone := TimeZone(builder.random(0, 12))

	return timeZone.String()
}
