package chainchronize

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"sync/atomic"
	"time"

	stream "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/throughput"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
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
	solicitDurationMSEC         = 2
	maximumAwaitingSolicitsEveryNode = 20
	solicitReissueMoments       = 30

	//
	//
	//
	//
	//
	//
	//
	minimumObtainFrequency = 128 * 1024 //

	//
	//
	//
	nodeLinkAwait = 3 * time.Second

	//
	//
	//
	minimumLedgersForeachUniqueSolicit = 50
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
type LedgerHub struct {
	facility.FoundationFacility
	initiateMoment   time.Time
	initiateAltitude int64

	mtx commitchronize.Exclusion
	//
	solicitors map[int64]*breakpointSolicitor
	altitude     int64 //
	//
	nodes         map[p2p.ID]*breakpointNode
	prohibitedNodes   map[p2p.ID]time.Time
	orderedNodes   []*breakpointNode //
	maximumNodeAltitude int64     //

	//
	countAwaiting int32 //

	solicitsStream chan<- LedgerSolicit
	faultsStream   chan<- nodeFailure
}

//
//
type LedgerSolicit struct {
	Altitude int64
	NodeUUID p2p.ID
}

//
//
func FreshLedgerHub(initiate int64, solicitsStream chan<- LedgerSolicit, faultsStream chan<- nodeFailure) *LedgerHub {
	bp := &LedgerHub{
		nodes:       make(map[p2p.ID]*breakpointNode),
		prohibitedNodes: make(map[p2p.ID]time.Time),
		solicitors:  make(map[int64]*breakpointSolicitor),
		altitude:      initiate,
		initiateAltitude: initiate,
		countAwaiting:  0,

		solicitsStream: solicitsStream,
		faultsStream:   faultsStream,
	}
	bp.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", bp)
	return bp
}

//
//
func (hub *LedgerHub) UponInitiate() error {
	hub.initiateMoment = time.Now()
	go hub.createSolicitorsProcedure()
	return nil
}

//
func (hub *LedgerHub) createSolicitorsProcedure() {
	for {
		if !hub.EqualsActive() {
			return
		}

		//
		//
		if time.Since(hub.initiateMoment) < nodeLinkAwait {
			//
			snoozeInterval := nodeLinkAwait - time.Since(hub.initiateMoment)
			time.Sleep(snoozeInterval)
		}

		hub.mtx.Lock()
		var (
			maximumSolicitorsSpawned = len(hub.solicitors) >= len(hub.nodes)*maximumAwaitingSolicitsEveryNode

			followingAltitude           = hub.altitude + int64(len(hub.solicitors))
			maximumNodeAltitudeAttained = followingAltitude > hub.maximumNodeAltitude
		)
		hub.mtx.Unlock()

		switch {
		case maximumSolicitorsSpawned: //
			time.Sleep(solicitDurationMSEC * time.Millisecond)
			hub.discardExpiredNodes()
		case maximumNodeAltitudeAttained: //
			time.Sleep(solicitDurationMSEC * time.Millisecond)
		default:
			//
			hub.createFollowingSolicitor(followingAltitude)
			//
			time.Sleep(solicitDurationMSEC * time.Millisecond)
		}
	}
}

func (hub *LedgerHub) discardExpiredNodes() {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	for _, node := range hub.nodes {
		if !node.actedDeadline && node.countAwaiting > 0 {
			currentFrequency := node.obtainOverseer.Condition().CurrentFrequency
			//
			if currentFrequency != 0 && currentFrequency < minimumObtainFrequency {
				err := errors.New("REDACTED")
				hub.transmitFailure(err, node.id)
				hub.Tracer.Failure("REDACTED", "REDACTED", node.id,
					"REDACTED", err,
					"REDACTED", fmt.Sprintf("REDACTED", currentFrequency/1024),
					"REDACTED", fmt.Sprintf("REDACTED", minimumObtainFrequency/1024))
				node.actedDeadline = true
			}

			node.currentFrequency = currentFrequency
		}

		if node.actedDeadline {
			hub.discardNode(node.id)
		}
	}

	for nodeUUID := range hub.prohibitedNodes {
		if !hub.equalsNodeProhibited(nodeUUID) {
			delete(hub.prohibitedNodes, nodeUUID)
		}
	}

	hub.arrangeNodes()
}

//
//
func (hub *LedgerHub) ObtainCondition() (altitude int64, countAwaiting int32, extentSolicitors int) {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	return hub.altitude, atomic.LoadInt32(&hub.countAwaiting), len(hub.solicitors)
}

//
//
func (hub *LedgerHub) EqualsSeizedActive() bool {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	//
	if len(hub.nodes) == 0 {
		hub.Tracer.Diagnose("REDACTED")
		return false
	}

	//
	//
	//
	//
	//
	acceptedLedgerEitherScheduledOutput := hub.altitude > 0 || time.Since(hub.initiateMoment) > 5*time.Second
	mineSuccessionEqualsGreatestBetweenNodes := hub.maximumNodeAltitude == 0 || hub.altitude >= (hub.maximumNodeAltitude-1)
	equalsSeizedActive := acceptedLedgerEitherScheduledOutput && mineSuccessionEqualsGreatestBetweenNodes
	return equalsSeizedActive
}

//
//
//
//
//
//
//
func (hub *LedgerHub) GlanceCoupleLedgers() (initial, ordinal *kinds.Ledger, initialAddnEndorse *kinds.ExpandedEndorse) {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	if r := hub.solicitors[hub.altitude]; r != nil {
		initial = r.obtainLedger()
		initialAddnEndorse = r.obtainExpandedEndorse()
	}
	if r := hub.solicitors[hub.altitude+1]; r != nil {
		ordinal = r.obtainLedger()
	}
	return
}

//
func (hub *LedgerHub) ExtractSolicit() {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	r := hub.solicitors[hub.altitude]
	if r == nil {
		panic(fmt.Sprintf("REDACTED", hub.altitude))
	}

	if err := r.Halt(); err != nil {
		hub.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	delete(hub.solicitors, hub.altitude)
	hub.altitude++

	//
	//
	for i := int64(0); i < minimumLedgersForeachUniqueSolicit && i < int64(len(hub.solicitors)); i++ {
		hub.solicitors[hub.altitude+i].freshAltitude(hub.altitude)
	}
}

//
//
//
func (hub *LedgerHub) DiscardNodeAlsoReiterateEveryNodeSolicits(altitude int64) p2p.ID {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	solicit := hub.solicitors[altitude]
	if solicit == nil {
		return "REDACTED"
	}

	nodeUUID := solicit.attainedLedgerOriginatingNodeUUID()
	//
	hub.discardNode(nodeUUID)
	hub.prohibitNode(nodeUUID)
	return nodeUUID
}

//
//
func (hub *LedgerHub) ReiterateSolicitOriginating(altitude int64, nodeUUID p2p.ID) {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	if solicitor, ok := hub.solicitors[altitude]; ok { //
		if solicitor.actedSolicitOriginating(nodeUUID) { //
			solicitor.reiterate(nodeUUID)
		}
	}
}

//
func (hub *LedgerHub) ReiterateSolicit(altitude int64) p2p.ID {
	return hub.DiscardNodeAlsoReiterateEveryNodeSolicits(altitude)
}

//
//
//
//
//
//
//
//
func (hub *LedgerHub) AppendLedger(nodeUUID p2p.ID, ledger *kinds.Ledger, addnEndorse *kinds.ExpandedEndorse, ledgerExtent int) error {
	if addnEndorse != nil && ledger.Altitude != addnEndorse.Altitude {
		err := fmt.Errorf("REDACTED", ledger.Altitude, addnEndorse.Altitude)
		//
		hub.transmitFailure(err, nodeUUID)
		return err
	}

	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	solicitor := hub.solicitors[ledger.Altitude]
	if solicitor == nil {
		//
		//
		//
		//
		if ledger.Altitude > hub.altitude || ledger.Altitude < hub.initiateAltitude {
			err := fmt.Errorf("REDACTED",
				ledger.Altitude, hub.altitude, hub.initiateAltitude)
			hub.transmitFailure(err, nodeUUID)
			return err
		}

		return fmt.Errorf("REDACTED", ledger.Altitude, nodeUUID)
	}

	if !solicitor.assignLedger(ledger, addnEndorse, nodeUUID) {
		err := fmt.Errorf("REDACTED", ledger.Altitude, solicitor.solicitedOriginating(), nodeUUID)
		hub.transmitFailure(err, nodeUUID)
		return err
	}

	atomic.AddInt32(&hub.countAwaiting, -1)
	node := hub.nodes[nodeUUID]
	if node != nil {
		node.decreaseAwaiting(ledgerExtent)
	}

	return nil
}

//
func (hub *LedgerHub) Altitude() int64 {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()
	return hub.altitude
}

//
func (hub *LedgerHub) MaximumNodeAltitude() int64 {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()
	return hub.maximumNodeAltitude
}

//
func (hub *LedgerHub) AssignNodeScope(nodeUUID p2p.ID, foundation int64, altitude int64) {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	node := hub.nodes[nodeUUID]
	if node != nil {
		if foundation < node.foundation || altitude < node.altitude {
			hub.Tracer.Details(
				"REDACTED",
				"REDACTED", nodeUUID,
				"REDACTED", altitude,
				"REDACTED", foundation,
				"REDACTED", node.altitude,
				"REDACTED", node.foundation,
			)

			//
			hub.discardNode(nodeUUID)
			hub.prohibitNode(nodeUUID)

			return
		}
		node.foundation = foundation
		node.altitude = altitude
	} else {
		if hub.equalsNodeProhibited(nodeUUID) {
			hub.Tracer.Diagnose("REDACTED", "REDACTED", nodeUUID)
			return
		}
		node = freshBreakpointNode(hub, nodeUUID, foundation, altitude)
		node.assignTracer(hub.Tracer.Using("REDACTED", nodeUUID))
		hub.nodes[nodeUUID] = node
		//
		//
		hub.orderedNodes = append([]*breakpointNode{node}, hub.orderedNodes...)
	}

	if altitude > hub.maximumNodeAltitude {
		hub.maximumNodeAltitude = altitude
	}
}

//
//
func (hub *LedgerHub) DiscardNode(nodeUUID p2p.ID) {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	hub.discardNode(nodeUUID)
}

//
func (hub *LedgerHub) discardNode(nodeUUID p2p.ID) {
	for _, solicitor := range hub.solicitors {
		if solicitor.actedSolicitOriginating(nodeUUID) {
			solicitor.reiterate(nodeUUID)
		}
	}

	node, ok := hub.nodes[nodeUUID]
	if ok {
		if node.deadline != nil {
			node.deadline.Stop()
		}

		delete(hub.nodes, nodeUUID)
		for i, p := range hub.orderedNodes {
			if p.id == nodeUUID {
				hub.orderedNodes = append(hub.orderedNodes[:i], hub.orderedNodes[i+1:]...)
				break
			}
		}

		//
		//
		if node.altitude == hub.maximumNodeAltitude {
			hub.reviseMaximumNodeAltitude()
		}
	}
}

//
func (hub *LedgerHub) reviseMaximumNodeAltitude() {
	var max int64
	for _, node := range hub.nodes {
		if node.altitude > max {
			max = node.altitude
		}
	}
	hub.maximumNodeAltitude = max
}

//
func (hub *LedgerHub) EqualsNodeProhibited(nodeUUID p2p.ID) bool {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()
	return hub.equalsNodeProhibited(nodeUUID)
}

//
func (hub *LedgerHub) equalsNodeProhibited(nodeUUID p2p.ID) bool {
	//
	return time.Since(hub.prohibitedNodes[nodeUUID]) < time.Second*60
}

//
func (hub *LedgerHub) prohibitNode(nodeUUID p2p.ID) {
	hub.Tracer.Diagnose("REDACTED", nodeUUID)
	hub.prohibitedNodes[nodeUUID] = committime.Now()
}

//
//
func (hub *LedgerHub) selectIncreaseAccessibleNode(altitude int64, omitNodeUUID p2p.ID) *breakpointNode {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	for _, node := range hub.orderedNodes {
		if node.id == omitNodeUUID {
			continue
		}
		if node.actedDeadline {
			hub.discardNode(node.id)
			continue
		}
		if node.countAwaiting >= maximumAwaitingSolicitsEveryNode {
			continue
		}
		if altitude < node.foundation || altitude > node.altitude {
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
func (hub *LedgerHub) arrangeNodes() {
	sort.Slice(hub.orderedNodes, func(i, j int) bool {
		return hub.orderedNodes[i].currentFrequency > hub.orderedNodes[j].currentFrequency
	})
}

func (hub *LedgerHub) createFollowingSolicitor(followingAltitude int64) {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	solicit := freshBreakpointSolicitor(hub, followingAltitude)

	hub.solicitors[followingAltitude] = solicit
	atomic.AddInt32(&hub.countAwaiting, 1)

	if err := solicit.Initiate(); err != nil {
		solicit.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

//
func (hub *LedgerHub) transmitSolicit(altitude int64, nodeUUID p2p.ID) {
	if !hub.EqualsActive() {
		return
	}
	hub.solicitsStream <- LedgerSolicit{altitude, nodeUUID}
}

//
func (hub *LedgerHub) transmitFailure(err error, nodeUUID p2p.ID) {
	if !hub.EqualsActive() {
		return
	}
	hub.faultsStream <- nodeFailure{err, nodeUUID}
}

//
//
//
func (hub *LedgerHub) diagnose() string {
	hub.mtx.Lock()
	defer hub.mtx.Unlock()

	str := "REDACTED"
	followingAltitude := hub.altitude + int64(len(hub.solicitors))
	for h := hub.altitude; h < followingAltitude; h++ {
		if hub.solicitors[h] == nil {
			str += fmt.Sprintf("REDACTED", h)
		} else {
			str += fmt.Sprintf("REDACTED", h)
			str += fmt.Sprintf("REDACTED", hub.solicitors[h].ledger != nil)
			str += fmt.Sprintf("REDACTED", hub.solicitors[h].addnEndorse != nil)
		}
	}
	return str
}

//

type breakpointNode struct {
	actedDeadline  bool
	currentFrequency     int64
	countAwaiting  int32
	altitude      int64
	foundation        int64
	hub        *LedgerHub
	id          p2p.ID
	obtainOverseer *stream.Overseer

	deadline *time.Timer

	tracer log.Tracer
}

func freshBreakpointNode(hub *LedgerHub, nodeUUID p2p.ID, foundation int64, altitude int64) *breakpointNode {
	node := &breakpointNode{
		hub:       hub,
		id:         nodeUUID,
		foundation:       foundation,
		altitude:     altitude,
		countAwaiting: 0,
		tracer:     log.FreshNooperationTracer(),
	}
	return node
}

func (node *breakpointNode) assignTracer(l log.Tracer) {
	node.tracer = l
}

func (node *breakpointNode) restoreOverseer() {
	node.obtainOverseer = stream.New(time.Second, time.Second*40)
	primaryItem := float64(minimumObtainFrequency) * math.E
	node.obtainOverseer.AssignRemain(primaryItem)
}

func (node *breakpointNode) restoreDeadline() {
	if node.deadline == nil {
		node.deadline = time.AfterFunc(nodeDeadline, node.uponDeadline)
	} else {
		node.deadline.Reset(nodeDeadline)
	}
}

func (node *breakpointNode) increaseAwaiting() {
	if node.countAwaiting == 0 {
		node.restoreOverseer()
		node.restoreDeadline()
	}
	node.countAwaiting++
}

func (node *breakpointNode) decreaseAwaiting(obtainExtent int) {
	node.countAwaiting--
	if node.countAwaiting == 0 {
		node.deadline.Stop()
	} else {
		node.obtainOverseer.Revise(obtainExtent)
		node.restoreDeadline()
	}
}

func (node *breakpointNode) uponDeadline() {
	node.hub.mtx.Lock()
	defer node.hub.mtx.Unlock()

	node.hub.transmitFailure(FaultNodeDeadline, node.id)
	node.tracer.Failure("REDACTED", "REDACTED", FaultNodeDeadline, "REDACTED", nodeDeadline)
	node.actedDeadline = true
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
type breakpointSolicitor struct {
	facility.FoundationFacility

	hub        *LedgerHub
	altitude      int64
	attainedLedgerStream  chan struct{}
	reiterateStream      chan p2p.ID //
	freshAltitudeStream chan int64

	mtx          commitchronize.Exclusion
	nodeUUID       p2p.ID
	ordinalNodeUUID p2p.ID //
	attainedLedgerOriginating p2p.ID
	ledger        *kinds.Ledger
	addnEndorse    *kinds.ExpandedEndorse
}

func freshBreakpointSolicitor(hub *LedgerHub, altitude int64) *breakpointSolicitor {
	bpr := &breakpointSolicitor{
		hub:        hub,
		altitude:      altitude,
		attainedLedgerStream:  make(chan struct{}, 1),
		reiterateStream:      make(chan p2p.ID, 1),
		freshAltitudeStream: make(chan int64, 1),

		nodeUUID:       "REDACTED",
		ordinalNodeUUID: "REDACTED",
		ledger:        nil,
	}
	bpr.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", bpr)
	return bpr
}

func (bpr *breakpointSolicitor) UponInitiate() error {
	go bpr.solicitProcedure()
	return nil
}

//
func (bpr *breakpointSolicitor) assignLedger(ledger *kinds.Ledger, addnEndorse *kinds.ExpandedEndorse, nodeUUID p2p.ID) bool {
	bpr.mtx.Lock()
	if bpr.nodeUUID != nodeUUID && bpr.ordinalNodeUUID != nodeUUID {
		bpr.mtx.Unlock()
		return false
	}
	if bpr.ledger != nil {
		bpr.mtx.Unlock()
		return true //
	}

	bpr.ledger = ledger
	bpr.addnEndorse = addnEndorse
	bpr.attainedLedgerOriginating = nodeUUID
	bpr.mtx.Unlock()

	select {
	case bpr.attainedLedgerStream <- struct{}{}:
	default:
	}
	return true
}

func (bpr *breakpointSolicitor) obtainLedger() *kinds.Ledger {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.ledger
}

func (bpr *breakpointSolicitor) obtainExpandedEndorse() *kinds.ExpandedEndorse {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.addnEndorse
}

//
func (bpr *breakpointSolicitor) solicitedOriginating() []p2p.ID {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	nodeIDXDstore := make([]p2p.ID, 0, 2)
	if bpr.nodeUUID != "REDACTED" {
		nodeIDXDstore = append(nodeIDXDstore, bpr.nodeUUID)
	}
	if bpr.ordinalNodeUUID != "REDACTED" {
		nodeIDXDstore = append(nodeIDXDstore, bpr.ordinalNodeUUID)
	}
	return nodeIDXDstore
}

//
func (bpr *breakpointSolicitor) actedSolicitOriginating(nodeUUID p2p.ID) bool {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.nodeUUID == nodeUUID || bpr.ordinalNodeUUID == nodeUUID
}

//
func (bpr *breakpointSolicitor) attainedLedgerOriginatingNodeUUID() p2p.ID {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()
	return bpr.attainedLedgerOriginating
}

//
func (bpr *breakpointSolicitor) restore(nodeUUID p2p.ID) (discardedLedger bool) {
	bpr.mtx.Lock()
	defer bpr.mtx.Unlock()

	//
	if bpr.attainedLedgerOriginating == nodeUUID {
		bpr.ledger = nil
		bpr.addnEndorse = nil
		bpr.attainedLedgerOriginating = "REDACTED"
		discardedLedger = true
		atomic.AddInt32(&bpr.hub.countAwaiting, 1)
	}

	if bpr.nodeUUID == nodeUUID {
		bpr.nodeUUID = "REDACTED"
	} else {
		bpr.ordinalNodeUUID = "REDACTED"
	}

	return discardedLedger
}

//
//
//
func (bpr *breakpointSolicitor) reiterate(nodeUUID p2p.ID) {
	select {
	case bpr.reiterateStream <- nodeUUID:
	default:
	}
}

func (bpr *breakpointSolicitor) selectNodeAlsoTransmitSolicit() {
	bpr.mtx.Lock()
	ordinalNodeUUID := bpr.ordinalNodeUUID
	bpr.mtx.Unlock()

	var node *breakpointNode
SELECT_NODE_CYCLE:
	for {
		if !bpr.EqualsActive() || !bpr.hub.EqualsActive() {
			return
		}
		node = bpr.hub.selectIncreaseAccessibleNode(bpr.altitude, ordinalNodeUUID)
		if node == nil {
			bpr.Tracer.Diagnose("REDACTED", "REDACTED", bpr.altitude)
			time.Sleep(solicitDurationMSEC * time.Millisecond)
			continue SELECT_NODE_CYCLE
		}
		break SELECT_NODE_CYCLE
	}
	bpr.mtx.Lock()
	bpr.nodeUUID = node.id
	bpr.mtx.Unlock()

	bpr.hub.transmitSolicit(bpr.altitude, node.id)
}

//
//
func (bpr *breakpointSolicitor) selectOrdinalNodeAlsoTransmitSolicit() (selected bool) {
	bpr.mtx.Lock()
	if bpr.ordinalNodeUUID != "REDACTED" {
		bpr.mtx.Unlock()
		return false
	}
	nodeUUID := bpr.nodeUUID
	bpr.mtx.Unlock()

	ordinalNode := bpr.hub.selectIncreaseAccessibleNode(bpr.altitude, nodeUUID)
	if ordinalNode != nil {
		bpr.mtx.Lock()
		bpr.ordinalNodeUUID = ordinalNode.id
		bpr.mtx.Unlock()

		bpr.hub.transmitSolicit(bpr.altitude, ordinalNode.id)
		return true
	}

	return false
}

//
func (bpr *breakpointSolicitor) freshAltitude(altitude int64) {
	select {
	case bpr.freshAltitudeStream <- altitude:
	default:
	}
}

//
//
func (bpr *breakpointSolicitor) solicitProcedure() {
	attainedLedger := false

EXTERNAL_CYCLE:
	for {
		bpr.selectNodeAlsoTransmitSolicit()

		hubAltitude := bpr.hub.Altitude()
		if bpr.altitude-hubAltitude < minimumLedgersForeachUniqueSolicit {
			bpr.selectOrdinalNodeAlsoTransmitSolicit()
		}

		reissueClock := time.NewTimer(solicitReissueMoments * time.Second)
		defer reissueClock.Stop()

		for {
			select {
			case <-bpr.hub.Exit():
				if err := bpr.Halt(); err != nil {
					bpr.Tracer.Failure("REDACTED", "REDACTED", err)
				}
				return
			case <-bpr.Exit():
				return
			case <-reissueClock.C:
				if !attainedLedger {
					bpr.Tracer.Diagnose("REDACTED", "REDACTED", bpr.altitude, "REDACTED", bpr.nodeUUID, "REDACTED", bpr.ordinalNodeUUID)
					bpr.restore(bpr.nodeUUID)
					bpr.restore(bpr.ordinalNodeUUID)
					continue EXTERNAL_CYCLE
				}
			case nodeUUID := <-bpr.reiterateStream:
				if bpr.actedSolicitOriginating(nodeUUID) {
					discardedLedger := bpr.restore(nodeUUID)
					if discardedLedger {
						attainedLedger = false
					}
				}
				//
				//
				if len(bpr.solicitedOriginating()) == 0 {
					reissueClock.Stop()
					continue EXTERNAL_CYCLE
				}
			case freshAltitude := <-bpr.freshAltitudeStream:
				if !attainedLedger && bpr.altitude-freshAltitude < minimumLedgersForeachUniqueSolicit {
					//
					//
					//
					//
					if selected := bpr.selectOrdinalNodeAlsoTransmitSolicit(); selected {
						if !reissueClock.Stop() {
							<-reissueClock.C
						}
						reissueClock.Reset(solicitReissueMoments * time.Second)
					}
				}
			case <-bpr.attainedLedgerStream:
				attainedLedger = true
				//
				//
			}
		}
	}
}
