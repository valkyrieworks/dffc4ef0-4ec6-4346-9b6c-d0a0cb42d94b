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

	"github.com/valkyrieworks/iface/instance/objectdepot"
	iface "github.com/valkyrieworks/iface/kinds"
	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	memprotocol "github.com/valkyrieworks/schema/consensuscore/txpool"
	"github.com/valkyrieworks/gateway"
	"github.com/valkyrieworks/kinds"
)

const (
	countTrans  = 1000
	deadline = 120 * time.Second //
)

type nodeStatus struct {
	level int64
}

func (ps nodeStatus) FetchLevel() int64 {
	return ps.level
}

//
//
func VerifyHandlerMulticastTransSignal(t *testing.T) {
	settings := cfg.VerifySettings()
	//
	//
	//
	//
	const N = 2
	handlers, _ := createAndEstablishHandlers(settings, N)
	defer func() {
		for _, r := range handlers {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()
	for _, r := range handlers {
		for _, node := range r.Router.Nodes().Clone() {
			node.Set(kinds.NodeStatusKey, nodeStatus{1})
		}
	}

	txs := appendArbitraryTrans(t, handlers[0].txpool, countTrans, UnclearNodeUID)
	waitForTransOnHandlers(t, txs, handlers)
}

//
func VerifyHandlerParallelism(t *testing.T) {
	settings := cfg.VerifySettings()
	const N = 2
	handlers, _ := createAndEstablishHandlers(settings, N)
	defer func() {
		for _, r := range handlers {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()
	for _, r := range handlers {
		for _, node := range r.Router.Nodes().Clone() {
			node.Set(kinds.NodeStatusKey, nodeStatus{1})
		}
	}
	var wg sync.WaitGroup

	const countTrans = 5

	for i := 0; i < 1000; i++ {
		wg.Add(2)

		//
		//
		txs := appendArbitraryTrans(t, handlers[0].txpool, countTrans, UnclearNodeUID)
		go func() {
			defer wg.Done()

			handlers[0].txpool.Secure()
			defer handlers[0].txpool.Release()

			transferReplies := make([]*iface.InvokeTransferOutcome, len(txs))
			for i := range txs {
				transferReplies[i] = &iface.InvokeTransferOutcome{Code: 0}
			}
			err := handlers[0].txpool.Modify(1, txs, transferReplies, nil, nil)
			assert.NoError(t, err)
		}()

		//
		//
		_ = appendArbitraryTrans(t, handlers[1].txpool, countTrans, UnclearNodeUID)
		go func() {
			defer wg.Done()

			handlers[1].txpool.Secure()
			defer handlers[1].txpool.Release()
			err := handlers[1].txpool.Modify(1, []kinds.Tx{}, make([]*iface.InvokeTransferOutcome, 0), nil, nil)
			assert.NoError(t, err)
		}()

		//
		handlers[1].txpool.Purge()
	}

	wg.Wait()
}

//
//
func VerifyHandlerNoMulticastToEmitter(t *testing.T) {
	settings := cfg.VerifySettings()
	const N = 2
	handlers, _ := createAndEstablishHandlers(settings, N)
	defer func() {
		for _, r := range handlers {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()
	for _, r := range handlers {
		for _, node := range r.Router.Nodes().Clone() {
			node.Set(kinds.NodeStatusKey, nodeStatus{1})
		}
	}

	const nodeUID = 1
	appendArbitraryTrans(t, handlers[0].txpool, countTrans, nodeUID)
	assureNoTrans(t, handlers[nodeUID], 100*time.Millisecond)
}

func VerifyTxpoolHandlerMaximumTransferOctets(t *testing.T) {
	settings := cfg.VerifySettings()

	const N = 2
	handlers, _ := createAndEstablishHandlers(settings, N)
	defer func() {
		for _, r := range handlers {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()
	for _, r := range handlers {
		for _, node := range r.Router.Nodes().Clone() {
			node.Set(kinds.NodeStatusKey, nodeStatus{1})
		}
	}

	//
	//
	tx1 := objectdepot.NewArbitraryTransfer(settings.Txpool.MaximumTransferOctets)
	err := handlers[0].txpool.InspectTransfer(tx1, func(reply *iface.ReplyInspectTransfer) {
		require.False(t, reply.IsErr())
	}, TransferDetails{EmitterUID: UnclearNodeUID})
	require.NoError(t, err)
	waitForTransOnHandlers(t, []kinds.Tx{tx1}, handlers)

	handlers[0].txpool.Purge()
	handlers[1].txpool.Purge()

	//
	//
	tx2 := objectdepot.NewArbitraryTransfer(settings.Txpool.MaximumTransferOctets + 1)
	err = handlers[0].txpool.InspectTransfer(tx2, func(reply *iface.ReplyInspectTransfer) {
		require.False(t, reply.IsErr())
	}, TransferDetails{EmitterUID: UnclearNodeUID})
	require.Error(t, err)
}

func VerifyMulticastTransferForNodeHaltsWhenNodeHalts(t *testing.T) {
	if testing.Short() {
		t.Skip("REDACTED")
	}

	settings := cfg.VerifySettings()
	const N = 2
	handlers, _ := createAndEstablishHandlers(settings, N)
	defer func() {
		for _, r := range handlers {
			if err := r.Halt(); err != nil {
				assert.NoError(t, err)
			}
		}
	}()

	//
	sw := handlers[1].Router
	sw.HaltNodeForFault(sw.Nodes().Clone()[0], errors.New("REDACTED"))

	//
	//
	leaktest.CheckTimeout(t, 10*time.Second)()
}

func VerifyMulticastTransferForNodeHaltsWhenHandlerHalts(t *testing.T) {
	if testing.Short() {
		t.Skip("REDACTED")
	}

	settings := cfg.VerifySettings()
	const N = 2
	_, routers := createAndEstablishHandlers(settings, N)

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
	return log.VerifyingTracerWithHueFn(func(keyvalues ...any) term.FgBgColor {
		for i := 0; i < len(keyvalues)-1; i += 2 {
			if keyvalues[i] == "REDACTED" {
				return term.FgBgColor{Fg: term.Color(uint8(keyvalues[i+1].(int) + 1))}
			}
		}
		return term.FgBgColor{}
	})
}

//
func createAndEstablishHandlers(settings *cfg.Settings, n int) ([]*Handler, []*p2p.Router) {
	handlers := make([]*Handler, n)
	tracer := txpoolTracer()
	for i := 0; i < n; i++ {
		app := objectdepot.NewInRamSoftware()
		cc := gateway.NewNativeCustomerOriginator(app)
		txpool, sanitize := newTxpoolWithApplication(cc)
		defer sanitize()

		handlers[i] = NewHandler(settings.Txpool, txpool, false) //
		handlers[i].AssignTracer(tracer.With("REDACTED", i))
	}

	routers := p2p.CreateLinkedRouters(settings.P2P, n, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", handlers[i])
		return s
	}, p2p.Connect2routers)
	return handlers, routers
}

func newDistinctTrans(n int) kinds.Txs {
	txs := make(kinds.Txs, n)
	for i := 0; i < n; i++ {
		txs[i] = objectdepot.NewTransferFromUID(i)
	}
	return txs
}

func waitForTransOnHandlers(t *testing.T, txs kinds.Txs, handlers []*Handler) {
	//
	wg := new(sync.WaitGroup)
	for i, handler := range handlers {
		wg.Add(1)
		go func(r *Handler, handlerOrdinal int) {
			defer wg.Done()
			inspectTransInSequence(t, txs, r, handlerOrdinal)
		}(handler, i)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	clock := time.After(deadline)
	select {
	case <-clock:
		t.Fatal("REDACTED")
	case <-done:
	}
}

//
func waitForCountTransInTxpool(countTrans int, txpool Txpool) {
	for txpool.Volume() < countTrans {
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
func inspectTransInSequence(t *testing.T, txs kinds.Txs, handler *Handler, handlerOrdinal int) {
	waitForCountTransInTxpool(len(txs), handler.txpool)

	//
	harvestedTrans := handler.txpool.HarvestMaximumTrans(len(txs))
	for i, tx := range txs {
		assert.Equalf(t, tx, harvestedTrans[i],
			"REDACTED", i, handlerOrdinal, tx, harvestedTrans[i])
	}
}

//
func assureNoTrans(t *testing.T, handler *Handler, deadline time.Duration) {
	time.Sleep(deadline) //
	assert.Zero(t, handler.txpool.Volume())
}

func VerifyTxpoolArrays(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel string
		tx       []byte
		expirationOctets string
	}{
		{"REDACTED", []byte{123}, "REDACTED"},
		{"REDACTED", []byte("REDACTED"), "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		msg := memprotocol.Signal{
			Sum: &memprotocol.Signal_Trans{
				Txs: &memprotocol.Txs{Txs: [][]byte{tc.tx}},
			},
		}
		bz, err := msg.Serialize()
		require.NoError(t, err, tc.verifyLabel)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyLabel)
	}
}
