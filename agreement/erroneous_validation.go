package agreement

import (
	"testing"
	"time"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/log"
	ctrng "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	enginecons "github.com/valkyrieworks/schema/consensuscore/agreement"
	ctschema "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//

//
//
func TestReactorInvalidPrecommit(t *testing.T) {
	N := 4
	css, cleanup := randConsensusNet(t, N, "REDACTED", newMockTickerFunc(true), newKVStore,
		func(c *cfg.Config) {
			c.Consensus.TimeoutPropose = 3000 * time.Millisecond
			c.Consensus.TimeoutPrevote = 1000 * time.Millisecond
			c.Consensus.TimeoutPrecommit = 1000 * time.Millisecond
		})
	defer cleanup()

	for i := 0; i < N; i++ {
		ticker := NewTimeoutTicker()
		ticker.SetLogger(css[i].Logger)
		css[i].SetTimeoutTicker(ticker)

	}

	reactors, blocksSubs, eventBuses := startConsensusNet(t, css, N)
	defer stopConsensusNet(log.TestingLogger(), reactors, eventBuses)

	//
	byzValIdx := N - 1
	byzVal := css[byzValIdx]
	byzR := reactors[byzValIdx]

	//
	//
	byzVal.mtx.Lock()
	pv := byzVal.privValidator
	byzVal.doPrevote = func(int64, int32) {
		invalidDoPrevoteFunc(t, byzVal, byzR.Switch, pv)
	}
	byzVal.mtx.Unlock()

	//
	//
	for i := 0; i < 10; i++ {
		timeoutWaitGroup(N, func(j int) {
			<-blocksSubs[j].Out()
		})
	}
}

func invalidDoPrevoteFunc(t *testing.T, cs *State, sw p2p.Switcher, pv kinds.PrivValidator) {
	//
	//
	//
	//
	go func() {
		cs.mtx.Lock()
		defer cs.mtx.Unlock()
		cs.privValidator = pv
		pubKey, err := cs.privValidator.GetPubKey()
		if err != nil {
			panic(err)
		}
		addr := pubKey.Address()
		valIndex, _ := cs.Validators.GetByAddress(addr)

		//
		blockHash := octets.HexBytes(ctrng.Bytes(32))
		precommit := &kinds.Vote{
			ValidatorAddress: addr,
			ValidatorIndex:   valIndex,
			Height:           cs.Height,
			Round:            cs.Round,
			Timestamp:        cs.voteTime(),
			Type:             ctschema.PrecommitType,
			BlockID: kinds.BlockID{
				Hash:          blockHash,
				PartSetHeader: kinds.PartSetHeader{Total: 1, Hash: ctrng.Bytes(32)},
			},
		}
		p := precommit.ToProto()
		err = cs.privValidator.SignVote(cs.state.ChainID, p)
		if err != nil {
			t.Error(err)
		}
		precommit.Signature = p.Signature
		precommit.ExtensionSignature = p.ExtensionSignature
		cs.privValidator = nil //

		peers := sw.Peers().Copy()
		for _, peer := range peers {
			cs.Logger.Info("REDACTED", "REDACTED", blockHash, "REDACTED", peer)
			peer.Send(p2p.Envelope{
				Message:   &enginecons.Vote{Vote: precommit.ToProto()},
				ChannelID: VoteChannel,
			})
		}
	}()
}
