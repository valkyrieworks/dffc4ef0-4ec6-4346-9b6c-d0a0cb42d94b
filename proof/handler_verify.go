package proofs_test

import (
	"encoding/hex"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/go-kit/log/term"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/proof"
	"github.com/valkyrieworks/proof/simulations"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	netpeersims "github.com/valkyrieworks/p2p/simulations"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

var (
	countProof = 10
	deadline     = 120 * time.Second //
)

//
//
//
//
func VerifyHandlerMulticastProof(t *testing.T) {
	settings := cfg.VerifySettings()
	N := 7

	//
	statusDSz := make([]sm.Depot, N)
	val := kinds.NewEmulatePV()
	//
	level := int64(countProof) + 10
	for i := 0; i < N; i++ {
		statusDSz[i] = bootstrapRatifierStatus(val, level)
	}

	//
	handlers, repositories := createAndLinkHandlersAndRepositories(settings, statusDSz)

	//
	for _, r := range handlers {
		for _, node := range r.Router.Nodes().Clone() {
			ps := nodeStatus{level}
			node.Set(kinds.NodeStatusKey, ps)
		}
	}

	//
	//
	evtCatalog := transmitProof(t, repositories[0], val, countProof)
	waitForProof(t, evtCatalog, repositories)
}

//
//
//
func VerifyHandlerDiscerningMulticast(t *testing.T) {
	settings := cfg.VerifySettings()

	val := kinds.NewEmulatePV()
	level1 := int64(countProof) + 10
	level2 := int64(countProof) / 2

	//
	statusDB1 := bootstrapRatifierStatus(val, level1)
	statusDB2 := bootstrapRatifierStatus(val, level2)

	//
	handlers, repositories := createAndLinkHandlersAndRepositories(settings, []sm.Depot{statusDB1, statusDB2})

	//
	for _, r := range handlers {
		for _, node := range r.Router.Nodes().Clone() {
			ps := nodeStatus{level1}
			node.Set(kinds.NodeStatusKey, ps)
		}
	}

	//
	node := handlers[0].Router.Nodes().Clone()[0]
	ps := nodeStatus{level2}
	node.Set(kinds.NodeStatusKey, ps)

	//
	evtCatalog := transmitProof(t, repositories[0], val, countProof)

	//
	waitForProof(t, evtCatalog[:countProof/2-1], []*proof.Depository{repositories[1]})

	//
	nodes := handlers[1].Router.Nodes().Clone()
	assert.Len(t, nodes, 1)
}

//
//
//
//
//
func VerifyHandlersGossipNoConfirmedProof(t *testing.T) {
	settings := cfg.VerifySettings()

	val := kinds.NewEmulatePV()
	var level int64 = 10

	//
	statusDB1 := bootstrapRatifierStatus(val, level-1)
	statusDB2 := bootstrapRatifierStatus(val, level-2)
	status, err := statusDB1.Import()
	require.NoError(t, err)
	status.FinalLedgerLevel++

	//
	handlers, repositories := createAndLinkHandlersAndRepositories(settings, []sm.Depot{statusDB1, statusDB2})

	evtCatalog := transmitProof(t, repositories[0], val, 2)
	repositories[0].Modify(status, evtCatalog)
	require.EqualValues(t, uint32(0), repositories[0].Volume())

	time.Sleep(100 * time.Millisecond)

	node := handlers[0].Router.Nodes().Clone()[0]
	ps := nodeStatus{level - 2}
	node.Set(kinds.NodeStatusKey, ps)

	node = handlers[1].Router.Nodes().Clone()[0]
	ps = nodeStatus{level}
	node.Set(kinds.NodeStatusKey, ps)

	//
	time.Sleep(300 * time.Millisecond)

	//
	assert.Equal(t, uint32(0), repositories[1].Volume(), "REDACTED")

	//
	evtCatalog = make([]kinds.Proof, 3)
	for i := 0; i < 3; i++ {
		ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level-3+int64(i),
			time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), val, status.LedgerUID)
		require.NoError(t, err)
		err = repositories[0].AppendProof(ev)
		require.NoError(t, err)
		evtCatalog[i] = ev
	}

	//
	time.Sleep(300 * time.Millisecond)

	//
	nodeEvt, _ := repositories[1].AwaitingProof(10000)
	assert.EqualValues(t, []kinds.Proof{evtCatalog[0]}, nodeEvt)

	//
	//
	//
	status.FinalLedgerLevel++
	repositories[0].Modify(status, []kinds.Proof{evtCatalog[2]})
	//
	require.EqualValues(t, uint32(2), repositories[0].Volume())

	//
	repositories[1].Modify(status, kinds.ProofCatalog{})
	node = handlers[0].Router.Nodes().Clone()[0]
	ps = nodeStatus{level}
	node.Set(kinds.NodeStatusKey, ps)

	//
	time.Sleep(300 * time.Millisecond)

	nodeEvt, _ = repositories[1].AwaitingProof(1000)
	assert.EqualValues(t, []kinds.Proof{evtCatalog[0], evtCatalog[1]}, nodeEvt)
}

func VerifyHandlerMulticastProofRamSeepage(t *testing.T) {
	proofTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	proofStore := dbm.NewMemoryStore()
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(
		&kinds.LedgerMeta{Heading: kinds.Heading{Time: proofTime}},
	)
	val := kinds.NewEmulatePV()
	statusDepot := bootstrapRatifierStatus(val, 1)
	depository, err := proof.NewDepository(proofStore, statusDepot, ledgerDepot)
	require.NoError(t, err)

	p := &netpeersims.Node{}

	p.On("REDACTED").Once().Return(true)
	p.On("REDACTED").Return(false)
	//
	//
	defer leaktest.CheckTimeout(t, 10*time.Second)()

	p.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Packet)
		return ok && e.StreamUID == proof.ProofConduit
	})).Return(false)
	exitChan := make(<-chan struct{})
	p.On("REDACTED").Return(exitChan)
	ps := nodeStatus{2}
	p.On("REDACTED", kinds.NodeStatusKey).Return(ps)
	p.On("REDACTED").Return("REDACTED")
	p.On("REDACTED").Return("REDACTED")

	r := proof.NewHandler(depository)
	r.AssignTracer(log.VerifyingTracer())
	r.AppendNode(p)

	_ = transmitProof(t, depository, val, 2)
}

//
//
func proofTracer() log.Tracer {
	return log.VerifyingTracerWithHueFn(func(keyvalues ...any) term.FgBgColor {
		for i := 0; i < len(keyvalues)-1; i += 2 {
			if keyvalues[i] == "REDACTED" {
				return term.FgBgColor{Fg: term.Color(uint8(keyvalues[i+1].(int) + 1))}
			}
		}
		return term.FgBgColor{}
	})
}

//
func createAndLinkHandlersAndRepositories(settings *cfg.Settings, statusShelves []sm.Depot) ([]*proof.Handler,
	[]*proof.Depository,
) {
	N := len(statusShelves)

	handlers := make([]*proof.Handler, N)
	repositories := make([]*proof.Depository, N)
	tracer := proofTracer()
	proofTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < N; i++ {
		proofStore := dbm.NewMemoryStore()
		ledgerDepot := &simulations.LedgerDepot{}
		ledgerDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(
			&kinds.LedgerMeta{Heading: kinds.Heading{Time: proofTime}},
		)
		depository, err := proof.NewDepository(proofStore, statusShelves[i], ledgerDepot)
		if err != nil {
			panic(err)
		}
		repositories[i] = depository
		handlers[i] = proof.NewHandler(depository)
		handlers[i].AssignTracer(tracer.With("REDACTED", i))
	}

	p2p.CreateLinkedRouters(settings.P2P, N, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", handlers[i])
		return s
	}, p2p.Connect2routers)

	return handlers, repositories
}

//
func waitForProof(t *testing.T, evs kinds.ProofCatalog, repositories []*proof.Depository) {
	//
	wg := new(sync.WaitGroup)
	for i := 0; i < len(repositories); i++ {
		wg.Add(1)
		go _waitforproof(t, wg, evs, i, repositories)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	clock := time.After(deadline)
	select {
	case <-clock:
		t.Fatal("REDACTED")
	case <-done:
	}
}

//
func _waitforproof(
	t *testing.T,
	wg *sync.WaitGroup,
	evs kinds.ProofCatalog,
	depositoryIdx int,
	repositories []*proof.Depository,
) {
	eventpool := repositories[depositoryIdx]
	var evtCatalog []kinds.Proof
	ongoingDepositoryVolume := 0
	for ongoingDepositoryVolume != len(evs) {
		evtCatalog, _ = eventpool.AwaitingProof(int64(len(evs) * 500)) //
		ongoingDepositoryVolume = len(evtCatalog)
		time.Sleep(time.Millisecond * 100)
	}

	//
	evtIndex := make(map[string]kinds.Proof)
	for _, e := range evtCatalog {
		evtIndex[string(e.Digest())] = e
	}
	for i, anticipatedEvt := range evs {
		acquiredEvt := evtIndex[string(anticipatedEvt.Digest())]
		assert.Equal(t, anticipatedEvt, acquiredEvt,
			fmt.Sprintf("REDACTED",
				i, depositoryIdx, anticipatedEvt, acquiredEvt))
	}

	wg.Done()
}

func transmitProof(t *testing.T, eventpool *proof.Depository, val kinds.PrivateRatifier, n int) kinds.ProofCatalog {
	evtCatalog := make([]kinds.Proof, n)
	for i := 0; i < n; i++ {
		ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(int64(i+1),
			time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), val, proofSeriesUID)
		require.NoError(t, err)
		err = eventpool.AppendProof(ev)
		require.NoError(t, err)
		evtCatalog[i] = ev
	}
	return evtCatalog
}

type nodeStatus struct {
	level int64
}

func (ps nodeStatus) FetchLevel() int64 {
	return ps.level
}

func instanceBallot(t byte) *kinds.Ballot {
	imprint, err := time.Parse(kinds.TimeLayout, "REDACTED")
	if err != nil {
		panic(err)
	}

	return &kinds.Ballot{
		Kind:      engineproto.AttestedMessageKind(t),
		Level:    3,
		Cycle:     2,
		Timestamp: imprint,
		LedgerUID: kinds.LedgerUID{
			Digest: comethash.Sum([]byte("REDACTED")),
			SegmentAssignHeading: kinds.SegmentAssignHeading{
				Sum: 1000000,
				Digest:  comethash.Sum([]byte("REDACTED")),
			},
		},
		RatifierLocation: vault.LocationDigest([]byte("REDACTED")),
		RatifierOrdinal:   56789,
	}
}

//
func VerifyProofArrays(t *testing.T) {
	val := &kinds.Ratifier{
		Location:     vault.LocationDigest([]byte("REDACTED")),
		PollingEnergy: 10,
	}

	valueCollection := kinds.NewRatifierCollection([]*kinds.Ratifier{val})

	repl, err := kinds.NewReplicatedBallotProof(
		instanceBallot(1),
		instanceBallot(2),
		standardProofTime,
		valueCollection,
	)
	require.NoError(t, err)

	verifyScenarios := []struct {
		verifyLabel     string
		proofCatalog []kinds.Proof
		expirationOctets     string
	}{
		{"REDACTED", []kinds.Proof{repl}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		evi := make([]engineproto.Proof, len(tc.proofCatalog))
		for i := 0; i < len(tc.proofCatalog); i++ {
			ev, err := kinds.ProofToSchema(tc.proofCatalog[i])
			require.NoError(t, err, tc.verifyLabel)
			evi[i] = *ev
		}

		epl := engineproto.ProofCatalog{
			Proof: evi,
		}

		bz, err := epl.Serialize()
		require.NoError(t, err, tc.verifyLabel)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyLabel)

	}
}
