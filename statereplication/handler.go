package statereplication

import (
	"context"
	"errors"
	"sort"
	"time"

	atci "github.com/valkyrieworks/atci/kinds"
	"github.com/valkyrieworks/settings"
	ctsync "github.com/valkyrieworks/utils/alignment"
	"github.com/valkyrieworks/p2p"
	sschema "github.com/valkyrieworks/schema/consensuscore/statereplication"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	SnapshotChannel = byte(0x60)
	//
	ChunkChannel = byte(0x61)
	//
	recentSnapshots = 10
)

//
//
type Reactor struct {
	p2p.BaseReactor

	cfg       settings.StateSyncConfig
	conn      gateway.AppConnSnapshot
	connQuery gateway.AppConnQuery
	tempDir   string
	metrics   *Metrics

	//
	//
	mtx    ctsync.RWMutex
	syncer *syncer
}

//
func NewReactor(
	cfg settings.StateSyncConfig,
	conn gateway.AppConnSnapshot,
	connQuery gateway.AppConnQuery,
	metrics *Metrics,
) *Reactor {
	r := &Reactor{
		cfg:       cfg,
		conn:      conn,
		connQuery: connQuery,
		metrics:   metrics,
	}
	r.BaseReactor = *p2p.NewBaseReactor("REDACTED", r)

	return r
}

//
func (r *Reactor) GetChannels() []*p2p.ChannelDescriptor {
	return []*p2p.ChannelDescriptor{
		{
			ID:                  SnapshotChannel,
			Priority:            5,
			SendQueueCapacity:   10,
			RecvMessageCapacity: snapshotMsgSize,
			MessageType:         &sschema.Message{},
		},
		{
			ID:                  ChunkChannel,
			Priority:            3,
			SendQueueCapacity:   10,
			RecvMessageCapacity: chunkMsgSize,
			MessageType:         &sschema.Message{},
		},
	}
}

//
func (r *Reactor) OnStart() error {
	return nil
}

//
func (r *Reactor) AddPeer(peer p2p.Peer) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if r.syncer != nil {
		r.syncer.AddPeer(peer)
	}
}

//
func (r *Reactor) RemovePeer(peer p2p.Peer, _ any) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if r.syncer != nil {
		r.syncer.RemovePeer(peer)
	}
}

//
func (r *Reactor) Receive(e p2p.Envelope) {
	if !r.IsRunning() {
		return
	}

	err := validateMsg(e.Message, r.cfg.MaxSnapshotChunks)
	if err != nil {
		if errors.Is(err, ErrExceedsMaxSnapshotChunks) {
			r.syncer.RejectPeer(e.Src)
		}
		r.Logger.Error("REDACTED", "REDACTED", e.Src, "REDACTED", e.Message, "REDACTED", err)
		r.Switch.StopPeerForError(e.Src, err)
		return
	}

	switch e.ChannelID {
	case SnapshotChannel:
		switch msg := e.Message.(type) {
		case *sschema.SnapshotsRequest:
			snapshots, err := r.recentSnapshots(recentSnapshots)
			if err != nil {
				r.Logger.Error("REDACTED", "REDACTED", err)
				return
			}
			for _, snapshot := range snapshots {
				r.Logger.Debug("REDACTED", "REDACTED", snapshot.Height,
					"REDACTED", snapshot.Format, "REDACTED", e.Src.ID())
				e.Src.Send(p2p.Envelope{
					ChannelID: e.ChannelID,
					Message: &sschema.SnapshotsResponse{
						Height:   snapshot.Height,
						Format:   snapshot.Format,
						Chunks:   snapshot.Chunks,
						Hash:     snapshot.Hash,
						Metadata: snapshot.Metadata,
					},
				})
			}

		case *sschema.SnapshotsResponse:
			r.mtx.RLock()
			defer r.mtx.RUnlock()
			if r.syncer == nil {
				r.Logger.Debug("REDACTED")
				return
			}
			r.Logger.Debug("REDACTED", "REDACTED", msg.Height, "REDACTED", msg.Format, "REDACTED", e.Src.ID())
			_, err := r.syncer.AddSnapshot(e.Src, &snapshot{
				Height:   msg.Height,
				Format:   msg.Format,
				Chunks:   msg.Chunks,
				Hash:     msg.Hash,
				Metadata: msg.Metadata,
			})
			//
			if err != nil {
				r.Logger.Error("REDACTED", "REDACTED", msg.Height, "REDACTED", msg.Format,
					"REDACTED", e.Src.ID(), "REDACTED", err)
				return
			}

		default:
			r.Logger.Error("REDACTED", msg)
		}

	case ChunkChannel:
		switch msg := e.Message.(type) {
		case *sschema.ChunkRequest:
			r.Logger.Debug("REDACTED", "REDACTED", msg.Height, "REDACTED", msg.Format,
				"REDACTED", msg.Index, "REDACTED", e.Src.ID())
			resp, err := r.conn.LoadSnapshotChunk(context.TODO(), &atci.RequestLoadSnapshotChunk{
				Height: msg.Height,
				Format: msg.Format,
				Chunk:  msg.Index,
			})
			if err != nil {
				r.Logger.Error("REDACTED", "REDACTED", msg.Height, "REDACTED", msg.Format,
					"REDACTED", msg.Index, "REDACTED", err)
				return
			}
			r.Logger.Debug("REDACTED", "REDACTED", msg.Height, "REDACTED", msg.Format,
				"REDACTED", msg.Index, "REDACTED", e.Src.ID())
			e.Src.Send(p2p.Envelope{
				ChannelID: ChunkChannel,
				Message: &sschema.ChunkResponse{
					Height:  msg.Height,
					Format:  msg.Format,
					Index:   msg.Index,
					Chunk:   resp.Chunk,
					Missing: resp.Chunk == nil,
				},
			})

		case *sschema.ChunkResponse:
			r.mtx.RLock()
			defer r.mtx.RUnlock()
			if r.syncer == nil {
				r.Logger.Debug("REDACTED", "REDACTED", e.Src.ID())
				return
			}
			r.Logger.Debug("REDACTED", "REDACTED", msg.Height, "REDACTED", msg.Format,
				"REDACTED", msg.Index, "REDACTED", e.Src.ID())
			_, err := r.syncer.AddChunk(&chunk{
				Height: msg.Height,
				Format: msg.Format,
				Index:  msg.Index,
				Chunk:  msg.Chunk,
				Sender: e.Src.ID(),
			})
			if err != nil {
				r.Logger.Error("REDACTED", "REDACTED", msg.Height, "REDACTED", msg.Format,
					"REDACTED", msg.Index, "REDACTED", err)
				return
			}

		default:
			r.Logger.Error("REDACTED", msg)
		}

	default:
		r.Logger.Error("REDACTED", e.ChannelID)
	}
}

//
func (r *Reactor) recentSnapshots(n uint32) ([]*snapshot, error) {
	resp, err := r.conn.ListSnapshots(context.TODO(), &atci.RequestListSnapshots{})
	if err != nil {
		return nil, err
	}
	sort.Slice(resp.Snapshots, func(i, j int) bool {
		a := resp.Snapshots[i]
		b := resp.Snapshots[j]
		switch {
		case a.Height > b.Height:
			return true
		case a.Height == b.Height && a.Format > b.Format:
			return true
		default:
			return false
		}
	})
	snapshots := make([]*snapshot, 0, n)
	for i, s := range resp.Snapshots {
		if i >= recentSnapshots {
			break
		}
		snapshots = append(snapshots, &snapshot{
			Height:   s.Height,
			Format:   s.Format,
			Chunks:   s.Chunks,
			Hash:     s.Hash,
			Metadata: s.Metadata,
		})
	}
	return snapshots, nil
}

//
//
func (r *Reactor) Sync(stateProvider StateProvider, discoveryTime time.Duration) (sm.State, *kinds.Commit, error) {
	r.mtx.Lock()
	if r.syncer != nil {
		r.mtx.Unlock()
		return sm.State{}, nil, errors.New("REDACTED")
	}
	r.metrics.Syncing.Set(1)
	r.syncer = newSyncer(r.cfg, r.Logger, r.conn, r.connQuery, stateProvider, r.tempDir)
	r.mtx.Unlock()

	hook := func() {
		r.Logger.Debug("REDACTED")
		//

		r.Switch.BroadcastAsync(p2p.Envelope{
			ChannelID: SnapshotChannel,
			Message:   &sschema.SnapshotsRequest{},
		})
	}

	hook()

	state, commit, err := r.syncer.SyncAny(discoveryTime, hook)

	r.mtx.Lock()
	r.syncer = nil
	r.metrics.Syncing.Set(0)
	r.mtx.Unlock()
	return state, commit, err
}
