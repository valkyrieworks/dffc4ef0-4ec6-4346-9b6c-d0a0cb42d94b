package peer

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
	nodestub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/simulate"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

func VerifyPeerInitiateHalt(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)

	//
	n, err := FallbackFreshPeer(settings, log.VerifyingTracer())
	require.NoError(t, err)
	err = n.Initiate()
	require.NoError(t, err)

	t.Logf("REDACTED", n.sw.PeerDetails())

	//
	ledgersUnder, err := n.IncidentChannel().Listen(context.Background(), "REDACTED", kinds.IncidentInquireFreshLedger)
	require.NoError(t, err)
	select {
	case <-ledgersUnder.Out():
	case <-ledgersUnder.Aborted():
		t.Fatal("REDACTED")
	case <-time.After(10 * time.Second):
		t.Fatal("REDACTED")
	}

	//
	go func() {
		err = n.Halt()
		require.NoError(t, err)
	}()

	select {
	case <-n.Exit():
	case <-time.After(5 * time.Second):
		pid := os.Getpid()
		p, err := os.FindProcess(pid)
		if err != nil {
			panic(err)
		}
		err = p.Signal(syscall.SIGABRT)
		fmt.Println(err)
		t.Fatal("REDACTED")
	}
}

func VerifyPartitionAlsoShaveBlank(t *testing.T) {
	verifyScenarios := []struct {
		s        string
		sep      string
		delimiters   string
		anticipated []string
	}{
		{"REDACTED", "REDACTED", "REDACTED", []string{"REDACTED", "REDACTED", "REDACTED"}},
		{"REDACTED", "REDACTED", "REDACTED", []string{"REDACTED", "REDACTED", "REDACTED"}},
		{"REDACTED", "REDACTED", "REDACTED", []string{"REDACTED", "REDACTED", "REDACTED"}},
		{"REDACTED", "REDACTED", "REDACTED", []string{"REDACTED"}},
		{"REDACTED", "REDACTED", "REDACTED", []string{}},
	}

	for _, tc := range verifyScenarios {
		assert.Equal(t, tc.anticipated, partitionAlsoShaveBlank(tc.s, tc.sep, tc.delimiters), "REDACTED", tc.s)
	}
}

func VerifyPeerPostponedInitiate(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	now := committime.Now()

	//
	n, err := FallbackFreshPeer(settings, log.VerifyingTracer())
	n.OriginPaper().OriginMoment = now.Add(2 * time.Second)
	require.NoError(t, err)

	err = n.Initiate()
	require.NoError(t, err)
	defer n.Halt() //

	initiateMoment := committime.Now()
	assert.Equal(t, true, initiateMoment.After(n.OriginPaper().OriginMoment))
}

func VerifyPeerAssignApplicationEdition(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)

	//
	n, err := FallbackFreshPeer(settings, log.VerifyingTracer())
	require.NoError(t, err)

	//
	platformEdition := statedepot.PlatformEdition

	//
	status, err := n.statusDepot.Fetch()
	require.NoError(t, err)
	assert.Equal(t, status.Edition.Agreement.App, platformEdition)

	//
	assert.Equal(t, n.peerDetails.(p2p.FallbackPeerDetails).SchemeEdition.App, platformEdition)
}

func VerifyProfilerDaemon(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	settings.RPC.ProfilerOverhearLocation = verifyLiberateLocation(t)

	//
	_, err := http.Get("REDACTED" + settings.RPC.ProfilerOverhearLocation) //
	assert.Error(t, err)

	n, err := FallbackFreshPeer(settings, log.VerifyingTracer())
	assert.NoError(t, err)
	assert.NoError(t, n.Initiate())
	defer func() {
		require.NoError(t, n.Halt())
	}()
	assert.NotNil(t, n.profilerDaemon)

	reply, err := http.Get("REDACTED" + settings.RPC.ProfilerOverhearLocation + "REDACTED")
	assert.NoError(t, err)
	defer reply.Body.Close()
	assert.Equal(t, 200, reply.StatusCode)
}

func VerifyPeerAssignPrivateItemTcpsocket(t *testing.T) {
	location := "REDACTED" + verifyLiberateLocation(t)

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	settings.PrivateAssessorOverhearLocation = location

	caller := privatevalue.CallStreamProc(location, 100*time.Millisecond, edwards25519.ProducePrivateToken())
	callerGateway := privatevalue.FreshEndorserCallerGateway(
		log.VerifyingTracer(),
		caller,
	)
	privatevalue.EndorserCallerGatewayDeadlineRetrievePersist(100 * time.Millisecond)(callerGateway)

	endorserDaemon := privatevalue.FreshEndorserDaemon(
		callerGateway,
		verify.FallbackVerifySuccessionUUID,
		kinds.FreshSimulatePRV(),
	)

	go func() {
		err := endorserDaemon.Initiate()
		if err != nil {
			panic(err)
		}
	}()
	defer endorserDaemon.Halt() //

	n, err := FallbackFreshPeer(settings, log.VerifyingTracer())
	require.NoError(t, err)
	assert.IsType(t, &privatevalue.ReissueEndorserCustomer{}, n.PrivateAssessor())
}

//
func VerifyPrivateAssessorOverhearLocationNegativeScheme(t *testing.T) {
	locationNegativeHeading := verifyLiberateLocation(t)

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	settings.PrivateAssessorOverhearLocation = locationNegativeHeading

	_, err := FallbackFreshPeer(settings, log.VerifyingTracer())
	assert.Error(t, err)
}

func VerifyPeerAssignPrivateItemProcess(t *testing.T) {
	tempfile := "REDACTED" + commitrand.Str(6) + "REDACTED"
	defer os.Remove(tempfile) //

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	settings.PrivateAssessorOverhearLocation = "REDACTED" + tempfile

	caller := privatevalue.CallPosixProc(tempfile)
	callerGateway := privatevalue.FreshEndorserCallerGateway(
		log.VerifyingTracer(),
		caller,
	)
	privatevalue.EndorserCallerGatewayDeadlineRetrievePersist(100 * time.Millisecond)(callerGateway)

	prvtcs := privatevalue.FreshEndorserDaemon(
		callerGateway,
		verify.FallbackVerifySuccessionUUID,
		kinds.FreshSimulatePRV(),
	)

	go func() {
		err := prvtcs.Initiate()
		require.NoError(t, err)
	}()
	defer prvtcs.Halt() //

	n, err := FallbackFreshPeer(settings, log.VerifyingTracer())
	require.NoError(t, err)
	assert.IsType(t, &privatevalue.ReissueEndorserCustomer{}, n.PrivateAssessor())
}

//
func verifyLiberateLocation(t *testing.T) string {
	ln, err := net.Listen("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer ln.Close()

	return fmt.Sprintf("REDACTED", ln.Addr().(*net.TCPAddr).Port)
}

//
//
func VerifyGenerateNominationLedger(t *testing.T) {
	ctx := t.Context()

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	cc := delegate.FreshRegionalCustomerOriginator(statedepot.FreshInsideRamPlatform())
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.Nil(t, err)
	defer delegatePlatform.Halt() //

	tracer := log.VerifyingTracer()

	var altitude int64 = 1
	status, statusDatastore, privateItems := status(1, altitude)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	var (
		fragmentExtent uint32 = 256
		maximumOctets int64  = 16384
	)
	maximumProofOctets := maximumOctets / 2
	status.AgreementSettings.Ledger.MaximumOctets = maximumOctets
	status.AgreementSettings.Proof.MaximumOctets = maximumProofOctets
	nominatorLocation, _ := status.Assessors.ObtainViaOrdinal(0)

	//
	txpoollTelemetry := txpooll.NooperationTelemetry()
	txpool := txpooll.FreshCNCatalogTxpool(settings.Txpool,
		delegatePlatform.Txpool(),
		status.FinalLedgerAltitude,
		txpooll.UsingTelemetry(txpoollTelemetry),
		txpooll.UsingPriorInspect(sm.TransferPriorInspect(status)),
		txpooll.UsingRelayInspect(sm.TransferRelayInspect(status)))

	//
	proofDatastore := dbm.FreshMemoryDatastore()
	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	proofHub, err := proof.FreshHub(proofDatastore, statusDepot, ledgerDepot)
	require.NoError(t, err)
	proofHub.AssignTracer(tracer)

	//
	//
	var prevailingOctets int64
	for prevailingOctets <= maximumProofOctets {
		ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude, time.Now(), privateItems[0], "REDACTED")
		require.NoError(t, err)
		prevailingOctets += int64(len(ev.Octets()))
		proofHub.DiscloseDiscordantBallots(ev.BallotAN, ev.BallotBYTE)
	}

	occurenceCatalog, extent := proofHub.AwaitingProof(status.AgreementSettings.Proof.MaximumOctets)
	require.Less(t, extent, status.AgreementSettings.Proof.MaximumOctets+1)
	occurenceData := &kinds.ProofData{Proof: occurenceCatalog}
	require.EqualValues(t, extent, occurenceData.OctetExtent())

	//
	//
	transferMagnitude := 100
	for i := 0; i <= int(maximumOctets)/transferMagnitude; i++ {
		tx := commitrand.Octets(transferMagnitude)
		err := txpool.InspectTransfer(tx, nil, txpooll.TransferDetails{})
		assert.NoError(t, err)
	}

	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		tracer,
		delegatePlatform.Agreement(),
		txpool,
		proofHub,
		ledgerDepot,
	)

	addnEndorse := &kinds.ExpandedEndorse{Altitude: altitude - 1}
	ledger, err := ledgerExecute.GenerateNominationLedger(
		ctx,
		altitude,
		status,
		addnEndorse,
		nominatorLocation,
	)
	require.NoError(t, err)

	//
	fragmentAssign, err := ledger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	assert.Less(t, fragmentAssign.OctetExtent(), maximumOctets)

	fragmentAssignOriginatingHeadline := kinds.FreshFragmentAssignOriginatingHeading(fragmentAssign.Heading())
	for fragmentAssignOriginatingHeadline.Tally() < fragmentAssignOriginatingHeadline.Sum() {
		appended, err := fragmentAssignOriginatingHeadline.AppendFragment(fragmentAssign.ObtainFragment(int(fragmentAssignOriginatingHeadline.Tally())))
		require.NoError(t, err)
		require.True(t, appended)
	}
	assert.EqualValues(t, fragmentAssignOriginatingHeadline.OctetExtent(), fragmentAssign.OctetExtent())

	err = ledgerExecute.CertifyLedger(status, ledger)
	assert.NoError(t, err)
}

func VerifyMaximumNominationLedgerExtent(t *testing.T) {
	ctx := t.Context()

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	cc := delegate.FreshRegionalCustomerOriginator(statedepot.FreshInsideRamPlatform())
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.Nil(t, err)
	defer delegatePlatform.Halt() //

	tracer := log.VerifyingTracer()

	var altitude int64 = 1
	status, statusDatastore, _ := status(1, altitude)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	var maximumOctets int64 = 16384
	var fragmentExtent uint32 = 256
	status.AgreementSettings.Ledger.MaximumOctets = maximumOctets
	nominatorLocation, _ := status.Assessors.ObtainViaOrdinal(0)

	//
	txpoollTelemetry := txpooll.NooperationTelemetry()
	txpool := txpooll.FreshCNCatalogTxpool(settings.Txpool,
		delegatePlatform.Txpool(),
		status.FinalLedgerAltitude,
		txpooll.UsingTelemetry(txpoollTelemetry),
		txpooll.UsingPriorInspect(sm.TransferPriorInspect(status)),
		txpooll.UsingRelayInspect(sm.TransferRelayInspect(status)))

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())

	//
	transferMagnitude := int(kinds.MaximumDataOctetsNegativeProof(maximumOctets, 1))
	tx := commitrand.Octets(transferMagnitude - 4) //
	err = txpool.InspectTransfer(tx, nil, txpooll.TransferDetails{})
	assert.NoError(t, err)

	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		tracer,
		delegatePlatform.Agreement(),
		txpool,
		sm.VoidProofHub{},
		ledgerDepot,
	)

	addnEndorse := &kinds.ExpandedEndorse{Altitude: altitude - 1}
	ledger, err := ledgerExecute.GenerateNominationLedger(
		ctx,
		altitude,
		status,
		addnEndorse,
		nominatorLocation,
	)
	require.NoError(t, err)

	pb, err := ledger.TowardSchema()
	require.NoError(t, err)
	assert.Less(t, int64(pb.Extent()), maximumOctets)

	//
	fragmentAssign, err := ledger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	assert.EqualValues(t, fragmentAssign.OctetExtent(), int64(pb.Extent()))
}

func VerifyPeerFreshPeerBespokeEngines(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)

	cr := nodestub.FreshHandler()
	cr.Conduits = []*link.ConduitDefinition{
		{
			ID:                  byte(0x31),
			Urgency:            5,
			TransmitStagingVolume:   100,
			ObtainSignalVolume: 100,
		},
	}
	bespokeChainchronizeHandler := nodestub.FreshHandler()

	peerToken, err := p2p.FetchEitherProducePeerToken(settings.PeerTokenRecord())
	require.NoError(t, err)

	n, err := FreshPeer(settings,
		privatevalue.FetchEitherProduceRecordPRV(settings.PrivateAssessorTokenRecord(), settings.PrivateAssessorStatusRecord()),
		peerToken,
		delegate.FallbackCustomerOriginator(settings.DelegateApplication, settings.Iface, settings.DatastorePath()),
		FallbackInaugurationPaperSupplierMethod(settings),
		cfg.FallbackDatastoreSupplier,
		FallbackTelemetrySupplier(settings.Telemetry),
		log.VerifyingTracer(),
		BespokeEngines(map[string]p2p.Handler{"REDACTED": cr, "REDACTED": bespokeChainchronizeHandler}),
	)
	require.NoError(t, err)

	err = n.Initiate()
	require.NoError(t, err)
	defer n.Halt() //

	assert.True(t, cr.EqualsActive())
	handler, ok := n.Router().Handler("REDACTED")
	assert.True(t, ok)
	assert.Equal(t, cr, handler)

	assert.True(t, bespokeChainchronizeHandler.EqualsActive())

	handler, ok = n.Router().Handler("REDACTED")
	assert.True(t, ok)
	assert.Equal(t, bespokeChainchronizeHandler, handler)

	conduits := n.PeerDetails().(p2p.FallbackPeerDetails).Conduits
	assert.Contains(t, conduits, txpooll.TxpoolConduit)
	assert.Contains(t, conduits, cr.Conduits[0].ID)
}

func status(nthValues int, altitude int64) (sm.Status, dbm.DB, []kinds.PrivateAssessor) {
	privateItems := make([]kinds.PrivateAssessor, nthValues)
	values := make([]kinds.OriginAssessor, nthValues)
	for i := 0; i < nthValues; i++ {
		privateItem := kinds.FreshSimulatePRV()
		privateItems[i] = privateItem
		values[i] = kinds.OriginAssessor{
			Location: privateItem.PrivateToken.PublicToken().Location(),
			PublicToken:  privateItem.PrivateToken.PublicToken(),
			Potency:   1000,
			Alias:    fmt.Sprintf("REDACTED", i),
		}
	}
	s, _ := sm.CreateInaugurationStatus(&kinds.OriginPaper{
		SuccessionUUID:    "REDACTED",
		Assessors: values,
		PlatformDigest:    nil,
	})

	//
	statusDatastore := dbm.FreshMemoryDatastore()
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	if err := statusDepot.Persist(s); err != nil {
		panic(err)
	}

	for i := 1; i < int(altitude); i++ {
		s.FinalLedgerAltitude++
		s.FinalAssessors = s.Assessors.Duplicate()
		if err := statusDepot.Persist(s); err != nil {
			panic(err)
		}
	}
	return s, statusDatastore, privateItems
}
