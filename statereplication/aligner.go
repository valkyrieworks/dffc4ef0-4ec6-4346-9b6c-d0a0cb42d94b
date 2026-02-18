package statereplication

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	atci "github.com/valkyrieworks/atci/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	ctsync "github.com/valkyrieworks/utils/alignment"
	"github.com/valkyrieworks/minimal"
	"github.com/valkyrieworks/p2p"
	sschema "github.com/valkyrieworks/schema/consensuscore/statereplication"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	chunkTimeout = 2 * time.Minute

	//
	//
	minimumDiscoveryTime = 5 * time.Second
)

var (
	//
	errAbort = errors.New("REDACTED")
	//
	errRetrySnapshot = errors.New("REDACTED")
	//
	errRejectSnapshot = errors.New("REDACTED")
	//
	errRejectFormat = errors.New("REDACTED")
	//
	errRejectSender = errors.New("REDACTED")
	//
	errVerifyFailed = errors.New("REDACTED")
	//
	errTimeout = errors.New("REDACTED")
	//
	errNoSnapshots = errors.New("REDACTED")
)

//
//
//
type syncer struct {
	logger        log.Logger
	stateProvider StateProvider
	conn          gateway.AppConnSnapshot
	connQuery     gateway.AppConnQuery
	snapshots     *snapshotPool
	tempDir       string
	chunkFetchers int32
	retryTimeout  time.Duration

	mtx    ctsync.RWMutex
	chunks *chunkQueue
}

//
func newSyncer(
	cfg settings.StateSyncConfig,
	logger log.Logger,
	conn gateway.AppConnSnapshot,
	connQuery gateway.AppConnQuery,
	stateProvider StateProvider,
	tempDir string,
) *syncer {
	return &syncer{
		logger:        logger,
		stateProvider: stateProvider,
		conn:          conn,
		connQuery:     connQuery,
		snapshots:     newSnapshotPool(),
		tempDir:       tempDir,
		chunkFetchers: cfg.ChunkFetchers,
		retryTimeout:  cfg.ChunkRequestTimeout,
	}
}

//
//
func (s *syncer) AddChunk(chunk *chunk) (bool, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if s.chunks == nil {
		return false, errors.New("REDACTED")
	}
	added, err := s.chunks.Add(chunk)
	if err != nil {
		return false, err
	}
	if added {
		s.logger.Debug("REDACTED", "REDACTED", chunk.Height, "REDACTED", chunk.Format,
			"REDACTED", chunk.Index)
	} else {
		s.logger.Debug("REDACTED", "REDACTED", chunk.Height, "REDACTED", chunk.Format,
			"REDACTED", chunk.Index)
	}
	return added, nil
}

//
//
func (s *syncer) AddSnapshot(peer p2p.Peer, snapshot *snapshot) (bool, error) {
	added, err := s.snapshots.Add(peer, snapshot)
	if err != nil {
		return false, err
	}
	if added {
		s.logger.Info("REDACTED", "REDACTED", snapshot.Height, "REDACTED", snapshot.Format,
			"REDACTED", log.NewLazySprintf("REDACTED", snapshot.Hash))
	}
	return added, nil
}

//
//
func (s *syncer) AddPeer(peer p2p.Peer) {
	s.logger.Debug("REDACTED", "REDACTED", peer.ID())
	e := p2p.Envelope{
		ChannelID: SnapshotChannel,
		Message:   &sschema.SnapshotsRequest{},
	}
	peer.Send(e)
}

//
func (s *syncer) RemovePeer(peer p2p.Peer) {
	s.logger.Debug("REDACTED", "REDACTED", peer.ID())
	s.snapshots.RemovePeer(peer.ID())
}

//
func (s *syncer) RejectPeer(peer p2p.Peer) {
	s.logger.Debug("REDACTED", "REDACTED", peer.ID())
	s.snapshots.RejectPeer(peer.ID())
}

//
//
//
func (s *syncer) SyncAny(discoveryTime time.Duration, retryHook func()) (sm.State, *kinds.Commit, error) {
	if discoveryTime != 0 && discoveryTime < minimumDiscoveryTime {
		discoveryTime = 5 * minimumDiscoveryTime
	}

	if discoveryTime > 0 {
		s.logger.Info("REDACTED", "REDACTED", discoveryTime)
		time.Sleep(discoveryTime)
	}

	//
	//
	var (
		snapshot *snapshot
		chunks   *chunkQueue
		err      error
	)
	for {
		//
		if snapshot == nil {
			snapshot = s.snapshots.Best()
			chunks = nil
		}
		if snapshot == nil {
			if discoveryTime == 0 {
				return sm.State{}, nil, errNoSnapshots
			}
			retryHook()
			s.logger.Info("REDACTED", "REDACTED", log.NewLazySprintf("REDACTED", discoveryTime))
			time.Sleep(discoveryTime)
			continue
		}
		if chunks == nil {
			chunks, err = newChunkQueue(snapshot, s.tempDir)
			if err != nil {
				return sm.State{}, nil, fmt.Errorf("REDACTED", err)
			}
			defer chunks.Close() //
		}

		newState, commit, err := s.Sync(snapshot, chunks)
		switch {
		case err == nil:
			return newState, commit, nil

		case errors.Is(err, errAbort):
			return sm.State{}, nil, err

		case errors.Is(err, errRetrySnapshot):
			chunks.RetryAll()
			s.logger.Info("REDACTED", "REDACTED", snapshot.Height, "REDACTED", snapshot.Format,
				"REDACTED", log.NewLazySprintf("REDACTED", snapshot.Hash))
			continue

		case errors.Is(err, errTimeout):
			s.snapshots.Reject(snapshot)
			s.logger.Error("REDACTED",
				"REDACTED", snapshot.Height, "REDACTED", snapshot.Format, "REDACTED", log.NewLazySprintf("REDACTED", snapshot.Hash))

		case errors.Is(err, errRejectSnapshot):
			s.snapshots.Reject(snapshot)
			s.logger.Info("REDACTED", "REDACTED", snapshot.Height, "REDACTED", snapshot.Format,
				"REDACTED", log.NewLazySprintf("REDACTED", snapshot.Hash))

		case errors.Is(err, errRejectFormat):
			s.snapshots.RejectFormat(snapshot.Format)
			s.logger.Info("REDACTED", "REDACTED", snapshot.Format)

		case errors.Is(err, errRejectSender):
			s.logger.Info("REDACTED", "REDACTED", snapshot.Height, "REDACTED", snapshot.Format,
				"REDACTED", log.NewLazySprintf("REDACTED", snapshot.Hash))
			for _, peer := range s.snapshots.GetPeers(snapshot) {
				s.snapshots.RejectPeer(peer.ID())
				s.logger.Info("REDACTED", "REDACTED", peer.ID())
			}

		case errors.Is(err, context.DeadlineExceeded):
			s.logger.Info("REDACTED", "REDACTED", snapshot.Height, "REDACTED", err)
			s.snapshots.Reject(snapshot)

		default:
			return sm.State{}, nil, fmt.Errorf("REDACTED", err)
		}

		//
		err = chunks.Close()
		if err != nil {
			s.logger.Error("REDACTED", "REDACTED", err)
		}
		snapshot = nil
		chunks = nil
	}
}

//
//
func (s *syncer) Sync(snapshot *snapshot, chunks *chunkQueue) (sm.State, *kinds.Commit, error) {
	s.mtx.Lock()
	if s.chunks != nil {
		s.mtx.Unlock()
		return sm.State{}, nil, errors.New("REDACTED")
	}
	s.chunks = chunks
	s.mtx.Unlock()
	defer func() {
		s.mtx.Lock()
		s.chunks = nil
		s.mtx.Unlock()
	}()

	hctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	appHash, err := s.stateProvider.AppHash(hctx, snapshot.Height)
	if err != nil {
		s.logger.Info("REDACTED", "REDACTED", err)
		if errors.Is(err, minimal.ErrNoWitnesses) {
			return sm.State{}, nil, err
		}
		return sm.State{}, nil, errRejectSnapshot
	}
	snapshot.trustedAppHash = appHash

	//
	err = s.offerSnapshot(snapshot)
	if err != nil {
		return sm.State{}, nil, err
	}

	//
	fetchCtx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	for i := int32(0); i < s.chunkFetchers; i++ {
		go s.fetchChunks(fetchCtx, snapshot, chunks)
	}

	pctx, pcancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer pcancel()

	//
	state, err := s.stateProvider.State(pctx, snapshot.Height)
	if err != nil {
		s.logger.Info("REDACTED", "REDACTED", err)
		if errors.Is(err, minimal.ErrNoWitnesses) {
			return sm.State{}, nil, err
		}
		return sm.State{}, nil, errRejectSnapshot
	}
	commit, err := s.stateProvider.Commit(pctx, snapshot.Height)
	if err != nil {
		s.logger.Info("REDACTED", "REDACTED", err)
		if errors.Is(err, minimal.ErrNoWitnesses) {
			return sm.State{}, nil, err
		}
		return sm.State{}, nil, errRejectSnapshot
	}

	//
	err = s.applyChunks(chunks)
	if err != nil {
		return sm.State{}, nil, err
	}

	//
	if err := s.verifyApp(snapshot, state.Version.Consensus.App); err != nil {
		return sm.State{}, nil, err
	}

	//
	s.logger.Info("REDACTED", "REDACTED", snapshot.Height, "REDACTED", snapshot.Format,
		"REDACTED", log.NewLazySprintf("REDACTED", snapshot.Hash))

	return state, commit, nil
}

//
//
func (s *syncer) offerSnapshot(snapshot *snapshot) error {
	s.logger.Info("REDACTED", "REDACTED", snapshot.Height,
		"REDACTED", snapshot.Format, "REDACTED", log.NewLazySprintf("REDACTED", snapshot.Hash))
	resp, err := s.conn.OfferSnapshot(context.TODO(), &atci.RequestOfferSnapshot{
		Snapshot: &atci.Snapshot{
			Height:   snapshot.Height,
			Format:   snapshot.Format,
			Chunks:   snapshot.Chunks,
			Hash:     snapshot.Hash,
			Metadata: snapshot.Metadata,
		},
		AppHash: snapshot.trustedAppHash,
	})
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	switch resp.Result {
	case atci.ResponseOfferSnapshot_ACCEPT:
		s.logger.Info("REDACTED", "REDACTED", snapshot.Height,
			"REDACTED", snapshot.Format, "REDACTED", log.NewLazySprintf("REDACTED", snapshot.Hash))
		return nil
	case atci.ResponseOfferSnapshot_ABORT:
		return errAbort
	case atci.ResponseOfferSnapshot_REJECT:
		return errRejectSnapshot
	case atci.ResponseOfferSnapshot_REJECT_FORMAT:
		return errRejectFormat
	case atci.ResponseOfferSnapshot_REJECT_SENDER:
		return errRejectSender
	default:
		return fmt.Errorf("REDACTED", resp.Result)
	}
}

//
//
func (s *syncer) applyChunks(chunks *chunkQueue) error {
	for {
		chunk, err := chunks.Next()
		if errors.Is(err, errDone) {
			return nil
		} else if err != nil {
			return fmt.Errorf("REDACTED", err)
		}

		resp, err := s.conn.ApplySnapshotChunk(context.TODO(), &atci.RequestApplySnapshotChunk{
			Index:  chunk.Index,
			Chunk:  chunk.Chunk,
			Sender: string(chunk.Sender),
		})
		if err != nil {
			return fmt.Errorf("REDACTED", chunk.Index, err)
		}
		s.logger.Info("REDACTED", "REDACTED", chunk.Height,
			"REDACTED", chunk.Format, "REDACTED", chunk.Index, "REDACTED", chunks.Size())

		//
		for _, index := range resp.RefetchChunks {
			err := chunks.Discard(index)
			if err != nil {
				return fmt.Errorf("REDACTED", index, err)
			}
		}

		//
		for _, sender := range resp.RejectSenders {
			if sender != "REDACTED" {
				s.snapshots.RejectPeer(p2p.ID(sender))
				err := chunks.DiscardSender(p2p.ID(sender))
				if err != nil {
					return fmt.Errorf("REDACTED", err)
				}
			}
		}

		switch resp.Result {
		case atci.ResponseApplySnapshotChunk_ACCEPT:
		case atci.ResponseApplySnapshotChunk_ABORT:
			return errAbort
		case atci.ResponseApplySnapshotChunk_RETRY:
			chunks.Retry(chunk.Index)
		case atci.ResponseApplySnapshotChunk_RETRY_SNAPSHOT:
			return errRetrySnapshot
		case atci.ResponseApplySnapshotChunk_REJECT_SNAPSHOT:
			return errRejectSnapshot
		default:
			return fmt.Errorf("REDACTED", resp.Result)
		}
	}
}

//
//
func (s *syncer) fetchChunks(ctx context.Context, snapshot *snapshot, chunks *chunkQueue) {
	var (
		next  = true
		index uint32
		err   error
	)

	for {
		if next {
			index, err = chunks.Allocate()
			if errors.Is(err, errDone) {
				//
				//
				select {
				case <-ctx.Done():
					return
				default:
				}
				time.Sleep(2 * time.Second)
				continue
			}
			if err != nil {
				s.logger.Error("REDACTED", "REDACTED", err)
				return
			}
		}
		s.logger.Info("REDACTED", "REDACTED", snapshot.Height,
			"REDACTED", snapshot.Format, "REDACTED", index, "REDACTED", chunks.Size())

		ticker := time.NewTicker(s.retryTimeout)
		defer ticker.Stop()

		s.requestChunk(snapshot, index)

		select {
		case <-chunks.WaitFor(index):
			next = true

		case <-ticker.C:
			next = false

		case <-ctx.Done():
			return
		}

		ticker.Stop()
	}
}

//
func (s *syncer) requestChunk(snapshot *snapshot, chunk uint32) {
	peer := s.snapshots.GetPeer(snapshot)
	if peer == nil {
		s.logger.Error("REDACTED", "REDACTED", snapshot.Height,
			"REDACTED", snapshot.Format, "REDACTED", log.NewLazySprintf("REDACTED", snapshot.Hash))
		return
	}
	s.logger.Debug("REDACTED", "REDACTED", snapshot.Height,
		"REDACTED", snapshot.Format, "REDACTED", chunk, "REDACTED", peer.ID())
	peer.Send(p2p.Envelope{
		ChannelID: ChunkChannel,
		Message: &sschema.ChunkRequest{
			Height: snapshot.Height,
			Format: snapshot.Format,
			Index:  chunk,
		},
	})
}

//
func (s *syncer) verifyApp(snapshot *snapshot, appVersion uint64) error {
	resp, err := s.connQuery.Info(context.TODO(), gateway.RequestInfo)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	//
	if resp.AppVersion != appVersion {
		//
		//
		return fmt.Errorf("REDACTED",
			appVersion, resp.AppVersion)
	}
	if !bytes.Equal(snapshot.trustedAppHash, resp.LastBlockAppHash) {
		s.logger.Error("REDACTED",
			"REDACTED", fmt.Sprintf("REDACTED", snapshot.trustedAppHash),
			"REDACTED", fmt.Sprintf("REDACTED", resp.LastBlockAppHash))
		return errVerifyFailed
	}
	if uint64(resp.LastBlockHeight) != snapshot.Height {
		s.logger.Error(
			"REDACTED",
			"REDACTED", snapshot.Height,
			"REDACTED", resp.LastBlockHeight,
		)
		return errVerifyFailed
	}

	s.logger.Info("REDACTED", "REDACTED", snapshot.Height, "REDACTED", log.NewLazySprintf("REDACTED", snapshot.trustedAppHash))
	return nil
}
