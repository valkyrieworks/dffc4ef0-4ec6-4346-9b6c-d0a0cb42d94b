package statuschronize

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"sort"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

//
type imageToken [sha256.Size]byte

//
type image struct {
	Altitude   uint64
	Layout   uint32
	Segments   uint32
	Digest     []byte
	Attributes []byte

	reliablePlatformDigest []byte //
}

//
//
//
func (s *image) Key() imageToken {
	//
	digester := sha256.New()
	fmt.Fprintf(digester, "REDACTED", s.Altitude, s.Layout, s.Segments)
	digester.Write(s.Digest)
	digester.Write(s.Attributes)
	var key imageToken
	copy(key[:], digester.Sum(nil))
	return key
}

//
type imageHub struct {
	commitchronize.Exclusion
	images     map[imageToken]*image
	imageNodes map[imageToken]map[p2p.ID]p2p.Node

	//
	layoutOrdinal map[uint32]map[imageToken]bool
	altitudeOrdinal map[uint64]map[imageToken]bool
	nodeOrdinal   map[p2p.ID]map[imageToken]bool

	//
	layoutDenylist   map[uint32]bool
	nodeDenylist     map[p2p.ID]bool
	imageDenylist map[imageToken]bool
}

//
func freshImageHub() *imageHub {
	return &imageHub{
		images:          make(map[imageToken]*image),
		imageNodes:      make(map[imageToken]map[p2p.ID]p2p.Node),
		layoutOrdinal:        make(map[uint32]map[imageToken]bool),
		altitudeOrdinal:        make(map[uint64]map[imageToken]bool),
		nodeOrdinal:          make(map[p2p.ID]map[imageToken]bool),
		layoutDenylist:   make(map[uint32]bool),
		nodeDenylist:     make(map[p2p.ID]bool),
		imageDenylist: make(map[imageToken]bool),
	}
}

//
//
//
func (p *imageHub) Add(node p2p.Node, image *image) (bool, error) {
	key := image.Key()

	p.Lock()
	defer p.Unlock()

	switch {
	case p.layoutDenylist[image.Layout]:
		return false, nil
	case p.nodeDenylist[node.ID()]:
		return false, nil
	case p.imageDenylist[key]:
		return false, nil
	case len(p.nodeOrdinal[node.ID()]) >= currentImages:
		return false, nil
	}

	if p.imageNodes[key] == nil {
		p.imageNodes[key] = make(map[p2p.ID]p2p.Node)
	}
	p.imageNodes[key][node.ID()] = node

	if p.nodeOrdinal[node.ID()] == nil {
		p.nodeOrdinal[node.ID()] = make(map[imageToken]bool)
	}
	p.nodeOrdinal[node.ID()][key] = true

	if p.images[key] != nil {
		return false, nil
	}
	p.images[key] = image

	if p.layoutOrdinal[image.Layout] == nil {
		p.layoutOrdinal[image.Layout] = make(map[imageToken]bool)
	}
	p.layoutOrdinal[image.Layout][key] = true

	if p.altitudeOrdinal[image.Altitude] == nil {
		p.altitudeOrdinal[image.Altitude] = make(map[imageToken]bool)
	}
	p.altitudeOrdinal[image.Altitude][key] = true

	return true, nil
}

//
func (p *imageHub) Optimal() *image {
	ordered := p.Ordered()
	if len(ordered) == 0 {
		return nil
	}
	return ordered[0]
}

//
func (p *imageHub) ObtainNode(image *image) p2p.Node {
	nodes := p.ObtainNodes(image)
	if len(nodes) == 0 {
		return nil
	}
	return nodes[rand.Intn(len(nodes))] //
}

//
func (p *imageHub) ObtainNodes(image *image) []p2p.Node {
	key := image.Key()
	p.Lock()
	defer p.Unlock()

	nodes := make([]p2p.Node, 0, len(p.imageNodes[key]))
	for _, node := range p.imageNodes[key] {
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
func (p *imageHub) Ordered() []*image {
	p.Lock()
	defer p.Unlock()

	nominees := make([]*image, 0, len(p.images))
	for key := range p.images {
		nominees = append(nominees, p.images[key])
	}

	sort.Slice(nominees, func(i, j int) bool {
		a := nominees[i]
		b := nominees[j]

		switch {
		case a.Altitude > b.Altitude:
			return true
		case a.Altitude < b.Altitude:
			return false
		case a.Layout > b.Layout:
			return true
		case a.Layout < b.Layout:
			return false
		case len(p.imageNodes[a.Key()]) > len(p.imageNodes[b.Key()]):
			return true
		default:
			return false
		}
	})

	return nominees
}

//
func (p *imageHub) Decline(image *image) {
	key := image.Key()
	p.Lock()
	defer p.Unlock()

	p.imageDenylist[key] = true
	p.discardImage(key)
}

//
func (p *imageHub) DeclineLayout(layout uint32) {
	p.Lock()
	defer p.Unlock()

	p.layoutDenylist[layout] = true
	for key := range p.layoutOrdinal[layout] {
		p.discardImage(key)
	}
}

//
func (p *imageHub) DeclineNode(nodeUUID p2p.ID) {
	if nodeUUID == "REDACTED" {
		return
	}
	p.Lock()
	defer p.Unlock()

	p.discardNode(nodeUUID)
	p.nodeDenylist[nodeUUID] = true
}

//
func (p *imageHub) DiscardNode(nodeUUID p2p.ID) {
	p.Lock()
	defer p.Unlock()
	p.discardNode(nodeUUID)
}

//
func (p *imageHub) discardNode(nodeUUID p2p.ID) {
	for key := range p.nodeOrdinal[nodeUUID] {
		delete(p.imageNodes[key], nodeUUID)
		if len(p.imageNodes[key]) == 0 {
			p.discardImage(key)
		}
	}
	delete(p.nodeOrdinal, nodeUUID)
}

//
func (p *imageHub) discardImage(key imageToken) {
	image := p.images[key]
	if image == nil {
		return
	}

	delete(p.images, key)
	delete(p.layoutOrdinal[image.Layout], key)
	delete(p.altitudeOrdinal[image.Altitude], key)
	for nodeUUID := range p.imageNodes[key] {
		delete(p.nodeOrdinal[nodeUUID], key)
	}
	delete(p.imageNodes, key)
}
