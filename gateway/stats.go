package gateway

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
	ProcedureCadenceMoments metrics.Histogram `metrics_bucketsizes:".0001,.0004,.002,.009,.02,.1,.65,2,6,25" metrics_labels:"procedure, type"`
}
