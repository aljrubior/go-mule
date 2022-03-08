package requests

func NewServerMetricRequestBuilder(time string) ServerMetricRequestBuilder {
	return ServerMetricRequestBuilder{
		time: time,
	}
}

type ServerMetricRequestBuilder struct {
	time string
}

func (t ServerMetricRequestBuilder) Build() ServerMetricRequest {
	return ServerMetricRequest{
		ClassLoadingUnloaded:          []Metric{NewRandomMetric(t.time, 5, 10)},
		MemoryUsage:                   []Metric{NewRandomMetric(t.time, 20, 80)},
		TenuredGenCommitted:           []Metric{NewRandomMetric(t.time, 256, 768)},
		ThreadCount:                   []Metric{NewRandomMetric(t.time, 20, 60)},
		CompressedClassSpaceUsage:     []Metric{NewRandomMetric(t.time, 10, 20)},
		ClassLoadingTotal:             []Metric{NewRandomMetric(t.time, 15000, 30000)},
		CodeCacheCommitted:            []Metric{NewRandomMetric(t.time, 15, 30)},
		SurvivorUsage:                 []Metric{NewRandomMetric(t.time, 0, 10)},
		CompressedClassSpaceCommitted: []Metric{NewRandomMetric(t.time, 15, 30)},
		GcMarkSweepCount:              []Metric{NewRandomMetric(t.time, 1, 5)},
		GcMarkSweepTime:               []Metric{NewRandomMetric(t.time, 1, 5)},
		SurvivorTotal:                 []Metric{NewRandomMetric(t.time, 25, 75)},
		TernuredGenTotal:              []Metric{NewRandomMetric(t.time, 256, 768)},
		MemoryCommitted:               []Metric{NewRandomMetric(t.time, 512, 1024)},
		CPUUsage:                      []Metric{NewRandomMetric(t.time, 10, 50)},
		GCParNewCount:                 []Metric{NewRandomMetric(t.time, 1, 10)},
		MetaspaceCommitted:            []Metric{NewRandomMetric(t.time, 128, 512)},
		CodeCacheTotal:                []Metric{NewRandomMetric(t.time, 128, 512)},
		GCParNewTime:                  []Metric{NewRandomMetric(t.time, 0, 1)},
		SurvivorCommitted:             []Metric{NewRandomMetric(t.time, 15, 75)},
		CodeCacheUsage:                []Metric{NewRandomMetric(t.time, 15, 75)},
		EdenTotal:                     []Metric{NewRandomMetric(t.time, 128, 512)},
		TernuredGenUsage:              []Metric{NewRandomMetric(t.time, 10, 50)},
		EdenCommitted:                 []Metric{NewRandomMetric(t.time, 128, 512)},
		MetaspaceTotal:                []Metric{NewRandomMetric(t.time, 128, 256)},
		LoadAverage:                   []Metric{NewRandomMetric(t.time, 0, 1)},
		MemoryTotal:                   []Metric{NewRandomMetric(t.time, 512, 1024)},
		ClassLoadingLoaded:            []Metric{NewRandomMetric(t.time, 15000, 30000)},
		MetaspaceUsage:                []Metric{NewRandomMetric(t.time, 128, 512)},
		EdenUsage:                     []Metric{NewRandomMetric(t.time, 15, 50)},
		CompressedClassSpaceTotal:     []Metric{NewRandomMetric(t.time, 128, 512)},
		AvailableProcessors:           []Metric{NewRandomMetric(t.time, 1, 10)},
	}

}
