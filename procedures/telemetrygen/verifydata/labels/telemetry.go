package labels

import "github.com/go-kit/kit/metrics"

//

type Metrics struct {
	WithLabels     metrics.Counter   `metrics_labels:"phase,time"`
	WithExpBuckets metrics.Histogram `metrics_buckettype:"exp" metrics_bucketsizes:".1,100,8"`
	WithBuckets    metrics.Histogram `metrics_bucketsizes:"1, 2, 3, 4, 5"`
	Named          metrics.Counter   `metrics_name:"measurement_using_alias"`
}
