package p2p

import (
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

//
type simulateNode struct {
	facility.FoundationFacility
	ip net.IP
	id ID
}

func (mp *simulateNode) PurgeHalt()               { mp.Halt() } //
func (mp *simulateNode) AttemptTransmit(Wrapper) bool    { return true }
func (mp *simulateNode) Transmit(Wrapper) bool       { return true }
func (mp *simulateNode) PeerDetails() PeerDetails       { return FallbackPeerDetails{} }
func (mp *simulateNode) Condition() LinkageCondition { return LinkageCondition{} }
func (mp *simulateNode) ID() ID                   { return mp.id }
func (mp *simulateNode) EqualsOutgoing() bool         { return false }
func (mp *simulateNode) EqualsEnduring() bool       { return true }
func (mp *simulateNode) Get(s string) any         { return s }
func (mp *simulateNode) Set(string, any)          {}
func (mp *simulateNode) DistantINET() net.IP         { return mp.ip }
func (mp *simulateNode) PortLocation() *NetworkLocator  { return nil }
func (mp *simulateNode) DistantLocation() net.Addr     { return &net.TCPAddr{IP: mp.ip, Port: 8800} }
func (mp *simulateNode) ShutdownLink() error         { return nil }
func (mp *simulateNode) AssignDeletionUnsuccessful()        {}
func (mp *simulateNode) ObtainDeletionUnsuccessful() bool   { return false }

//
func freshSimulateNode(ip net.IP) *simulateNode {
	if ip == nil {
		ip = net.IP{127, 0, 0, 1}
	}
	peerToken := PeerToken{PrivateToken: edwards25519.ProducePrivateToken()}
	return &simulateNode{
		ip: ip,
		id: peerToken.ID(),
	}
}

func VerifyNodeAssignAppendDiscardSingle(t *testing.T) {
	t.Parallel()

	nodeAssign := FreshNodeAssign()

	var nodeCatalog []Node
	for i := 0; i < 5; i++ {
		p := freshSimulateNode(net.IP{127, 0, 0, byte(i)})
		if err := nodeAssign.Add(p); err != nil {
			t.Error(err)
		}
		nodeCatalog = append(nodeCatalog, p)
	}

	n := len(nodeCatalog)
	//
	for i, nodeLocatedLeading := range nodeCatalog {
		discarded := nodeAssign.Discard(nodeLocatedLeading)
		assert.True(t, discarded)
		desireExtent := n - i - 1
		for j := 0; j < 2; j++ {
			assert.Equal(t, false, nodeAssign.Has(nodeLocatedLeading.ID()), "REDACTED", i, j)
			assert.Equal(t, desireExtent, nodeAssign.Extent(), "REDACTED", i, j)
			//
			discarded := nodeAssign.Discard(nodeLocatedLeading)
			assert.False(t, discarded)
		}
	}

	//
	//
	for _, node := range nodeCatalog {
		if err := nodeAssign.Add(node); err != nil {
			t.Error(err)
		}
	}

	//
	for i := n - 1; i >= 0; i-- {
		nodeLocatedTerminate := nodeCatalog[i]
		discarded := nodeAssign.Discard(nodeLocatedTerminate)
		assert.True(t, discarded)
		assert.Equal(t, false, nodeAssign.Has(nodeLocatedTerminate.ID()), "REDACTED", i)
		assert.Equal(t, i, nodeAssign.Extent(), "REDACTED", i)
	}
}

func VerifyNodeAssignAppendDiscardMultiple(t *testing.T) {
	t.Parallel()
	nodeAssign := FreshNodeAssign()

	nodes := []Node{}
	N := 100
	for i := 0; i < N; i++ {
		node := freshSimulateNode(net.IP{127, 0, 0, byte(i)})
		if err := nodeAssign.Add(node); err != nil {
			t.Errorf("REDACTED")
		}
		if nodeAssign.Extent() != i+1 {
			t.Errorf("REDACTED")
		}
		nodes = append(nodes, node)
	}

	for i, node := range nodes {
		discarded := nodeAssign.Discard(node)
		assert.True(t, discarded)
		if nodeAssign.Has(node.ID()) {
			t.Errorf("REDACTED")
		}
		if nodeAssign.Extent() != len(nodes)-i-1 {
			t.Errorf("REDACTED")
		}
	}
}

func VerifyNodeAssignAppendReplicated(t *testing.T) {
	t.Parallel()
	nodeAssign := FreshNodeAssign()
	node := freshSimulateNode(nil)

	n := 20
	errorsChn := make(chan error)
	//
	//
	//
	//
	//
	for i := 0; i < n; i++ {
		go func() {
			errorsChn <- nodeAssign.Add(node)
		}()
	}

	//
	errorsCalculation := make(map[string]int)
	for i := 0; i < n; i++ {
		err := <-errorsChn

		switch err.(type) {
		case FaultRouterReplicatedNodeUUID:
			errorsCalculation["REDACTED"]++
		default:
			errorsCalculation["REDACTED"]++
		}
	}

	//
	//
	desireFaultTally, attainedFaultTally := n-1, errorsCalculation["REDACTED"]
	assert.Equal(t, desireFaultTally, attainedFaultTally, "REDACTED")

	desireVoidFaultTally, attainedVoidFaultTally := 1, errorsCalculation["REDACTED"]
	assert.Equal(t, desireVoidFaultTally, attainedVoidFaultTally, "REDACTED")
}

func VerifyNodeAssignObtain(t *testing.T) {
	t.Parallel()

	var (
		nodeAssign = FreshNodeAssign()
		node    = freshSimulateNode(nil)
	)

	assert.Nil(t, nodeAssign.Get(node.ID()), "REDACTED")

	if err := nodeAssign.Add(node); err != nil {
		t.Fatalf("REDACTED", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//
		//
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			possess, desire := nodeAssign.Get(node.ID()), node
			assert.Equal(t, possess, desire, "REDACTED", i, possess, desire)
		}(i)
	}
	wg.Wait()
}
