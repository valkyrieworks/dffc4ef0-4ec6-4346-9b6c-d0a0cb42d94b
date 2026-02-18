package agreement

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/go-kit/log/term"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	appconn "github.com/valkyrieworks/atci/requester"
	"github.com/valkyrieworks/atci/instance/dbstore"
	atci "github.com/valkyrieworks/atci/kinds"
	cfg "github.com/valkyrieworks/settings"
	statetypes "github.com/valkyrieworks/agreement/kinds"
	"github.com/valkyrieworks/intrinsic/validation"
	enginebytes "github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/log"
	ctsystem "github.com/valkyrieworks/utils/os"
	enginepubsub "github.com/valkyrieworks/utils/broadcast"
	ctsync "github.com/valkyrieworks/utils/alignment"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/authkey"
	ctschema "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	cttime "github.com/valkyrieworks/kinds/moment"
)

const (
	testSubscriber = "REDACTED"
)

//
//
type cleanupFunc func()

//
var (
	config                *cfg.Config //
	consensusReplayConfig *cfg.Config
	ensureTimeout         = time.Millisecond * 200
)

func ensureDir(dir string, mode os.FileMode) {
	if err := ctsystem.EnsureDir(dir, mode); err != nil {
		panic(err)
	}
}

func ResetConfig(name string) *cfg.Config {
	return validation.ResetTestRoot(name)
}

//
//

type validatorStub struct {
	Index  int32 //
	Height int64
	Round  int32
	kinds.PrivValidator
	VotingPower int64
	lastVote    *kinds.Vote
}

var testMinPower int64 = 10

func newValidatorStub(privValidator kinds.PrivValidator, valIndex int32) *validatorStub {
	return &validatorStub{
		Index:         valIndex,
		PrivValidator: privValidator,
		VotingPower:   testMinPower,
	}
}

func (vs *validatorStub) signVote(
	voteType ctschema.SignedMsgType,
	hash []byte,
	header kinds.PartSetHeader,
	voteExtension []byte,
	extEnabled bool,
) (*kinds.Vote, error) {
	pubKey, err := vs.GetPubKey()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	vote := &kinds.Vote{
		Type:             voteType,
		Height:           vs.Height,
		Round:            vs.Round,
		BlockID:          kinds.BlockID{Hash: hash, PartSetHeader: header},
		Timestamp:        cttime.Now(),
		ValidatorAddress: pubKey.Address(),
		ValidatorIndex:   vs.Index,
		Extension:        voteExtension,
	}
	v := vote.ToProto()
	if err = vs.SignVote(validation.DefaultTestChainID, v); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	//
	if signDataIsEqual(vs.lastVote, v) {
		v.Signature = vs.lastVote.Signature
		v.Timestamp = vs.lastVote.Timestamp
		v.ExtensionSignature = vs.lastVote.ExtensionSignature
	}

	vote.Signature = v.Signature
	vote.Timestamp = v.Timestamp
	vote.ExtensionSignature = v.ExtensionSignature

	if !extEnabled {
		vote.ExtensionSignature = nil
	}

	return vote, err
}

//
func signVote(vs *validatorStub, voteType ctschema.SignedMsgType, hash []byte, header kinds.PartSetHeader, extEnabled bool) *kinds.Vote {
	var ext []byte
	//
	if extEnabled {
		if voteType != ctschema.PrecommitType {
			panic(fmt.Errorf("REDACTED"))
		}
		if len(hash) != 0 || !header.IsZero() {
			ext = []byte("REDACTED")
		}
	}
	v, err := vs.signVote(voteType, hash, header, ext, extEnabled)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	vs.lastVote = v

	return v
}

func signVotes(
	voteType ctschema.SignedMsgType,
	hash []byte,
	header kinds.PartSetHeader,
	extEnabled bool,
	vss ...*validatorStub,
) []*kinds.Vote {
	votes := make([]*kinds.Vote, len(vss))
	for i, vs := range vss {
		votes[i] = signVote(vs, voteType, hash, header, extEnabled)
	}
	return votes
}

func incrementHeight(vss ...*validatorStub) {
	for _, vs := range vss {
		vs.Height++
	}
}

func incrementRound(vss ...*validatorStub) {
	for _, vs := range vss {
		vs.Round++
	}
}

type ValidatorStubsByPower []*validatorStub

func (vss ValidatorStubsByPower) Len() int {
	return len(vss)
}

func (vss ValidatorStubsByPower) Less(i, j int) bool {
	vssi, err := vss[i].GetPubKey()
	if err != nil {
		panic(err)
	}
	vssj, err := vss[j].GetPubKey()
	if err != nil {
		panic(err)
	}

	if vss[i].VotingPower == vss[j].VotingPower {
		return bytes.Compare(vssi.Address(), vssj.Address()) == -1
	}
	return vss[i].VotingPower > vss[j].VotingPower
}

func (vss ValidatorStubsByPower) Swap(i, j int) {
	it := vss[i]
	vss[i] = vss[j]
	vss[i].Index = int32(i)
	vss[j] = it
	vss[j].Index = int32(j)
}

//
//

func startTestRound(cs *State, height int64, round int32) {
	cs.enterNewRound(height, round)
	cs.startRoutines(0)
}

//
func decideProposal(
	ctx context.Context,
	t *testing.T,
	cs1 *State,
	vs *validatorStub,
	height int64,
	round int32,
) (*kinds.Proposal, *kinds.Block) {
	cs1.mtx.Lock()
	block, err := cs1.createProposalBlock(ctx)
	require.NoError(t, err)
	blockParts, err := block.MakePartSet(kinds.BlockPartSizeBytes)
	require.NoError(t, err)
	validRound := cs1.ValidRound
	chainID := cs1.state.ChainID
	cs1.mtx.Unlock()
	if block == nil {
		panic("REDACTED")
	}

	//
	polRound, propBlockID := validRound, kinds.BlockID{Hash: block.Hash(), PartSetHeader: blockParts.Header()}
	proposal := kinds.NewProposal(height, round, polRound, propBlockID)
	p := proposal.ToProto()
	if err := vs.SignProposal(chainID, p); err != nil {
		panic(err)
	}

	proposal.Signature = p.Signature

	return proposal, block
}

func addVotes(to *State, votes ...*kinds.Vote) {
	for _, vote := range votes {
		to.peerMsgQueue <- msgInfo{Msg: &VoteMessage{vote}}
	}
}

func signAddVotes(
	to *State,
	voteType ctschema.SignedMsgType,
	hash []byte,
	header kinds.PartSetHeader,
	extEnabled bool,
	vss ...*validatorStub,
) {
	votes := signVotes(voteType, hash, header, extEnabled, vss...)
	addVotes(to, votes...)
}

func validatePrevote(t *testing.T, cs *State, round int32, privVal *validatorStub, blockHash []byte) {
	prevotes := cs.Votes.Prevotes(round)
	pubKey, err := privVal.GetPubKey()
	require.NoError(t, err)
	address := pubKey.Address()
	var vote *kinds.Vote
	if vote = prevotes.GetByAddress(address); vote == nil {
		panic("REDACTED")
	}
	if blockHash == nil {
		if vote.BlockID.Hash != nil {
			panic(fmt.Sprintf("REDACTED", vote.BlockID.Hash))
		}
	} else {
		if !bytes.Equal(vote.BlockID.Hash, blockHash) {
			panic(fmt.Sprintf("REDACTED", blockHash, vote.BlockID.Hash))
		}
	}
}

func validateLastPrecommit(t *testing.T, cs *State, privVal *validatorStub, blockHash []byte) {
	votes := cs.LastCommit
	pv, err := privVal.GetPubKey()
	require.NoError(t, err)
	address := pv.Address()
	var vote *kinds.Vote
	if vote = votes.GetByAddress(address); vote == nil {
		panic("REDACTED")
	}
	if !bytes.Equal(vote.BlockID.Hash, blockHash) {
		panic(fmt.Sprintf("REDACTED", blockHash, vote.BlockID.Hash))
	}
}

func validatePrecommit(
	t *testing.T,
	cs *State,
	thisRound,
	lockRound int32,
	privVal *validatorStub,
	votedBlockHash,
	lockedBlockHash []byte,
) {
	precommits := cs.Votes.Precommits(thisRound)
	pv, err := privVal.GetPubKey()
	require.NoError(t, err)
	address := pv.Address()
	var vote *kinds.Vote
	if vote = precommits.GetByAddress(address); vote == nil {
		panic("REDACTED")
	}

	if votedBlockHash == nil {
		if vote.BlockID.Hash != nil {
			panic("REDACTED")
		}
	} else {
		if !bytes.Equal(vote.BlockID.Hash, votedBlockHash) {
			panic("REDACTED")
		}
	}

	rs := cs.GetRoundState()
	if lockedBlockHash == nil {
		if rs.LockedRound != lockRound || rs.LockedBlock != nil {
			panic(fmt.Sprintf(
				"REDACTED",
				lockRound,
				rs.LockedRound,
				rs.LockedBlock))
		}
	} else {
		if rs.LockedRound != lockRound || !bytes.Equal(rs.LockedBlock.Hash(), lockedBlockHash) {
			panic(fmt.Sprintf(
				"REDACTED",
				lockRound,
				rs.LockedRound,
				rs.LockedBlock.Hash(),
				lockedBlockHash))
		}
	}
}

func subscribeToVoter(cs *State, addr []byte) <-chan enginepubsub.Message {
	votesSub, err := cs.eventBus.SubscribeUnbuffered(context.Background(), testSubscriber, kinds.EventQueryVote)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", testSubscriber, kinds.EventQueryVote))
	}
	ch := make(chan enginepubsub.Message)
	go func() {
		for msg := range votesSub.Out() {
			vote := msg.Data().(kinds.EventDataVote)
			//
			if bytes.Equal(addr, vote.Vote.ValidatorAddress) {
				ch <- msg
			}
		}
	}()
	return ch
}

//
//

func newState(state sm.State, pv kinds.PrivValidator, app atci.Application) *State {
	config := validation.ResetTestRoot("REDACTED")
	return newStateWithConfig(config, state, pv, app)
}

func newStateWithConfig(
	thisConfig *cfg.Config,
	state sm.State,
	pv kinds.PrivValidator,
	app atci.Application,
) *State {
	blockDB := dbm.NewMemDB()
	return newStateWithConfigAndBlockStore(thisConfig, state, pv, app, blockDB)
}

func newStateWithConfigAndBlockStore(
	thisConfig *cfg.Config,
	state sm.State,
	pv kinds.PrivValidator,
	app atci.Application,
	blockDB dbm.DB,
) *State {
	//
	blockStore := depot.NewBlockStore(blockDB)

	//
	mtx := new(ctsync.Mutex)

	proxyAppConnCon := gateway.NewAppConnConsensus(appconn.NewLocalClient(mtx, app), gateway.NopMetrics())
	proxyAppConnMem := gateway.NewAppConnMempool(appconn.NewLocalClient(mtx, app), gateway.NopMetrics())
	//
	memplMetrics := txpool.NopMetrics()

	//
	mempool := txpool.NewCListMempool(config.Mempool,
		proxyAppConnMem,
		state.LastBlockHeight,
		txpool.WithMetrics(memplMetrics),
		txpool.WithPreCheck(sm.TxPreCheck(state)),
		txpool.WithPostCheck(sm.TxPostCheck(state)))

	if thisConfig.Consensus.WaitForTxs() {
		mempool.EnableTxsAvailable()
	}

	evpool := sm.EmptyEvidencePool{}

	//
	stateDB := blockDB
	stateStore := sm.NewStore(stateDB, sm.StoreOptions{
		DiscardABCIResponses: false,
	})

	if err := stateStore.Save(state); err != nil { //
		panic(err)
	}

	blockExec := sm.NewBlockExecutor(stateStore, log.TestingLogger(), proxyAppConnCon, mempool, evpool, blockStore)
	cs := NewState(thisConfig.Consensus, state, blockExec, blockStore, mempool, evpool)
	cs.SetLogger(log.TestingLogger().With("REDACTED", "REDACTED"))
	cs.SetPrivValidator(pv)

	eventBus := kinds.NewEventBus()
	eventBus.SetLogger(log.TestingLogger().With("REDACTED", "REDACTED"))
	err := eventBus.Start()
	if err != nil {
		panic(err)
	}
	cs.SetEventBus(eventBus)
	return cs
}

func loadPrivValidator(config *cfg.Config) *authkey.FilePV {
	privValidatorKeyFile := config.PrivValidatorKeyFile()
	ensureDir(filepath.Dir(privValidatorKeyFile), 0o700)
	privValidatorStateFile := config.PrivValidatorStateFile()
	privValidator := authkey.LoadOrGenFilePV(privValidatorKeyFile, privValidatorStateFile)
	privValidator.Reset()
	return privValidator
}

func randState(nValidators int) (*State, []*validatorStub) {
	return randStateWithApp(nValidators, dbstore.NewInMemoryApplication())
}

func randStateWithAppWithHeight(
	nValidators int,
	app atci.Application,
	height int64,
) (*State, []*validatorStub) {
	c := validation.ConsensusParams()
	c.ABCI.VoteExtensionsEnableHeight = height
	return randStateWithAppImpl(nValidators, app, c)
}

func randStateWithApp(nValidators int, app atci.Application) (*State, []*validatorStub) {
	c := validation.ConsensusParams()
	return randStateWithAppImpl(nValidators, app, c)
}

func randStateWithAppImpl(
	nValidators int,
	app atci.Application,
	consensusParams *kinds.ConsensusParams,
) (*State, []*validatorStub) {
	//
	state, privVals := randGenesisState(nValidators, false, 10, consensusParams)

	vss := make([]*validatorStub, nValidators)

	cs := newState(state, privVals[0], app)

	for i := 0; i < nValidators; i++ {
		vss[i] = newValidatorStub(privVals[i], int32(i))
	}
	//
	incrementHeight(vss[1:]...)

	return cs, vss
}

//

func ensureNoNewEvent(ch <-chan enginepubsub.Message, timeout time.Duration,
	errorMessage string,
) {
	select {
	case <-time.After(timeout):
		break
	case <-ch:
		panic(errorMessage)
	}
}

func ensureNoNewEventOnChannel(ch <-chan enginepubsub.Message) {
	ensureNoNewEvent(
		ch,
		ensureTimeout*8/10, //
		"REDACTED")
}

func ensureNoNewRoundStep(stepCh <-chan enginepubsub.Message) {
	ensureNoNewEvent(
		stepCh,
		ensureTimeout,
		"REDACTED")
}

func ensureNoNewUnlock(unlockCh <-chan enginepubsub.Message) {
	ensureNoNewEvent(
		unlockCh,
		ensureTimeout,
		"REDACTED")
}

func ensureNoNewTimeout(stepCh <-chan enginepubsub.Message, timeout int64) {
	timeoutDuration := time.Duration(timeout*10) * time.Nanosecond
	ensureNoNewEvent(
		stepCh,
		timeoutDuration,
		"REDACTED")
}

func ensureNewEvent(ch <-chan enginepubsub.Message, height int64, round int32, timeout time.Duration, errorMessage string) {
	select {
	case <-time.After(timeout):
		panic(errorMessage)
	case msg := <-ch:
		roundStateEvent, ok := msg.Data().(kinds.EventDataRoundState)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if roundStateEvent.Height != height {
			panic(fmt.Sprintf("REDACTED", height, roundStateEvent.Height))
		}
		if roundStateEvent.Round != round {
			panic(fmt.Sprintf("REDACTED", round, roundStateEvent.Round))
		}
		//
	}
}

func ensureNewRound(roundCh <-chan enginepubsub.Message, height int64, round int32) {
	select {
	case <-time.After(ensureTimeout):
		panic("REDACTED")
	case msg := <-roundCh:
		newRoundEvent, ok := msg.Data().(kinds.EventDataNewRound)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if newRoundEvent.Height != height {
			panic(fmt.Sprintf("REDACTED", height, newRoundEvent.Height))
		}
		if newRoundEvent.Round != round {
			panic(fmt.Sprintf("REDACTED", round, newRoundEvent.Round))
		}
	}
}

func ensureNewTimeout(timeoutCh <-chan enginepubsub.Message, height int64, round int32, timeout int64) {
	timeoutDuration := time.Duration(timeout*10) * time.Nanosecond
	ensureNewEvent(timeoutCh, height, round, timeoutDuration,
		"REDACTED")
}

func ensureNewProposal(proposalCh <-chan enginepubsub.Message, height int64, round int32) {
	select {
	case <-time.After(ensureTimeout):
		panic("REDACTED")
	case msg := <-proposalCh:
		proposalEvent, ok := msg.Data().(kinds.EventDataCompleteProposal)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if proposalEvent.Height != height {
			panic(fmt.Sprintf("REDACTED", height, proposalEvent.Height))
		}
		if proposalEvent.Round != round {
			panic(fmt.Sprintf("REDACTED", round, proposalEvent.Round))
		}
	}
}

func ensureNewValidBlock(validBlockCh <-chan enginepubsub.Message, height int64, round int32) {
	ensureNewEvent(validBlockCh, height, round, ensureTimeout,
		"REDACTED")
}

func ensureNewBlock(blockCh <-chan enginepubsub.Message, height int64) {
	select {
	case <-time.After(ensureTimeout):
		panic("REDACTED")
	case msg := <-blockCh:
		blockEvent, ok := msg.Data().(kinds.EventDataNewBlock)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if blockEvent.Block.Height != height {
			panic(fmt.Sprintf("REDACTED", height, blockEvent.Block.Height))
		}
	}
}

func ensureNewBlockHeader(blockCh <-chan enginepubsub.Message, height int64, blockHash enginebytes.HexBytes) {
	select {
	case <-time.After(ensureTimeout):
		panic("REDACTED")
	case msg := <-blockCh:
		blockHeaderEvent, ok := msg.Data().(kinds.EventDataNewBlockHeader)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if blockHeaderEvent.Header.Height != height {
			panic(fmt.Sprintf("REDACTED", height, blockHeaderEvent.Header.Height))
		}
		if !bytes.Equal(blockHeaderEvent.Header.Hash(), blockHash) {
			panic(fmt.Sprintf("REDACTED", blockHash, blockHeaderEvent.Header.Hash()))
		}
	}
}

func ensureNewUnlock(unlockCh <-chan enginepubsub.Message, height int64, round int32) {
	ensureNewEvent(unlockCh, height, round, ensureTimeout,
		"REDACTED")
}

func ensureProposal(proposalCh <-chan enginepubsub.Message, height int64, round int32, propID kinds.BlockID) {
	select {
	case <-time.After(ensureTimeout):
		panic("REDACTED")
	case msg := <-proposalCh:
		proposalEvent, ok := msg.Data().(kinds.EventDataCompleteProposal)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if proposalEvent.Height != height {
			panic(fmt.Sprintf("REDACTED", height, proposalEvent.Height))
		}
		if proposalEvent.Round != round {
			panic(fmt.Sprintf("REDACTED", round, proposalEvent.Round))
		}
		if !proposalEvent.BlockID.Equals(propID) {
			panic(fmt.Sprintf("REDACTED", proposalEvent.BlockID, propID))
		}
	}
}

func ensurePrecommit(voteCh <-chan enginepubsub.Message, height int64, round int32) {
	ensureVote(voteCh, height, round, ctschema.PrecommitType)
}

func ensurePrevote(voteCh <-chan enginepubsub.Message, height int64, round int32) {
	ensureVote(voteCh, height, round, ctschema.PrevoteType)
}

func ensureVote(voteCh <-chan enginepubsub.Message, height int64, round int32,
	voteType ctschema.SignedMsgType,
) {
	select {
	case <-time.After(ensureTimeout):
		panic("REDACTED")
	case msg := <-voteCh:
		voteEvent, ok := msg.Data().(kinds.EventDataVote)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		vote := voteEvent.Vote
		if vote.Height != height {
			panic(fmt.Sprintf("REDACTED", height, vote.Height))
		}
		if vote.Round != round {
			panic(fmt.Sprintf("REDACTED", round, vote.Round))
		}
		if vote.Type != voteType {
			panic(fmt.Sprintf("REDACTED", voteType, vote.Type))
		}
	}
}

func ensurePrevoteMatch(t *testing.T, voteCh <-chan enginepubsub.Message, height int64, round int32, hash []byte) {
	t.Helper()
	ensureVoteMatch(t, voteCh, height, round, hash, ctschema.PrevoteType)
}

func ensurePrecommitMatch(t *testing.T, voteCh <-chan enginepubsub.Message, height int64, round int32, hash []byte) {
	t.Helper()
	ensureVoteMatch(t, voteCh, height, round, hash, ctschema.PrecommitType)
}

func ensureVoteMatch(t *testing.T, voteCh <-chan enginepubsub.Message, height int64, round int32, hash []byte, voteType ctschema.SignedMsgType) {
	t.Helper()
	select {
	case <-time.After(ensureTimeout):
		t.Fatal("REDACTED")
	case msg := <-voteCh:
		voteEvent, ok := msg.Data().(kinds.EventDataVote)
		require.True(t, ok, "REDACTED",
			msg.Data())

		vote := voteEvent.Vote
		assert.Equal(t, height, vote.Height, "REDACTED", height, vote.Height)
		assert.Equal(t, round, vote.Round, "REDACTED", round, vote.Round)
		assert.Equal(t, voteType, vote.Type, "REDACTED", voteType, vote.Type)
		if hash == nil {
			require.Nil(t, vote.BlockID.Hash, "REDACTED", vote.BlockID.Hash)
		} else {
			require.True(t, bytes.Equal(vote.BlockID.Hash, hash), "REDACTED", hash, vote.BlockID.Hash)
		}
	}
}

func ensurePrecommitTimeout(ch <-chan enginepubsub.Message) {
	select {
	case <-time.After(ensureTimeout):
		panic("REDACTED")
	case <-ch:
	}
}

func ensureNewEventOnChannel(ch <-chan enginepubsub.Message) {
	select {
	case <-time.After(ensureTimeout * 12 / 10): //
		panic("REDACTED")
	case <-ch:
	}
}

//
//

//
//
func consensusLogger() log.Logger {
	return log.TestingLoggerWithColorFn(func(keyvals ...any) term.FgBgColor {
		for i := 0; i < len(keyvals)-1; i += 2 {
			if keyvals[i] == "REDACTED" {
				return term.FgBgColor{Fg: term.Color(uint8(keyvals[i+1].(int) + 1))}
			}
		}
		return term.FgBgColor{}
	}).With("REDACTED", "REDACTED")
}

func randConsensusNet(t *testing.T, nValidators int, testName string, tickerFunc func() TimeoutTicker,
	appFunc func() atci.Application, configOpts ...func(*cfg.Config),
) ([]*State, cleanupFunc) {
	t.Helper()
	genDoc, privVals := randGenesisDoc(nValidators, false, 30, nil)
	css := make([]*State, nValidators)
	logger := consensusLogger()
	configRootDirs := make([]string, 0, nValidators)
	for i := 0; i < nValidators; i++ {
		stateDB := dbm.NewMemDB() //
		stateStore := sm.NewStore(stateDB, sm.StoreOptions{
			DiscardABCIResponses: false,
		})
		state, _ := stateStore.LoadFromDBOrGenesisDoc(genDoc)
		thisConfig := ResetConfig(fmt.Sprintf("REDACTED", testName, i))
		configRootDirs = append(configRootDirs, thisConfig.RootDir)
		for _, opt := range configOpts {
			opt(thisConfig)
		}
		ensureDir(filepath.Dir(thisConfig.Consensus.WalFile()), 0o700) //
		app := appFunc()
		vals := kinds.TM2PB.ValidatorUpdates(state.Validators)
		_, err := app.InitChain(context.Background(), &atci.RequestInitChain{Validators: vals})
		require.NoError(t, err)

		css[i] = newStateWithConfigAndBlockStore(thisConfig, state, privVals[i], app, stateDB)
		css[i].SetTimeoutTicker(tickerFunc())
		css[i].SetLogger(logger.With("REDACTED", i, "REDACTED", "REDACTED"))
	}
	return css, func() {
		for _, dir := range configRootDirs {
			os.RemoveAll(dir)
		}
	}
}

//
func randConsensusNetWithPeers(
	t *testing.T,
	nValidators,
	nPeers int,
	testName string,
	tickerFunc func() TimeoutTicker,
	appFunc func(string) atci.Application,
) ([]*State, *kinds.GenesisDoc, *cfg.Config, cleanupFunc) {
	c := validation.ConsensusParams()
	genDoc, privVals := randGenesisDoc(nValidators, false, testMinPower, c)
	css := make([]*State, nPeers)
	logger := consensusLogger()
	var peer0Config *cfg.Config
	configRootDirs := make([]string, 0, nPeers)
	for i := 0; i < nPeers; i++ {
		stateDB := dbm.NewMemDB() //
		stateStore := sm.NewStore(stateDB, sm.StoreOptions{
			DiscardABCIResponses: false,
		})
		t.Cleanup(func() { _ = stateStore.Close() })
		state, _ := stateStore.LoadFromDBOrGenesisDoc(genDoc)
		thisConfig := ResetConfig(fmt.Sprintf("REDACTED", testName, i))
		configRootDirs = append(configRootDirs, thisConfig.RootDir)
		ensureDir(filepath.Dir(thisConfig.Consensus.WalFile()), 0o700) //
		if i == 0 {
			peer0Config = thisConfig
		}
		var privVal kinds.PrivValidator
		if i < nValidators {
			privVal = privVals[i]
		} else {
			tempKeyFile, err := os.CreateTemp("REDACTED", "REDACTED")
			if err != nil {
				panic(err)
			}
			tempStateFile, err := os.CreateTemp("REDACTED", "REDACTED")
			if err != nil {
				panic(err)
			}

			privVal = authkey.GenFilePV(tempKeyFile.Name(), tempStateFile.Name())
		}

		app := appFunc(path.Join(config.DBDir(), fmt.Sprintf("REDACTED", testName, i)))
		vals := kinds.TM2PB.ValidatorUpdates(state.Validators)
		if _, ok := app.(*dbstore.Application); ok {
			//
			state.Version.Consensus.App = dbstore.AppVersion
		}
		_, err := app.InitChain(context.Background(), &atci.RequestInitChain{Validators: vals})
		require.NoError(t, err)

		css[i] = newStateWithConfig(thisConfig, state, privVal, app)
		css[i].SetTimeoutTicker(tickerFunc())
		css[i].SetLogger(logger.With("REDACTED", i, "REDACTED", "REDACTED"))
	}
	return css, genDoc, peer0Config, func() {
		for _, dir := range configRootDirs {
			os.RemoveAll(dir)
		}
	}
}

func getSwitchIndex(switches []*p2p.Switch, peer p2p.Peer) int {
	for i, s := range switches {
		if peer.NodeInfo().ID() == s.NodeInfo().ID() {
			return i
		}
	}
	panic("REDACTED")
}

//
//

func randGenesisDoc(numValidators int,
	randPower bool,
	minPower int64,
	consensusParams *kinds.ConsensusParams,
) (*kinds.GenesisDoc, []kinds.PrivValidator) {
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

	return &kinds.GenesisDoc{
		GenesisTime:     cttime.Now(),
		InitialHeight:   1,
		ChainID:         validation.DefaultTestChainID,
		Validators:      validators,
		ConsensusParams: consensusParams,
	}, privValidators
}

func randGenesisState(
	numValidators int,
	randPower bool,
	minPower int64,
	consensusParams *kinds.ConsensusParams,
) (sm.State, []kinds.PrivValidator) {
	genDoc, privValidators := randGenesisDoc(numValidators, randPower, minPower, consensusParams)
	s0, _ := sm.MakeGenesisState(genDoc)
	return s0, privValidators
}

//
//

func newMockTickerFunc(onlyOnce bool) func() TimeoutTicker {
	return func() TimeoutTicker {
		return &mockTicker{
			c:        make(chan timeoutInfo, 10),
			onlyOnce: onlyOnce,
		}
	}
}

//
//
type mockTicker struct {
	c chan timeoutInfo

	mtx      sync.Mutex
	onlyOnce bool
	fired    bool
}

func (m *mockTicker) Start() error {
	return nil
}

func (m *mockTicker) Stop() error {
	return nil
}

func (m *mockTicker) ScheduleTimeout(ti timeoutInfo) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if m.onlyOnce && m.fired {
		return
	}
	if ti.Step == statetypes.RoundStepNewHeight {
		m.c <- ti
		m.fired = true
	}
}

func (m *mockTicker) Chan() <-chan timeoutInfo {
	return m.c
}

func (*mockTicker) SetLogger(log.Logger) {}

func newPersistentKVStore() atci.Application {
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}
	return dbstore.NewPersistentApplication(dir)
}

func newKVStore() atci.Application {
	return dbstore.NewInMemoryApplication()
}

func newPersistentKVStoreWithPath(dbDir string) atci.Application {
	return dbstore.NewPersistentApplication(dbDir)
}

func signDataIsEqual(v1 *kinds.Vote, v2 *ctschema.Vote) bool {
	if v1 == nil || v2 == nil {
		return false
	}

	return v1.Type == v2.Type &&
		bytes.Equal(v1.BlockID.Hash, v2.BlockID.GetHash()) &&
		v1.Height == v2.GetHeight() &&
		v1.Round == v2.Round &&
		bytes.Equal(v1.ValidatorAddress.Bytes(), v2.GetValidatorAddress()) &&
		v1.ValidatorIndex == v2.GetValidatorIndex() &&
		bytes.Equal(v1.Extension, v2.Extension)
}
