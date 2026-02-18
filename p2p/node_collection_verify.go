package p2p

import (
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/daemon"
)

//
type emulateNode struct {
	daemon.RootDaemon
	ip net.IP
	id ID
}

func (mp *emulateNode) PurgeHalt()               { mp.Halt() } //
func (mp *emulateNode) AttemptTransmit(Packet) bool    { return true }
func (mp *emulateNode) Transmit(Packet) bool       { return true }
func (mp *emulateNode) MemberDetails() MemberDetails       { return StandardMemberDetails{} }
func (mp *emulateNode) Status() LinkageState { return LinkageState{} }
func (mp *emulateNode) ID() ID                   { return mp.id }
func (mp *emulateNode) IsOutgoing() bool         { return false }
func (mp *emulateNode) IsDurable() bool       { return true }
func (mp *emulateNode) Get(s string) any         { return s }
func (mp *emulateNode) Set(string, any)          {}
func (mp *emulateNode) DistantIP() net.IP         { return mp.ip }
func (mp *emulateNode) SocketAddress() *NetLocation  { return nil }
func (mp *emulateNode) DistantAddress() net.Addr     { return &net.TCPAddr{IP: mp.ip, Port: 8800} }
func (mp *emulateNode) EndLink() error         { return nil }
func (mp *emulateNode) CollectionDeletionErrored()        {}
func (mp *emulateNode) FetchDeletionErrored() bool   { return false }

//
func newEmulateNode(ip net.IP) *emulateNode {
	if ip == nil {
		ip = net.IP{127, 0, 0, 1}
	}
	memberKey := MemberKey{PrivateKey: ed25519.GeneratePrivateKey()}
	return &emulateNode{
		ip: ip,
		id: memberKey.ID(),
	}
}

func VerifyNodeCollectionAppendDeleteOne(t *testing.T) {
	t.Parallel()

	nodeCollection := NewNodeCollection()

	var nodeCatalog []Node
	for i := 0; i < 5; i++ {
		p := newEmulateNode(net.IP{127, 0, 0, byte(i)})
		if err := nodeCollection.Add(p); err != nil {
			t.Error(err)
		}
		nodeCatalog = append(nodeCatalog, p)
	}

	n := len(nodeCatalog)
	//
	for i, nodeAtHead := range nodeCatalog {
		deleted := nodeCollection.Delete(nodeAtHead)
		assert.True(t, deleted)
		desireVolume := n - i - 1
		for j := 0; j < 2; j++ {
			assert.Equal(t, false, nodeCollection.Has(nodeAtHead.ID()), "REDACTED", i, j)
			assert.Equal(t, desireVolume, nodeCollection.Volume(), "REDACTED", i, j)
			//
			deleted := nodeCollection.Delete(nodeAtHead)
			assert.False(t, deleted)
		}
	}

	//
	//
	for _, node := range nodeCatalog {
		if err := nodeCollection.Add(node); err != nil {
			t.Error(err)
		}
	}

	//
	for i := n - 1; i >= 0; i-- {
		nodeAtTerminate := nodeCatalog[i]
		deleted := nodeCollection.Delete(nodeAtTerminate)
		assert.True(t, deleted)
		assert.Equal(t, false, nodeCollection.Has(nodeAtTerminate.ID()), "REDACTED", i)
		assert.Equal(t, i, nodeCollection.Volume(), "REDACTED", i)
	}
}

func VerifyNodeCollectionAppendDeleteNumerous(t *testing.T) {
	t.Parallel()
	nodeCollection := NewNodeCollection()

	nodes := []Node{}
	N := 100
	for i := 0; i < N; i++ {
		node := newEmulateNode(net.IP{127, 0, 0, byte(i)})
		if err := nodeCollection.Add(node); err != nil {
			t.Errorf("REDACTED")
		}
		if nodeCollection.Volume() != i+1 {
			t.Errorf("REDACTED")
		}
		nodes = append(nodes, node)
	}

	for i, node := range nodes {
		deleted := nodeCollection.Delete(node)
		assert.True(t, deleted)
		if nodeCollection.Has(node.ID()) {
			t.Errorf("REDACTED")
		}
		if nodeCollection.Volume() != len(nodes)-i-1 {
			t.Errorf("REDACTED")
		}
	}
}

func VerifyNodeCollectionAppendReplicated(t *testing.T) {
	t.Parallel()
	nodeCollection := NewNodeCollection()
	node := newEmulateNode(nil)

	n := 20
	faultsChannel := make(chan error)
	//
	//
	//
	//
	//
	for i := 0; i < n; i++ {
		go func() {
			faultsChannel <- nodeCollection.Add(node)
		}()
	}

	//
	faultsCount := make(map[string]int)
	for i := 0; i < n; i++ {
		err := <-faultsChannel

		switch err.(type) {
		case ErrRouterReplicatedNodeUID:
			faultsCount["REDACTED"]++
		default:
			faultsCount["REDACTED"]++
		}
	}

	//
	//
	desireErrNumber, acquiredErrNumber := n-1, faultsCount["REDACTED"]
	assert.Equal(t, desireErrNumber, acquiredErrNumber, "REDACTED")

	desireNullErrNumber, acquiredNullErrNumber := 1, faultsCount["REDACTED"]
	assert.Equal(t, desireNullErrNumber, acquiredNullErrNumber, "REDACTED")
}

func VerifyNodeCollectionFetch(t *testing.T) {
	t.Parallel()

	var (
		nodeCollection = NewNodeCollection()
		node    = newEmulateNode(nil)
	)

	assert.Nil(t, nodeCollection.Get(node.ID()), "REDACTED")

	if err := nodeCollection.Add(node); err != nil {
		t.Fatalf("REDACTED", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//
		//
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			possess, desire := nodeCollection.Get(node.ID()), node
			assert.Equal(t, possess, desire, "REDACTED", i, possess, desire)
		}(i)
	}
	wg.Wait()
}
