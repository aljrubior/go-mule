package application

type CronScheduler struct {
	Kind,
	Name,
	FlowName,
	Expression,
	TimeZone string
	Enabled bool
}
