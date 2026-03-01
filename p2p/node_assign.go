package p2p

import (
	"net"

	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
type IDXNodeAssign interface {
	//
	Has(key ID) bool

	//
	Get(key ID) Node
	//
	Duplicate() []Node
	//
	Extent() int
	//
	ForeachEvery(node func(Node))
	//
	Unpredictable() Node
}

//

//
type NodeAssign struct {
	mtx    commitchronize.Exclusion
	search map[ID]*nodeAssignElement
	catalog   []Node
}

type nodeAssignElement struct {
	node  Node
	ordinal int
}

//
func FreshNodeAssign() *NodeAssign {
	return &NodeAssign{
		search: make(map[ID]*nodeAssignElement),
		catalog:   make([]Node, 0, 256),
	}
}

//
//
func (ps *NodeAssign) Add(node Node) error {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.search[node.ID()] != nil {
		return FaultRouterReplicatedNodeUUID{node.ID()}
	}
	if node.ObtainDeletionUnsuccessful() {
		return FaultNodeDeletion{}
	}

	ordinal := len(ps.catalog)
	//
	//
	ps.catalog = append(ps.catalog, node)
	ps.search[node.ID()] = &nodeAssignElement{node, ordinal}
	return nil
}

//
//
func (ps *NodeAssign) Has(nodeToken ID) bool {
	ps.mtx.Lock()
	_, ok := ps.search[nodeToken]
	ps.mtx.Unlock()
	return ok
}

//
//
func (ps *NodeAssign) OwnsINET(nodeINET net.IP) bool {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	for _, node := range ps.catalog {
		if node.DistantINET().Equal(nodeINET) {
			return true
		}
	}

	return false
}

//
//
func (ps *NodeAssign) Get(nodeToken ID) Node {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	record, ok := ps.search[nodeToken]
	if ok {
		return record.node
	}
	return nil
}

//
func (ps *NodeAssign) Discard(node Node) bool {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	record, ok := ps.search[node.ID()]
	if !ok || len(ps.catalog) == 0 {
		//
		//
		//
		//
		//
		node.AssignDeletionUnsuccessful()
		return false
	}
	ordinal := record.ordinal

	//
	delete(ps.search, node.ID())

	//
	if ordinal != len(ps.catalog)-1 {
		//
		finalNode := ps.catalog[len(ps.catalog)-1]
		record := ps.search[finalNode.ID()]
		record.ordinal = ordinal
		ps.catalog[ordinal] = record.node
	}

	//
	ps.catalog[len(ps.catalog)-1] = nil //
	ps.catalog = ps.catalog[:len(ps.catalog)-1]

	return true
}

//
func (ps *NodeAssign) Extent() int {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return len(ps.catalog)
}

//
//
//
//
//
//
//
func (ps *NodeAssign) Catalog() []Node {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.catalog
}

//
//
//
func (ps *NodeAssign) Duplicate() []Node {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	c := make([]Node, len(ps.catalog))
	copy(c, ps.catalog)
	return c
}

//
func (ps *NodeAssign) ForeachEvery(fn func(node Node)) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	for _, record := range ps.search {
		fn(record.node)
	}
}

//
func (ps *NodeAssign) Unpredictable() Node {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if len(ps.catalog) == 0 {
		return nil
	}

	return ps.catalog[commitrand.Int()%len(ps.catalog)]
}
