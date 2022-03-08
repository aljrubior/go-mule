package strategies

func NewNoActionStrategy() *NoActionStrategy {
	return &NoActionStrategy{}
}

type NoActionStrategy struct {
}

func (t *NoActionStrategy) Execute() {
}
