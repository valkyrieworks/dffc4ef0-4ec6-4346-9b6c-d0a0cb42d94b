package statusconnect

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"sort"

	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p"
)

//
type mirrorKey [sha256.Size]byte

//
type mirror struct {
	Level   uint64
	Layout   uint32
	Segments   uint32
	Digest     []byte
	Metainfo []byte

	validatedApplicationDigest []byte //
}

//
//
//
func (s *mirror) Key() mirrorKey {
	//
	digester := sha256.New()
	fmt.Fprintf(digester, "REDACTED", s.Level, s.Layout, s.Segments)
	digester.Write(s.Digest)
	digester.Write(s.Metainfo)
	var key mirrorKey
	copy(key[:], digester.Sum(nil))
	return key
}

//
type mirrorDepository struct {
	engineconnect.Lock
	mirrors     map[mirrorKey]*mirror
	mirrorNodes map[mirrorKey]map[p2p.ID]p2p.Node

	//
	layoutOrdinal map[uint32]map[mirrorKey]bool
	levelOrdinal map[uint64]map[mirrorKey]bool
	nodeOrdinal   map[p2p.ID]map[mirrorKey]bool

	//
	layoutDenylist   map[uint32]bool
	nodeDenylist     map[p2p.ID]bool
	mirrorDenylist map[mirrorKey]bool
}

//
func newMirrorDepository() *mirrorDepository {
	return &mirrorDepository{
		mirrors:          make(map[mirrorKey]*mirror),
		mirrorNodes:      make(map[mirrorKey]map[p2p.ID]p2p.Node),
		layoutOrdinal:        make(map[uint32]map[mirrorKey]bool),
		levelOrdinal:        make(map[uint64]map[mirrorKey]bool),
		nodeOrdinal:          make(map[p2p.ID]map[mirrorKey]bool),
		layoutDenylist:   make(map[uint32]bool),
		nodeDenylist:     make(map[p2p.ID]bool),
		mirrorDenylist: make(map[mirrorKey]bool),
	}
}

//
//
//
func (p *mirrorDepository) Add(node p2p.Node, mirror *mirror) (bool, error) {
	key := mirror.Key()

	p.Lock()
	defer p.Unlock()

	switch {
	case p.layoutDenylist[mirror.Layout]:
		return false, nil
	case p.nodeDenylist[node.ID()]:
		return false, nil
	case p.mirrorDenylist[key]:
		return false, nil
	case len(p.nodeOrdinal[node.ID()]) >= currentMirrors:
		return false, nil
	}

	if p.mirrorNodes[key] == nil {
		p.mirrorNodes[key] = make(map[p2p.ID]p2p.Node)
	}
	p.mirrorNodes[key][node.ID()] = node

	if p.nodeOrdinal[node.ID()] == nil {
		p.nodeOrdinal[node.ID()] = make(map[mirrorKey]bool)
	}
	p.nodeOrdinal[node.ID()][key] = true

	if p.mirrors[key] != nil {
		return false, nil
	}
	p.mirrors[key] = mirror

	if p.layoutOrdinal[mirror.Layout] == nil {
		p.layoutOrdinal[mirror.Layout] = make(map[mirrorKey]bool)
	}
	p.layoutOrdinal[mirror.Layout][key] = true

	if p.levelOrdinal[mirror.Level] == nil {
		p.levelOrdinal[mirror.Level] = make(map[mirrorKey]bool)
	}
	p.levelOrdinal[mirror.Level][key] = true

	return true, nil
}

//
func (p *mirrorDepository) Optimal() *mirror {
	rated := p.Rated()
	if len(rated) == 0 {
		return nil
	}
	return rated[0]
}

//
func (p *mirrorDepository) FetchNode(mirror *mirror) p2p.Node {
	nodes := p.FetchNodes(mirror)
	if len(nodes) == 0 {
		return nil
	}
	return nodes[rand.Intn(len(nodes))] //
}

//
func (p *mirrorDepository) FetchNodes(mirror *mirror) []p2p.Node {
	key := mirror.Key()
	p.Lock()
	defer p.Unlock()

	nodes := make([]p2p.Node, 0, len(p.mirrorNodes[key]))
	for _, node := range p.mirrorNodes[key] {
		nodes = append(nodes, node)
	}
	//
	sort.Slice(nodes, func(a int, b int) bool {
		return nodes[a].ID() < nodes[b].ID()
	})
	return nodes
}

//
//
//
func (p *mirrorDepository) Rated() []*mirror {
	p.Lock()
	defer p.Unlock()

	hopefuls := make([]*mirror, 0, len(p.mirrors))
	for key := range p.mirrors {
		hopefuls = append(hopefuls, p.mirrors[key])
	}

	sort.Slice(hopefuls, func(i, j int) bool {
		a := hopefuls[i]
		b := hopefuls[j]

		switch {
		case a.Level > b.Level:
			return true
		case a.Level < b.Level:
			return false
		case a.Layout > b.Layout:
			return true
		case a.Layout < b.Layout:
			return false
		case len(p.mirrorNodes[a.Key()]) > len(p.mirrorNodes[b.Key()]):
			return true
		default:
			return false
		}
	})

	return hopefuls
}

//
func (p *mirrorDepository) Decline(mirror *mirror) {
	key := mirror.Key()
	p.Lock()
	defer p.Unlock()

	p.mirrorDenylist[key] = true
	p.deleteMirror(key)
}

//
func (p *mirrorDepository) DeclineLayout(layout uint32) {
	p.Lock()
	defer p.Unlock()

	p.layoutDenylist[layout] = true
	for key := range p.layoutOrdinal[layout] {
		p.deleteMirror(key)
	}
}

//
func (p *mirrorDepository) DeclineNode(nodeUID p2p.ID) {
	if nodeUID == "REDACTED" {
		return
	}
	p.Lock()
	defer p.Unlock()

	p.deleteNode(nodeUID)
	p.nodeDenylist[nodeUID] = true
}

//
func (p *mirrorDepository) DeleteNode(nodeUID p2p.ID) {
	p.Lock()
	defer p.Unlock()
	p.deleteNode(nodeUID)
}

//
func (p *mirrorDepository) deleteNode(nodeUID p2p.ID) {
	for key := range p.nodeOrdinal[nodeUID] {
		delete(p.mirrorNodes[key], nodeUID)
		if len(p.mirrorNodes[key]) == 0 {
			p.deleteMirror(key)
		}
	}
	delete(p.nodeOrdinal, nodeUID)
}

//
func (p *mirrorDepository) deleteMirror(key mirrorKey) {
	mirror := p.mirrors[key]
	if mirror == nil {
		return
	}

	delete(p.mirrors, key)
	delete(p.layoutOrdinal[mirror.Layout], key)
	delete(p.levelOrdinal[mirror.Level], key)
	for nodeUID := range p.mirrorNodes[key] {
		delete(p.nodeOrdinal[nodeUID], key)
	}
	delete(p.mirrorNodes, key)
}
