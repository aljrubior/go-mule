package application

func NewFlow(name string) Flow {
	return Flow{
		Name: name,
	}
}

type Flow struct {
	Name string
}
