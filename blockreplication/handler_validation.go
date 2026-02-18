package blockreplication

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"

	bcschema "github.com/valkyrieworks/schema/consensuscore/blockreplication"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	atci "github.com/valkyrieworks/atci/kinds"
	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/intrinsic/validation"
	"github.com/valkyrieworks/utils/log"
	tpmocks "github.com/valkyrieworks/txpool/simulations"
	"github.com/valkyrieworks/p2p"
	ctschema "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	cttime "github.com/valkyrieworks/kinds/moment"
)

var config *cfg.Config

func randGenesisDoc(numValidators int, randPower bool, minPower int64) (*kinds.GenesisDoc, []kinds.PrivValidator) {
	validators := make([]kinds.GenesisValidator, numValidators)
	privValidators := make([]kinds.PrivValidator, numValidators)
	for i := 0; i < numValidators; i++ {
		val, privVal := kinds.RandValidator(randPower, minPower)
		validators[i] = kinds.GenesisValidator{
			PubKey: val.PubKey,
			Power:  val.VotingPower,
		}
		privValidators[i] = privVal
	}
	sort.Sort(kinds.PrivValidatorsByAddress(privValidators))

	consPar := kinds.DefaultConsensusParams()
	consPar.ABCI.VoteExtensionsEnableHeight = 1
	return &kinds.GenesisDoc{
		GenesisTime:     cttime.Now(),
		ChainID:         validation.DefaultTestChainID,
		Validators:      validators,
		ConsensusParams: consPar,
	}, privValidators
}

type ReactorPair struct {
	reactor *ByzantineReactor
	app     gateway.AppConns
}

func newReactor(
	t *testing.T,
	logger log.Logger,
	genDoc *kinds.GenesisDoc,
	privVals []kinds.PrivValidator,
	maxBlockHeight int64,
	incorrectData ...int64,
) ReactorPair {
	if len(privVals) != 1 {
		panic("REDACTED")
	}
	var incorrectBlock int64 = 0
	if len(incorrectData) > 0 {
		incorrectBlock = incorrectData[0]
	}

	app := atci.NewBaseApplication()
	cc := gateway.NewLocalClientCreator(app)
	proxyApp := gateway.NewAppConns(cc, gateway.NopMetrics())
	err := proxyApp.Start()
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	blockDB := dbm.NewMemDB()
	stateDB := dbm.NewMemDB()
	stateStore := sm.NewStore(stateDB, sm.StoreOptions{
		DiscardABCIResponses: false,
	})
	blockStore := depot.NewBlockStore(blockDB)

	state, err := stateStore.LoadFromDBOrGenesisDoc(genDoc)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	mp := &tpmocks.Mempool{}
	mp.On("REDACTED").Return()
	mp.On("REDACTED").Return()
	mp.On("REDACTED", mock.Anything).Return(nil)
	mp.On("REDACTED",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(nil)

	//
	//
	//
	blockSync := true
	db := dbm.NewMemDB()
	stateStore = sm.NewStore(db, sm.StoreOptions{
		DiscardABCIResponses: false,
	})
	blockExec := sm.NewBlockExecutor(stateStore, log.TestingLogger(), proxyApp.Consensus(),
		mp, sm.EmptyEvidencePool{}, blockStore)
	if err = stateStore.Save(state); err != nil {
		panic(err)
	}

	//
	seenExtCommit := &kinds.ExtendedCommit{}

	pubKey, err := privVals[0].GetPubKey()
	if err != nil {
		panic(err)
	}
	addr := pubKey.Address()
	idx, _ := state.Validators.GetByAddress(addr)

	//
	for blockHeight := int64(1); blockHeight <= maxBlockHeight; blockHeight++ {
		voteExtensionIsEnabled := genDoc.ConsensusParams.ABCI.VoteExtensionsEnabled(blockHeight)

		lastExtCommit := seenExtCommit.Clone()

		thisBlock, err := state.MakeBlock(blockHeight, nil, lastExtCommit.ToCommit(), nil, state.Validators.Proposer.Address)
		require.NoError(t, err)

		thisParts, err := thisBlock.MakePartSet(kinds.BlockPartSizeBytes)
		require.NoError(t, err)
		blockID := kinds.BlockID{Hash: thisBlock.Hash(), PartSetHeader: thisParts.Header()}

		//
		vote, err := kinds.MakeVote(
			privVals[0],
			thisBlock.ChainID,
			idx,
			thisBlock.Height,
			0,
			ctschema.PrecommitType,
			blockID,
			time.Now(),
		)
		if err != nil {
			panic(err)
		}
		seenExtCommit = &kinds.ExtendedCommit{
			Height:             vote.Height,
			Round:              vote.Round,
			BlockID:            blockID,
			ExtendedSignatures: []kinds.ExtendedCommitSig{vote.ExtendedCommitSig()},
		}

		state, err = blockExec.ApplyBlock(state, blockID, thisBlock)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}

		saveCorrectVoteExtensions := blockHeight != incorrectBlock
		if saveCorrectVoteExtensions == voteExtensionIsEnabled {
			blockStore.SaveBlockWithExtendedCommit(thisBlock, thisParts, seenExtCommit)
		} else {
			blockStore.SaveBlock(thisBlock, thisParts, seenExtCommit.ToCommit())
		}
	}

	r := NewReactor(blockSync, false, state.Copy(), blockExec, blockStore, nil, 0, NopMetrics())
	bcReactor := NewByzantineReactor(incorrectBlock, r)
	bcReactor.SetLogger(logger.With("REDACTED", "REDACTED"))

	return ReactorPair{bcReactor, proxyApp}
}

func TestNoBlockResponse(t *testing.T) {
	config = validation.ResetTestRoot("REDACTED")
	defer os.RemoveAll(config.RootDir)
	genDoc, privVals := randGenesisDoc(1, false, 30)

	maxBlockHeight := int64(65)

	reactorPairs := make([]ReactorPair, 2)

	reactorPairs[0] = newReactor(t, log.TestingLogger(), genDoc, privVals, maxBlockHeight)
	reactorPairs[1] = newReactor(t, log.TestingLogger(), genDoc, privVals, 0)

	p2p.MakeConnectedSwitches(config.P2P, 2, func(i int, s *p2p.Switch) *p2p.Switch {
		s.AddReactor("REDACTED", reactorPairs[i].reactor)
		return s
	}, p2p.Connect2Switches)

	defer func() {
		for _, r := range reactorPairs {
			err := r.reactor.Stop()
			require.NoError(t, err)
			err = r.app.Stop()
			require.NoError(t, err)
		}
	}()

	tests := []struct {
		height   int64
		existent bool
	}{
		{maxBlockHeight + 2, false},
		{10, true},
		{1, true},
		{100, false},
	}

	for !reactorPairs[1].reactor.pool.IsCaughtUp() {
		time.Sleep(10 * time.Millisecond)
	}

	assert.Equal(t, maxBlockHeight, reactorPairs[0].reactor.store.Height())

	for _, tt := range tests {
		block := reactorPairs[1].reactor.store.LoadBlock(tt.height)
		if tt.existent {
			assert.True(t, block != nil)
		} else {
			assert.True(t, block == nil)
		}
	}
}

//
//
//
//
//
func TestBadBlockStopsPeer(t *testing.T) {
	config = validation.ResetTestRoot("REDACTED")
	defer os.RemoveAll(config.RootDir)
	genDoc, privVals := randGenesisDoc(1, false, 30)

	maxBlockHeight := int64(148)

	//
	otherGenDoc, otherPrivVals := randGenesisDoc(1, false, 30)
	otherChain := newReactor(t, log.TestingLogger(), otherGenDoc, otherPrivVals, maxBlockHeight)

	defer func() {
		err := otherChain.reactor.Stop()
		require.Error(t, err)
		err = otherChain.app.Stop()
		require.NoError(t, err)
	}()

	reactorPairs := make([]ReactorPair, 4)

	reactorPairs[0] = newReactor(t, log.TestingLogger(), genDoc, privVals, maxBlockHeight)
	reactorPairs[1] = newReactor(t, log.TestingLogger(), genDoc, privVals, 0)
	reactorPairs[2] = newReactor(t, log.TestingLogger(), genDoc, privVals, 0)
	reactorPairs[3] = newReactor(t, log.TestingLogger(), genDoc, privVals, 0)

	switches := p2p.MakeConnectedSwitches(config.P2P, 4, func(i int, s *p2p.Switch) *p2p.Switch {
		s.AddReactor("REDACTED", reactorPairs[i].reactor)
		return s
	}, p2p.Connect2Switches)

	defer func() {
		for _, r := range reactorPairs {
			err := r.reactor.Stop()
			require.NoError(t, err)

			err = r.app.Stop()
			require.NoError(t, err)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		caughtUp := true
		for _, r := range reactorPairs {
			if !r.reactor.pool.IsCaughtUp() {
				caughtUp = false
			}
		}
		if caughtUp {
			break
		}
	}

	//
	assert.Equal(t, 3, reactorPairs[1].reactor.Switch.Peers().Size())

	//
	//
	reactorPairs[3].reactor.store = otherChain.reactor.store

	lastReactorPair := newReactor(t, log.TestingLogger(), genDoc, privVals, 0)
	reactorPairs = append(reactorPairs, lastReactorPair)

	switches = append(switches, p2p.MakeConnectedSwitches(config.P2P, 1, func(i int, s *p2p.Switch) *p2p.Switch {
		s.AddReactor("REDACTED", reactorPairs[len(reactorPairs)-1].reactor)
		return s
	}, p2p.Connect2Switches)...)

	for i := 0; i < len(reactorPairs)-1; i++ {
		p2p.Connect2Switches(switches, i, len(reactorPairs)-1)
	}

	for !lastReactorPair.reactor.pool.IsCaughtUp() && lastReactorPair.reactor.Switch.Peers().Size() != 0 {
		time.Sleep(1 * time.Second)
	}

	assert.True(t, lastReactorPair.reactor.Switch.Peers().Size() < len(reactorPairs)-1)
}

func TestCheckSwitchToConsensusLastHeightZero(t *testing.T) {
	const maxBlockHeight = int64(45)

	config = validation.ResetTestRoot("REDACTED")
	defer os.RemoveAll(config.RootDir)
	genDoc, privVals := randGenesisDoc(1, false, 30)

	reactorPairs := make([]ReactorPair, 1, 2)
	reactorPairs[0] = newReactor(t, log.TestingLogger(), genDoc, privVals, 0)
	reactorPairs[0].reactor.intervalSwitchToConsensus = 50 * time.Millisecond
	defer func() {
		for _, r := range reactorPairs {
			err := r.reactor.Stop()
			require.NoError(t, err)
			err = r.app.Stop()
			require.NoError(t, err)
		}
	}()

	reactorPairs = append(reactorPairs, newReactor(t, log.TestingLogger(), genDoc, privVals, maxBlockHeight))

	var switches []*p2p.Switch
	for _, r := range reactorPairs {
		switches = append(switches, p2p.MakeConnectedSwitches(config.P2P, 1, func(i int, s *p2p.Switch) *p2p.Switch {
			s.AddReactor("REDACTED", r.reactor)
			return s
		}, p2p.Connect2Switches)...)
	}

	time.Sleep(60 * time.Millisecond)

	//
	p2p.Connect2Switches(switches, 0, 1)

	startTime := time.Now()
	for {
		time.Sleep(20 * time.Millisecond)
		caughtUp := true
		for _, r := range reactorPairs {
			if !r.reactor.pool.IsCaughtUp() {
				caughtUp = false
				break
			}
		}
		if caughtUp {
			break
		}
		if time.Since(startTime) > 90*time.Second {
			msg := "REDACTED"
			for i, r := range reactorPairs {
				h, p, lr := r.reactor.pool.GetStatus()
				c := r.reactor.pool.IsCaughtUp()
				msg += fmt.Sprintf("REDACTED", i, h, p, lr, c)
			}
			require.Fail(t, msg)
		}
	}

	//
	//
	//
	const maxDiff = 3
	for _, r := range reactorPairs {
		assert.GreaterOrEqual(t, r.reactor.store.Height(), maxBlockHeight-maxDiff)
	}
}

func ExtendedCommitNetworkHelper(t *testing.T, maxBlockHeight int64, enableVoteExtensionAt int64, invalidBlockHeightAt int64) {
	config = validation.ResetTestRoot("REDACTED")
	defer os.RemoveAll(config.RootDir)
	genDoc, privVals := randGenesisDoc(1, false, 30)
	genDoc.ConsensusParams.ABCI.VoteExtensionsEnableHeight = enableVoteExtensionAt

	reactorPairs := make([]ReactorPair, 1, 2)
	reactorPairs[0] = newReactor(t, log.TestingLogger(), genDoc, privVals, 0)
	reactorPairs[0].reactor.intervalSwitchToConsensus = 50 * time.Millisecond
	defer func() {
		for _, r := range reactorPairs {
			err := r.reactor.Stop()
			require.NoError(t, err)
			err = r.app.Stop()
			require.NoError(t, err)
		}
	}()

	reactorPairs = append(reactorPairs, newReactor(t, log.TestingLogger(), genDoc, privVals, maxBlockHeight, invalidBlockHeightAt))

	var switches []*p2p.Switch
	for _, r := range reactorPairs {
		switches = append(switches, p2p.MakeConnectedSwitches(config.P2P, 1, func(i int, s *p2p.Switch) *p2p.Switch {
			s.AddReactor("REDACTED", r.reactor)
			return s
		}, p2p.Connect2Switches)...)
	}

	time.Sleep(60 * time.Millisecond)

	//
	p2p.Connect2Switches(switches, 0, 1)

	startTime := time.Now()
	for {
		time.Sleep(20 * time.Millisecond)
		//
		require.False(t, reactorPairs[0].reactor.pool.IsCaughtUp(), "REDACTED")
		//
		if time.Since(startTime) > 5*time.Second {
			assert.Equal(t, 0, reactorPairs[0].reactor.Switch.Peers().Size(), "REDACTED")
			assert.Equal(t, 0, reactorPairs[1].reactor.Switch.Peers().Size(), "REDACTED")
			break
		}
	}
}

//
func TestCheckExtendedCommitExtra(t *testing.T) {
	const maxBlockHeight = 10
	const enableVoteExtension = 5
	const invalidBlockHeight = 3

	ExtendedCommitNetworkHelper(t, maxBlockHeight, enableVoteExtension, invalidBlockHeight)
}

//
func TestCheckExtendedCommitMissing(t *testing.T) {
	const maxBlockHeight = 10
	const enableVoteExtension = 5
	const invalidBlockHeight = 8

	ExtendedCommitNetworkHelper(t, maxBlockHeight, enableVoteExtension, invalidBlockHeight)
}

//
//
//
//
type ByzantineReactor struct {
	*Reactor
	corruptedBlock int64
}

func NewByzantineReactor(invalidBlock int64, conR *Reactor) *ByzantineReactor {
	return &ByzantineReactor{
		Reactor:        conR,
		corruptedBlock: invalidBlock,
	}
}

//
//
//
func (bcR *ByzantineReactor) respondToPeer(msg *bcschema.BlockRequest, src p2p.Peer) (queued bool) {
	block := bcR.store.LoadBlock(msg.Height)
	if block == nil {
		bcR.Logger.Info("REDACTED", "REDACTED", src, "REDACTED", msg.Height)
		return src.TrySend(p2p.Envelope{
			ChannelID: BlocksyncChannel,
			Message:   &bcschema.NoBlockResponse{Height: msg.Height},
		})
	}

	state, err := bcR.blockExec.Store().Load()
	if err != nil {
		bcR.Logger.Error("REDACTED", "REDACTED", err)
		return false
	}
	var extCommit *kinds.ExtendedCommit
	voteExtensionEnabled := state.ConsensusParams.ABCI.VoteExtensionsEnabled(msg.Height)
	incorrectBlock := bcR.corruptedBlock == msg.Height
	if voteExtensionEnabled && !incorrectBlock || !voteExtensionEnabled && incorrectBlock {
		extCommit = bcR.store.LoadBlockExtendedCommit(msg.Height)
		if extCommit == nil {
			bcR.Logger.Error("REDACTED", "REDACTED", block)
			return false
		}
	}

	bl, err := block.ToProto()
	if err != nil {
		bcR.Logger.Error("REDACTED", "REDACTED", err)
		return false
	}

	return src.TrySend(p2p.Envelope{
		ChannelID: BlocksyncChannel,
		Message: &bcschema.BlockResponse{
			Block:     bl,
			ExtCommit: extCommit.ToProto(),
		},
	})
}

//
//
func (bcR *ByzantineReactor) Receive(e p2p.Envelope) {
	if err := ValidateMsg(e.Message); err != nil {
		bcR.Logger.Error("REDACTED", "REDACTED", e.Src, "REDACTED", e.Message, "REDACTED", err)
		bcR.Switch.StopPeerForError(e.Src, err)
		return
	}

	bcR.Logger.Debug("REDACTED", "REDACTED", e.Src, "REDACTED", e.ChannelID, "REDACTED", e.Message)

	switch msg := e.Message.(type) {
	case *bcschema.BlockRequest:
		bcR.respondToPeer(msg, e.Src)
	case *bcschema.BlockResponse:
		bi, err := kinds.BlockFromProto(msg.Block)
		if err != nil {
			bcR.Logger.Error("REDACTED", "REDACTED", e.Src, "REDACTED", e.Message, "REDACTED", err)
			bcR.Switch.StopPeerForError(e.Src, err)
			return
		}
		var extCommit *kinds.ExtendedCommit
		if msg.ExtCommit != nil {
			var err error
			extCommit, err = kinds.ExtendedCommitFromProto(msg.ExtCommit)
			if err != nil {
				bcR.Logger.Error("REDACTED",
					"REDACTED", e.Src,
					"REDACTED", err)
				bcR.Switch.StopPeerForError(e.Src, err)
				return
			}
		}

		if err := bcR.pool.AddBlock(e.Src.ID(), bi, extCommit, msg.Block.Size()); err != nil {
			bcR.Logger.Error("REDACTED", "REDACTED", e.Src, "REDACTED", err)
		}
	case *bcschema.StatusRequest:
		//
		e.Src.TrySend(p2p.Envelope{
			ChannelID: BlocksyncChannel,
			Message: &bcschema.StatusResponse{
				Height: bcR.store.Height(),
				Base:   bcR.store.Base(),
			},
		})
	case *bcschema.StatusResponse:
		//
		bcR.pool.SetPeerRange(e.Src.ID(), msg.Base, msg.Height)
	case *bcschema.NoBlockResponse:
		bcR.Logger.Debug("REDACTED", "REDACTED", e.Src, "REDACTED", msg.Height)
		bcR.pool.RedoRequestFrom(msg.Height, e.Src.ID())
	default:
		bcR.Logger.Error(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
	}
}
