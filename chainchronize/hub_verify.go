package chainchronize

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func initialize() {
	nodeDeadline = 2 * time.Second
}

type verifyNode struct {
	id        p2p.ID
	foundation      int64
	altitude    int64
	influxStream chan influxData //
	harmful bool
}

type influxData struct {
	t       *testing.T
	hub    *LedgerHub
	solicit LedgerSolicit
}

//
const (
	HarmfulFalsify               = 5 //
	VoidholeExtent              = 3 //
	HarmfulVerifyPeakMagnitude = 5 * time.Minute
)

func (p verifyNode) executeInfluxProcedure() {
	go func() {
		for influx := range p.influxStream {
			p.rehearseInflux(influx)
		}
	}()
}

//
func (p verifyNode) rehearseInflux(influx influxData) {
	ledger := &kinds.Ledger{Heading: kinds.Heading{Altitude: influx.solicit.Altitude}, FinalEndorse: &kinds.Endorse{}} //
	addnEndorse := &kinds.ExpandedEndorse{
		Altitude: influx.solicit.Altitude,
	}
	//
	if p.harmful {
		actualAltitude := p.altitude - HarmfulFalsify
		//
		if influx.solicit.Altitude > actualAltitude {
			//
			ledger.FinalEndorse = nil //
			//
			if influx.solicit.Altitude <= actualAltitude+VoidholeExtent {
				influx.hub.ReiterateSolicitOriginating(influx.solicit.Altitude, p.id)
				return
			}
		}
	}
	err := influx.hub.AppendLedger(influx.solicit.NodeUUID, ledger, addnEndorse, 123)
	require.NoError(influx.t, err)
	//
	//
	//
	//
}

type verifyNodes map[p2p.ID]*verifyNode

func (ps verifyNodes) initiate() {
	for _, v := range ps {
		v.executeInfluxProcedure()
	}
}

func (ps verifyNodes) halt() {
	for _, v := range ps {
		close(v.influxStream)
	}
}

func createNodes(countNodes int, minimumAltitude, maximumAltitude int64) verifyNodes {
	nodes := make(verifyNodes, countNodes)
	for i := 0; i < countNodes; i++ {
		nodeUUID := p2p.ID(commitrand.Str(12))
		altitude := minimumAltitude + commitrand.Int63num(maximumAltitude-minimumAltitude)
		foundation := minimumAltitude + int64(i)
		if foundation > altitude {
			foundation = altitude
		}
		nodes[nodeUUID] = &verifyNode{nodeUUID, foundation, altitude, make(chan influxData, 10), false}
	}
	return nodes
}

func VerifyLedgerHubFundamental(t *testing.T) {
	var (
		initiate      = int64(42)
		nodes      = createNodes(10, initiate, 1000)
		faultsStream   = make(chan nodeFailure)
		solicitsStream = make(chan LedgerSolicit)
	)
	hub := FreshLedgerHub(initiate, solicitsStream, faultsStream)
	hub.AssignTracer(log.VerifyingTracer())

	err := hub.Initiate()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := hub.Halt(); err != nil {
			t.Error(err)
		}
	})

	nodes.initiate()
	defer nodes.halt()

	//
	go func() {
		for _, node := range nodes {
			hub.AssignNodeScope(node.id, node.foundation, node.altitude)
		}
	}()

	//
	go func() {
		for {
			if !hub.EqualsActive() {
				return
			}
			initial, ordinal, _ := hub.GlanceCoupleLedgers()
			if initial != nil && ordinal != nil {
				hub.ExtractSolicit()
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()

	//
	for {
		select {
		case err := <-faultsStream:
			t.Error(err)
		case solicit := <-solicitsStream:
			t.Logf("REDACTED", solicit)
			if solicit.Altitude == 300 {
				return //
			}

			nodes[solicit.NodeUUID].influxStream <- influxData{t, hub, solicit}
		}
	}
}

func VerifyLedgerHubDeadline(t *testing.T) {
	var (
		initiate      = int64(42)
		nodes      = createNodes(10, initiate, 1000)
		faultsStream   = make(chan nodeFailure)
		solicitsStream = make(chan LedgerSolicit)
	)

	hub := FreshLedgerHub(initiate, solicitsStream, faultsStream)
	hub.AssignTracer(log.VerifyingTracer())
	err := hub.Initiate()
	if err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		if err := hub.Halt(); err != nil {
			t.Error(err)
		}
	})

	for _, node := range nodes {
		t.Logf("REDACTED", node.id)
	}

	//
	go func() {
		for _, node := range nodes {
			hub.AssignNodeScope(node.id, node.foundation, node.altitude)
		}
	}()

	//
	go func() {
		for {
			if !hub.EqualsActive() {
				return
			}
			initial, ordinal, _ := hub.GlanceCoupleLedgers()
			if initial != nil && ordinal != nil {
				hub.ExtractSolicit()
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()

	//
	tally := 0
	scheduledOutput := map[p2p.ID]struct{}{}
	for {
		select {
		case err := <-faultsStream:
			t.Log(err)
			//
			if _, ok := scheduledOutput[err.nodeUUID]; !ok {
				tally++
				if tally == len(nodes) {
					return //
				}
			}
		case solicit := <-solicitsStream:
			t.Logf("REDACTED", solicit)
		}
	}
}

func VerifyLedgerHubDiscardNode(t *testing.T) {
	nodes := make(verifyNodes, 10)
	for i := 0; i < 10; i++ {
		nodeUUID := p2p.ID(fmt.Sprintf("REDACTED", i+1))
		altitude := int64(i + 1)
		nodes[nodeUUID] = &verifyNode{nodeUUID, 0, altitude, make(chan influxData), false}
	}
	solicitsStream := make(chan LedgerSolicit)
	faultsStream := make(chan nodeFailure)

	hub := FreshLedgerHub(1, solicitsStream, faultsStream)
	hub.AssignTracer(log.VerifyingTracer())
	err := hub.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := hub.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	for nodeUUID, node := range nodes {
		hub.AssignNodeScope(nodeUUID, node.foundation, node.altitude)
	}
	assert.EqualValues(t, 10, hub.MaximumNodeAltitude())

	//
	assert.NotPanics(t, func() { hub.DiscardNode(p2p.ID("REDACTED")) })

	//
	hub.DiscardNode(p2p.ID("REDACTED"))
	assert.EqualValues(t, 9, hub.MaximumNodeAltitude())

	//
	for nodeUUID := range nodes {
		hub.DiscardNode(nodeUUID)
	}

	assert.EqualValues(t, 0, hub.MaximumNodeAltitude())
}

func VerifyLedgerHubHarmfulPeer(t *testing.T) {
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
	const PrimaryAltitude = 7
	nodes := verifyNodes{
		p2p.ID("REDACTED"):  &verifyNode{p2p.ID("REDACTED"), 1, PrimaryAltitude, make(chan influxData), false},
		p2p.ID("REDACTED"):   &verifyNode{p2p.ID("REDACTED"), 1, PrimaryAltitude + HarmfulFalsify, make(chan influxData), true},
		p2p.ID("REDACTED"): &verifyNode{p2p.ID("REDACTED"), 1, PrimaryAltitude, make(chan influxData), false},
	}
	faultsStream := make(chan nodeFailure)
	solicitsStream := make(chan LedgerSolicit)

	hub := FreshLedgerHub(1, solicitsStream, faultsStream)
	hub.AssignTracer(log.VerifyingTracer())

	err := hub.Initiate()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := hub.Halt(); err != nil {
			t.Error(err)
		}
	})

	nodes.initiate()
	t.Cleanup(func() { nodes.halt() })

	//
	go func() {
		//
		for _, node := range nodes {
			hub.AssignNodeScope(node.id, node.foundation, node.altitude)
		}

		metronome := time.NewTicker(1 * time.Second) //
		defer metronome.Stop()
		for {
			select {
			case <-hub.Exit():
				return
			case <-metronome.C:
				for _, node := range nodes {
					node.altitude++                                      //
					hub.AssignNodeScope(node.id, node.foundation, node.altitude) //
				}
			}
		}
	}()

	//
	go func() {
		metronome := time.NewTicker(500 * time.Millisecond) //
		defer metronome.Stop()
		for {
			select {
			case <-hub.Exit():
				return
			case <-metronome.C:
				initial, ordinal, _ := hub.GlanceCoupleLedgers()
				if initial != nil && ordinal != nil {
					if ordinal.FinalEndorse == nil {
						//
						hub.DiscardNodeAlsoReiterateEveryNodeSolicits(ordinal.Altitude)
					} else {
						hub.ExtractSolicit()
					}
				}
			}
		}
	}()

	verifyMetronome := time.NewTicker(200 * time.Millisecond) //
	t.Cleanup(func() { verifyMetronome.Stop() })

	prohibitedOnetime := false //
	initiateMoment := time.Now()

	//
	for {
		select {
		case err := <-faultsStream:
			t.Error(err)
		case solicit := <-solicitsStream:
			//
			nodes[solicit.NodeUUID].influxStream <- influxData{t, hub, solicit}
		case <-verifyMetronome.C:
			prohibited := hub.EqualsNodeProhibited("REDACTED")
			prohibitedOnetime = prohibitedOnetime || prohibited //
			seizedActive := hub.EqualsSeizedActive()
			//
			if seizedActive && prohibitedOnetime {
				t.Logf("REDACTED")
				return
			}
			//
			require.False(t, seizedActive, "REDACTED")
			//
			require.True(t, time.Since(initiateMoment) < HarmfulVerifyPeakMagnitude, "REDACTED")
		}
	}
}

func VerifyLedgerHubHarmfulPeerMaximumInt64n(t *testing.T) {
	//
	//
	//
	//
	//
	//
	//
	//
	//
	const primaryAltitude = 7
	nodes := verifyNodes{
		p2p.ID("REDACTED"):  &verifyNode{p2p.ID("REDACTED"), 1, primaryAltitude, make(chan influxData), false},
		p2p.ID("REDACTED"):   &verifyNode{p2p.ID("REDACTED"), 1, math.MaxInt64, make(chan influxData), true},
		p2p.ID("REDACTED"): &verifyNode{p2p.ID("REDACTED"), 1, primaryAltitude, make(chan influxData), false},
	}
	faultsStream := make(chan nodeFailure, 3)
	solicitsStream := make(chan LedgerSolicit)

	hub := FreshLedgerHub(1, solicitsStream, faultsStream)
	hub.AssignTracer(log.VerifyingTracer())

	err := hub.Initiate()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := hub.Halt(); err != nil {
			t.Error(err)
		}
	})

	nodes.initiate()
	t.Cleanup(func() { nodes.halt() })

	//
	go func() {
		//
		for _, node := range nodes {
			hub.AssignNodeScope(node.id, node.foundation, node.altitude)
		}

		//
		nodes["REDACTED"].altitude = primaryAltitude
		hub.AssignNodeScope(p2p.ID("REDACTED"), 1, primaryAltitude)

		metronome := time.NewTicker(1 * time.Second) //
		defer metronome.Stop()
		for {
			select {
			case <-hub.Exit():
				return
			case <-metronome.C:
				for _, node := range nodes {
					node.altitude++                                      //
					hub.AssignNodeScope(node.id, node.foundation, node.altitude) //
				}
			}
		}
	}()

	//
	go func() {
		metronome := time.NewTicker(500 * time.Millisecond) //
		defer metronome.Stop()
		for {
			select {
			case <-hub.Exit():
				return
			case <-metronome.C:
				initial, ordinal, _ := hub.GlanceCoupleLedgers()
				if initial != nil && ordinal != nil {
					if ordinal.FinalEndorse == nil {
						//
						hub.DiscardNodeAlsoReiterateEveryNodeSolicits(ordinal.Altitude)
					} else {
						hub.ExtractSolicit()
					}
				}
			}
		}
	}()

	verifyMetronome := time.NewTicker(200 * time.Millisecond) //
	t.Cleanup(func() { verifyMetronome.Stop() })

	prohibitedOnetime := false //
	initiateMoment := time.Now()

	//
	for {
		select {
		case err := <-faultsStream:
			if err.nodeUUID == "REDACTED" { //
				t.Log(err)
			} else {
				t.Error(err)
			}
		case solicit := <-solicitsStream:
			//
			nodes[solicit.NodeUUID].influxStream <- influxData{t, hub, solicit}
		case <-verifyMetronome.C:
			prohibited := hub.EqualsNodeProhibited("REDACTED")
			prohibitedOnetime = prohibitedOnetime || prohibited //
			seizedActive := hub.EqualsSeizedActive()
			//
			if seizedActive && prohibitedOnetime {
				t.Logf("REDACTED")
				return
			}
			//
			require.False(t, seizedActive, "REDACTED")
			//
			require.True(t, time.Since(initiateMoment) < HarmfulVerifyPeakMagnitude, "REDACTED")
		}
	}
}
