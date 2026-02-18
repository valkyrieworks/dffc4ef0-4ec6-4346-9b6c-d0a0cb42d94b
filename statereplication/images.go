package statereplication

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"sort"

	ctsync "github.com/valkyrieworks/utils/alignment"
	"github.com/valkyrieworks/p2p"
)

//
type snapshotKey [sha256.Size]byte

//
type snapshot struct {
	Height   uint64
	Format   uint32
	Chunks   uint32
	Hash     []byte
	Metadata []byte

	trustedAppHash []byte //
}

//
//
//
func (s *snapshot) Key() snapshotKey {
	//
	hasher := sha256.New()
	fmt.Fprintf(hasher, "REDACTED", s.Height, s.Format, s.Chunks)
	hasher.Write(s.Hash)
	hasher.Write(s.Metadata)
	var key snapshotKey
	copy(key[:], hasher.Sum(nil))
	return key
}

//
type snapshotPool struct {
	ctsync.Mutex
	snapshots     map[snapshotKey]*snapshot
	snapshotPeers map[snapshotKey]map[p2p.ID]p2p.Peer

	//
	formatIndex map[uint32]map[snapshotKey]bool
	heightIndex map[uint64]map[snapshotKey]bool
	peerIndex   map[p2p.ID]map[snapshotKey]bool

	//
	formatRejectlist   map[uint32]bool
	peerRejectlist     map[p2p.ID]bool
	snapshotRejectlist map[snapshotKey]bool
}

//
func newSnapshotPool() *snapshotPool {
	return &snapshotPool{
		snapshots:          make(map[snapshotKey]*snapshot),
		snapshotPeers:      make(map[snapshotKey]map[p2p.ID]p2p.Peer),
		formatIndex:        make(map[uint32]map[snapshotKey]bool),
		heightIndex:        make(map[uint64]map[snapshotKey]bool),
		peerIndex:          make(map[p2p.ID]map[snapshotKey]bool),
		formatRejectlist:   make(map[uint32]bool),
		peerRejectlist:     make(map[p2p.ID]bool),
		snapshotRejectlist: make(map[snapshotKey]bool),
	}
}

//
//
//
func (p *snapshotPool) Add(peer p2p.Peer, snapshot *snapshot) (bool, error) {
	key := snapshot.Key()

	p.Lock()
	defer p.Unlock()

	switch {
	case p.formatRejectlist[snapshot.Format]:
		return false, nil
	case p.peerRejectlist[peer.ID()]:
		return false, nil
	case p.snapshotRejectlist[key]:
		return false, nil
	case len(p.peerIndex[peer.ID()]) >= recentSnapshots:
		return false, nil
	}

	if p.snapshotPeers[key] == nil {
		p.snapshotPeers[key] = make(map[p2p.ID]p2p.Peer)
	}
	p.snapshotPeers[key][peer.ID()] = peer

	if p.peerIndex[peer.ID()] == nil {
		p.peerIndex[peer.ID()] = make(map[snapshotKey]bool)
	}
	p.peerIndex[peer.ID()][key] = true

	if p.snapshots[key] != nil {
		return false, nil
	}
	p.snapshots[key] = snapshot

	if p.formatIndex[snapshot.Format] == nil {
		p.formatIndex[snapshot.Format] = make(map[snapshotKey]bool)
	}
	p.formatIndex[snapshot.Format][key] = true

	if p.heightIndex[snapshot.Height] == nil {
		p.heightIndex[snapshot.Height] = make(map[snapshotKey]bool)
	}
	p.heightIndex[snapshot.Height][key] = true

	return true, nil
}

//
func (p *snapshotPool) Best() *snapshot {
	ranked := p.Ranked()
	if len(ranked) == 0 {
		return nil
	}
	return ranked[0]
}

//
func (p *snapshotPool) GetPeer(snapshot *snapshot) p2p.Peer {
	peers := p.GetPeers(snapshot)
	if len(peers) == 0 {
		return nil
	}
	return peers[rand.Intn(len(peers))] //
}

//
func (p *snapshotPool) GetPeers(snapshot *snapshot) []p2p.Peer {
	key := snapshot.Key()
	p.Lock()
	defer p.Unlock()

	peers := make([]p2p.Peer, 0, len(p.snapshotPeers[key]))
	for _, peer := range p.snapshotPeers[key] {
		peers = append(peers, peer)
	}
	//
	sort.Slice(peers, func(a int, b int) bool {
		return peers[a].ID() < peers[b].ID()
	})
	return peers
}

//
//
//
func (p *snapshotPool) Ranked() []*snapshot {
	p.Lock()
	defer p.Unlock()

	candidates := make([]*snapshot, 0, len(p.snapshots))
	for key := range p.snapshots {
		candidates = append(candidates, p.snapshots[key])
	}

	sort.Slice(candidates, func(i, j int) bool {
		a := candidates[i]
		b := candidates[j]

		switch {
		case a.Height > b.Height:
			return true
		case a.Height < b.Height:
			return false
		case a.Format > b.Format:
			return true
		case a.Format < b.Format:
			return false
		case len(p.snapshotPeers[a.Key()]) > len(p.snapshotPeers[b.Key()]):
			return true
		default:
			return false
		}
	})

	return candidates
}

//
func (p *snapshotPool) Reject(snapshot *snapshot) {
	key := snapshot.Key()
	p.Lock()
	defer p.Unlock()

	p.snapshotRejectlist[key] = true
	p.removeSnapshot(key)
}

//
func (p *snapshotPool) RejectFormat(format uint32) {
	p.Lock()
	defer p.Unlock()

	p.formatRejectlist[format] = true
	for key := range p.formatIndex[format] {
		p.removeSnapshot(key)
	}
}

//
func (p *snapshotPool) RejectPeer(peerID p2p.ID) {
	if peerID == "REDACTED" {
		return
	}
	p.Lock()
	defer p.Unlock()

	p.removePeer(peerID)
	p.peerRejectlist[peerID] = true
}

//
func (p *snapshotPool) RemovePeer(peerID p2p.ID) {
	p.Lock()
	defer p.Unlock()
	p.removePeer(peerID)
}

//
func (p *snapshotPool) removePeer(peerID p2p.ID) {
	for key := range p.peerIndex[peerID] {
		delete(p.snapshotPeers[key], peerID)
		if len(p.snapshotPeers[key]) == 0 {
			p.removeSnapshot(key)
		}
	}
	delete(p.peerIndex, peerID)
}

//
func (p *snapshotPool) removeSnapshot(key snapshotKey) {
	snapshot := p.snapshots[key]
	if snapshot == nil {
		return
	}

	delete(p.snapshots, key)
	delete(p.formatIndex[snapshot.Format], key)
	delete(p.heightIndex[snapshot.Height], key)
	for peerID := range p.snapshotPeers[key] {
		delete(p.peerIndex[peerID], key)
	}
	delete(p.snapshotPeers, key)
}
