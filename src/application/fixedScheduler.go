package application

type FixedScheduler struct {
	Kind,
	Name,
	FlowName,
	TimeUnit string
	Frequency,
	StartDelay int
	Enabled bool
}
