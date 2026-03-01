package txpool

import (
	"encoding/hex"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/go-kit/log/term"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	memoryschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	countTrans  = 1000
	deadline = 120 * time.Second //
)

type nodeStatus struct {
	altitude int64
}

func (ps nodeStatus) ObtainAltitude() int64 {
	return ps.altitude
}

//
//
func VerifyHandlerMulticastTransArtifact(t *testing.T) {
	settings := cfg.VerifySettings()
	//
	//
	//
	//
	const N = 2
	engines, _ := createAlsoRelateEngines(settings, N)
	defer func() {
		for _, r := range engines {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()
	for _, r := range engines {
		for _, node := range r.Router.Nodes().Duplicate() {
			node.Set(kinds.NodeStatusToken, nodeStatus{1})
		}
	}

	txs := appendUnpredictableTrans(t, engines[0].txpool, countTrans, UnfamiliarNodeUUID)
	pauseForeachTransUponEngines(t, txs, engines)
}

//
func VerifyHandlerParallelism(t *testing.T) {
	settings := cfg.VerifySettings()
	const N = 2
	engines, _ := createAlsoRelateEngines(settings, N)
	defer func() {
		for _, r := range engines {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()
	for _, r := range engines {
		for _, node := range r.Router.Nodes().Duplicate() {
			node.Set(kinds.NodeStatusToken, nodeStatus{1})
		}
	}
	var wg sync.WaitGroup

	const countTrans = 5

	for i := 0; i < 1000; i++ {
		wg.Add(2)

		//
		//
		txs := appendUnpredictableTrans(t, engines[0].txpool, countTrans, UnfamiliarNodeUUID)
		go func() {
			defer wg.Done()

			engines[0].txpool.Secure()
			defer engines[0].txpool.Release()

			transferReplies := make([]*iface.InvokeTransferOutcome, len(txs))
			for i := range txs {
				transferReplies[i] = &iface.InvokeTransferOutcome{Cipher: 0}
			}
			err := engines[0].txpool.Revise(1, txs, transferReplies, nil, nil)
			assert.NoError(t, err)
		}()

		//
		//
		_ = appendUnpredictableTrans(t, engines[1].txpool, countTrans, UnfamiliarNodeUUID)
		go func() {
			defer wg.Done()

			engines[1].txpool.Secure()
			defer engines[1].txpool.Release()
			err := engines[1].txpool.Revise(1, []kinds.Tx{}, make([]*iface.InvokeTransferOutcome, 0), nil, nil)
			assert.NoError(t, err)
		}()

		//
		engines[1].txpool.Purge()
	}

	wg.Wait()
}

//
//
func VerifyHandlerNegativeMulticastTowardOriginator(t *testing.T) {
	settings := cfg.VerifySettings()
	const N = 2
	engines, _ := createAlsoRelateEngines(settings, N)
	defer func() {
		for _, r := range engines {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()
	for _, r := range engines {
		for _, node := range r.Router.Nodes().Duplicate() {
			node.Set(kinds.NodeStatusToken, nodeStatus{1})
		}
	}

	const nodeUUID = 1
	appendUnpredictableTrans(t, engines[0].txpool, countTrans, nodeUUID)
	assureNegativeTrans(t, engines[nodeUUID], 100*time.Millisecond)
}

func VerifyTxpoolHandlerMaximumTransferOctets(t *testing.T) {
	settings := cfg.VerifySettings()

	const N = 2
	engines, _ := createAlsoRelateEngines(settings, N)
	defer func() {
		for _, r := range engines {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()
	for _, r := range engines {
		for _, node := range r.Router.Nodes().Duplicate() {
			node.Set(kinds.NodeStatusToken, nodeStatus{1})
		}
	}

	//
	//
	tx1 := statedepot.FreshUnpredictableTransfer(settings.Txpool.MaximumTransferOctets)
	err := engines[0].txpool.InspectTransfer(tx1, func(reply *iface.ReplyInspectTransfer) {
		require.False(t, reply.EqualsFault())
	}, TransferDetails{OriginatorUUID: UnfamiliarNodeUUID})
	require.NoError(t, err)
	pauseForeachTransUponEngines(t, []kinds.Tx{tx1}, engines)

	engines[0].txpool.Purge()
	engines[1].txpool.Purge()

	//
	//
	tx2 := statedepot.FreshUnpredictableTransfer(settings.Txpool.MaximumTransferOctets + 1)
	err = engines[0].txpool.InspectTransfer(tx2, func(reply *iface.ReplyInspectTransfer) {
		require.False(t, reply.EqualsFault())
	}, TransferDetails{OriginatorUUID: UnfamiliarNodeUUID})
	require.Error(t, err)
}

func VerifyMulticastTransferForeachNodeHaltsWheneverNodeHalts(t *testing.T) {
	if testing.Short() {
		t.Skip("REDACTED")
	}

	settings := cfg.VerifySettings()
	const N = 2
	engines, _ := createAlsoRelateEngines(settings, N)
	defer func() {
		for _, r := range engines {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()

	//
	sw := engines[1].Router
	sw.HaltNodeForeachFailure(sw.Nodes().Duplicate()[0], errors.New("REDACTED"))

	//
	//
	leaktest.CheckTimeout(t, 10*time.Second)()
}

func VerifyMulticastTransferForeachNodeHaltsWheneverHandlerHalts(t *testing.T) {
	if testing.Short() {
		t.Skip("REDACTED")
	}

	settings := cfg.VerifySettings()
	const N = 2
	_, routers := createAlsoRelateEngines(settings, N)

	//
	for _, s := range routers {
		assert.NoError(t, s.Halt())
	}

	//
	//
	leaktest.CheckTimeout(t, 10*time.Second)()
}

//
//
func txpoolTracer() log.Tracer {
	return log.VerifyingTracerUsingHueProc(func(tokvals ...any) term.FgBgColor {
		for i := 0; i < len(tokvals)-1; i += 2 {
			if tokvals[i] == "REDACTED" {
				return term.FgBgColor{Fg: term.Color(uint8(tokvals[i+1].(int) + 1))}
			}
		}
		return term.FgBgColor{}
	})
}

//
func createAlsoRelateEngines(settings *cfg.Settings, n int) ([]*Handler, []*p2p.Router) {
	engines := make([]*Handler, n)
	tracer := txpoolTracer()
	for i := 0; i < n; i++ {
		app := statedepot.FreshInsideRamPlatform()
		cc := delegate.FreshRegionalCustomerOriginator(app)
		txpool, sanitize := freshTxpoolUsingApplication(cc)
		defer sanitize()

		engines[i] = FreshHandler(settings.Txpool, txpool, false) //
		engines[i].AssignTracer(tracer.Using("REDACTED", i))
	}

	routers := p2p.CreateAssociatedRouters(settings.P2P, n, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", engines[i])
		return s
	}, p2p.Connect2routers)
	return engines, routers
}

func freshDistinctTrans(n int) kinds.Txs {
	txs := make(kinds.Txs, n)
	for i := 0; i < n; i++ {
		txs[i] = statedepot.FreshTransferOriginatingUUID(i)
	}
	return txs
}

func pauseForeachTransUponEngines(t *testing.T, txs kinds.Txs, engines []*Handler) {
	//
	wg := new(sync.WaitGroup)
	for i, handler := range engines {
		wg.Add(1)
		go func(r *Handler, handlerPosition int) {
			defer wg.Done()
			inspectTransInsideSequence(t, txs, r, handlerPosition)
		}(handler, i)
	}

	complete := make(chan struct{})
	go func() {
		wg.Wait()
		close(complete)
	}()

	clock := time.After(deadline)
	select {
	case <-clock:
		t.Fatal("REDACTED")
	case <-complete:
	}
}

//
func pauseForeachCountTransInsideTxpool(countTrans int, txpool Txpool) {
	for txpool.Extent() < countTrans {
		time.Sleep(time.Millisecond * 100)
	}
}

//
//
//
//
//

//
//
//
//

//
//
func inspectTransInsideSequence(t *testing.T, txs kinds.Txs, handler *Handler, handlerPosition int) {
	pauseForeachCountTransInsideTxpool(len(txs), handler.txpool)

	//
	harvestedTrans := handler.txpool.HarvestMaximumTrans(len(txs))
	for i, tx := range txs {
		assert.Equalf(t, tx, harvestedTrans[i],
			"REDACTED", i, handlerPosition, tx, harvestedTrans[i])
	}
}

//
func assureNegativeTrans(t *testing.T, handler *Handler, deadline time.Duration) {
	time.Sleep(deadline) //
	assert.Zero(t, handler.txpool.Extent())
}

func VerifyTxpoolArrays(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias string
		tx       []byte
		expirationOctets string
	}{
		{"REDACTED", []byte{123}, "REDACTED"},
		{"REDACTED", []byte("REDACTED"), "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		msg := memoryschema.Signal{
			Sum: &memoryschema.Artifact_Trans{
				Txs: &memoryschema.Txs{Txs: [][]byte{tc.tx}},
			},
		}
		bz, err := msg.Serialize()
		require.NoError(t, err, tc.verifyAlias)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyAlias)
	}
}
