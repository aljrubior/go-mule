package application

func NewApplication(
	name string,
	flows []Flow,
	fixedSchedulers []FixedScheduler,
	cronSchedulers []CronScheduler) *Application {

	return &Application{
		Name:            name,
		Flows:           flows,
		FixedSchedulers: fixedSchedulers,
		CronSchedulers:  cronSchedulers,
	}
}

type Application struct {
	Name            string
	Flows           []Flow
	FixedSchedulers []FixedScheduler
	CronSchedulers  []CronScheduler
}
