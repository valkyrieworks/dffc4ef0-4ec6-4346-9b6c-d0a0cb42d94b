package blockreplication

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"sync/atomic"
	"time"

	stream "github.com/valkyrieworks/utils/velocity"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/facility"
	ctsync "github.com/valkyrieworks/utils/alignment"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/kinds"
	cttime "github.com/valkyrieworks/kinds/moment"
)

/**
s
0
s
e
s
B
s

n
*/

const (
	requestIntervalMS         = 2
	maxPendingRequestsPerPeer = 20
	requestRetrySeconds       = 30

	//
	//
	//
	//
	//
	//
	//
	minRecvRate = 128 * 1024 //

	//
	//
	//
	peerConnWait = 3 * time.Second

	//
	//
	//
	minBlocksForSingleRequest = 50
)

var peerTimeout = 15 * time.Second //

/**
.
s
.
.

l
e
r
*/

//
type BlockPool struct {
	facility.BaseService
	startTime   time.Time
	startHeight int64

	mtx ctsync.Mutex
	//
	requesters map[int64]*bpRequester
	height     int64 //
	//
	peers         map[p2p.ID]*bpPeer
	bannedPeers   map[p2p.ID]time.Time
	sortedPeers   []*bpPeer //
	maxPeerHeight int64     //

	//
	numPending int32 //

	requestsCh chan<- BlockRequest
	errorsCh   chan<- peerError
}

//
//
type BlockRequest struct {
	Height int64
	PeerID p2p.ID
}

//
//
func NewBlockPool(start int64, requestsCh chan<- BlockRequest, errorsCh chan<- peerError) *BlockPool {
	bp := &BlockPool{
		peers:       make(map[p2p.ID]*bpPeer),
		bannedPeers: make(map[p2p.ID]time.Time),
		requesters:  make(map[int64]*bpRequester),
		height:      start,
		startHeight: start,
		numPending:  0,

		requestsCh: requestsCh,
		errorsCh:   errorsCh,
	}
	bp.BaseService = *facility.NewBaseService(nil, "REDACTED", bp)
	return bp
}

//
//
func (pool *BlockPool) OnStart() error {
	pool.startTime = time.Now()
	go pool.makeRequestersRoutine()
	return nil
}

//
func (pool *BlockPool) makeRequestersRoutine() {
	for {
		if !pool.IsRunning() {
			return
		}

		//
		//
		if time.Since(pool.startTime) < peerConnWait {
			//
			sleepDuration := peerConnWait - time.Since(pool.startTime)
			time.Sleep(sleepDuration)
		}

		pool.mtx.Lock()
		var (
			maxRequestersCreated = len(pool.requesters) >= len(pool.peers)*maxPendingRequestsPerPeer

			nextHeight           = pool.height + int64(len(pool.requesters))
			maxPeerHeightReached = nextHeight > pool.maxPeerHeight
		)
		pool.mtx.Unlock()

		switch {
		case maxRequestersCreated: //
			time.Sleep(requestIntervalMS * time.Millisecond)
			pool.removeTimedoutPeers()
		case maxPeerHeightReached: //
			time.Sleep(requestIntervalMS * time.Millisecond)
		default:
			//
			pool.makeNextRequester(nextHeight)
			//
			time.Sleep(requestIntervalMS * time.Millisecond)
		}
	}
}

func (pool *BlockPool) removeTimedoutPeers() {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	for _, peer := range pool.peers {
		if !peer.didTimeout && peer.numPending > 0 {
			curRate := peer.recvMonitor.Status().CurRate
			//
			if curRate != 0 && curRate < minRecvRate {
				err := errors.New("REDACTED")
				pool.sendError(err, peer.id)
				pool.Logger.Error("REDACTED", "REDACTED", peer.id,
					"REDACTED", err,
					"REDACTED", fmt.Sprintf("REDACTED", curRate/1024),
					"REDACTED", fmt.Sprintf("REDACTED", minRecvRate/1024))
				peer.didTimeout = true
			}

			peer.curRate = curRate
		}

		if peer.didTimeout {
			pool.removePeer(peer.id)
		}
	}

	for peerID := range pool.bannedPeers {
		if !pool.isPeerBanned(peerID) {
			delete(pool.bannedPeers, peerID)
		}
	}

	pool.sortPeers()
}

//
//
func (pool *BlockPool) GetStatus() (height int64, numPending int32, lenRequesters int) {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	return pool.height, atomic.LoadInt32(&pool.numPending), len(pool.requesters)
}

//
//
func (pool *BlockPool) IsCaughtUp() bool {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	//
	if len(pool.peers) == 0 {
		pool.Logger.Debug("REDACTED")
		return false
	}

	//
	//
	//
	//
	//
	receivedBlockOrTimedOut := pool.height > 0 || time.Since(pool.startTime) > 5*time.Second
	ourChainIsLongestAmongPeers := pool.maxPeerHeight == 0 || pool.height >= (pool.maxPeerHeight-1)
	isCaughtUp := receivedBlockOrTimedOut && ourChainIsLongestAmongPeers
	return isCaughtUp
}

//
//
//
//
//
//
//
func (pool *BlockPool) PeekTwoBlocks() (first, second *kinds.Block, firstExtCommit *kinds.ExtendedCommit) {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	if r := pool.requesters[pool.height]; r != nil {
		first = r.getBlock()
		firstExtCommit = r.getExtendedCommit()
	}
	if r := pool.requesters[pool.height+1]; r != nil {
		second = r.getBlock()
	}
	return
}

//
func (pool *BlockPool) PopRequest() {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	r := pool.requesters[pool.height]
	if r == nil {
		panic(fmt.Sprintf("REDACTED", pool.height))
	}

	if err := r.Stop(); err != nil {
		pool.Logger.Error("REDACTED", "REDACTED", err)
	}
	delete(pool.requesters, pool.height)
	pool.height++

	//
	//
	for i := int64(0); i < minBlocksForSingleRequest && i < int64(len(pool.requesters)); i++ {
		pool.requesters[pool.height+i].newHeight(pool.height)
	}
}

//
//
//
func (pool *BlockPool) RemovePeerAndRedoAllPeerRequests(height int64) p2p.ID {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	request := pool.requesters[height]
	peerID := request.gotBlockFromPeerID()
	//
	pool.removePeer(peerID)
	pool.banPeer(peerID)
	return peerID
}

//
//
func (pool *BlockPool) RedoRequestFrom(height int64, peerID p2p.ID) {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	if requester, ok := pool.requesters[height]; ok { //
		if requester.didRequestFrom(peerID) { //
			requester.redo(peerID)
		}
	}
}

//
func (pool *BlockPool) RedoRequest(height int64) p2p.ID {
	return pool.RemovePeerAndRedoAllPeerRequests(height)
}

//
//
//
//
//
//
//
//
func (pool *BlockPool) AddBlock(peerID p2p.ID, block *kinds.Block, extCommit *kinds.ExtendedCommit, blockSize int) error {
	if extCommit != nil && block.Height != extCommit.Height {
		err := fmt.Errorf("REDACTED", block.Height, extCommit.Height)
		//
		pool.sendError(err, peerID)
		return err
	}

	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	requester := pool.requesters[block.Height]
	if requester == nil {
		//
		//
		//
		//
		if block.Height > pool.height || block.Height < pool.startHeight {
			err := fmt.Errorf("REDACTED",
				block.Height, pool.height, pool.startHeight)
			pool.sendError(err, peerID)
			return err
		}

		return fmt.Errorf("REDACTED", block.Height, peerID)
	}

	if !requester.setBlock(block, extCommit, peerID) {
		err := fmt.Errorf("REDACTED", block.Height, requester.requestedFrom(), peerID)
		pool.sendError(err, peerID)
		return err
	}

	atomic.AddInt32(&pool.numPending, -1)
	peer := pool.peers[peerID]
	if peer != nil {
		peer.decrPending(blockSize)
	}

	return nil
}

//
func (pool *BlockPool) Height() int64 {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()
	return pool.height
}

//
func (pool *BlockPool) MaxPeerHeight() int64 {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()
	return pool.maxPeerHeight
}

//
func (pool *BlockPool) SetPeerRange(peerID p2p.ID, base int64, height int64) {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	peer := pool.peers[peerID]
	if peer != nil {
		if base < peer.base || height < peer.height {
			pool.Logger.Info(
				"REDACTED",
				"REDACTED", peerID,
				"REDACTED", height,
				"REDACTED", base,
				"REDACTED", peer.height,
				"REDACTED", peer.base,
			)

			//
			pool.removePeer(peerID)
			pool.banPeer(peerID)

			return
		}
		peer.base = base
		peer.height = height
	} else {
		if pool.isPeerBanned(peerID) {
			pool.Logger.Debug("REDACTED", "REDACTED", peerID)
			return
		}
		peer = newBPPeer(pool, peerID, base, height)
		peer.setLogger(pool.Logger.With("REDACTED", peerID))
		pool.peers[peerID] = peer
		//
		//
		pool.sortedPeers = append([]*bpPeer{peer}, pool.sortedPeers...)
	}

	if height > pool.maxPeerHeight {
		pool.maxPeerHeight = height
	}
}

//
//
func (pool *BlockPool) RemovePeer(peerID p2p.ID) {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	pool.removePeer(peerID)
}

//
func (pool *BlockPool) removePeer(peerID p2p.ID) {
	for _, requester := range pool.requesters {
		if requester.didRequestFrom(peerID) {
			requester.redo(peerID)
		}
	}

	peer, ok := pool.peers[peerID]
	if ok {
		if peer.timeout != nil {
			peer.timeout.Stop()
		}

		delete(pool.peers, peerID)
		for i, p := range pool.sortedPeers {
			if p.id == peerID {
				pool.sortedPeers = append(pool.sortedPeers[:i], pool.sortedPeers[i+1:]...)
				break
			}
		}

		//
		//
		if peer.height == pool.maxPeerHeight {
			pool.updateMaxPeerHeight()
		}
	}
}

//
func (pool *BlockPool) updateMaxPeerHeight() {
	var max int64
	for _, peer := range pool.peers {
		if peer.height > max {
			max = peer.height
		}
	}
	pool.maxPeerHeight = max
}

//
func (pool *BlockPool) IsPeerBanned(peerID p2p.ID) bool {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()
	return pool.isPeerBanned(peerID)
}

//
func (pool *BlockPool) isPeerBanned(peerID p2p.ID) bool {
	//
	return time.Since(pool.bannedPeers[peerID]) < time.Second*60
}

//
func (pool *BlockPool) banPeer(peerID p2p.ID) {
	pool.Logger.Debug("REDACTED", peerID)
	pool.bannedPeers[peerID] = cttime.Now()
}

//
//
func (pool *BlockPool) pickIncrAvailablePeer(height int64, excludePeerID p2p.ID) *bpPeer {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	for _, peer := range pool.sortedPeers {
		if peer.id == excludePeerID {
			continue
		}
		if peer.didTimeout {
			pool.removePeer(peer.id)
			continue
		}
		if peer.numPending >= maxPendingRequestsPerPeer {
			continue
		}
		if height < peer.base || height > peer.height {
			continue
		}
		peer.incrPending()
		return peer
	}

	return nil
}

//
//
//
func (pool *BlockPool) sortPeers() {
	sort.Slice(pool.sortedPeers, func(i, j int) bool {
		return pool.sortedPeers[i].curRate > pool.sortedPeers[j].curRate
	})
}

func (pool *BlockPool) makeNextRequester(nextHeight int64) {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	request := newBPRequester(pool, nextHeight)

	pool.requesters[nextHeight] = request
	atomic.AddInt32(&pool.numPending, 1)

	if err := request.Start(); err != nil {
		request.Logger.Error("REDACTED", "REDACTED", err)
	}
}

//
func (pool *BlockPool) sendRequest(height int64, peerID p2p.ID) {
	if !pool.IsRunning() {
		return
	}
	pool.requestsCh <- BlockRequest{height, peerID}
}

//
func (pool *BlockPool) sendError(err error, peerID p2p.ID) {
	if !pool.IsRunning() {
		return
	}
	pool.errorsCh <- peerError{err, peerID}
}

//
//
//
func (pool *BlockPool) debug() string {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	str := "REDACTED"
	nextHeight := pool.height + int64(len(pool.requesters))
	for h := pool.height; h < nextHeight; h++ {
		if pool.requesters[h] == nil {
			str += fmt.Sprintf("REDACTED", h)
		} else {
			str += fmt.Sprintf("REDACTED", h)
			str += fmt.Sprintf("REDACTED", pool.requesters[h].block != nil)
			str += fmt.Sprintf("REDACTED", pool.requesters[h].extCommit != nil)
		}
	}
	return str
}

//

type bpPeer struct {
	didTimeout  bool
	curRate     int64
	numPending  int32
	height      int64
	base        int64
	pool        *BlockPool
	id          p2p.ID
	recvMonitor *stream.Monitor

	timeout *time.Timer

	logger log.Logger
}

func newBPPeer(pool *BlockPool, peerID p2p.ID, base int64, height int64) *bpPeer {
	peer := &bpPeer{
		pool:       pool,
		id:         peerID,
		base:       base,
		height:     height,
		numPending: 0,
		logger:     log.NewNopLogger(),
	}
	return peer
}

func (peer *bpPeer) setLogger(l log.Logger) {
	peer.logger = l
}

func (peer *bpPeer) resetMonitor() {
	peer.recvMonitor = stream.New(time.Second, time.Second*40)
	initialValue := float64(minRecvRate) * math.E
	peer.recvMonitor.SetREMA(initialValue)
}

func (peer *bpPeer) resetTimeout() {
	if peer.timeout == nil {
		peer.timeout = time.AfterFunc(peerTimeout, peer.onTimeout)
	} else {
		peer.timeout.Reset(peerTimeout)
	}
}

func (peer *bpPeer) incrPending() {
	if peer.numPending == 0 {
		peer.resetMonitor()
		peer.resetTimeout()
	}
	peer.numPending++
}

func (peer *bpPeer) decrPending(recvSize int) {
	peer.numPending--
	if peer.numPending == 0 {
		peer.timeout.Stop()
	} else {
		peer.recvMonitor.Update(recvSize)
		peer.resetTimeout()
	}
}

func (peer *bpPeer) onTimeout() {
	peer.pool.mtx.Lock()
	defer peer.pool.mtx.Unlock()

	peer.pool.sendError(ErrPeerTimeout, peer.id)
	peer.logger.Error("REDACTED", "REDACTED", ErrPeerTimeout, "REDACTED", peerTimeout)
	peer.didTimeout = true
}

//

//
//
//
//
//
//
//
//
type bpRequester struct {
	facility.BaseService

	pool        *BlockPool
	height      int64
	gotBlockCh  chan struct{}
	redoCh      chan p2p.ID //
	newHeightCh chan int64

	mtx          ctsync.Mutex
	peerID       p2p.ID
	secondPeerID p2p.ID //
	gotBlockFrom p2p.ID
	block        *kinds.Block
	extCommit    *kinds.ExtendedCommit
}

func newBPRequester(pool *BlockPool, height int64) *bpRequester {
	bpr := &bpRequester{
		pool:        pool,
		height:      height,
		gotBlockCh:  make(chan struct{}, 1),
		redoCh:      make(chan p2p.ID, 1),
		newHeightCh: make(chan int64, 1),

		peerID:       "REDACTED",
		secondPeerID: "REDACTED",
		block:        nil,
	}
	bpr.BaseService = *facility.NewBaseService(nil, "REDACTED", bpr)
	return bpr
}

func (bpr *bpRequester) OnStart() error {
	go bpr.requestRoutine()
	return nil
}

//
func (bpr *bpRequester) setBlock(block *kinds.Block, extCommit *kinds.ExtendedCommit, peerID p2p.ID) bool {
	bpr.mtx.Lock()
	if bpr.peerID != peerID && bpr.secondPeerID != peerID {
		bpr.mtx.Unlock()
		return false
	}
	if bpr.block != nil {
		bpr.mtx.Unlock()
		return true //
	}

	bpr.block = block
	bpr.extCommit = extCommit
	bpr.gotBlockFrom = peerID
	bpr.mtx.Unlock()

	select {
	case bpr.gotBlockCh <- struct{}{}:
	default:
	}
	return true
}

func (bpr *bpRequester) getBlock() *kinds.Block {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.block
}

func (bpr *bpRequester) getExtendedCommit() *kinds.ExtendedCommit {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.extCommit
}

//
func (bpr *bpRequester) requestedFrom() []p2p.ID {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	peerIDs := make([]p2p.ID, 0, 2)
	if bpr.peerID != "REDACTED" {
		peerIDs = append(peerIDs, bpr.peerID)
	}
	if bpr.secondPeerID != "REDACTED" {
		peerIDs = append(peerIDs, bpr.secondPeerID)
	}
	return peerIDs
}

//
func (bpr *bpRequester) didRequestFrom(peerID p2p.ID) bool {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.peerID == peerID || bpr.secondPeerID == peerID
}

//
func (bpr *bpRequester) gotBlockFromPeerID() p2p.ID {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.gotBlockFrom
}

//
func (bpr *bpRequester) reset(peerID p2p.ID) (removedBlock bool) {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()

	//
	if bpr.gotBlockFrom == peerID {
		bpr.block = nil
		bpr.extCommit = nil
		bpr.gotBlockFrom = "REDACTED"
		removedBlock = true
		atomic.AddInt32(&bpr.pool.numPending, 1)
	}

	if bpr.peerID == peerID {
		bpr.peerID = "REDACTED"
	} else {
		bpr.secondPeerID = "REDACTED"
	}

	return removedBlock
}

//
//
//
func (bpr *bpRequester) redo(peerID p2p.ID) {
	select {
	case bpr.redoCh <- peerID:
	default:
	}
}

func (bpr *bpRequester) pickPeerAndSendRequest() {
	bpr.mtx.Lock()
	secondPeerID := bpr.secondPeerID
	bpr.mtx.Unlock()

	var peer *bpPeer
PICK_PEER_LOOP:
	for {
		if !bpr.IsRunning() || !bpr.pool.IsRunning() {
			return
		}
		peer = bpr.pool.pickIncrAvailablePeer(bpr.height, secondPeerID)
		if peer == nil {
			bpr.Logger.Debug("REDACTED", "REDACTED", bpr.height)
			time.Sleep(requestIntervalMS * time.Millisecond)
			continue PICK_PEER_LOOP
		}
		break PICK_PEER_LOOP
	}
	bpr.mtx.Lock()
	bpr.peerID = peer.id
	bpr.mtx.Unlock()

	bpr.pool.sendRequest(bpr.height, peer.id)
}

//
//
func (bpr *bpRequester) pickSecondPeerAndSendRequest() (picked bool) {
	bpr.mtx.Lock()
	if bpr.secondPeerID != "REDACTED" {
		bpr.mtx.Unlock()
		return false
	}
	peerID := bpr.peerID
	bpr.mtx.Unlock()

	secondPeer := bpr.pool.pickIncrAvailablePeer(bpr.height, peerID)
	if secondPeer != nil {
		bpr.mtx.Lock()
		bpr.secondPeerID = secondPeer.id
		bpr.mtx.Unlock()

		bpr.pool.sendRequest(bpr.height, secondPeer.id)
		return true
	}

	return false
}

//
func (bpr *bpRequester) newHeight(height int64) {
	select {
	case bpr.newHeightCh <- height:
	default:
	}
}

//
//
func (bpr *bpRequester) requestRoutine() {
	gotBlock := false

OUTER_LOOP:
	for {
		bpr.pickPeerAndSendRequest()

		poolHeight := bpr.pool.Height()
		if bpr.height-poolHeight < minBlocksForSingleRequest {
			bpr.pickSecondPeerAndSendRequest()
		}

		retryTimer := time.NewTimer(requestRetrySeconds * time.Second)
		defer retryTimer.Stop()

		for {
			select {
			case <-bpr.pool.Quit():
				if err := bpr.Stop(); err != nil {
					bpr.Logger.Error("REDACTED", "REDACTED", err)
				}
				return
			case <-bpr.Quit():
				return
			case <-retryTimer.C:
				if !gotBlock {
					bpr.Logger.Debug("REDACTED", "REDACTED", bpr.height, "REDACTED", bpr.peerID, "REDACTED", bpr.secondPeerID)
					bpr.reset(bpr.peerID)
					bpr.reset(bpr.secondPeerID)
					continue OUTER_LOOP
				}
			case peerID := <-bpr.redoCh:
				if bpr.didRequestFrom(peerID) {
					removedBlock := bpr.reset(peerID)
					if removedBlock {
						gotBlock = false
					}
				}
				//
				//
				if len(bpr.requestedFrom()) == 0 {
					retryTimer.Stop()
					continue OUTER_LOOP
				}
			case newHeight := <-bpr.newHeightCh:
				if !gotBlock && bpr.height-newHeight < minBlocksForSingleRequest {
					//
					//
					//
					//
					if picked := bpr.pickSecondPeerAndSendRequest(); picked {
						if !retryTimer.Stop() {
							<-retryTimer.C
						}
						retryTimer.Reset(requestRetrySeconds * time.Second)
					}
				}
			case <-bpr.gotBlockCh:
				gotBlock = true
				//
				//
			}
		}
	}
}
