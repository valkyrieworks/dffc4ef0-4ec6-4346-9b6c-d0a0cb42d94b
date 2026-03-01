package investigate_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	ifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/scrutinize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	webcustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	indexsimulate "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/simulations"
	machinestubs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
	transindexsimulate "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func VerifyScrutinizeInitializer(t *testing.T) {
	cfg := verify.RestoreVerifyOrigin("REDACTED")
	t.Cleanup(leaktest.Check(t))
	defer func() { _ = os.RemoveAll(cfg.OriginPath) }()
	t.Run("REDACTED", func(t *testing.T) {
		d, err := scrutinize.FreshOriginatingSettings(cfg)
		require.NoError(t, err)
		require.NotNil(t, d)
	})
}

func VerifyScrutinizeExecute(t *testing.T) {
	cfg := verify.RestoreVerifyOrigin("REDACTED")
	t.Cleanup(leaktest.Check(t))
	defer func() { _ = os.RemoveAll(cfg.OriginPath) }()
	t.Run("REDACTED", func(t *testing.T) {
		d, err := scrutinize.FreshOriginatingSettings(cfg)
		require.NoError(t, err)
		ctx, abort := context.WithCancel(context.Background())
		haltedGroup := &sync.WaitGroup{}
		haltedGroup.Add(1)
		go func() {
			require.NoError(t, d.Run(ctx))
			haltedGroup.Done()
		}()
		abort()
		haltedGroup.Wait()
	})
}

func VerifyLedger(t *testing.T) {
	verifyAltitude := int64(1)
	verifyLedger := new(kinds.Ledger)
	verifyLedger.Altitude = verifyAltitude
	verifyLedger.FinalEndorseDigest = []byte("REDACTED")
	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)

	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(verifyAltitude)
	ledgerDepotSimulate.On("REDACTED").Return(int64(0))
	ledgerDepotSimulate.On("REDACTED", verifyAltitude).Return(&kinds.LedgerSummary{})
	ledgerDepotSimulate.On("REDACTED", verifyAltitude).Return(verifyLedger)
	ledgerDepotSimulate.On("REDACTED").Return(nil)

	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}

	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)
	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)
	outcomeLedger, err := cli.Ledger(context.Background(), &verifyAltitude)
	require.NoError(t, err)
	require.Equal(t, verifyLedger.Altitude, outcomeLedger.Ledger.Altitude)
	require.Equal(t, verifyLedger.FinalEndorseDigest, outcomeLedger.Ledger.FinalEndorseDigest)
	abort()
	wg.Wait()

	ledgerDepotSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
}

func VerifyTransferLookup(t *testing.T) {
	verifyDigest := []byte("REDACTED")
	verifyTransfer := []byte("REDACTED")
	verifyInquire := fmt.Sprintf("REDACTED", string(verifyDigest))
	verifyTransferOutcome := &ifacetypes.TransferOutcome{
		Altitude: 1,
		Ordinal:  100,
		Tx:     verifyTransfer,
	}

	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(nil)
	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}
	transferOrdinalizerSimulate.On("REDACTED", mock.Anything,
		mock.MatchedBy(func(q *inquire.Inquire) bool {
			return verifyInquire == strings.ReplaceAll(q.Text(), "REDACTED", "REDACTED")
		})).
		Return([]*ifacetypes.TransferOutcome{verifyTransferOutcome}, nil)

	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)
	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)

	screen := 1
	outcomeTransferLookup, err := cli.TransferLookup(context.Background(), verifyInquire, false, &screen, &screen, "REDACTED")
	require.NoError(t, err)
	require.Len(t, outcomeTransferLookup.Txs, 1)
	require.Equal(t, kinds.Tx(verifyTransfer), outcomeTransferLookup.Txs[0].Tx)

	abort()
	wg.Wait()

	transferOrdinalizerSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
	ledgerDepotSimulate.AssertExpectations(t)
}

func VerifyTransfer(t *testing.T) {
	verifyDigest := []byte("REDACTED")
	verifyTransfer := []byte("REDACTED")

	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(nil)
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}
	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	transferOrdinalizerSimulate.On("REDACTED", verifyDigest).Return(&ifacetypes.TransferOutcome{
		Tx: verifyTransfer,
	}, nil)

	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)
	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)

	res, err := cli.Tx(context.Background(), verifyDigest, false)
	require.NoError(t, err)
	require.Equal(t, kinds.Tx(verifyTransfer), res.Tx)

	abort()
	wg.Wait()

	transferOrdinalizerSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
	ledgerDepotSimulate.AssertExpectations(t)
}

func VerifyAgreementParameters(t *testing.T) {
	verifyAltitude := int64(1)
	verifyMaximumFuel := int64(55)
	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate.On("REDACTED").Return(verifyAltitude)
	ledgerDepotSimulate.On("REDACTED").Return(int64(0))
	statusDepotSimulate.On("REDACTED", verifyAltitude).Return(kinds.AgreementSettings{
		Ledger: kinds.LedgerParameters{
			MaximumFuel: verifyMaximumFuel,
		},
	}, nil)
	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}
	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)

	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)
	parameters, err := cli.AgreementSettings(context.Background(), &verifyAltitude)
	require.NoError(t, err)
	require.Equal(t, parameters.AgreementSettings.Ledger.MaximumFuel, verifyMaximumFuel)

	abort()
	wg.Wait()

	ledgerDepotSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
}

func VerifyLedgerOutcomes(t *testing.T) {
	verifyAltitude := int64(1)
	verifyFuelUtilized := int64(100)
	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)
	//
	statusDepotSimulate.On("REDACTED", verifyAltitude).Return(&ifacetypes.ReplyCulminateLedger{
		TransferOutcomes: []*ifacetypes.InvokeTransferOutcome{
			{
				FuelUtilized: verifyFuelUtilized,
			},
		},
	}, nil)
	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate.On("REDACTED").Return(int64(0))
	ledgerDepotSimulate.On("REDACTED").Return(verifyAltitude)
	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}
	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)

	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)
	res, err := cli.LedgerOutcomes(context.Background(), &verifyAltitude)
	require.NoError(t, err)
	require.Equal(t, res.TransOutcomes[0].FuelUtilized, verifyFuelUtilized)

	abort()
	wg.Wait()

	ledgerDepotSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
}

func VerifyEndorse(t *testing.T) {
	verifyAltitude := int64(1)
	verifyIteration := int32(101)
	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate.On("REDACTED").Return(int64(0))
	ledgerDepotSimulate.On("REDACTED").Return(verifyAltitude)
	ledgerDepotSimulate.On("REDACTED", verifyAltitude).Return(&kinds.LedgerSummary{}, nil)
	ledgerDepotSimulate.On("REDACTED", verifyAltitude).Return(&kinds.Endorse{
		Altitude: verifyAltitude,
		Iteration:  verifyIteration,
	}, nil)
	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}
	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)

	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)
	res, err := cli.Endorse(context.Background(), &verifyAltitude)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Endorse.Iteration, verifyIteration)

	abort()
	wg.Wait()

	ledgerDepotSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
}

func VerifyLedgerViaDigest(t *testing.T) {
	verifyAltitude := int64(1)
	verifyDigest := []byte("REDACTED")
	verifyLedger := new(kinds.Ledger)
	verifyLedger.Altitude = verifyAltitude
	verifyLedger.FinalEndorseDigest = verifyDigest
	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate.On("REDACTED", verifyAltitude).Return(&kinds.LedgerSummary{
		LedgerUUID: kinds.LedgerUUID{
			Digest: verifyDigest,
		},
		Heading: kinds.Heading{
			Altitude: verifyAltitude,
		},
	}, nil)
	ledgerDepotSimulate.On("REDACTED", verifyDigest).Return(verifyLedger, nil)
	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}
	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)

	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)
	res, err := cli.LedgerViaDigest(context.Background(), verifyDigest)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, []byte(res.LedgerUUID.Digest), verifyDigest)

	abort()
	wg.Wait()

	ledgerDepotSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
}

func VerifyLedgerchain(t *testing.T) {
	verifyAltitude := int64(1)
	verifyLedger := new(kinds.Ledger)
	verifyLedgerDigest := []byte("REDACTED")
	verifyLedger.Altitude = verifyAltitude
	verifyLedger.FinalEndorseDigest = verifyLedgerDigest
	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)

	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate.On("REDACTED").Return(verifyAltitude)
	ledgerDepotSimulate.On("REDACTED").Return(int64(0))
	ledgerDepotSimulate.On("REDACTED", verifyAltitude).Return(&kinds.LedgerSummary{
		LedgerUUID: kinds.LedgerUUID{
			Digest: verifyLedgerDigest,
		},
	})
	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}
	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)

	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)
	res, err := cli.LedgerchainDetails(context.Background(), 0, 100)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, verifyLedgerDigest, []byte(res.LedgerMetadata[0].LedgerUUID.Digest))

	abort()
	wg.Wait()

	ledgerDepotSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
}

func VerifyAssessors(t *testing.T) {
	verifyAltitude := int64(1)
	verifyBallotingPotency := int64(100)
	verifyAssessors := kinds.AssessorAssign{
		Assessors: []*kinds.Assessor{
			{
				BallotingPotency: verifyBallotingPotency,
			},
		},
	}
	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)
	statusDepotSimulate.On("REDACTED", verifyAltitude).Return(&verifyAssessors, nil)

	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(nil)
	ledgerDepotSimulate.On("REDACTED").Return(verifyAltitude)
	ledgerDepotSimulate.On("REDACTED").Return(int64(0))
	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}
	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)

	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)

	verifyScreen := 1
	verifyEveryScreen := 100
	res, err := cli.Assessors(context.Background(), &verifyAltitude, &verifyScreen, &verifyEveryScreen)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, verifyBallotingPotency, res.Assessors[0].BallotingPotency)

	abort()
	wg.Wait()

	ledgerDepotSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
}

func VerifyLedgerLookup(t *testing.T) {
	verifyAltitude := int64(1)
	verifyLedgerDigest := []byte("REDACTED")
	verifyInquire := "REDACTED"
	statusDepotSimulate := &machinestubs.Depot{}
	statusDepotSimulate.On("REDACTED").Return(nil)

	ledgerDepotSimulate := &machinestubs.LedgerDepot{}
	ledgerDepotSimulate.On("REDACTED").Return(nil)

	transferOrdinalizerSimulate := &transindexsimulate.TransferOrdinalizer{}
	ldgOffsetSimulate := &indexsimulate.LedgerOrdinalizer{}
	ledgerDepotSimulate.On("REDACTED", verifyAltitude).Return(&kinds.Ledger{
		Heading: kinds.Heading{
			Altitude: verifyAltitude,
		},
	}, nil)
	ledgerDepotSimulate.On("REDACTED", verifyAltitude).Return(&kinds.LedgerSummary{
		LedgerUUID: kinds.LedgerUUID{
			Digest: verifyLedgerDigest,
		},
	})
	ldgOffsetSimulate.On("REDACTED", mock.Anything,
		mock.MatchedBy(func(q *inquire.Inquire) bool { return verifyInquire == q.Text() })).
		Return([]int64{verifyAltitude}, nil)
	remoteSettings := settings.VerifyRemoteSettings()
	d := scrutinize.New(remoteSettings, ledgerDepotSimulate, statusDepotSimulate, transferOrdinalizerSimulate, ldgOffsetSimulate)

	ctx, abort := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	initiatedGroup := &sync.WaitGroup{}
	initiatedGroup.Add(1)
	go func() {
		initiatedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	initiatedGroup.Wait()
	demandRelate(t, remoteSettings.OverhearLocation, 20)
	cli, err := webcustomer.New(remoteSettings.OverhearLocation, "REDACTED")
	require.NoError(t, err)

	verifyScreen := 1
	verifyEveryScreen := 100
	verifySequenceVia := "REDACTED"
	res, err := cli.LedgerLookup(context.Background(), verifyInquire, &verifyScreen, &verifyEveryScreen, verifySequenceVia)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, verifyLedgerDigest, []byte(res.Ledgers[0].LedgerUUID.Digest))

	abort()
	wg.Wait()

	ledgerDepotSimulate.AssertExpectations(t)
	statusDepotSimulate.AssertExpectations(t)
}

func demandRelate(t testing.TB, location string, attempts int) {
	fragments := strings.SplitN(location, "REDACTED", 2)
	if len(fragments) != 2 {
		t.Fatalf("REDACTED", location)
	}
	var err error
	for i := 0; i < attempts; i++ {
		var link net.Conn
		link, err = net.Dial(fragments[0], fragments[1])
		if err == nil {
			link.Close()
			return
		}
		//
		time.Sleep(time.Microsecond * 100)
	}
	t.Fatalf("REDACTED", location, attempts, err)
}
