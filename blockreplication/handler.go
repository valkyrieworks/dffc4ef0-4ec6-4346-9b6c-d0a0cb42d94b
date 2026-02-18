package blockreplication

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valkyrieworks/security"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/lnet"
	"github.com/valkyrieworks/p2p"
	bcschema "github.com/valkyrieworks/schema/consensuscore/blockreplication"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
)

//
const BlocksyncChannel = byte(0x40)

const (
	defaultIntervalStatusUpdate  = 10 * time.Second
	followerIntervalStatusUpdate = 1 * time.Second

	defaultIntervalSwitchToConsensus = 1 * time.Second

	//
	intervalTrySync = 10 * time.Millisecond
)

type consensusReactor interface {
	//
	//
	SwitchToConsensus(state sm.State, skipWAL bool)
}

type mempoolReactor interface {
	//
	EnableInOutTxs()
}

type peerError struct {
	err    error
	peerID p2p.ID
}

func (e peerError) Error() string {
	return fmt.Sprintf("REDACTED", e.peerID, e.err.Error())
}

//
type Reactor struct {
	p2p.BaseReactor

	//
	initialState sm.State

	//
	//
	enabled *atomic.Bool

	//
	followerMode bool

	blockExec     *sm.BlockExecutor
	store         sm.BlockStore
	pool          *BlockPool
	localAddr     security.Address
	poolRoutineWg sync.WaitGroup

	requestsCh <-chan BlockRequest
	errorsCh   <-chan peerError

	//
	intervalSwitchToConsensus time.Duration

	//
	intervalStatusUpdate time.Duration

	metrics *Metrics
}

//
func NewReactor(
	enabled bool,
	followerMode bool,
	state sm.State,
	blockExec *sm.BlockExecutor,
	store *depot.BlockStore,
	localAddr security.Address,
	offlineStateSyncHeight int64,
	metrics *Metrics,
) *Reactor {
	storeHeight := depot.Height()
	if storeHeight == 0 {
		//
		//
		//
		//
		//
		//
		storeHeight = offlineStateSyncHeight
	}

	if state.LastBlockHeight != storeHeight {
		panic(fmt.Sprintf(
			"REDACTED",
			state.LastBlockHeight,
			storeHeight,
		))
	}

	//
	//
	requestsCh := make(chan BlockRequest)

	const capacity = 1000                      //
	errorsCh := make(chan peerError, capacity) //

	startHeight := storeHeight + 1
	if startHeight == 1 {
		startHeight = state.InitialHeight
	}
	pool := NewBlockPool(startHeight, requestsCh, errorsCh)

	enabledFlag := &atomic.Bool{}
	enabledFlag.Store(enabled)

	intervalStatusUpdate := defaultIntervalStatusUpdate
	if followerMode {
		intervalStatusUpdate = followerIntervalStatusUpdate
	}

	r := &Reactor{
		initialState:              state,
		blockExec:                 blockExec,
		store:                     store,
		pool:                      pool,
		enabled:                   enabledFlag,
		followerMode:              followerMode,
		localAddr:                 localAddr,
		requestsCh:                requestsCh,
		errorsCh:                  errorsCh,
		metrics:                   metrics,
		intervalSwitchToConsensus: defaultIntervalSwitchToConsensus,
		intervalStatusUpdate:      intervalStatusUpdate,
	}

	r.BaseReactor = *p2p.NewBaseReactor("REDACTED", r)

	return r
}

//
func (r *Reactor) SetLogger(l log.Logger) {
	r.Logger = l
	r.pool.Logger = l
}

//
func (r *Reactor) OnStart() error {
	//
	if !r.enabled.Load() {
		return nil
	}

	return r.runPool(false)
}

func (r *Reactor) runPool(stateSynced bool) error {
	if err := r.pool.Start(); err != nil {
		return err
	}

	r.poolRoutineWg.Add(1)
	go func() {
		defer r.poolRoutineWg.Done()
		r.poolRoutine(stateSynced)
	}()

	return nil
}

//
func (r *Reactor) Enable(state sm.State) error {
	if !r.enabled.CompareAndSwap(false, true) {
		return ErrAlreadyEnabled
	}

	r.initialState = state
	r.pool.height = state.LastBlockHeight + 1

	return r.runPool(true)
}

//
func (r *Reactor) OnStop() {
	if !r.enabled.Load() {
		return
	}

	if err := r.pool.Stop(); err != nil {
		r.Logger.Error("REDACTED", "REDACTED", err)
	}

	r.poolRoutineWg.Wait()
}

//
func (r *Reactor) GetChannels() []*p2p.ChannelDescriptor {
	return []*p2p.ChannelDescriptor{
		{
			ID:                  BlocksyncChannel,
			Priority:            5,
			SendQueueCapacity:   1000,
			RecvBufferCapacity:  50 * 4096,
			RecvMessageCapacity: MaxMsgSize,
			MessageType:         &bcschema.Message{},
		},
	}
}

//
func (r *Reactor) AddPeer(peer p2p.Peer) {
	peer.Send(p2p.Envelope{
		ChannelID: BlocksyncChannel,
		Message: &bcschema.StatusResponse{
			Base:   r.store.Base(),
			Height: r.store.Height(),
		},
	})
	//

	//
	//
}

//
func (r *Reactor) RemovePeer(peer p2p.Peer, _ any) {
	r.pool.RemovePeer(peer.ID())
}

//
//
func (r *Reactor) respondToPeer(msg *bcschema.BlockRequest, src p2p.Peer) {
	block := r.store.LoadBlock(msg.Height)
	if block == nil {
		r.Logger.Info("REDACTED", "REDACTED", src, "REDACTED", msg.Height)
		src.TrySend(p2p.Envelope{
			ChannelID: BlocksyncChannel,
			Message:   &bcschema.NoBlockResponse{Height: msg.Height},
		})

		return
	}

	state, err := r.blockExec.Store().Load()
	if err != nil {
		r.Logger.Error("REDACTED", "REDACTED", err)
		return
	}

	var extCommit *kinds.ExtendedCommit
	if state.ConsensusParams.ABCI.VoteExtensionsEnabled(msg.Height) {
		extCommit = r.store.LoadBlockExtendedCommit(msg.Height)
		if extCommit == nil {
			r.Logger.Error("REDACTED", "REDACTED", block)
			return
		}
	}

	bl, err := block.ToProto()
	if err != nil {
		r.Logger.Error("REDACTED", "REDACTED", err)
		return
	}

	src.TrySend(p2p.Envelope{
		ChannelID: BlocksyncChannel,
		Message: &bcschema.BlockResponse{
			Block:     bl,
			ExtCommit: extCommit.ToProto(),
		},
	})
}

func (r *Reactor) handlePeerResponse(msg *bcschema.BlockResponse, src p2p.Peer) {
	bi, err := kinds.BlockFromProto(msg.Block)
	if err != nil {
		r.Logger.Error("REDACTED", "REDACTED", src, "REDACTED", msg, "REDACTED", err)
		r.stopPeerForError(src, err)
		return
	}

	var extCommit *kinds.ExtendedCommit
	if msg.ExtCommit != nil {
		extCommit, err = kinds.ExtendedCommitFromProto(msg.ExtCommit)
		if err != nil {
			r.Logger.Error("REDACTED", "REDACTED", src, "REDACTED", err)
			r.stopPeerForError(src, err)
			return
		}
	}

	if err := r.pool.AddBlock(src.ID(), bi, extCommit, msg.Block.Size()); err != nil {
		r.Logger.Error("REDACTED", "REDACTED", src, "REDACTED", err)
	}
}

//
func (r *Reactor) Receive(e p2p.Envelope) {
	if err := ValidateMsg(e.Message); err != nil {
		r.Logger.Error("REDACTED", "REDACTED", e.Src, "REDACTED", e.Message, "REDACTED", err)
		r.stopPeerForError(e.Src, err)
		return
	}

	r.Logger.Debug("REDACTED", "REDACTED", e.Src, "REDACTED", e.ChannelID, "REDACTED", e.Message)

	switch msg := e.Message.(type) {
	case *bcschema.BlockRequest:
		//
		r.respondToPeer(msg, e.Src)
	case *bcschema.BlockResponse:
		//
		go r.handlePeerResponse(msg, e.Src)
	case *bcschema.StatusRequest:
		//
		e.Src.TrySend(p2p.Envelope{
			ChannelID: BlocksyncChannel,
			Message: &bcschema.StatusResponse{
				Height: r.store.Height(),
				Base:   r.store.Base(),
			},
		})
	case *bcschema.StatusResponse:
		//
		r.pool.SetPeerRange(e.Src.ID(), msg.Base, msg.Height)
	case *bcschema.NoBlockResponse:
		r.Logger.Debug("REDACTED", "REDACTED", e.Src, "REDACTED", msg.Height)
		r.pool.RedoRequestFrom(msg.Height, e.Src.ID())
	default:
		r.Logger.Error(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
	}
}

func (r *Reactor) localNodeBlocksTheChain(state sm.State) bool {
	_, val := state.Validators.GetByAddress(r.localAddr)
	if val == nil {
		return false
	}
	total := state.Validators.TotalVotingPower()
	return val.VotingPower >= total/3
}

//
//
func (r *Reactor) poolRoutine(stateSynced bool) {
	r.metrics.Syncing.Set(1)
	defer r.metrics.Syncing.Set(0)

	trySyncTicker := time.NewTicker(intervalTrySync)
	defer trySyncTicker.Stop()

	statusUpdateTicker := time.NewTicker(r.intervalStatusUpdate)
	defer statusUpdateTicker.Stop()

	switchToConsensusTicker := time.NewTicker(r.intervalSwitchToConsensus)
	defer switchToConsensusTicker.Stop()

	go r.poolEventsRoutine(statusUpdateTicker)

	var (
		chainID                    = r.initialState.ChainID
		state                      = r.initialState
		initialCommitHasExtensions = state.LastBlockHeight > 0 &&
			r.store.LoadBlockExtendedCommit(state.LastBlockHeight) != nil

		didProcessCh = make(chan struct{}, 1)

		//
		blocksSynced = 0
		lastHundred  = time.Now()
		lastRate     = 0.0
	)

FOR_LOOP:
	for {
		select {
		case <-r.Quit():
			break FOR_LOOP
		case <-r.pool.Quit():
			break FOR_LOOP
		case <-switchToConsensusTicker.C:
			height, numPending, lenRequesters := r.pool.GetStatus()
			outbound, inbound, _ := r.Switch.NumPeers()

			r.Logger.Debug(
				"REDACTED",
				"REDACTED", numPending,
				"REDACTED", lenRequesters,
				"REDACTED", outbound,
				"REDACTED", inbound,
				"REDACTED", state.LastBlockHeight,
			)

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
			//
			//
			//
			//
			//
			//
			missingExtension := true
			if state.LastBlockHeight == 0 ||
				!state.ConsensusParams.ABCI.VoteExtensionsEnabled(state.LastBlockHeight) ||
				blocksSynced > 0 ||
				initialCommitHasExtensions {
				missingExtension = false
			}

			//
			if missingExtension {
				r.Logger.Info(
					"REDACTED",
					"REDACTED", height,
					"REDACTED", state.LastBlockHeight,
					"REDACTED", state.InitialHeight,
					"REDACTED", r.pool.MaxPeerHeight(),
				)
				continue FOR_LOOP
			}

			//
			if !r.pool.IsCaughtUp() && !r.localNodeBlocksTheChain(state) {
				continue FOR_LOOP
			}

			if r.followerMode {
				r.Logger.Debug("REDACTED", "REDACTED", state.LastBlockHeight)
				continue FOR_LOOP
			}

			r.Logger.Info("REDACTED", "REDACTED", height)
			if err := r.pool.Stop(); err != nil {
				r.Logger.Error("REDACTED", "REDACTED", err)
			}

			memR, exists := r.Switch.Reactor("REDACTED")
			if exists {
				if memR, ok := memR.(mempoolReactor); ok {
					memR.EnableInOutTxs()
				}
			}

			conR, exists := r.Switch.Reactor("REDACTED")
			if exists {
				if conR, ok := conR.(consensusReactor); ok {
					conR.SwitchToConsensus(state, blocksSynced > 0 || stateSynced)
				}
			}

			break FOR_LOOP
		case <-trySyncTicker.C:
			select {
			case didProcessCh <- struct{}{}:
			default:
			}
		case <-didProcessCh:
			//
			//
			//
			//
			//
			//
			//

			//
			first, second, extCommit := r.pool.PeekTwoBlocks()
			if first == nil || second == nil {
				//
				//
				continue FOR_LOOP
			}
			//
			if state.LastBlockHeight > 0 && state.LastBlockHeight+1 != first.Height {
				//
				panic(fmt.Errorf("REDACTED", state.LastBlockHeight+1, first.Height))
			}
			if first.Height+1 != second.Height {
				//
				panic(fmt.Errorf("REDACTED", state.LastBlockHeight, first.Height))
			}

			//
			//
			//
			//
			if !r.IsRunning() || !r.pool.IsRunning() {
				break FOR_LOOP
			}
			//
			didProcessCh <- struct{}{}

			firstParts, err := first.MakePartSet(kinds.BlockPartSizeBytes)
			if err != nil {
				r.Logger.Error("REDACTED", "REDACTED", first.Height, "REDACTED", err.Error())
				break FOR_LOOP
			}

			firstPartSetHeader := firstParts.Header()
			firstID := kinds.BlockID{Hash: first.Hash(), PartSetHeader: firstPartSetHeader}

			//
			//
			//
			//
			//
			err = state.Validators.VerifyCommitLight(chainID, firstID, first.Height, second.LastCommit)

			if err == nil {
				//
				err = r.blockExec.ValidateBlock(state, first)
			}
			presentExtCommit := extCommit != nil
			extensionsEnabled := state.ConsensusParams.ABCI.VoteExtensionsEnabled(first.Height)
			if presentExtCommit != extensionsEnabled {
				err = fmt.Errorf("REDACTED"+
					"REDACTED",
					first.Height, presentExtCommit, extensionsEnabled,
				)
			}
			if err == nil && extensionsEnabled {
				//
				err = extCommit.EnsureExtensions(true)
			}
			if err != nil {
				r.Logger.Error("REDACTED", "REDACTED", err)
				peerID := r.pool.RemovePeerAndRedoAllPeerRequests(first.Height)
				peer := r.Switch.Peers().Get(peerID)
				if peer != nil {
					//
					//
					r.stopPeerForError(peer, ErrReactorValidation{Err: err})
				}
				peerID2 := r.pool.RemovePeerAndRedoAllPeerRequests(second.Height)
				peer2 := r.Switch.Peers().Get(peerID2)
				if peer2 != nil && peer2 != peer {
					//
					//
					r.stopPeerForError(peer2, ErrReactorValidation{Err: err})
				}
				continue FOR_LOOP
			}

			r.pool.PopRequest()

			//
			if extensionsEnabled {
				r.store.SaveBlockWithExtendedCommit(first, firstParts, extCommit)
			} else {
				//
				//
				//
				//
				r.store.SaveBlock(first, firstParts, second.LastCommit)
			}

			//
			//
			state, err = r.blockExec.ApplyVerifiedBlock(state, firstID, first)
			if err != nil {
				//
				panic(fmt.Sprintf("REDACTED", first.Height, first.Hash(), err))
			}

			r.metrics.recordBlockMetrics(first)
			blocksSynced++

			if blocksSynced%100 == 0 {
				lastRate = 0.9*lastRate + 0.1*(100/time.Since(lastHundred).Seconds())
				lastHundred = time.Now()
				r.Logger.Info(
					"REDACTED",
					"REDACTED", r.pool.height,
					"REDACTED", r.pool.MaxPeerHeight(),
					"REDACTED", lastRate,
				)
			}

			continue FOR_LOOP
		}
	}
}

func (r *Reactor) poolEventsRoutine(statusUpdateTicker *time.Ticker) {
	for {
		select {
		case <-r.Quit():
			return
		case <-r.pool.Quit():
			return
		case request := <-r.requestsCh:
			//
			peer := r.Switch.Peers().Get(request.PeerID)
			if peer == nil {
				continue
			}

			queued := peer.TrySend(p2p.Envelope{
				ChannelID: BlocksyncChannel,
				Message:   &bcschema.BlockRequest{Height: request.Height},
			})

			if !queued {
				r.Logger.Debug("REDACTED", "REDACTED", peer.ID(), "REDACTED", request.Height)
			}
		case err := <-r.errorsCh:
			//
			if peer := r.Switch.Peers().Get(err.peerID); peer != nil {
				r.stopPeerForError(peer, err.err)
			}
		case <-statusUpdateTicker.C:
			//
			r.Switch.BroadcastAsync(p2p.Envelope{
				ChannelID: BlocksyncChannel,
				Message:   &bcschema.StatusRequest{},
			})
		}
	}
}

func (r *Reactor) stopPeerForError(peer p2p.Peer, err error) {
	if r.followerMode && shouldBeReconnected(err) {
		err = &lnet.ErrorTransient{Err: err}
	}

	r.Switch.StopPeerForError(peer, err)
}

//
//
//
//
func shouldBeReconnected(err error) bool {
	//
	return errors.Is(err, ErrPeerTimeout)
}
