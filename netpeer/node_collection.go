package netpeer

import (
	"sort"
	"sync"

	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"
)

//
//
type NodeCollection struct {
	machine *Machine

	nodes map[peer.ID]*Node
	mu    sync.RWMutex

	stats *p2p.Stats
	tracer  log.Tracer
}

var _ p2p.IDXNodeCollection = (*NodeCollection)(nil)

var (
	ErrNodePresent = errors.New("REDACTED")
	ErrEgoNode   = errors.New("REDACTED")
)

//
func NewNodeCollection(machine *Machine, stats *p2p.Stats, tracer log.Tracer) *NodeCollection {
	return &NodeCollection{
		machine:    machine,
		nodes:   make(map[peer.ID]*Node),
		stats: stats,
		tracer:  tracer,
	}
}

func (ps *NodeCollection) Has(key p2p.ID) bool {
	_, present := ps.fetchByKey(key)

	return present
}

func (ps *NodeCollection) Get(key p2p.ID) p2p.Node {
	node, present := ps.fetchByKey(key)
	if !present {
		return nil
	}

	return node
}

//
//
type NodeAppendSettings struct {
	Internal       bool
	Durable    bool
	Absolute bool
	OnPriorBegin func(p *Node)
	OnAfterBegin  func(p *Node)
	OnBeginErrored func(p *Node, cause any)
}

//
//
//
//
func (ps *NodeCollection) Add(addressDetails peer.AddrInfo, opts NodeAppendSettings) (*Node, error) {
	id := addressDetails.ID

	switch {
	case id == ps.machine.ID():
		return nil, ErrEgoNode
	case len(addressDetails.Addrs) == 0:
		return nil, errors.New("REDACTED")
	}

	ps.tracer.Details("REDACTED", "REDACTED", id.String())

	p, err := NewNode(ps.machine, addressDetails, ps.stats, opts.Internal, opts.Durable, opts.Absolute)
	if err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	set := ps.set(id, p)
	if !set {
		return nil, ErrNodePresent
	}

	if opts.OnPriorBegin != nil {
		opts.OnPriorBegin(p)
	}

	if err := p.Begin(); err != nil {
		ps.clear(id)
		if opts.OnBeginErrored != nil {
			opts.OnBeginErrored(p, err)
		}
		return nil, errors.Wrap(err, "REDACTED")
	}

	if opts.OnAfterBegin != nil {
		opts.OnAfterBegin(p)
	}

	ps.stats.Nodes.Add(1)

	return p, nil
}

//
//
//
type NodeDeletionSettings struct {
	Cause      any
	OnAfterHalt func(p *Node, cause any)
}

func (ps *NodeCollection) Delete(key p2p.ID, opts NodeDeletionSettings) error {
	id := ps.keyToNodeUID(key)
	if id == "REDACTED" {
		return errors.New("REDACTED")
	}

	if id == ps.machine.ID() {
		return ErrEgoNode
	}

	ps.tracer.Details("REDACTED", "REDACTED", id.String(), "REDACTED", opts.Cause)

	p, ok := ps.clear(id)
	if !ok {
		return errors.New("REDACTED")
	}

	if err := p.Halt(); err != nil {
		return errors.Wrap(err, "REDACTED")
	}

	if opts.OnAfterHalt != nil {
		opts.OnAfterHalt(p, opts.Cause)
	}

	ps.stats.Nodes.Add(-1)

	return nil
}

func (ps *NodeCollection) DeleteAll(opts NodeDeletionSettings) {
	nodes := ps.Clone()

	for _, node := range nodes {
		id := peer.ID()
		if err := ps.Delete(id, opts); err != nil {
			ps.tracer.Fault("REDACTED", "REDACTED", id, "REDACTED", err)
		}
	}
}

func (ps *NodeCollection) ForEach(fn func(p2p.Node)) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, p := range ps.nodes {
		fn(p)
	}
}

//
//
//
//
func (ps *NodeCollection) Arbitrary() p2p.Node {
	nodes := ps.Clone()
	if len(nodes) == 0 {
		return nil
	}

	idx := engineseed.Intn(len(nodes))

	return nodes[idx]
}

//
//
//
func (ps *NodeCollection) Clone() []p2p.Node {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	outcomes := make([]p2p.Node, 0, len(ps.nodes))
	for _, p := range ps.nodes {
		outcomes = append(outcomes, p)
	}

	sort.Slice(outcomes, func(i, j int) bool {
		return outcomes[i].ID() < outcomes[j].ID()
	})

	return outcomes
}

//
func (ps *NodeCollection) Volume() int {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	return len(ps.nodes)
}

func (ps *NodeCollection) fetchByKey(key p2p.ID) (*Node, bool) {
	id := ps.keyToNodeUID(key)
	if id == "REDACTED" {
		return nil, false
	}

	ps.mu.RLock()
	defer ps.mu.RUnlock()

	p, ok := ps.nodes[id]
	if !ok {
		return nil, false
	}

	return p, true
}

//
func (ps *NodeCollection) set(id peer.ID, p *Node) bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.nodes[id]; ok {
		return false
	}

	ps.nodes[id] = p

	return true
}

//
//
func (ps *NodeCollection) clear(id peer.ID) (*Node, bool) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	p, ok := ps.nodes[id]
	if !ok {
		return nil, false
	}

	delete(ps.nodes, id)

	return p, true
}

func (ps *NodeCollection) keyToNodeUID(key p2p.ID) peer.ID {
	if key == "REDACTED" {
		return "REDACTED"
	}

	b, err := base58.Decode(string(key))
	if err != nil {
		ps.tracer.Fault("REDACTED", "REDACTED", key, "REDACTED", err)
		return "REDACTED"
	}

	id, err := peer.IDFromBytes(b)
	if err != nil {
		ps.tracer.Fault("REDACTED", "REDACTED", key, "REDACTED", err)
		return "REDACTED"
	}

	return id
}

//
//
func nodeUIDToKey(id peer.ID) p2p.ID {
	return p2p.ID(id.String())
}
