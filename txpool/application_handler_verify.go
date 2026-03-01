package txpool

import (
	"context"
	"sync"
	"testing"
	"time"

	ifacemimic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer/simulations"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func VerifyApplicationHandler(t *testing.T) {
	const (
		deadline  = 5 * time.Second
		duration = 200 * time.Millisecond
	)

	sooner := func(fn func() bool) {
		require.Eventually(t, fn, deadline, duration)
	}

	//
	//
	var (
		peerAN = freshApplicationHandlerPeer(t, "REDACTED")
		peerBYTE = freshApplicationHandlerPeer(t, "REDACTED")
		peerCN = freshApplicationHandlerPeer(t, "REDACTED")
		peers = []*applicationHandlerPeer{peerAN, peerBYTE, peerCN}
	)

	//
	uponInitiate := func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", peers[i].handler)
		return s
	}

	routers := p2p.CreateAssociatedRouters(settings.VerifySettings().P2P, len(peers), uponInitiate, p2p.Connect2routers)

	for i, peer := range peers {
		peer.sw = routers[i]
		peer.handler.ActivateInsideOutputTrans()
	}

	defer func() {
		for _, peer := range peers {
			if err := peer.sw.Halt(); err != nil {
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
		err := peerAN.txpool.AppendTransfer(tx)
		require.NoError(t, err, "REDACTED", tx)
	}

	//
	//
	sooner(func() bool {
		accepted := peerBYTE.fetchAcceptedTrans()
		return transHold(accepted, trans1)
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
		err := peerBYTE.txpool.AppendTransfer(tx)
		require.NoError(t, err, "REDACTED", tx)
	}

	//
	//
	sooner(func() bool {
		accepted := peerAN.fetchAcceptedTrans()
		return transHold(accepted, trans2)
	})

	//
	//
	//
	everyTrans := append(trans1, trans2...)
	sooner(func() bool {
		acceptedAN := peerAN.fetchAcceptedTrans()
		acceptedBYTE := peerBYTE.fetchAcceptedTrans()
		acceptedCN := peerCN.fetchAcceptedTrans()
		return transHold(acceptedAN, everyTrans) &&
			transHold(acceptedBYTE, everyTrans) &&
			transHold(acceptedCN, everyTrans)
	})

	//
	require.False(t, ownsReplicas(peerAN.fetchAcceptedTrans()))
	require.False(t, ownsReplicas(peerBYTE.fetchAcceptedTrans()))
	require.False(t, ownsReplicas(peerCN.fetchAcceptedTrans()))
}

func VerifySegmentTrans(t *testing.T) {
	createTransfer := func(extent int) kinds.Tx {
		return kinds.Tx(arbitrary.Octets(extent))
	}

	towardTrans := func(extents []int) kinds.Txs {
		txs := make([]kinds.Tx, 0, len(extents))
		for _, extent := range extents {
			txs = append(txs, createTransfer(extent))
		}
		return txs
	}

	for _, tt := range []struct {
		alias   string
		influx  []int
		extent   int
		emission [][]int
	}{
		{
			alias:   "REDACTED",
			influx:  []int{100},
			extent:   200,
			emission: [][]int{{100}},
		},
		{
			alias:   "REDACTED",
			influx:  []int{100},
			extent:   50,
			emission: [][]int{{100}},
		},
		{
			alias:   "REDACTED",
			influx:  []int{100, 100, 100},
			extent:   200,
			emission: [][]int{{100, 100}, {100}},
		},
		{
			alias:   "REDACTED",
			influx:  []int{100, 100, 100},
			extent:   100,
			emission: [][]int{{100}, {100}, {100}},
		},
		{
			alias:   "REDACTED",
			influx:  []int{101, 20, 30, 50, 2, 102, 3},
			extent:   100,
			emission: [][]int{{101}, {20, 30, 50}, {2}, {102}, {3}},
		},
	} {
		t.Run(tt.alias, func(t *testing.T) {
			//
			influx := towardTrans(tt.influx)

			anticipated := make([]kinds.Txs, 0, len(tt.emission))
			for _, segment := range tt.emission {
				anticipated = append(anticipated, towardTrans(segment))
			}

			//
			existing := segmentTrans(influx, tt.extent)

			//
			require.Equal(t, len(anticipated), len(existing), "REDACTED")

			for i, segment := range existing {
				require.Equal(t, len(anticipated[i]), len(segment), "REDACTED", i)
			}
		})
	}
}

type applicationHandlerPeer struct {
	t    *testing.T
	alias string

	app     *ifacemimic.Customer
	txpool *ApplicationTxpool
	handler *ApplicationHandler
	sw      *p2p.Router

	txpoolTrans  kinds.Txs
	acceptedTrans kinds.Txs
	mu          sync.Mutex

	tracer log.Tracer
}

func freshApplicationHandlerPeer(t *testing.T, alias string) *applicationHandlerPeer {
	settings := settings.VerifySettings()
	tracer := log.VerifyingTracer().Using("REDACTED", alias)
	app := ifacemimic.FreshCustomer(t)

	txpool := FreshApplicationTxpool(
		settings.Txpool,
		app,
		UsingMorningTracer(tracer.Using("REDACTED", "REDACTED")),
	)

	handler := FreshApplicationHandler(settings.Txpool, txpool, true)
	handler.AssignTracer(tracer.Using("REDACTED", "REDACTED"))

	ts := &applicationHandlerPeer{
		t:       t,
		alias:    alias,
		app:     app,
		txpool: txpool,
		handler: handler,
		tracer:  tracer,
	}

	ts.configureApplicationSimulate()

	return ts
}

func (ts *applicationHandlerPeer) appendTransfer(tx kinds.Tx) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.tracer.Details("REDACTED", "REDACTED", string(tx))

	ts.txpoolTrans = append(ts.txpoolTrans, tx)
	ts.acceptedTrans = append(ts.acceptedTrans, tx)
}

func (ts *applicationHandlerPeer) harvestTrans() kinds.Txs {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.tracer.Details("REDACTED")

	out := make(kinds.Txs, 0, len(ts.txpoolTrans))
	out = append(out, ts.txpoolTrans...)

	ts.txpoolTrans = ts.txpoolTrans[:0]

	return out
}

func (ts *applicationHandlerPeer) fetchAcceptedTrans() kinds.Txs {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	out := make(kinds.Txs, 0, len(ts.acceptedTrans))
	out = append(out, ts.acceptedTrans...)

	return out
}

func (ts *applicationHandlerPeer) configureApplicationSimulate() {
	simulateGrps := func(procedure string, fn any) *mock.Call {
		return ts.app.On(procedure, mock.Anything, mock.Anything).Return(fn).Maybe()
	}

	simulateGrps("REDACTED", func(_ context.Context, req *iface.SolicitAppendTransfer) (*iface.ReplyAppendTransfer, error) {
		ts.appendTransfer(req.Tx)
		return &iface.ReplyAppendTransfer{
			Cipher: iface.CipherKindOKAY,
		}, nil
	})

	simulateGrps("REDACTED", func(_ context.Context, req *iface.SolicitHarvestTrans) (*iface.ReplyHarvestTrans, error) {
		out := ts.harvestTrans()

		return &iface.ReplyHarvestTrans{Txs: out.TowardSegmentBelongingOctets()}, nil
	})
}

func transHold(set, segment kinds.Txs) bool {
	stash := make(map[string]struct{})

	for _, tx := range set {
		stash[tx.Text()] = struct{}{}
	}

	for _, tx := range segment {
		if _, ok := stash[tx.Text()]; !ok {
			return false
		}
	}

	return true
}

func ownsReplicas(txs kinds.Txs) bool {
	stash := make(map[string]struct{})

	for _, tx := range txs {
		if _, ok := stash[tx.Text()]; ok {
			return true
		}

		stash[tx.Text()] = struct{}{}
	}

	return false
}
