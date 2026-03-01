package proof_test

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

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	netmocks "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/simulations"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
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
	statusDeltaBytes := make([]sm.Depot, N)
	val := kinds.FreshSimulatePRV()
	//
	altitude := int64(countProof) + 10
	for i := 0; i < N; i++ {
		statusDeltaBytes[i] = bootstrapAssessorStatus(val, altitude)
	}

	//
	engines, clusters := createAlsoRelateEnginesAlsoClusters(settings, statusDeltaBytes)

	//
	for _, r := range engines {
		for _, node := range r.Router.Nodes().Duplicate() {
			ps := nodeStatus{altitude}
			node.Set(kinds.NodeStatusToken, ps)
		}
	}

	//
	//
	occurenceCatalog := transmitProof(t, clusters[0], val, countProof)
	pauseForeachProof(t, occurenceCatalog, clusters)
}

//
//
//
func VerifyHandlerDiscerningMulticast(t *testing.T) {
	settings := cfg.VerifySettings()

	val := kinds.FreshSimulatePRV()
	altitude1 := int64(countProof) + 10
	altitude2 := int64(countProof) / 2

	//
	statusDepot1 := bootstrapAssessorStatus(val, altitude1)
	statusDatabase2 := bootstrapAssessorStatus(val, altitude2)

	//
	engines, clusters := createAlsoRelateEnginesAlsoClusters(settings, []sm.Depot{statusDepot1, statusDatabase2})

	//
	for _, r := range engines {
		for _, node := range r.Router.Nodes().Duplicate() {
			ps := nodeStatus{altitude1}
			node.Set(kinds.NodeStatusToken, ps)
		}
	}

	//
	node := engines[0].Router.Nodes().Duplicate()[0]
	ps := nodeStatus{altitude2}
	node.Set(kinds.NodeStatusToken, ps)

	//
	occurenceCatalog := transmitProof(t, clusters[0], val, countProof)

	//
	pauseForeachProof(t, occurenceCatalog[:countProof/2-1], []*proof.Hub{clusters[1]})

	//
	nodes := engines[1].Router.Nodes().Duplicate()
	assert.Len(t, nodes, 1)
}

//
//
//
//
//
func VerifyEnginesBroadcastNegativeRatifiedProof(t *testing.T) {
	settings := cfg.VerifySettings()

	val := kinds.FreshSimulatePRV()
	var altitude int64 = 10

	//
	statusDepot1 := bootstrapAssessorStatus(val, altitude-1)
	statusDatabase2 := bootstrapAssessorStatus(val, altitude-2)
	status, err := statusDepot1.Fetch()
	require.NoError(t, err)
	status.FinalLedgerAltitude++

	//
	engines, clusters := createAlsoRelateEnginesAlsoClusters(settings, []sm.Depot{statusDepot1, statusDatabase2})

	occurenceCatalog := transmitProof(t, clusters[0], val, 2)
	clusters[0].Revise(status, occurenceCatalog)
	require.EqualValues(t, uint32(0), clusters[0].Extent())

	time.Sleep(100 * time.Millisecond)

	node := engines[0].Router.Nodes().Duplicate()[0]
	ps := nodeStatus{altitude - 2}
	node.Set(kinds.NodeStatusToken, ps)

	node = engines[1].Router.Nodes().Duplicate()[0]
	ps = nodeStatus{altitude}
	node.Set(kinds.NodeStatusToken, ps)

	//
	time.Sleep(300 * time.Millisecond)

	//
	assert.Equal(t, uint32(0), clusters[1].Extent(), "REDACTED")

	//
	occurenceCatalog = make([]kinds.Proof, 3)
	for i := 0; i < 3; i++ {
		ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude-3+int64(i),
			time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), val, status.SuccessionUUID)
		require.NoError(t, err)
		err = clusters[0].AppendProof(ev)
		require.NoError(t, err)
		occurenceCatalog[i] = ev
	}

	//
	time.Sleep(300 * time.Millisecond)

	//
	nodeOccurence, _ := clusters[1].AwaitingProof(10000)
	assert.EqualValues(t, []kinds.Proof{occurenceCatalog[0]}, nodeOccurence)

	//
	//
	//
	status.FinalLedgerAltitude++
	clusters[0].Revise(status, []kinds.Proof{occurenceCatalog[2]})
	//
	require.EqualValues(t, uint32(2), clusters[0].Extent())

	//
	clusters[1].Revise(status, kinds.ProofCatalog{})
	node = engines[0].Router.Nodes().Duplicate()[0]
	ps = nodeStatus{altitude}
	node.Set(kinds.NodeStatusToken, ps)

	//
	time.Sleep(300 * time.Millisecond)

	nodeOccurence, _ = clusters[1].AwaitingProof(1000)
	assert.EqualValues(t, []kinds.Proof{occurenceCatalog[0], occurenceCatalog[1]}, nodeOccurence)
}

func VerifyHandlerMulticastProofRamSeepage(t *testing.T) {
	proofMoment := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	proofDatastore := dbm.FreshMemoryDatastore()
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(
		&kinds.LedgerSummary{Heading: kinds.Heading{Moment: proofMoment}},
	)
	val := kinds.FreshSimulatePRV()
	statusDepot := bootstrapAssessorStatus(val, 1)
	hub, err := proof.FreshHub(proofDatastore, statusDepot, ledgerDepot)
	require.NoError(t, err)

	p := &netmocks.Node{}

	p.On("REDACTED").Once().Return(true)
	p.On("REDACTED").Return(false)
	//
	//
	defer leaktest.CheckTimeout(t, 10*time.Second)()

	p.On("REDACTED", mock.MatchedBy(func(i any) bool {
		e, ok := i.(p2p.Wrapper)
		return ok && e.ConduitUUID == proof.ProofConduit
	})).Return(false)
	exitChnl := make(<-chan struct{})
	p.On("REDACTED").Return(exitChnl)
	ps := nodeStatus{2}
	p.On("REDACTED", kinds.NodeStatusToken).Return(ps)
	p.On("REDACTED").Return("REDACTED")
	p.On("REDACTED").Return("REDACTED")

	r := proof.FreshHandler(hub)
	r.AssignTracer(log.VerifyingTracer())
	r.AppendNode(p)

	_ = transmitProof(t, hub, val, 2)
}

//
//
func proofTracer() log.Tracer {
	return log.VerifyingTracerUsingHueProc(func(tokvals ...any) term.FgBgColor {
		for i := 0; i < len(tokvals)-1; i += 2 {
			if tokvals[i] == "REDACTED" {
				return term.FgBgColor{Fg: term.Color(uint8(tokvals[i+1].(int) + 1))}
			}
		}
		return term.FgBgColor{}
	})
}

//
func createAlsoRelateEnginesAlsoClusters(settings *cfg.Settings, statusCaches []sm.Depot) ([]*proof.Handler,
	[]*proof.Hub,
) {
	N := len(statusCaches)

	engines := make([]*proof.Handler, N)
	clusters := make([]*proof.Hub, N)
	tracer := proofTracer()
	proofMoment := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < N; i++ {
		proofDatastore := dbm.FreshMemoryDatastore()
		ledgerDepot := &simulations.LedgerDepot{}
		ledgerDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(
			&kinds.LedgerSummary{Heading: kinds.Heading{Moment: proofMoment}},
		)
		hub, err := proof.FreshHub(proofDatastore, statusCaches[i], ledgerDepot)
		if err != nil {
			panic(err)
		}
		clusters[i] = hub
		engines[i] = proof.FreshHandler(hub)
		engines[i].AssignTracer(tracer.Using("REDACTED", i))
	}

	p2p.CreateAssociatedRouters(settings.P2P, N, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", engines[i])
		return s
	}, p2p.Connect2routers)

	return engines, clusters
}

//
func pauseForeachProof(t *testing.T, evs kinds.ProofCatalog, clusters []*proof.Hub) {
	//
	wg := new(sync.WaitGroup)
	for i := 0; i < len(clusters); i++ {
		wg.Add(1)
		go _waitforproof(t, wg, evs, i, clusters)
	}

	complete := make(chan struct{})
	go func() {
		wg.Wait()
		close(complete)
	}()

	clock := time.After(deadline)
	select {
	case <-clock:
		t.Fatal("REDACTED")
	case <-complete:
	}
}

//
func _waitforproof(
	t *testing.T,
	wg *sync.WaitGroup,
	evs kinds.ProofCatalog,
	hubOffset int,
	clusters []*proof.Hub,
) {
	incidentpool := clusters[hubOffset]
	var occurenceCatalog []kinds.Proof
	prevailingHubExtent := 0
	for prevailingHubExtent != len(evs) {
		occurenceCatalog, _ = incidentpool.AwaitingProof(int64(len(evs) * 500)) //
		prevailingHubExtent = len(occurenceCatalog)
		time.Sleep(time.Millisecond * 100)
	}

	//
	occurenceIndex := make(map[string]kinds.Proof)
	for _, e := range occurenceCatalog {
		occurenceIndex[string(e.Digest())] = e
	}
	for i, anticipatedOccurence := range evs {
		attainedOccurence := occurenceIndex[string(anticipatedOccurence.Digest())]
		assert.Equal(t, anticipatedOccurence, attainedOccurence,
			fmt.Sprintf("REDACTED",
				i, hubOffset, anticipatedOccurence, attainedOccurence))
	}

	wg.Done()
}

func transmitProof(t *testing.T, incidentpool *proof.Hub, val kinds.PrivateAssessor, n int) kinds.ProofCatalog {
	occurenceCatalog := make([]kinds.Proof, n)
	for i := 0; i < n; i++ {
		ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(int64(i+1),
			time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), val, proofSuccessionUUID)
		require.NoError(t, err)
		err = incidentpool.AppendProof(ev)
		require.NoError(t, err)
		occurenceCatalog[i] = ev
	}
	return occurenceCatalog
}

type nodeStatus struct {
	altitude int64
}

func (ps nodeStatus) ObtainAltitude() int64 {
	return ps.altitude
}

func instanceBallot(t byte) *kinds.Ballot {
	imprint, err := time.Parse(kinds.MomentLayout, "REDACTED")
	if err != nil {
		panic(err)
	}

	return &kinds.Ballot{
		Kind:      commitchema.AttestedSignalKind(t),
		Altitude:    3,
		Iteration:     2,
		Timestamp: imprint,
		LedgerUUID: kinds.LedgerUUID{
			Digest: tenderminthash.Sum([]byte("REDACTED")),
			FragmentAssignHeading: kinds.FragmentAssignHeading{
				Sum: 1000000,
				Digest:  tenderminthash.Sum([]byte("REDACTED")),
			},
		},
		AssessorLocation: security.LocatorDigest([]byte("REDACTED")),
		AssessorOrdinal:   56789,
	}
}

//
func VerifyProofArrays(t *testing.T) {
	val := &kinds.Assessor{
		Location:     security.LocatorDigest([]byte("REDACTED")),
		BallotingPotency: 10,
	}

	itemAssign := kinds.FreshAssessorAssign([]*kinds.Assessor{val})

	repetition, err := kinds.FreshReplicatedBallotProof(
		instanceBallot(1),
		instanceBallot(2),
		fallbackProofMoment,
		itemAssign,
	)
	require.NoError(t, err)

	verifyScenarios := []struct {
		verifyAlias     string
		proofCatalog []kinds.Proof
		expirationOctets     string
	}{
		{"REDACTED", []kinds.Proof{repetition}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		evi := make([]commitchema.Proof, len(tc.proofCatalog))
		for i := 0; i < len(tc.proofCatalog); i++ {
			ev, err := kinds.ProofTowardSchema(tc.proofCatalog[i])
			require.NoError(t, err, tc.verifyAlias)
			evi[i] = *ev
		}

		epl := commitchema.ProofCatalog{
			Proof: evi,
		}

		bz, err := epl.Serialize()
		require.NoError(t, err, tc.verifyAlias)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyAlias)

	}
}
