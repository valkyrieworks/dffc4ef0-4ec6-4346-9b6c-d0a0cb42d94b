package netp2p

import (
	"sort"
	"sync"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"
)

//
//
type NodeAssign struct {
	machine *Machine

	nodes map[peer.ID]*Node
	mu    sync.RWMutex

	telemetry *p2p.Telemetry
	tracer  log.Tracer
}

var _ p2p.IDXNodeAssign = (*NodeAssign)(nil)

var (
	FaultNodePresent = errors.New("REDACTED")
	FaultEgoNode   = errors.New("REDACTED")
)

//
func FreshNodeAssign(machine *Machine, telemetry *p2p.Telemetry, tracer log.Tracer) *NodeAssign {
	return &NodeAssign{
		machine:    machine,
		nodes:   make(map[peer.ID]*Node),
		telemetry: telemetry,
		tracer:  tracer,
	}
}

func (ps *NodeAssign) Has(key p2p.ID) bool {
	_, present := ps.obtainViaToken(key)

	return present
}

func (ps *NodeAssign) Get(key p2p.ID) p2p.Node {
	node, present := ps.obtainViaToken(key)
	if !present {
		return nil
	}

	return node
}

//
//
type NodeAppendChoices struct {
	Secluded       bool
	Enduring    bool
	Absolute bool
	UponPriorInitiate func(p *Node)
	UponSubsequentInitiate  func(p *Node)
	UponInitiateUnsuccessful func(p *Node, rationale any)
}

//
//
//
//
func (ps *NodeAssign) Add(locationDetails peer.AddrInfo, choices NodeAppendChoices) (*Node, error) {
	id := locationDetails.ID

	switch {
	case id == ps.machine.ID():
		return nil, FaultEgoNode
	case len(locationDetails.Addrs) == 0:
		return nil, errors.New("REDACTED")
	}

	ps.tracer.Details("REDACTED", "REDACTED", id.String())

	p, err := FreshNode(ps.machine, locationDetails, ps.telemetry, choices.Secluded, choices.Enduring, choices.Absolute)
	if err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	set := ps.set(id, p)
	if !set {
		return nil, FaultNodePresent
	}

	if choices.UponPriorInitiate != nil {
		choices.UponPriorInitiate(p)
	}

	if err := p.Initiate(); err != nil {
		ps.deassign(id)
		if choices.UponInitiateUnsuccessful != nil {
			choices.UponInitiateUnsuccessful(p, err)
		}
		return nil, errors.Wrap(err, "REDACTED")
	}

	if choices.UponSubsequentInitiate != nil {
		choices.UponSubsequentInitiate(p)
	}

	ps.telemetry.Nodes.Add(1)

	return p, nil
}

//
//
//
type NodeDeletionChoices struct {
	Rationale      any
	UponSubsequentHalt func(p *Node, rationale any)
}

func (ps *NodeAssign) Discard(key p2p.ID, choices NodeDeletionChoices) error {
	id := ps.tokenTowardNodeUUID(key)
	if id == "REDACTED" {
		return errors.New("REDACTED")
	}

	if id == ps.machine.ID() {
		return FaultEgoNode
	}

	ps.tracer.Details("REDACTED", "REDACTED", id.String(), "REDACTED", choices.Rationale)

	p, ok := ps.deassign(id)
	if !ok {
		return errors.New("REDACTED")
	}

	if err := p.Halt(); err != nil {
		return errors.Wrap(err, "REDACTED")
	}

	if choices.UponSubsequentHalt != nil {
		choices.UponSubsequentHalt(p, choices.Rationale)
	}

	ps.telemetry.Nodes.Add(-1)

	return nil
}

func (ps *NodeAssign) DiscardEvery(choices NodeDeletionChoices) {
	nodes := ps.Duplicate()

	for _, node := range nodes {
		id := peer.ID()
		if err := ps.Discard(id, choices); err != nil {
			ps.tracer.Failure("REDACTED", "REDACTED", id, "REDACTED", err)
		}
	}
}

func (ps *NodeAssign) ForeachEvery(fn func(p2p.Node)) {
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
func (ps *NodeAssign) Unpredictable() p2p.Node {
	nodes := ps.Duplicate()
	if len(nodes) == 0 {
		return nil
	}

	idx := commitrand.Integern(len(nodes))

	return nodes[idx]
}

//
//
//
func (ps *NodeAssign) Duplicate() []p2p.Node {
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
func (ps *NodeAssign) Extent() int {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	return len(ps.nodes)
}

func (ps *NodeAssign) obtainViaToken(key p2p.ID) (*Node, bool) {
	id := ps.tokenTowardNodeUUID(key)
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
func (ps *NodeAssign) set(id peer.ID, p *Node) bool {
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
func (ps *NodeAssign) deassign(id peer.ID) (*Node, bool) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	p, ok := ps.nodes[id]
	if !ok {
		return nil, false
	}

	delete(ps.nodes, id)

	return p, true
}

func (ps *NodeAssign) tokenTowardNodeUUID(key p2p.ID) peer.ID {
	if key == "REDACTED" {
		return "REDACTED"
	}

	b, err := base58.Decode(string(key))
	if err != nil {
		ps.tracer.Failure("REDACTED", "REDACTED", key, "REDACTED", err)
		return "REDACTED"
	}

	id, err := peer.IDFromBytes(b)
	if err != nil {
		ps.tracer.Failure("REDACTED", "REDACTED", key, "REDACTED", err)
		return "REDACTED"
	}

	return id
}

//
//
func nodeUUIDTowardToken(id peer.ID) p2p.ID {
	return p2p.ID(id.String())
}
