package member

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

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/proof"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/p2p/link"
	p2pemulator "github.com/valkyrieworks/p2p/emulate"
	"github.com/valkyrieworks/privatekey"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

func VerifyMemberBeginHalt(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)

	//
	n, err := StandardNewMember(settings, log.VerifyingTracer())
	require.NoError(t, err)
	err = n.Begin()
	require.NoError(t, err)

	t.Logf("REDACTED", n.sw.MemberDetails())

	//
	ledgersSubtract, err := n.EventBus().Enrol(context.Background(), "REDACTED", kinds.EventInquireNewLedger)
	require.NoError(t, err)
	select {
	case <-ledgersSubtract.Out():
	case <-ledgersSubtract.Revoked():
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

func VerifyDivideAndShaveEmpty(t *testing.T) {
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
		assert.Equal(t, tc.anticipated, divideAndShaveEmpty(tc.s, tc.sep, tc.delimiters), "REDACTED", tc.s)
	}
}

func VerifyMemberDeferredBegin(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	now := engineclock.Now()

	//
	n, err := StandardNewMember(settings, log.VerifyingTracer())
	n.OriginPaper().OriginMoment = now.Add(2 * time.Second)
	require.NoError(t, err)

	err = n.Begin()
	require.NoError(t, err)
	defer n.Halt() //

	beginMoment := engineclock.Now()
	assert.Equal(t, true, beginMoment.After(n.OriginPaper().OriginMoment))
}

func VerifyMemberCollectionApplicationRelease(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)

	//
	n, err := StandardNewMember(settings, log.VerifyingTracer())
	require.NoError(t, err)

	//
	applicationRelease := objectdepot.ApplicationRelease

	//
	status, err := n.statusDepot.Import()
	require.NoError(t, err)
	assert.Equal(t, status.Release.Agreement.App, applicationRelease)

	//
	assert.Equal(t, n.memberDetails.(p2p.StandardMemberDetails).ProtocolRelease.App, applicationRelease)
}

func VerifyPprofHost(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	settings.RPC.PprofAcceptLocation = verifyReleaseAddress(t)

	//
	_, err := http.Get("REDACTED" + settings.RPC.PprofAcceptLocation) //
	assert.Error(t, err)

	n, err := StandardNewMember(settings, log.VerifyingTracer())
	assert.NoError(t, err)
	assert.NoError(t, n.Begin())
	defer func() {
		require.NoError(t, n.Halt())
	}()
	assert.NotNil(t, n.pprofSvc)

	reply, err := http.Get("REDACTED" + settings.RPC.PprofAcceptLocation + "REDACTED")
	assert.NoError(t, err)
	defer reply.Body.Close()
	assert.Equal(t, 200, reply.StatusCode)
}

func VerifyMemberCollectionPrivateValueTCP(t *testing.T) {
	address := "REDACTED" + verifyReleaseAddress(t)

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	settings.PrivateRatifierAcceptAddress = address

	caller := privatekey.CallTCPFn(address, 100*time.Millisecond, ed25519.GeneratePrivateKey())
	callerTerminus := privatekey.NewNotaryCallerGateway(
		log.VerifyingTracer(),
		caller,
	)
	privatekey.NotaryCallerTerminusDeadlineFetchRecord(100 * time.Millisecond)(callerTerminus)

	notaryHost := privatekey.NewNotaryHost(
		callerTerminus,
		verify.StandardVerifyLedgerUID,
		kinds.NewEmulatePV(),
	)

	go func() {
		err := notaryHost.Begin()
		if err != nil {
			panic(err)
		}
	}()
	defer notaryHost.Halt() //

	n, err := StandardNewMember(settings, log.VerifyingTracer())
	require.NoError(t, err)
	assert.IsType(t, &privatekey.ReprocessNotaryCustomer{}, n.PrivateRatifier())
}

//
func VerifyPrivateRatifierObserveAddressNoProtocol(t *testing.T) {
	addressNoPrefix := verifyReleaseAddress(t)

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	settings.PrivateRatifierAcceptAddress = addressNoPrefix

	_, err := StandardNewMember(settings, log.VerifyingTracer())
	assert.Error(t, err)
}

func VerifyMemberCollectionPrivateValueIPC(t *testing.T) {
	tempfile := "REDACTED" + engineseed.Str(6) + "REDACTED"
	defer os.Remove(tempfile) //

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	settings.PrivateRatifierAcceptAddress = "REDACTED" + tempfile

	caller := privatekey.CallUnixFn(tempfile)
	callerTerminus := privatekey.NewNotaryCallerGateway(
		log.VerifyingTracer(),
		caller,
	)
	privatekey.NotaryCallerTerminusDeadlineFetchRecord(100 * time.Millisecond)(callerTerminus)

	pvsc := privatekey.NewNotaryHost(
		callerTerminus,
		verify.StandardVerifyLedgerUID,
		kinds.NewEmulatePV(),
	)

	go func() {
		err := pvsc.Begin()
		require.NoError(t, err)
	}()
	defer pvsc.Halt() //

	n, err := StandardNewMember(settings, log.VerifyingTracer())
	require.NoError(t, err)
	assert.IsType(t, &privatekey.ReprocessNotaryCustomer{}, n.PrivateRatifier())
}

//
func verifyReleaseAddress(t *testing.T) string {
	ln, err := net.Listen("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer ln.Close()

	return fmt.Sprintf("REDACTED", ln.Addr().(*net.TCPAddr).Port)
}

//
//
func VerifyInstantiateNominationLedger(t *testing.T) {
	ctx := t.Context()

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	cc := gateway.NewNativeCustomerOriginator(objectdepot.NewInRamSoftware())
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.Nil(t, err)
	defer gatewayApplication.Halt() //

	tracer := log.VerifyingTracer()

	var level int64 = 1
	status, statusStore, privateValues := status(1, level)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	var (
		segmentVolume uint32 = 256
		maximumOctets int64  = 16384
	)
	maximumProofOctets := maximumOctets / 2
	status.AgreementOptions.Ledger.MaximumOctets = maximumOctets
	status.AgreementOptions.Proof.MaximumOctets = maximumProofOctets
	recommenderAddress, _ := status.Ratifiers.FetchByOrdinal(0)

	//
	memplStats := txpool.NoopStats()
	txpool := txpool.NewCCatalogTxpool(settings.Txpool,
		gatewayApplication.Txpool(),
		status.FinalLedgerLevel,
		txpool.WithStats(memplStats),
		txpool.WithPreInspect(sm.TransferPreInspect(status)),
		txpool.WithSubmitInspect(sm.TransferSubmitInspect(status)))

	//
	proofStore := dbm.NewMemoryStore()
	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	proofDepository, err := proof.NewDepository(proofStore, statusDepot, ledgerDepot)
	require.NoError(t, err)
	proofDepository.AssignTracer(tracer)

	//
	//
	var ongoingOctets int64
	for ongoingOctets <= maximumProofOctets {
		ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level, time.Now(), privateValues[0], "REDACTED")
		require.NoError(t, err)
		ongoingOctets += int64(len(ev.Octets()))
		proofDepository.NotifyClashingBallots(ev.BallotA, ev.BallotBYTE)
	}

	evtCatalog, volume := proofDepository.AwaitingProof(status.AgreementOptions.Proof.MaximumOctets)
	require.Less(t, volume, status.AgreementOptions.Proof.MaximumOctets+1)
	evtData := &kinds.ProofData{Proof: evtCatalog}
	require.EqualValues(t, volume, evtData.OctetVolume())

	//
	//
	transferExtent := 100
	for i := 0; i <= int(maximumOctets)/transferExtent; i++ {
		tx := engineseed.Octets(transferExtent)
		err := txpool.InspectTransfer(tx, nil, txpool.TransferDetails{})
		assert.NoError(t, err)
	}

	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		tracer,
		gatewayApplication.Agreement(),
		txpool,
		proofDepository,
		ledgerDepot,
	)

	extensionEndorse := &kinds.ExpandedEndorse{Level: level - 1}
	ledger, err := ledgerExecute.InstantiateNominationLedger(
		ctx,
		level,
		status,
		extensionEndorse,
		recommenderAddress,
	)
	require.NoError(t, err)

	//
	sectionCollection, err := ledger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	assert.Less(t, sectionCollection.OctetVolume(), maximumOctets)

	sectionCollectionFromHeading := kinds.NewSegmentCollectionFromHeading(sectionCollection.Heading())
	for sectionCollectionFromHeading.Number() < sectionCollectionFromHeading.Sum() {
		appended, err := sectionCollectionFromHeading.AppendSegment(sectionCollection.FetchSegment(int(sectionCollectionFromHeading.Number())))
		require.NoError(t, err)
		require.True(t, appended)
	}
	assert.EqualValues(t, sectionCollectionFromHeading.OctetVolume(), sectionCollection.OctetVolume())

	err = ledgerExecute.CertifyLedger(status, ledger)
	assert.NoError(t, err)
}

func VerifyMaximumNominationLedgerVolume(t *testing.T) {
	ctx := t.Context()

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	cc := gateway.NewNativeCustomerOriginator(objectdepot.NewInRamSoftware())
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.Nil(t, err)
	defer gatewayApplication.Halt() //

	tracer := log.VerifyingTracer()

	var level int64 = 1
	status, statusStore, _ := status(1, level)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	var maximumOctets int64 = 16384
	var segmentVolume uint32 = 256
	status.AgreementOptions.Ledger.MaximumOctets = maximumOctets
	recommenderAddress, _ := status.Ratifiers.FetchByOrdinal(0)

	//
	memplStats := txpool.NoopStats()
	txpool := txpool.NewCCatalogTxpool(settings.Txpool,
		gatewayApplication.Txpool(),
		status.FinalLedgerLevel,
		txpool.WithStats(memplStats),
		txpool.WithPreInspect(sm.TransferPreInspect(status)),
		txpool.WithSubmitInspect(sm.TransferSubmitInspect(status)))

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())

	//
	transferExtent := int(kinds.MaximumDataOctetsNoProof(maximumOctets, 1))
	tx := engineseed.Octets(transferExtent - 4) //
	err = txpool.InspectTransfer(tx, nil, txpool.TransferDetails{})
	assert.NoError(t, err)

	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		tracer,
		gatewayApplication.Agreement(),
		txpool,
		sm.EmptyProofDepository{},
		ledgerDepot,
	)

	extensionEndorse := &kinds.ExpandedEndorse{Level: level - 1}
	ledger, err := ledgerExecute.InstantiateNominationLedger(
		ctx,
		level,
		status,
		extensionEndorse,
		recommenderAddress,
	)
	require.NoError(t, err)

	pb, err := ledger.ToSchema()
	require.NoError(t, err)
	assert.Less(t, int64(pb.Volume()), maximumOctets)

	//
	sectionCollection, err := ledger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	assert.EqualValues(t, sectionCollection.OctetVolume(), int64(pb.Volume()))
}

func VerifyMemberNewMemberBespokeHandlers(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)

	cr := p2pemulator.NewHandler()
	cr.Streams = []*link.StreamDefinition{
		{
			ID:                  byte(0x31),
			Urgency:            5,
			TransmitBufferVolume:   100,
			AcceptSignalVolume: 100,
		},
	}
	bespokeChainconnectHandler := p2pemulator.NewHandler()

	memberKey, err := p2p.ImportOrGenerateMemberKey(settings.MemberKeyEntry())
	require.NoError(t, err)

	n, err := NewMember(settings,
		privatekey.ImportOrGenerateEntryPV(settings.PrivateRatifierKeyEntry(), settings.PrivateRatifierStatusEntry()),
		memberKey,
		gateway.StandardCustomerOriginator(settings.GatewayApplication, settings.Iface, settings.StoreFolder()),
		StandardOriginPaperSourceFunction(settings),
		cfg.StandardStoreSource,
		StandardStatsSource(settings.Telemetry),
		log.VerifyingTracer(),
		BespokeHandlers(map[string]p2p.Handler{"REDACTED": cr, "REDACTED": bespokeChainconnectHandler}),
	)
	require.NoError(t, err)

	err = n.Begin()
	require.NoError(t, err)
	defer n.Halt() //

	assert.True(t, cr.IsActive())
	handler, ok := n.Router().Handler("REDACTED")
	assert.True(t, ok)
	assert.Equal(t, cr, handler)

	assert.True(t, bespokeChainconnectHandler.IsActive())

	handler, ok = n.Router().Handler("REDACTED")
	assert.True(t, ok)
	assert.Equal(t, bespokeChainconnectHandler, handler)

	streams := n.MemberDetails().(p2p.StandardMemberDetails).Streams
	assert.Contains(t, streams, txpool.TxpoolConduit)
	assert.Contains(t, streams, cr.Streams[0].ID)
}

func status(nValues int, level int64) (sm.Status, dbm.DB, []kinds.PrivateRatifier) {
	privateValues := make([]kinds.PrivateRatifier, nValues)
	values := make([]kinds.OriginRatifier, nValues)
	for i := 0; i < nValues; i++ {
		privateValue := kinds.NewEmulatePV()
		privateValues[i] = privateValue
		values[i] = kinds.OriginRatifier{
			Location: privateValue.PrivateKey.PublicKey().Location(),
			PublicKey:  privateValue.PrivateKey.PublicKey(),
			Energy:   1000,
			Label:    fmt.Sprintf("REDACTED", i),
		}
	}
	s, _ := sm.CreateOriginStatus(&kinds.OriginPaper{
		LedgerUID:    "REDACTED",
		Ratifiers: values,
		ApplicationDigest:    nil,
	})

	//
	statusStore := dbm.NewMemoryStore()
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	if err := statusDepot.Persist(s); err != nil {
		panic(err)
	}

	for i := 1; i < int(level); i++ {
		s.FinalLedgerLevel++
		s.FinalRatifiers = s.Ratifiers.Clone()
		if err := statusDepot.Persist(s); err != nil {
			panic(err)
		}
	}
	return s, statusStore, privateValues
}
