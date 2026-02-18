package chainconnect

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/kinds"
)

func init() {
	nodeDeadline = 2 * time.Second
}

type verifyNode struct {
	id        p2p.ID
	root      int64
	level    int64
	influxChan chan influxData //
	harmful bool
}

type influxData struct {
	t       *testing.T
	depository    *LedgerDepository
	query LedgerQuery
}

//
const (
	HarmfulMislead               = 5 //
	VoidVolume              = 3 //
	HarmfulVerifyGreatestExtent = 5 * time.Minute
)

func (p verifyNode) executeInfluxProcess() {
	go func() {
		for influx := range p.influxChan {
			p.mimicInflux(influx)
		}
	}()
}

//
func (p verifyNode) mimicInflux(influx influxData) {
	ledger := &kinds.Ledger{Heading: kinds.Heading{Level: influx.query.Level}, FinalEndorse: &kinds.Endorse{}} //
	extensionEndorse := &kinds.ExpandedEndorse{
		Level: influx.query.Level,
	}
	//
	if p.harmful {
		actualLevel := p.level - HarmfulMislead
		//
		if influx.query.Level > actualLevel {
			//
			ledger.FinalEndorse = nil //
			//
			if influx.query.Level <= actualLevel+VoidVolume {
				influx.depository.ReworkQueryFrom(influx.query.Level, p.id)
				return
			}
		}
	}
	err := influx.depository.AppendLedger(influx.query.NodeUID, ledger, extensionEndorse, 123)
	require.NoError(influx.t, err)
	//
	//
	//
	//
}

type verifyNodes map[p2p.ID]*verifyNode

func (ps verifyNodes) begin() {
	for _, v := range ps {
		v.executeInfluxProcess()
	}
}

func (ps verifyNodes) halt() {
	for _, v := range ps {
		close(v.influxChan)
	}
}

func createNodes(countNodes int, minimumLevel, maximumLevel int64) verifyNodes {
	nodes := make(verifyNodes, countNodes)
	for i := 0; i < countNodes; i++ {
		nodeUID := p2p.ID(engineseed.Str(12))
		level := minimumLevel + engineseed.Int64count(maximumLevel-minimumLevel)
		root := minimumLevel + int64(i)
		if root > level {
			root = level
		}
		nodes[nodeUID] = &verifyNode{nodeUID, root, level, make(chan influxData, 10), false}
	}
	return nodes
}

func VerifyLedgerDepositorySimple(t *testing.T) {
	var (
		begin      = int64(42)
		nodes      = createNodes(10, begin, 1000)
		faultsChan   = make(chan nodeFault)
		queriesChan = make(chan LedgerQuery)
	)
	depository := NewLedgerDepository(begin, queriesChan, faultsChan)
	depository.AssignTracer(log.VerifyingTracer())

	err := depository.Begin()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := depository.Halt(); err != nil {
			t.Error(err)
		}
	})

	nodes.begin()
	defer nodes.halt()

	//
	go func() {
		for _, node := range nodes {
			depository.AssignNodeScope(node.id, node.root, node.level)
		}
	}()

	//
	go func() {
		for {
			if !depository.IsActive() {
				return
			}
			initial, moment, _ := depository.GlanceDualLedgers()
			if initial != nil && moment != nil {
				depository.EjectQuery()
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()

	//
	for {
		select {
		case err := <-faultsChan:
			t.Error(err)
		case query := <-queriesChan:
			t.Logf("REDACTED", query)
			if query.Level == 300 {
				return //
			}

			nodes[query.NodeUID].influxChan <- influxData{t, depository, query}
		}
	}
}

func VerifyLedgerDepositoryDeadline(t *testing.T) {
	var (
		begin      = int64(42)
		nodes      = createNodes(10, begin, 1000)
		faultsChan   = make(chan nodeFault)
		queriesChan = make(chan LedgerQuery)
	)

	depository := NewLedgerDepository(begin, queriesChan, faultsChan)
	depository.AssignTracer(log.VerifyingTracer())
	err := depository.Begin()
	if err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		if err := depository.Halt(); err != nil {
			t.Error(err)
		}
	})

	for _, node := range nodes {
		t.Logf("REDACTED", node.id)
	}

	//
	go func() {
		for _, node := range nodes {
			depository.AssignNodeScope(node.id, node.root, node.level)
		}
	}()

	//
	go func() {
		for {
			if !depository.IsActive() {
				return
			}
			initial, moment, _ := depository.GlanceDualLedgers()
			if initial != nil && moment != nil {
				depository.EjectQuery()
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()

	//
	tally := 0
	scheduledOut := map[p2p.ID]struct{}{}
	for {
		select {
		case err := <-faultsChan:
			t.Log(err)
			//
			if _, ok := scheduledOut[err.nodeUID]; !ok {
				tally++
				if tally == len(nodes) {
					return //
				}
			}
		case query := <-queriesChan:
			t.Logf("REDACTED", query)
		}
	}
}

func VerifyLedgerDepositoryDeleteNode(t *testing.T) {
	nodes := make(verifyNodes, 10)
	for i := 0; i < 10; i++ {
		nodeUID := p2p.ID(fmt.Sprintf("REDACTED", i+1))
		level := int64(i + 1)
		nodes[nodeUID] = &verifyNode{nodeUID, 0, level, make(chan influxData), false}
	}
	queriesChan := make(chan LedgerQuery)
	faultsChan := make(chan nodeFault)

	depository := NewLedgerDepository(1, queriesChan, faultsChan)
	depository.AssignTracer(log.VerifyingTracer())
	err := depository.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := depository.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	for nodeUID, node := range nodes {
		depository.AssignNodeScope(nodeUID, node.root, node.level)
	}
	assert.EqualValues(t, 10, depository.MaximumNodeLevel())

	//
	assert.NotPanics(t, func() { depository.DeleteNode(p2p.ID("REDACTED")) })

	//
	depository.DeleteNode(p2p.ID("REDACTED"))
	assert.EqualValues(t, 9, depository.MaximumNodeLevel())

	//
	for nodeUID := range nodes {
		depository.DeleteNode(nodeUID)
	}

	assert.EqualValues(t, 0, depository.MaximumNodeLevel())
}

func VerifyLedgerDepositoryHarmfulMember(t *testing.T) {
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
	const PrimaryLevel = 7
	nodes := verifyNodes{
		p2p.ID("REDACTED"):  &verifyNode{p2p.ID("REDACTED"), 1, PrimaryLevel, make(chan influxData), false},
		p2p.ID("REDACTED"):   &verifyNode{p2p.ID("REDACTED"), 1, PrimaryLevel + HarmfulMislead, make(chan influxData), true},
		p2p.ID("REDACTED"): &verifyNode{p2p.ID("REDACTED"), 1, PrimaryLevel, make(chan influxData), false},
	}
	faultsChan := make(chan nodeFault)
	queriesChan := make(chan LedgerQuery)

	depository := NewLedgerDepository(1, queriesChan, faultsChan)
	depository.AssignTracer(log.VerifyingTracer())

	err := depository.Begin()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := depository.Halt(); err != nil {
			t.Error(err)
		}
	})

	nodes.begin()
	t.Cleanup(func() { nodes.halt() })

	//
	go func() {
		//
		for _, node := range nodes {
			depository.AssignNodeScope(node.id, node.root, node.level)
		}

		timer := time.NewTicker(1 * time.Second) //
		defer timer.Stop()
		for {
			select {
			case <-depository.Exit():
				return
			case <-timer.C:
				for _, node := range nodes {
					node.level++                                      //
					depository.AssignNodeScope(node.id, node.root, node.level) //
				}
			}
		}
	}()

	//
	go func() {
		timer := time.NewTicker(500 * time.Millisecond) //
		defer timer.Stop()
		for {
			select {
			case <-depository.Exit():
				return
			case <-timer.C:
				initial, moment, _ := depository.GlanceDualLedgers()
				if initial != nil && moment != nil {
					if moment.FinalEndorse == nil {
						//
						depository.DeleteNodeAndReworkAllNodeQueries(moment.Level)
					} else {
						depository.EjectQuery()
					}
				}
			}
		}
	}()

	verifyTimer := time.NewTicker(200 * time.Millisecond) //
	t.Cleanup(func() { verifyTimer.Stop() })

	prohibitedOnce := false //
	beginMoment := time.Now()

	//
	for {
		select {
		case err := <-faultsChan:
			t.Error(err)
		case query := <-queriesChan:
			//
			nodes[query.NodeUID].influxChan <- influxData{t, depository, query}
		case <-verifyTimer.C:
			prohibited := depository.IsNodeProhibited("REDACTED")
			prohibitedOnce = prohibitedOnce || prohibited //
			seizedUp := depository.IsSeizedUp()
			//
			if seizedUp && prohibitedOnce {
				t.Logf("REDACTED")
				return
			}
			//
			require.False(t, seizedUp, "REDACTED")
			//
			require.True(t, time.Since(beginMoment) < HarmfulVerifyGreatestExtent, "REDACTED")
		}
	}
}

func VerifyLedgerDepositoryHarmfulMemberMaximumInt64(t *testing.T) {
	//
	//
	//
	//
	//
	//
	//
	//
	//
	const primaryLevel = 7
	nodes := verifyNodes{
		p2p.ID("REDACTED"):  &verifyNode{p2p.ID("REDACTED"), 1, primaryLevel, make(chan influxData), false},
		p2p.ID("REDACTED"):   &verifyNode{p2p.ID("REDACTED"), 1, math.MaxInt64, make(chan influxData), true},
		p2p.ID("REDACTED"): &verifyNode{p2p.ID("REDACTED"), 1, primaryLevel, make(chan influxData), false},
	}
	faultsChan := make(chan nodeFault, 3)
	queriesChan := make(chan LedgerQuery)

	depository := NewLedgerDepository(1, queriesChan, faultsChan)
	depository.AssignTracer(log.VerifyingTracer())

	err := depository.Begin()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := depository.Halt(); err != nil {
			t.Error(err)
		}
	})

	nodes.begin()
	t.Cleanup(func() { nodes.halt() })

	//
	go func() {
		//
		for _, node := range nodes {
			depository.AssignNodeScope(node.id, node.root, node.level)
		}

		//
		nodes["REDACTED"].level = primaryLevel
		depository.AssignNodeScope(p2p.ID("REDACTED"), 1, primaryLevel)

		timer := time.NewTicker(1 * time.Second) //
		defer timer.Stop()
		for {
			select {
			case <-depository.Exit():
				return
			case <-timer.C:
				for _, node := range nodes {
					node.level++                                      //
					depository.AssignNodeScope(node.id, node.root, node.level) //
				}
			}
		}
	}()

	//
	go func() {
		timer := time.NewTicker(500 * time.Millisecond) //
		defer timer.Stop()
		for {
			select {
			case <-depository.Exit():
				return
			case <-timer.C:
				initial, moment, _ := depository.GlanceDualLedgers()
				if initial != nil && moment != nil {
					if moment.FinalEndorse == nil {
						//
						depository.DeleteNodeAndReworkAllNodeQueries(moment.Level)
					} else {
						depository.EjectQuery()
					}
				}
			}
		}
	}()

	verifyTimer := time.NewTicker(200 * time.Millisecond) //
	t.Cleanup(func() { verifyTimer.Stop() })

	prohibitedOnce := false //
	beginMoment := time.Now()

	//
	for {
		select {
		case err := <-faultsChan:
			if err.nodeUID == "REDACTED" { //
				t.Log(err)
			} else {
				t.Error(err)
			}
		case query := <-queriesChan:
			//
			nodes[query.NodeUID].influxChan <- influxData{t, depository, query}
		case <-verifyTimer.C:
			prohibited := depository.IsNodeProhibited("REDACTED")
			prohibitedOnce = prohibitedOnce || prohibited //
			seizedUp := depository.IsSeizedUp()
			//
			if seizedUp && prohibitedOnce {
				t.Logf("REDACTED")
				return
			}
			//
			require.False(t, seizedUp, "REDACTED")
			//
			require.True(t, time.Since(beginMoment) < HarmfulVerifyGreatestExtent, "REDACTED")
		}
	}
}
