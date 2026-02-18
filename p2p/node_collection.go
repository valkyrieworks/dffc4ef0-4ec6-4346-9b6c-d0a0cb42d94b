package p2p

import (
	"net"

	engineseed "github.com/valkyrieworks/utils/random"
	engineconnect "github.com/valkyrieworks/utils/align"
)

//
type IDXNodeCollection interface {
	//
	Has(key ID) bool

	//
	Get(key ID) Node
	//
	Clone() []Node
	//
	Volume() int
	//
	ForEach(node func(Node))
	//
	Arbitrary() Node
}

//

//
type NodeCollection struct {
	mtx    engineconnect.Lock
	search map[ID]*nodeCollectionItem
	catalog   []Node
}

type nodeCollectionItem struct {
	node  Node
	ordinal int
}

//
func NewNodeCollection() *NodeCollection {
	return &NodeCollection{
		search: make(map[ID]*nodeCollectionItem),
		catalog:   make([]Node, 0, 256),
	}
}

//
//
func (ps *NodeCollection) Add(node Node) error {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.search[node.ID()] != nil {
		return ErrRouterReplicatedNodeUID{node.ID()}
	}
	if node.FetchDeletionErrored() {
		return ErrNodeDeletion{}
	}

	ordinal := len(ps.catalog)
	//
	//
	ps.catalog = append(ps.catalog, node)
	ps.search[node.ID()] = &nodeCollectionItem{node, ordinal}
	return nil
}

//
//
func (ps *NodeCollection) Has(nodeKey ID) bool {
	ps.mtx.Lock()
	_, ok := ps.search[nodeKey]
	ps.mtx.Unlock()
	return ok
}

//
//
func (ps *NodeCollection) HasIP(nodeIP net.IP) bool {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	for _, node := range ps.catalog {
		if node.DistantIP().Equal(nodeIP) {
			return true
		}
	}

	return false
}

//
//
func (ps *NodeCollection) Get(nodeKey ID) Node {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	item, ok := ps.search[nodeKey]
	if ok {
		return item.node
	}
	return nil
}

//
func (ps *NodeCollection) Delete(node Node) bool {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	item, ok := ps.search[node.ID()]
	if !ok || len(ps.catalog) == 0 {
		//
		//
		//
		//
		//
		node.CollectionDeletionErrored()
		return false
	}
	ordinal := item.ordinal

	//
	delete(ps.search, node.ID())

	//
	if ordinal != len(ps.catalog)-1 {
		//
		finalNode := ps.catalog[len(ps.catalog)-1]
		item := ps.search[finalNode.ID()]
		item.ordinal = ordinal
		ps.catalog[ordinal] = item.node
	}

	//
	ps.catalog[len(ps.catalog)-1] = nil //
	ps.catalog = ps.catalog[:len(ps.catalog)-1]

	return true
}

//
func (ps *NodeCollection) Volume() int {
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
func (ps *NodeCollection) Catalog() []Node {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.catalog
}

//
//
//
func (ps *NodeCollection) Clone() []Node {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	c := make([]Node, len(ps.catalog))
	copy(c, ps.catalog)
	return c
}

//
func (ps *NodeCollection) ForEach(fn func(node Node)) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	for _, item := range ps.search {
		fn(item.node)
	}
}

//
func (ps *NodeCollection) Arbitrary() Node {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if len(ps.catalog) == 0 {
		return nil
	}

	return ps.catalog[engineseed.Int()%len(ps.catalog)]
}
