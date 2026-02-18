package blockreplication

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"
	ctrng "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/kinds"
)

func init() {
	peerTimeout = 2 * time.Second
}

type testPeer struct {
	id        p2p.ID
	base      int64
	height    int64
	inputChan chan inputData //
	malicious bool
}

type inputData struct {
	t       *testing.T
	pool    *BlockPool
	request BlockRequest
}

//
const (
	MaliciousLie               = 5 //
	BlackholeSize              = 3 //
	MaliciousTestMaximumLength = 5 * time.Minute
)

func (p testPeer) runInputRoutine() {
	go func() {
		for input := range p.inputChan {
			p.simulateInput(input)
		}
	}()
}

//
func (p testPeer) simulateInput(input inputData) {
	block := &kinds.Block{Header: kinds.Header{Height: input.request.Height}, LastCommit: &kinds.Commit{}} //
	extCommit := &kinds.ExtendedCommit{
		Height: input.request.Height,
	}
	//
	if p.malicious {
		realHeight := p.height - MaliciousLie
		//
		if input.request.Height > realHeight {
			//
			block.LastCommit = nil //
			//
			if input.request.Height <= realHeight+BlackholeSize {
				input.pool.RedoRequestFrom(input.request.Height, p.id)
				return
			}
		}
	}
	err := input.pool.AddBlock(input.request.PeerID, block, extCommit, 123)
	require.NoError(input.t, err)
	//
	//
	//
	//
}

type testPeers map[p2p.ID]*testPeer

func (ps testPeers) start() {
	for _, v := range ps {
		v.runInputRoutine()
	}
}

func (ps testPeers) stop() {
	for _, v := range ps {
		close(v.inputChan)
	}
}

func makePeers(numPeers int, minHeight, maxHeight int64) testPeers {
	peers := make(testPeers, numPeers)
	for i := 0; i < numPeers; i++ {
		peerID := p2p.ID(ctrng.Str(12))
		height := minHeight + ctrng.Int63n(maxHeight-minHeight)
		base := minHeight + int64(i)
		if base > height {
			base = height
		}
		peers[peerID] = &testPeer{peerID, base, height, make(chan inputData, 10), false}
	}
	return peers
}

func TestBlockPoolBasic(t *testing.T) {
	var (
		start      = int64(42)
		peers      = makePeers(10, start, 1000)
		errorsCh   = make(chan peerError)
		requestsCh = make(chan BlockRequest)
	)
	pool := NewBlockPool(start, requestsCh, errorsCh)
	pool.SetLogger(log.TestingLogger())

	err := pool.Start()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := pool.Stop(); err != nil {
			t.Error(err)
		}
	})

	peers.start()
	defer peers.stop()

	//
	go func() {
		for _, peer := range peers {
			pool.SetPeerRange(peer.id, peer.base, peer.height)
		}
	}()

	//
	go func() {
		for {
			if !pool.IsRunning() {
				return
			}
			first, second, _ := pool.PeekTwoBlocks()
			if first != nil && second != nil {
				pool.PopRequest()
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()

	//
	for {
		select {
		case err := <-errorsCh:
			t.Error(err)
		case request := <-requestsCh:
			t.Logf("REDACTED", request)
			if request.Height == 300 {
				return //
			}

			peers[request.PeerID].inputChan <- inputData{t, pool, request}
		}
	}
}

func TestBlockPoolTimeout(t *testing.T) {
	var (
		start      = int64(42)
		peers      = makePeers(10, start, 1000)
		errorsCh   = make(chan peerError)
		requestsCh = make(chan BlockRequest)
	)

	pool := NewBlockPool(start, requestsCh, errorsCh)
	pool.SetLogger(log.TestingLogger())
	err := pool.Start()
	if err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		if err := pool.Stop(); err != nil {
			t.Error(err)
		}
	})

	for _, peer := range peers {
		t.Logf("REDACTED", peer.id)
	}

	//
	go func() {
		for _, peer := range peers {
			pool.SetPeerRange(peer.id, peer.base, peer.height)
		}
	}()

	//
	go func() {
		for {
			if !pool.IsRunning() {
				return
			}
			first, second, _ := pool.PeekTwoBlocks()
			if first != nil && second != nil {
				pool.PopRequest()
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()

	//
	counter := 0
	timedOut := map[p2p.ID]struct{}{}
	for {
		select {
		case err := <-errorsCh:
			t.Log(err)
			//
			if _, ok := timedOut[err.peerID]; !ok {
				counter++
				if counter == len(peers) {
					return //
				}
			}
		case request := <-requestsCh:
			t.Logf("REDACTED", request)
		}
	}
}

func TestBlockPoolRemovePeer(t *testing.T) {
	peers := make(testPeers, 10)
	for i := 0; i < 10; i++ {
		peerID := p2p.ID(fmt.Sprintf("REDACTED", i+1))
		height := int64(i + 1)
		peers[peerID] = &testPeer{peerID, 0, height, make(chan inputData), false}
	}
	requestsCh := make(chan BlockRequest)
	errorsCh := make(chan peerError)

	pool := NewBlockPool(1, requestsCh, errorsCh)
	pool.SetLogger(log.TestingLogger())
	err := pool.Start()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := pool.Stop(); err != nil {
			t.Error(err)
		}
	})

	//
	for peerID, peer := range peers {
		pool.SetPeerRange(peerID, peer.base, peer.height)
	}
	assert.EqualValues(t, 10, pool.MaxPeerHeight())

	//
	assert.NotPanics(t, func() { pool.RemovePeer(p2p.ID("REDACTED")) })

	//
	pool.RemovePeer(p2p.ID("REDACTED"))
	assert.EqualValues(t, 9, pool.MaxPeerHeight())

	//
	for peerID := range peers {
		pool.RemovePeer(peerID)
	}

	assert.EqualValues(t, 0, pool.MaxPeerHeight())
}

func TestBlockPoolMaliciousNode(t *testing.T) {
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	const InitialHeight = 7
	peers := testPeers{
		p2p.ID("REDACTED"):  &testPeer{p2p.ID("REDACTED"), 1, InitialHeight, make(chan inputData), false},
		p2p.ID("REDACTED"):   &testPeer{p2p.ID("REDACTED"), 1, InitialHeight + MaliciousLie, make(chan inputData), true},
		p2p.ID("REDACTED"): &testPeer{p2p.ID("REDACTED"), 1, InitialHeight, make(chan inputData), false},
	}
	errorsCh := make(chan peerError)
	requestsCh := make(chan BlockRequest)

	pool := NewBlockPool(1, requestsCh, errorsCh)
	pool.SetLogger(log.TestingLogger())

	err := pool.Start()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := pool.Stop(); err != nil {
			t.Error(err)
		}
	})

	peers.start()
	t.Cleanup(func() { peers.stop() })

	//
	go func() {
		//
		for _, peer := range peers {
			pool.SetPeerRange(peer.id, peer.base, peer.height)
		}

		ticker := time.NewTicker(1 * time.Second) //
		defer ticker.Stop()
		for {
			select {
			case <-pool.Quit():
				return
			case <-ticker.C:
				for _, peer := range peers {
					peer.height++                                      //
					pool.SetPeerRange(peer.id, peer.base, peer.height) //
				}
			}
		}
	}()

	//
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond) //
		defer ticker.Stop()
		for {
			select {
			case <-pool.Quit():
				return
			case <-ticker.C:
				first, second, _ := pool.PeekTwoBlocks()
				if first != nil && second != nil {
					if second.LastCommit == nil {
						//
						pool.RemovePeerAndRedoAllPeerRequests(second.Height)
					} else {
						pool.PopRequest()
					}
				}
			}
		}
	}()

	testTicker := time.NewTicker(200 * time.Millisecond) //
	t.Cleanup(func() { testTicker.Stop() })

	bannedOnce := false //
	startTime := time.Now()

	//
	for {
		select {
		case err := <-errorsCh:
			t.Error(err)
		case request := <-requestsCh:
			//
			peers[request.PeerID].inputChan <- inputData{t, pool, request}
		case <-testTicker.C:
			banned := pool.IsPeerBanned("REDACTED")
			bannedOnce = bannedOnce || banned //
			caughtUp := pool.IsCaughtUp()
			//
			if caughtUp && bannedOnce {
				t.Logf("REDACTED")
				return
			}
			//
			require.False(t, caughtUp, "REDACTED")
			//
			require.True(t, time.Since(startTime) < MaliciousTestMaximumLength, "REDACTED")
		}
	}
}

func TestBlockPoolMaliciousNodeMaxInt64(t *testing.T) {
	//
	//
	//
	//
	//
	//
	//
	//
	//
	const initialHeight = 7
	peers := testPeers{
		p2p.ID("REDACTED"):  &testPeer{p2p.ID("REDACTED"), 1, initialHeight, make(chan inputData), false},
		p2p.ID("REDACTED"):   &testPeer{p2p.ID("REDACTED"), 1, math.MaxInt64, make(chan inputData), true},
		p2p.ID("REDACTED"): &testPeer{p2p.ID("REDACTED"), 1, initialHeight, make(chan inputData), false},
	}
	errorsCh := make(chan peerError, 3)
	requestsCh := make(chan BlockRequest)

	pool := NewBlockPool(1, requestsCh, errorsCh)
	pool.SetLogger(log.TestingLogger())

	err := pool.Start()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := pool.Stop(); err != nil {
			t.Error(err)
		}
	})

	peers.start()
	t.Cleanup(func() { peers.stop() })

	//
	go func() {
		//
		for _, peer := range peers {
			pool.SetPeerRange(peer.id, peer.base, peer.height)
		}

		//
		peers["REDACTED"].height = initialHeight
		pool.SetPeerRange(p2p.ID("REDACTED"), 1, initialHeight)

		ticker := time.NewTicker(1 * time.Second) //
		defer ticker.Stop()
		for {
			select {
			case <-pool.Quit():
				return
			case <-ticker.C:
				for _, peer := range peers {
					peer.height++                                      //
					pool.SetPeerRange(peer.id, peer.base, peer.height) //
				}
			}
		}
	}()

	//
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond) //
		defer ticker.Stop()
		for {
			select {
			case <-pool.Quit():
				return
			case <-ticker.C:
				first, second, _ := pool.PeekTwoBlocks()
				if first != nil && second != nil {
					if second.LastCommit == nil {
						//
						pool.RemovePeerAndRedoAllPeerRequests(second.Height)
					} else {
						pool.PopRequest()
					}
				}
			}
		}
	}()

	testTicker := time.NewTicker(200 * time.Millisecond) //
	t.Cleanup(func() { testTicker.Stop() })

	bannedOnce := false //
	startTime := time.Now()

	//
	for {
		select {
		case err := <-errorsCh:
			if err.peerID == "REDACTED" { //
				t.Log(err)
			} else {
				t.Error(err)
			}
		case request := <-requestsCh:
			//
			peers[request.PeerID].inputChan <- inputData{t, pool, request}
		case <-testTicker.C:
			banned := pool.IsPeerBanned("REDACTED")
			bannedOnce = bannedOnce || banned //
			caughtUp := pool.IsCaughtUp()
			//
			if caughtUp && bannedOnce {
				t.Logf("REDACTED")
				return
			}
			//
			require.False(t, caughtUp, "REDACTED")
			//
			require.True(t, time.Since(startTime) < MaliciousTestMaximumLength, "REDACTED")
		}
	}
}
