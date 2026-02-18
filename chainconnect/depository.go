package chainconnect

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"sync/atomic"
	"time"

	stream "github.com/valkyrieworks/utils/pace"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
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
	queryCadenceMillis         = 2
	maximumAwaitingQueriesEachNode = 20
	queryReprocessMoments       = 30

	//
	//
	//
	//
	//
	//
	//
	minimumAcceptRatio = 128 * 1024 //

	//
	//
	//
	nodeLinkWait = 3 * time.Second

	//
	//
	//
	minimumLedgersForUniqueQuery = 50
)

var nodeDeadline = 15 * time.Second //

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
type LedgerDepository struct {
	daemon.RootDaemon
	beginMoment   time.Time
	beginLevel int64

	mtx engineconnect.Lock
	//
	inquirers map[int64]*breakpointInquirer
	level     int64 //
	//
	nodes         map[p2p.ID]*breakpointNode
	prohibitedNodes   map[p2p.ID]time.Time
	orderedNodes   []*breakpointNode //
	maximumNodeLevel int64     //

	//
	countAwaiting int32 //

	queriesChan chan<- LedgerQuery
	faultsChan   chan<- nodeFault
}

//
//
type LedgerQuery struct {
	Level int64
	NodeUID p2p.ID
}

//
//
func NewLedgerDepository(begin int64, queriesChan chan<- LedgerQuery, faultsChan chan<- nodeFault) *LedgerDepository {
	bp := &LedgerDepository{
		nodes:       make(map[p2p.ID]*breakpointNode),
		prohibitedNodes: make(map[p2p.ID]time.Time),
		inquirers:  make(map[int64]*breakpointInquirer),
		level:      begin,
		beginLevel: begin,
		countAwaiting:  0,

		queriesChan: queriesChan,
		faultsChan:   faultsChan,
	}
	bp.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", bp)
	return bp
}

//
//
func (depository *LedgerDepository) OnBegin() error {
	depository.beginMoment = time.Now()
	go depository.createInquirersProcess()
	return nil
}

//
func (depository *LedgerDepository) createInquirersProcess() {
	for {
		if !depository.IsActive() {
			return
		}

		//
		//
		if time.Since(depository.beginMoment) < nodeLinkWait {
			//
			pausePeriod := nodeLinkWait - time.Since(depository.beginMoment)
			time.Sleep(pausePeriod)
		}

		depository.mtx.Lock()
		var (
			maximumInquirersSpawned = len(depository.inquirers) >= len(depository.nodes)*maximumAwaitingQueriesEachNode

			followingLevel           = depository.level + int64(len(depository.inquirers))
			maximumNodeLevelAttained = followingLevel > depository.maximumNodeLevel
		)
		depository.mtx.Unlock()

		switch {
		case maximumInquirersSpawned: //
			time.Sleep(queryCadenceMillis * time.Millisecond)
			depository.deleteExpiredNodes()
		case maximumNodeLevelAttained: //
			time.Sleep(queryCadenceMillis * time.Millisecond)
		default:
			//
			depository.createFollowingInquirer(followingLevel)
			//
			time.Sleep(queryCadenceMillis * time.Millisecond)
		}
	}
}

func (depository *LedgerDepository) deleteExpiredNodes() {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	for _, node := range depository.nodes {
		if !node.didDeadline && node.countAwaiting > 0 {
			currentRatio := node.acceptAuditor.Status().CurrentRatio
			//
			if currentRatio != 0 && currentRatio < minimumAcceptRatio {
				err := errors.New("REDACTED")
				depository.transmitFault(err, node.id)
				depository.Tracer.Fault("REDACTED", "REDACTED", node.id,
					"REDACTED", err,
					"REDACTED", fmt.Sprintf("REDACTED", currentRatio/1024),
					"REDACTED", fmt.Sprintf("REDACTED", minimumAcceptRatio/1024))
				node.didDeadline = true
			}

			node.currentRatio = currentRatio
		}

		if node.didDeadline {
			depository.deleteNode(node.id)
		}
	}

	for nodeUID := range depository.prohibitedNodes {
		if !depository.isNodeProhibited(nodeUID) {
			delete(depository.prohibitedNodes, nodeUID)
		}
	}

	depository.arrangeNodes()
}

//
//
func (depository *LedgerDepository) FetchStatus() (level int64, countAwaiting int32, sizeInquirers int) {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	return depository.level, atomic.LoadInt32(&depository.countAwaiting), len(depository.inquirers)
}

//
//
func (depository *LedgerDepository) IsSeizedUp() bool {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	//
	if len(depository.nodes) == 0 {
		depository.Tracer.Diagnose("REDACTED")
		return false
	}

	//
	//
	//
	//
	//
	acceptedLedgerOrScheduledOut := depository.level > 0 || time.Since(depository.beginMoment) > 5*time.Second
	ourLedgerIsGreatestBetweenNodes := depository.maximumNodeLevel == 0 || depository.level >= (depository.maximumNodeLevel-1)
	isSeizedUp := acceptedLedgerOrScheduledOut && ourLedgerIsGreatestBetweenNodes
	return isSeizedUp
}

//
//
//
//
//
//
//
func (depository *LedgerDepository) GlanceDualLedgers() (initial, moment *kinds.Ledger, initialExtensionEndorse *kinds.ExpandedEndorse) {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	if r := depository.inquirers[depository.level]; r != nil {
		initial = r.fetchLedger()
		initialExtensionEndorse = r.fetchExpandedEndorse()
	}
	if r := depository.inquirers[depository.level+1]; r != nil {
		moment = r.fetchLedger()
	}
	return
}

//
func (depository *LedgerDepository) EjectQuery() {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	r := depository.inquirers[depository.level]
	if r == nil {
		panic(fmt.Sprintf("REDACTED", depository.level))
	}

	if err := r.Halt(); err != nil {
		depository.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	delete(depository.inquirers, depository.level)
	depository.level++

	//
	//
	for i := int64(0); i < minimumLedgersForUniqueQuery && i < int64(len(depository.inquirers)); i++ {
		depository.inquirers[depository.level+i].newLevel(depository.level)
	}
}

//
//
//
func (depository *LedgerDepository) DeleteNodeAndReworkAllNodeQueries(level int64) p2p.ID {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	query := depository.inquirers[level]
	nodeUID := query.acquiredLedgerFromNodeUID()
	//
	depository.deleteNode(nodeUID)
	depository.prohibitNode(nodeUID)
	return nodeUID
}

//
//
func (depository *LedgerDepository) ReworkQueryFrom(level int64, nodeUID p2p.ID) {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	if inquirer, ok := depository.inquirers[level]; ok { //
		if inquirer.didQueryFrom(nodeUID) { //
			inquirer.rework(nodeUID)
		}
	}
}

//
func (depository *LedgerDepository) ReworkQuery(level int64) p2p.ID {
	return depository.DeleteNodeAndReworkAllNodeQueries(level)
}

//
//
//
//
//
//
//
//
func (depository *LedgerDepository) AppendLedger(nodeUID p2p.ID, ledger *kinds.Ledger, extensionEndorse *kinds.ExpandedEndorse, ledgerVolume int) error {
	if extensionEndorse != nil && ledger.Level != extensionEndorse.Level {
		err := fmt.Errorf("REDACTED", ledger.Level, extensionEndorse.Level)
		//
		depository.transmitFault(err, nodeUID)
		return err
	}

	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	inquirer := depository.inquirers[ledger.Level]
	if inquirer == nil {
		//
		//
		//
		//
		if ledger.Level > depository.level || ledger.Level < depository.beginLevel {
			err := fmt.Errorf("REDACTED",
				ledger.Level, depository.level, depository.beginLevel)
			depository.transmitFault(err, nodeUID)
			return err
		}

		return fmt.Errorf("REDACTED", ledger.Level, nodeUID)
	}

	if !inquirer.assignLedger(ledger, extensionEndorse, nodeUID) {
		err := fmt.Errorf("REDACTED", ledger.Level, inquirer.queriedFrom(), nodeUID)
		depository.transmitFault(err, nodeUID)
		return err
	}

	atomic.AddInt32(&depository.countAwaiting, -1)
	node := depository.nodes[nodeUID]
	if node != nil {
		node.reduceAwaiting(ledgerVolume)
	}

	return nil
}

//
func (depository *LedgerDepository) Level() int64 {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()
	return depository.level
}

//
func (depository *LedgerDepository) MaximumNodeLevel() int64 {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()
	return depository.maximumNodeLevel
}

//
func (depository *LedgerDepository) AssignNodeScope(nodeUID p2p.ID, root int64, level int64) {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	node := depository.nodes[nodeUID]
	if node != nil {
		if root < node.root || level < node.level {
			depository.Tracer.Details(
				"REDACTED",
				"REDACTED", nodeUID,
				"REDACTED", level,
				"REDACTED", root,
				"REDACTED", node.level,
				"REDACTED", node.root,
			)

			//
			depository.deleteNode(nodeUID)
			depository.prohibitNode(nodeUID)

			return
		}
		node.root = root
		node.level = level
	} else {
		if depository.isNodeProhibited(nodeUID) {
			depository.Tracer.Diagnose("REDACTED", "REDACTED", nodeUID)
			return
		}
		node = newBreakpointNode(depository, nodeUID, root, level)
		node.assignTracer(depository.Tracer.With("REDACTED", nodeUID))
		depository.nodes[nodeUID] = node
		//
		//
		depository.orderedNodes = append([]*breakpointNode{node}, depository.orderedNodes...)
	}

	if level > depository.maximumNodeLevel {
		depository.maximumNodeLevel = level
	}
}

//
//
func (depository *LedgerDepository) DeleteNode(nodeUID p2p.ID) {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	depository.deleteNode(nodeUID)
}

//
func (depository *LedgerDepository) deleteNode(nodeUID p2p.ID) {
	for _, inquirer := range depository.inquirers {
		if inquirer.didQueryFrom(nodeUID) {
			inquirer.rework(nodeUID)
		}
	}

	node, ok := depository.nodes[nodeUID]
	if ok {
		if node.deadline != nil {
			node.deadline.Stop()
		}

		delete(depository.nodes, nodeUID)
		for i, p := range depository.orderedNodes {
			if p.id == nodeUID {
				depository.orderedNodes = append(depository.orderedNodes[:i], depository.orderedNodes[i+1:]...)
				break
			}
		}

		//
		//
		if node.level == depository.maximumNodeLevel {
			depository.modifyMaximumNodeLevel()
		}
	}
}

//
func (depository *LedgerDepository) modifyMaximumNodeLevel() {
	var max int64
	for _, node := range depository.nodes {
		if node.level > max {
			max = node.level
		}
	}
	depository.maximumNodeLevel = max
}

//
func (depository *LedgerDepository) IsNodeProhibited(nodeUID p2p.ID) bool {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()
	return depository.isNodeProhibited(nodeUID)
}

//
func (depository *LedgerDepository) isNodeProhibited(nodeUID p2p.ID) bool {
	//
	return time.Since(depository.prohibitedNodes[nodeUID]) < time.Second*60
}

//
func (depository *LedgerDepository) prohibitNode(nodeUID p2p.ID) {
	depository.Tracer.Diagnose("REDACTED", nodeUID)
	depository.prohibitedNodes[nodeUID] = engineclock.Now()
}

//
//
func (depository *LedgerDepository) selectIncreaseAccessibleNode(level int64, omitNodeUID p2p.ID) *breakpointNode {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	for _, node := range depository.orderedNodes {
		if node.id == omitNodeUID {
			continue
		}
		if node.didDeadline {
			depository.deleteNode(node.id)
			continue
		}
		if node.countAwaiting >= maximumAwaitingQueriesEachNode {
			continue
		}
		if level < node.root || level > node.level {
			continue
		}
		node.increaseAwaiting()
		return node
	}

	return nil
}

//
//
//
func (depository *LedgerDepository) arrangeNodes() {
	sort.Slice(depository.orderedNodes, func(i, j int) bool {
		return depository.orderedNodes[i].currentRatio > depository.orderedNodes[j].currentRatio
	})
}

func (depository *LedgerDepository) createFollowingInquirer(followingLevel int64) {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	query := newBreakpointInquirer(depository, followingLevel)

	depository.inquirers[followingLevel] = query
	atomic.AddInt32(&depository.countAwaiting, 1)

	if err := query.Begin(); err != nil {
		query.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}

//
func (depository *LedgerDepository) transmitQuery(level int64, nodeUID p2p.ID) {
	if !depository.IsActive() {
		return
	}
	depository.queriesChan <- LedgerQuery{level, nodeUID}
}

//
func (depository *LedgerDepository) transmitFault(err error, nodeUID p2p.ID) {
	if !depository.IsActive() {
		return
	}
	depository.faultsChan <- nodeFault{err, nodeUID}
}

//
//
//
func (depository *LedgerDepository) diagnose() string {
	depository.mtx.Lock()
	defer depository.mtx.Unlock()

	str := "REDACTED"
	followingLevel := depository.level + int64(len(depository.inquirers))
	for h := depository.level; h < followingLevel; h++ {
		if depository.inquirers[h] == nil {
			str += fmt.Sprintf("REDACTED", h)
		} else {
			str += fmt.Sprintf("REDACTED", h)
			str += fmt.Sprintf("REDACTED", depository.inquirers[h].ledger != nil)
			str += fmt.Sprintf("REDACTED", depository.inquirers[h].extensionEndorse != nil)
		}
	}
	return str
}

//

type breakpointNode struct {
	didDeadline  bool
	currentRatio     int64
	countAwaiting  int32
	level      int64
	root        int64
	depository        *LedgerDepository
	id          p2p.ID
	acceptAuditor *stream.Auditor

	deadline *time.Timer

	tracer log.Tracer
}

func newBreakpointNode(depository *LedgerDepository, nodeUID p2p.ID, root int64, level int64) *breakpointNode {
	node := &breakpointNode{
		depository:       depository,
		id:         nodeUID,
		root:       root,
		level:     level,
		countAwaiting: 0,
		tracer:     log.NewNoopTracer(),
	}
	return node
}

func (node *breakpointNode) assignTracer(l log.Tracer) {
	node.tracer = l
}

func (node *breakpointNode) restoreAuditor() {
	node.acceptAuditor = stream.New(time.Second, time.Second*40)
	primaryItem := float64(minimumAcceptRatio) * math.E
	node.acceptAuditor.AssignRemain(primaryItem)
}

func (node *breakpointNode) restoreDeadline() {
	if node.deadline == nil {
		node.deadline = time.AfterFunc(nodeDeadline, node.onDeadline)
	} else {
		node.deadline.Reset(nodeDeadline)
	}
}

func (node *breakpointNode) increaseAwaiting() {
	if node.countAwaiting == 0 {
		node.restoreAuditor()
		node.restoreDeadline()
	}
	node.countAwaiting++
}

func (node *breakpointNode) reduceAwaiting(acceptVolume int) {
	node.countAwaiting--
	if node.countAwaiting == 0 {
		node.deadline.Stop()
	} else {
		node.acceptAuditor.Modify(acceptVolume)
		node.restoreDeadline()
	}
}

func (node *breakpointNode) onDeadline() {
	node.depository.mtx.Lock()
	defer node.depository.mtx.Unlock()

	node.depository.transmitFault(ErrNodeDeadline, node.id)
	node.tracer.Fault("REDACTED", "REDACTED", ErrNodeDeadline, "REDACTED", nodeDeadline)
	node.didDeadline = true
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
type breakpointInquirer struct {
	daemon.RootDaemon

	depository        *LedgerDepository
	level      int64
	acquiredLedgerChan  chan struct{}
	reworkChan      chan p2p.ID //
	newLevelChan chan int64

	mtx          engineconnect.Lock
	nodeUID       p2p.ID
	momentNodeUID p2p.ID //
	acquiredLedgerFrom p2p.ID
	ledger        *kinds.Ledger
	extensionEndorse    *kinds.ExpandedEndorse
}

func newBreakpointInquirer(depository *LedgerDepository, level int64) *breakpointInquirer {
	bpr := &breakpointInquirer{
		depository:        depository,
		level:      level,
		acquiredLedgerChan:  make(chan struct{}, 1),
		reworkChan:      make(chan p2p.ID, 1),
		newLevelChan: make(chan int64, 1),

		nodeUID:       "REDACTED",
		momentNodeUID: "REDACTED",
		ledger:        nil,
	}
	bpr.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", bpr)
	return bpr
}

func (bpr *breakpointInquirer) OnBegin() error {
	go bpr.queryProcess()
	return nil
}

//
func (bpr *breakpointInquirer) assignLedger(ledger *kinds.Ledger, extensionEndorse *kinds.ExpandedEndorse, nodeUID p2p.ID) bool {
	bpr.mtx.Lock()
	if bpr.nodeUID != nodeUID && bpr.momentNodeUID != nodeUID {
		bpr.mtx.Unlock()
		return false
	}
	if bpr.ledger != nil {
		bpr.mtx.Unlock()
		return true //
	}

	bpr.ledger = ledger
	bpr.extensionEndorse = extensionEndorse
	bpr.acquiredLedgerFrom = nodeUID
	bpr.mtx.Unlock()

	select {
	case bpr.acquiredLedgerChan <- struct{}{}:
	default:
	}
	return true
}

func (bpr *breakpointInquirer) fetchLedger() *kinds.Ledger {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.ledger
}

func (bpr *breakpointInquirer) fetchExpandedEndorse() *kinds.ExpandedEndorse {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.extensionEndorse
}

//
func (bpr *breakpointInquirer) queriedFrom() []p2p.ID {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	nodeIDXDatastore := make([]p2p.ID, 0, 2)
	if bpr.nodeUID != "REDACTED" {
		nodeIDXDatastore = append(nodeIDXDatastore, bpr.nodeUID)
	}
	if bpr.momentNodeUID != "REDACTED" {
		nodeIDXDatastore = append(nodeIDXDatastore, bpr.momentNodeUID)
	}
	return nodeIDXDatastore
}

//
func (bpr *breakpointInquirer) didQueryFrom(nodeUID p2p.ID) bool {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.nodeUID == nodeUID || bpr.momentNodeUID == nodeUID
}

//
func (bpr *breakpointInquirer) acquiredLedgerFromNodeUID() p2p.ID {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.acquiredLedgerFrom
}

//
func (bpr *breakpointInquirer) restore(nodeUID p2p.ID) (deletedLedger bool) {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()

	//
	if bpr.acquiredLedgerFrom == nodeUID {
		bpr.ledger = nil
		bpr.extensionEndorse = nil
		bpr.acquiredLedgerFrom = "REDACTED"
		deletedLedger = true
		atomic.AddInt32(&bpr.depository.countAwaiting, 1)
	}

	if bpr.nodeUID == nodeUID {
		bpr.nodeUID = "REDACTED"
	} else {
		bpr.momentNodeUID = "REDACTED"
	}

	return deletedLedger
}

//
//
//
func (bpr *breakpointInquirer) rework(nodeUID p2p.ID) {
	select {
	case bpr.reworkChan <- nodeUID:
	default:
	}
}

func (bpr *breakpointInquirer) selectNodeAndTransmitQuery() {
	bpr.mtx.Lock()
	momentNodeUID := bpr.momentNodeUID
	bpr.mtx.Unlock()

	var node *breakpointNode
SELECT_NODE_CYCLE:
	for {
		if !bpr.IsActive() || !bpr.depository.IsActive() {
			return
		}
		node = bpr.depository.selectIncreaseAccessibleNode(bpr.level, momentNodeUID)
		if node == nil {
			bpr.Tracer.Diagnose("REDACTED", "REDACTED", bpr.level)
			time.Sleep(queryCadenceMillis * time.Millisecond)
			continue SELECT_NODE_CYCLE
		}
		break SELECT_NODE_CYCLE
	}
	bpr.mtx.Lock()
	bpr.nodeUID = node.id
	bpr.mtx.Unlock()

	bpr.depository.transmitQuery(bpr.level, node.id)
}

//
//
func (bpr *breakpointInquirer) selectMomentNodeAndTransmitQuery() (chosen bool) {
	bpr.mtx.Lock()
	if bpr.momentNodeUID != "REDACTED" {
		bpr.mtx.Unlock()
		return false
	}
	nodeUID := bpr.nodeUID
	bpr.mtx.Unlock()

	momentNode := bpr.depository.selectIncreaseAccessibleNode(bpr.level, nodeUID)
	if momentNode != nil {
		bpr.mtx.Lock()
		bpr.momentNodeUID = momentNode.id
		bpr.mtx.Unlock()

		bpr.depository.transmitQuery(bpr.level, momentNode.id)
		return true
	}

	return false
}

//
func (bpr *breakpointInquirer) newLevel(level int64) {
	select {
	case bpr.newLevelChan <- level:
	default:
	}
}

//
//
func (bpr *breakpointInquirer) queryProcess() {
	acquiredLedger := false

EXTERNAL_CYCLE:
	for {
		bpr.selectNodeAndTransmitQuery()

		depositoryLevel := bpr.depository.Level()
		if bpr.level-depositoryLevel < minimumLedgersForUniqueQuery {
			bpr.selectMomentNodeAndTransmitQuery()
		}

		reprocessClock := time.NewTimer(queryReprocessMoments * time.Second)
		defer reprocessClock.Stop()

		for {
			select {
			case <-bpr.depository.Exit():
				if err := bpr.Halt(); err != nil {
					bpr.Tracer.Fault("REDACTED", "REDACTED", err)
				}
				return
			case <-bpr.Exit():
				return
			case <-reprocessClock.C:
				if !acquiredLedger {
					bpr.Tracer.Diagnose("REDACTED", "REDACTED", bpr.level, "REDACTED", bpr.nodeUID, "REDACTED", bpr.momentNodeUID)
					bpr.restore(bpr.nodeUID)
					bpr.restore(bpr.momentNodeUID)
					continue EXTERNAL_CYCLE
				}
			case nodeUID := <-bpr.reworkChan:
				if bpr.didQueryFrom(nodeUID) {
					deletedLedger := bpr.restore(nodeUID)
					if deletedLedger {
						acquiredLedger = false
					}
				}
				//
				//
				if len(bpr.queriedFrom()) == 0 {
					reprocessClock.Stop()
					continue EXTERNAL_CYCLE
				}
			case newLevel := <-bpr.newLevelChan:
				if !acquiredLedger && bpr.level-newLevel < minimumLedgersForUniqueQuery {
					//
					//
					//
					//
					if chosen := bpr.selectMomentNodeAndTransmitQuery(); chosen {
						if !reprocessClock.Stop() {
							<-reprocessClock.C
						}
						reprocessClock.Reset(queryReprocessMoments * time.Second)
					}
				}
			case <-bpr.acquiredLedgerChan:
				acquiredLedger = true
				//
				//
			}
		}
	}
}
