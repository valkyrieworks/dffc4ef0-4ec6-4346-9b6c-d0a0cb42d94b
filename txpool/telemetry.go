package txpool

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
//
type Telemetry struct {
	//
	Extent metrics.Gauge

	//
	ExtentOctets metrics.Gauge

	//
	TransferExtentOctets metrics.Histogram `metrics_buckettype:"exp" metrics_bucketsizes:"1,3,7"`

	//
	//
	//
	//
	UnsuccessfulTrans metrics.Counter

	//
	//
	//
	//
	DeclinedTrans metrics.Counter

	//
	//
	//
	//
	ExpelledTrans metrics.Counter

	//
	ReinspectMultiples metrics.Counter

	//
	//
	DynamicOutgoingLinkages metrics.Gauge

	//
	//
	EarlierAcceptedTrans metrics.Counter

	//
	ClusterExtent metrics.Histogram `metrics_labels:"dir" metrics_bucketsizes:"1,2,5,10,30,50,100,200,300"`

	//
	HarvestedTrans metrics.Counter
}
