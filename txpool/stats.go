package txpool

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
//
type Stats struct {
	//
	Volume metrics.Gauge

	//
	VolumeOctets metrics.Gauge

	//
	TransferVolumeOctets metrics.Histogram `metrics_buckettype:"exp" metrics_bucketsizes:"1,3,7"`

	//
	//
	//
	//
	ErroredTrans metrics.Counter

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
	RevalidateInstances metrics.Counter

	//
	//
	EnabledOutgoingLinkages metrics.Gauge

	//
	//
	YetAcceptedTrans metrics.Counter

	//
	GroupVolume metrics.Histogram `metrics_labels:"dir" metrics_bucketsizes:"1,2,5,10,30,50,100,200,300"`

	//
	HarvestedTrans metrics.Counter
}
