package txpool

import (
	"context"
	"sync"
	"testing"
	"time"

	abciemulate "github.com/valkyrieworks/iface/customer/simulations"
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/kinds"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func VerifyApplicationHandler(t *testing.T) {
	const (
		deadline  = 5 * time.Second
		cadence = 200 * time.Millisecond
	)

	sooner := func(fn func() bool) {
		require.Eventually(t, fn, deadline, cadence)
	}

	//
	//
	var (
		memberA = newApplicationHandlerMember(t, "REDACTED")
		memberBYTE = newApplicationHandlerMember(t, "REDACTED")
		memberC = newApplicationHandlerMember(t, "REDACTED")
		instances = []*applicationHandlerMember{memberA, memberBYTE, memberC}
	)

	//
	onBegin := func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", instances[i].handler)
		return s
	}

	routers := p2p.CreateLinkedRouters(settings.VerifySettings().P2P, len(instances), onBegin, p2p.Connect2routers)

	for i, member := range instances {
		member.sw = routers[i]
		member.handler.ActivateInOutTrans()
	}

	defer func() {
		for _, member := range instances {
			if err := member.sw.Halt(); err != nil {
				require.NoError(t, err)
			}
		}
	}()

	//
	//
	trans1 := []kinds.Tx{
		kinds.Tx("REDACTED"),
		kinds.Tx("REDACTED"),
		kinds.Tx("REDACTED"),
		kinds.Tx("REDACTED"),
		kinds.Tx("REDACTED"),
	}
	for _, tx := range trans1 {
		err := memberA.txpool.EmbedTransfer(tx)
		require.NoError(t, err, "REDACTED", tx)
	}

	//
	//
	sooner(func() bool {
		accepted := memberBYTE.fetchAcceptedTrans()
		return transInclude(accepted, trans1)
	})

	//
	//
	trans2 := []kinds.Tx{
		kinds.Tx("REDACTED"),
		kinds.Tx("REDACTED"),
		kinds.Tx("REDACTED"),
		kinds.Tx("REDACTED"),
		kinds.Tx("REDACTED"),
	}
	for _, tx := range trans2 {
		err := memberBYTE.txpool.EmbedTransfer(tx)
		require.NoError(t, err, "REDACTED", tx)
	}

	//
	//
	sooner(func() bool {
		accepted := memberA.fetchAcceptedTrans()
		return transInclude(accepted, trans2)
	})

	//
	//
	//
	allTrans := append(trans1, trans2...)
	sooner(func() bool {
		acceptedA := memberA.fetchAcceptedTrans()
		acceptedBYTE := memberBYTE.fetchAcceptedTrans()
		acceptedC := memberC.fetchAcceptedTrans()
		return transInclude(acceptedA, allTrans) &&
			transInclude(acceptedBYTE, allTrans) &&
			transInclude(acceptedC, allTrans)
	})

	//
	require.False(t, hasReplicates(memberA.fetchAcceptedTrans()))
	require.False(t, hasReplicates(memberBYTE.fetchAcceptedTrans()))
	require.False(t, hasReplicates(memberC.fetchAcceptedTrans()))
}

func VerifySegmentTrans(t *testing.T) {
	createTransfer := func(volume int) kinds.Tx {
		return kinds.Tx(random.Octets(volume))
	}

	toTrans := func(extents []int) kinds.Txs {
		txs := make([]kinds.Tx, 0, len(extents))
		for _, volume := range extents {
			txs = append(txs, createTransfer(volume))
		}
		return txs
	}

	for _, tt := range []struct {
		label   string
		influx  []int
		volume   int
		result [][]int
	}{
		{
			label:   "REDACTED",
			influx:  []int{100},
			volume:   200,
			result: [][]int{{100}},
		},
		{
			label:   "REDACTED",
			influx:  []int{100},
			volume:   50,
			result: [][]int{{100}},
		},
		{
			label:   "REDACTED",
			influx:  []int{100, 100, 100},
			volume:   200,
			result: [][]int{{100, 100}, {100}},
		},
		{
			label:   "REDACTED",
			influx:  []int{100, 100, 100},
			volume:   100,
			result: [][]int{{100}, {100}, {100}},
		},
		{
			label:   "REDACTED",
			influx:  []int{101, 20, 30, 50, 2, 102, 3},
			volume:   100,
			result: [][]int{{101}, {20, 30, 50}, {2}, {102}, {3}},
		},
	} {
		t.Run(tt.label, func(t *testing.T) {
			//
			influx := toTrans(tt.influx)

			anticipated := make([]kinds.Txs, 0, len(tt.result))
			for _, segment := range tt.result {
				anticipated = append(anticipated, toTrans(segment))
			}

			//
			factual := segmentTrans(influx, tt.volume)

			//
			require.Equal(t, len(anticipated), len(factual), "REDACTED")

			for i, segment := range factual {
				require.Equal(t, len(anticipated[i]), len(segment), "REDACTED", i)
			}
		})
	}
}

type applicationHandlerMember struct {
	t    *testing.T
	label string

	app     *abciemulate.Customer
	txpool *ApplicationTxpool
	handler *ApplicationHandler
	sw      *p2p.Router

	txpoolTrans  kinds.Txs
	acceptedTrans kinds.Txs
	mu          sync.Mutex

	tracer log.Tracer
}

func newApplicationHandlerMember(t *testing.T, label string) *applicationHandlerMember {
	settings := settings.VerifySettings()
	tracer := log.VerifyingTracer().With("REDACTED", label)
	app := abciemulate.NewCustomer(t)

	txpool := NewApplicationTxpool(
		settings.Txpool,
		app,
		WithMorningTracer(tracer.With("REDACTED", "REDACTED")),
	)

	handler := NewApplicationHandler(settings.Txpool, txpool, true)
	handler.AssignTracer(tracer.With("REDACTED", "REDACTED"))

	ts := &applicationHandlerMember{
		t:       t,
		label:    label,
		app:     app,
		txpool: txpool,
		handler: handler,
		tracer:  tracer,
	}

	ts.configureApplicationEmulate()

	return ts
}

func (ts *applicationHandlerMember) embedTransfer(tx kinds.Tx) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.tracer.Details("REDACTED", "REDACTED", string(tx))

	ts.txpoolTrans = append(ts.txpoolTrans, tx)
	ts.acceptedTrans = append(ts.acceptedTrans, tx)
}

func (ts *applicationHandlerMember) harvestTrans() kinds.Txs {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.tracer.Details("REDACTED")

	out := make(kinds.Txs, 0, len(ts.txpoolTrans))
	out = append(out, ts.txpoolTrans...)

	ts.txpoolTrans = ts.txpoolTrans[:0]

	return out
}

func (ts *applicationHandlerMember) fetchAcceptedTrans() kinds.Txs {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	out := make(kinds.Txs, 0, len(ts.acceptedTrans))
	out = append(out, ts.acceptedTrans...)

	return out
}

func (ts *applicationHandlerMember) configureApplicationEmulate() {
	emulateGrpc := func(procedure string, fn any) *mock.Call {
		return ts.app.On(procedure, mock.Anything, mock.Anything).Return(fn).Maybe()
	}

	emulateGrpc("REDACTED", func(_ context.Context, req *iface.QueryEmbedTransfer) (*iface.ReplyEmbedTransfer, error) {
		ts.embedTransfer(req.Tx)
		return &iface.ReplyEmbedTransfer{
			Code: iface.CodeKindSuccess,
		}, nil
	})

	emulateGrpc("REDACTED", func(_ context.Context, req *iface.QueryHarvestTrans) (*iface.ReplyHarvestTrans, error) {
		out := ts.harvestTrans()

		return &iface.ReplyHarvestTrans{Txs: out.ToSegmentOfOctets()}, nil
	})
}

func transInclude(set, segment kinds.Txs) bool {
	repository := make(map[string]struct{})

	for _, tx := range set {
		repository[tx.String()] = struct{}{}
	}

	for _, tx := range segment {
		if _, ok := repository[tx.String()]; !ok {
			return false
		}
	}

	return true
}

func hasReplicates(txs kinds.Txs) bool {
	repository := make(map[string]struct{})

	for _, tx := range txs {
		if _, ok := repository[tx.String()]; ok {
			return true
		}

		repository[tx.String()] = struct{}{}
	}

	return false
}
