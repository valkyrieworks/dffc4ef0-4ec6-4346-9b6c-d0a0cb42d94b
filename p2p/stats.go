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
	StatsComponent = "REDACTED"
)

//
//
//
var itemToTagPattern = regexp.MustCompile("REDACTED")

//

//
type Stats struct {
	//
	Nodes metrics.Gauge
	//
	NodeAcceptOctetsSum metrics.Counter `metrics_labels:"node_uid,chID"`
	//
	NodeTransmitOctetsSum metrics.Counter `metrics_labels:"node_uid,chID"`
	//
	NodeAwaitingTransmitOctets metrics.Gauge `metrics_labels:"node_uid"`
	//
	CountTrans metrics.Gauge `metrics_labels:"node_uid"`
	//
	SignalAcceptOctetsSum metrics.Counter `metrics_labels:"signal_kind"`
	//
	SignalTransmitOctetsSum metrics.Counter `metrics_labels:"signal_kind"`
	//
	SignalsAccepted metrics.Counter `metrics_labels:"signal_kind,reactor"`
	//
	SignalsHandlerInJourney metrics.Gauge `metrics_labels:"signal_kind,reactor"`
	//
	SignalsHandlerAwaitingPeriod metrics.Histogram `metrics_labels:"signal_kind,reactor"`
	//
	SignalHandlerAcceptPeriod metrics.Histogram `metrics_labels:"signal_kind,reactor"`
	//
	SignalHandlerBufferParallelism metrics.Gauge `metrics_labels:"handler"`
}

type statsTagRepository struct {
	mtx               *sync.RWMutex
	signalTagLabels map[reflect.Type]string
}

//
//
//
//
func (m *statsTagRepository) ItemToIndicatorTag(i any) string {
	t := reflect.TypeOf(i)
	m.mtx.RLock()

	if s, ok := m.signalTagLabels[t]; ok {
		m.mtx.RUnlock()
		return s
	}
	m.mtx.RUnlock()

	s := t.String()
	ss := itemToTagPattern.FindStringSubmatch(s)
	l := fmt.Sprintf("REDACTED", ss[1], ss[2])
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.signalTagLabels[t] = l
	return l
}

func newStatsTagRepository() *statsTagRepository {
	return &statsTagRepository{
		mtx:               &sync.RWMutex{},
		signalTagLabels: map[reflect.Type]string{},
	}
}
