package requests

type ServerMetricRequest struct {
	ClassLoadingUnloaded          []Metric `json:"class-loading-unloaded"`
	MemoryUsage                   []Metric `json:"memory-usage"`
	TenuredGenCommitted           []Metric `json:"tenured-gen-committed"`
	ThreadCount                   []Metric `json:"thread-count"`
	CompressedClassSpaceUsage     []Metric `json:"compressed-class-space-usage"`
	ClassLoadingTotal             []Metric `json:"class-loading-total"`
	CodeCacheCommitted            []Metric `json:"code-cache-committed"`
	SurvivorUsage                 []Metric `json:"survivor-usage"`
	CompressedClassSpaceCommitted []Metric `json:"compressed-class-space-committed"`
	GcMarkSweepCount              []Metric `json:"gc-mark-sweep-count"`
	GcMarkSweepTime               []Metric `json:"gc-mark-sweep-time"`
	SurvivorTotal                 []Metric `json:"survivor-total"`
	TernuredGenTotal              []Metric `json:"tenured-gen-total"`
	MemoryCommitted               []Metric `json:"memory-committed"`
	CPUUsage                      []Metric `json:"cpu-usage"`
	GCParNewCount                 []Metric `json:"gc-par-new-count"`
	MetaspaceCommitted            []Metric `json:"metaspace-committed"`
	CodeCacheTotal                []Metric `json:"code-cache-total"`
	GCParNewTime                  []Metric `json:"gc-par-new-time"`
	SurvivorCommitted             []Metric `json:"survivor-committed"`
	CodeCacheUsage                []Metric `json:"code-cache-usage"`
	EdenTotal                     []Metric `json:"eden-total"`
	TernuredGenUsage              []Metric `json:"tenured-gen-usage"`
	EdenCommitted                 []Metric `json:"eden-committed"`
	MetaspaceTotal                []Metric `json:"metaspace-total"`
	LoadAverage                   []Metric `json:"load-average"`
	MemoryTotal                   []Metric `json:"memory-total"`
	ClassLoadingLoaded            []Metric `json:"class-loading-loaded"`
	MetaspaceUsage                []Metric `json:"metaspace-usage"`
	EdenUsage                     []Metric `json:"eden-usage"`
	CompressedClassSpaceTotal     []Metric `json:"compressed-class-space-total"`
	AvailableProcessors           []Metric `json:"available-processors"`
}
