package status

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
	LedgerHandlingMoment metrics.Histogram `metrics_buckettype:"lin" metrics_bucketsizes:"1, 10, 10"`

	//
	//
	//
	AgreementArgumentRevisions metrics.Counter

	//
	//
	//
	AssessorAssignRevisions metrics.Counter
}
