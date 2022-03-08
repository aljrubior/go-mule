package application

type TimeUnit int

const (
	MILLISECONDS TimeUnit = iota
	SECONDS
	MINUTES
	HOURS
	DAYS
)

func (timeUnit TimeUnit) String() string {
	return [...]string{
		"MILLISECONDS",
		"SECONDS",
		"MINUTES",
		"HOURS",
		"DAYS"}[timeUnit]
}
