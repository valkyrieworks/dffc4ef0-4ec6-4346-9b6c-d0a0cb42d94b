package p2p

import (
	"fmt"
	"reflect"
	"regexp"
	"sync"

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
var datumTowardTagPattern = regexp.MustCompile("REDACTED")

//

//
type Telemetry struct {
	//
	Nodes metrics.Gauge
	//
	NodeAcceptOctetsSum metrics.Counter `metrics_labels:"node_uuid,chID"`
	//
	NodeTransmitOctetsSum metrics.Counter `metrics_labels:"node_uuid,chID"`
	//
	NodeTransmitStagingExtent metrics.Gauge `metrics_labels:"node_uuid"`
	//
	CountTrans metrics.Gauge `metrics_labels:"node_uuid"`
	//
	ArtifactAcceptOctetsSum metrics.Counter `metrics_labels:"artifact_kind"`
	//
	ArtifactTransmitOctetsSum metrics.Counter `metrics_labels:"artifact_kind"`
	//
	SignalsAccepted metrics.Counter `metrics_labels:"artifact_kind,reactor"`
	//
	SignalsHandlerInsideAirtime metrics.Gauge `metrics_labels:"artifact_kind,reactor"`
	//
	SignalsHandlerAwaitingInterval metrics.Histogram `metrics_labels:"artifact_kind,reactor"`
	//
	ArtifactHandlerAcceptInterval metrics.Histogram `metrics_labels:"artifact_kind,reactor"`
	//
	ArtifactHandlerStagingParallelism metrics.Gauge `metrics_labels:"handler"`
}

type telemetryTagStash struct {
	mtx               *sync.RWMutex
	artifactTagIdentifiers map[reflect.Type]string
}

//
//
//
//
func (m *telemetryTagStash) DatumTowardIndicatorTag(i any) string {
	t := reflect.TypeOf(i)
	m.mtx.RLock()

	if s, ok := m.artifactTagIdentifiers[t]; ok {
		m.mtx.RUnlock()
		return s
	}
	m.mtx.RUnlock()

	s := t.String()
	ss := datumTowardTagPattern.FindStringSubmatch(s)
	l := fmt.Sprintf("REDACTED", ss[1], ss[2])
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.artifactTagIdentifiers[t] = l
	return l
}

func freshTelemetryTagStash() *telemetryTagStash {
	return &telemetryTagStash{
		mtx:               &sync.RWMutex{},
		artifactTagIdentifiers: map[reflect.Type]string{},
	}
}
