package requests

import "math/rand"

func NewRandomMetric(time string, min, max int) Metric {

	value := rand.Intn(max-min) + min

	return NewMetric(time, value)
}

func NewMetric(time string, value int) Metric {

	return Metric{
		Time:  time,
		Min:   value,
		Max:   value,
		Sum:   value,
		Avg:   value,
		Count: 1,
	}
}

type Metric struct {
	Time  string `json:"time"`
	Min   int    `json:"min"`
	Max   int    `json:"max"`
	Sum   int    `json:"sum"`
	Avg   int    `json:"avg"`
	Count int    `json:"count"`
}
