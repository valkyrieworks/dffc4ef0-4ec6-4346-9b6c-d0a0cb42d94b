package agreement

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/atci/instance/dbstore"
	atci "github.com/valkyrieworks/atci/kinds"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

//
func assertMempool(txn txNotifier) txpool.Mempool {
	return txn.(txpool.Mempool)
}

func TestMempoolNoProgressUntilTxsAvailable(t *testing.T) {
	config := ResetConfig("REDACTED")
	defer os.RemoveAll(config.RootDir)
	config.Consensus.CreateEmptyBlocks = false
	state, privVals := randGenesisState(1, false, 10, nil)
	app := dbstore.NewInMemoryApplication()
	resp, err := app.Info(context.Background(), gateway.RequestInfo)
	require.NoError(t, err)
	state.AppHash = resp.LastBlockAppHash
	cs := newStateWithConfig(config, state, privVals[0], app)
	assertMempool(cs.txNotifier).EnableTxsAvailable()
	height, round := cs.Height, cs.Round
	newBlockCh := subscribe(cs.eventBus, kinds.EventQueryNewBlock)
	startTestRound(cs, height, round)

	ensureNewEventOnChannel(newBlockCh) //
	ensureNoNewEventOnChannel(newBlockCh)
	deliverTxsRange(t, cs, 0, 1)
	ensureNewEventOnChannel(newBlockCh) //
	ensureNewEventOnChannel(newBlockCh) //
	ensureNoNewEventOnChannel(newBlockCh)
}

func TestMempoolProgressAfterCreateEmptyBlocksInterval(t *testing.T) {
	config := ResetConfig("REDACTED")
	defer os.RemoveAll(config.RootDir)

	config.Consensus.CreateEmptyBlocksInterval = ensureTimeout
	state, privVals := randGenesisState(1, false, 10, nil)
	app := dbstore.NewInMemoryApplication()
	resp, err := app.Info(context.Background(), gateway.RequestInfo)
	require.NoError(t, err)
	state.AppHash = resp.LastBlockAppHash
	cs := newStateWithConfig(config, state, privVals[0], app)

	assertMempool(cs.txNotifier).EnableTxsAvailable()

	newBlockCh := subscribe(cs.eventBus, kinds.EventQueryNewBlock)
	startTestRound(cs, cs.Height, cs.Round)

	ensureNewEventOnChannel(newBlockCh)   //
	ensureNoNewEventOnChannel(newBlockCh) //
	ensureNewEventOnChannel(newBlockCh)   //
}

func TestMempoolProgressInHigherRound(t *testing.T) {
	config := ResetConfig("REDACTED")
	defer os.RemoveAll(config.RootDir)
	config.Consensus.CreateEmptyBlocks = false
	state, privVals := randGenesisState(1, false, 10, nil)
	cs := newStateWithConfig(config, state, privVals[0], dbstore.NewInMemoryApplication())
	assertMempool(cs.txNotifier).EnableTxsAvailable()
	height, round := cs.Height, cs.Round
	newBlockCh := subscribe(cs.eventBus, kinds.EventQueryNewBlock)
	newRoundCh := subscribe(cs.eventBus, kinds.EventQueryNewRound)
	timeoutCh := subscribe(cs.eventBus, kinds.EventQueryTimeoutPropose)
	cs.setProposal = func(proposal *kinds.Proposal) error {
		if cs.Height == 2 && cs.Round == 0 {
			//
			//
			cs.Logger.Info("REDACTED")
			return nil
		}
		return cs.defaultSetProposal(proposal)
	}
	startTestRound(cs, height, round)

	ensureNewRound(newRoundCh, height, round) //
	ensureNewEventOnChannel(newBlockCh)       //

	height++ //
	round = 0

	ensureNewRound(newRoundCh, height, round) //
	deliverTxsRange(t, cs, 0, 1)              //
	ensureNewTimeout(timeoutCh, height, round, cs.config.TimeoutPropose.Nanoseconds())

	round++                                   //
	ensureNewRound(newRoundCh, height, round) //
	ensureNewEventOnChannel(newBlockCh)       //
}

func deliverTxsRange(t *testing.T, cs *State, start, end int) {
	//
	for i := start; i < end; i++ {
		err := assertMempool(cs.txNotifier).CheckTx(dbstore.NewTx(fmt.Sprintf("REDACTED", i), "REDACTED"), nil, txpool.TxInfo{})
		require.NoError(t, err)
	}
}

func TestMempoolTxConcurrentWithCommit(t *testing.T) {
	state, privVals := randGenesisState(1, false, 10, nil)
	blockDB := dbm.NewMemDB()
	stateStore := sm.NewStore(blockDB, sm.StoreOptions{DiscardABCIResponses: false})
	cs := newStateWithConfigAndBlockStore(config, state, privVals[0], dbstore.NewInMemoryApplication(), blockDB)
	err := stateStore.Save(state)
	require.NoError(t, err)
	newBlockEventsCh := subscribe(cs.eventBus, kinds.EventQueryNewBlockEvents)

	const numTxs int64 = 3000
	go deliverTxsRange(t, cs, 0, int(numTxs))

	startTestRound(cs, cs.Height, cs.Round)
	for n := int64(0); n < numTxs; {
		select {
		case msg := <-newBlockEventsCh:
			event := msg.Data().(kinds.EventDataNewBlockEvents)
			n += event.NumTxs
			t.Log("REDACTED", "REDACTED", event.NumTxs, "REDACTED", n)
		case <-time.After(30 * time.Second):
			t.Fatal("REDACTED")
		}
	}
}

func TestMempoolRmBadTx(t *testing.T) {
	state, privVals := randGenesisState(1, false, 10, nil)
	app := dbstore.NewInMemoryApplication()
	blockDB := dbm.NewMemDB()
	stateStore := sm.NewStore(blockDB, sm.StoreOptions{DiscardABCIResponses: false})
	cs := newStateWithConfigAndBlockStore(config, state, privVals[0], app, blockDB)
	err := stateStore.Save(state)
	require.NoError(t, err)

	//
	txBytes := dbstore.NewTx("REDACTED", "REDACTED")
	res, err := app.FinalizeBlock(context.Background(), &atci.RequestFinalizeBlock{Txs: [][]byte{txBytes}})
	require.NoError(t, err)
	assert.False(t, res.TxResults[0].IsErr())
	assert.True(t, len(res.AppHash) > 0)

	_, err = app.Commit(context.Background(), &atci.RequestCommit{})
	require.NoError(t, err)

	emptyMempoolCh := make(chan struct{})
	checkTxRespCh := make(chan struct{})
	go func() {
		//
		//
		//
		invalidTx := []byte("REDACTED")
		err := assertMempool(cs.txNotifier).CheckTx(invalidTx, func(r *atci.ResponseCheckTx) {
			if r.Code != dbstore.CodeTypeInvalidTxFormat {
				t.Errorf("REDACTED", r)
				return
			}
			checkTxRespCh <- struct{}{}
		}, txpool.TxInfo{})
		if err != nil {
			t.Errorf("REDACTED", err)
			return
		}

		//
		for {
			txs := assertMempool(cs.txNotifier).ReapMaxBytesMaxGas(int64(len(invalidTx)), -1)
			if len(txs) == 0 {
				emptyMempoolCh <- struct{}{}
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	//
	ticker := time.After(time.Second * 5)
	select {
	case <-checkTxRespCh:
		//
	case <-ticker:
		t.Errorf("REDACTED")
		return
	}

	//
	ticker = time.After(time.Second * 5)
	select {
	case <-emptyMempoolCh:
		//
	case <-ticker:
		t.Errorf("REDACTED")
		return
	}
}
