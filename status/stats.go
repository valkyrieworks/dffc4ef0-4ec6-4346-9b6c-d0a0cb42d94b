package status

import (
	"github.com/go-kit/kit/metrics"
)

const (
	//
	//
	StatsComponent = "REDACTED"
)

//

//
type Stats struct {
	//
	LedgerExecutionTime metrics.Histogram `metrics_buckettype:"lin" metrics_bucketsizes:"1, 10, 10"`

	//
	//
	//
	AgreementArgumentRefreshes metrics.Counter

	//
	//
	//
	RatifierCollectionRefreshes metrics.Counter
}
