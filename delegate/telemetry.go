package delegate

import (
	"github.com/go-kit/kit/metrics"
)

const (
	//
	//
	TelemetryComponent = "REDACTED"
)

//

//
type Telemetry struct {
	//
	ProcedureScheduleMoments metrics.Histogram `metrics_bucketsizes:"point0001,.0004,.002,.009,.02,.1,.65,2,6,25" metrics_labels:"procedure, type"`
}
