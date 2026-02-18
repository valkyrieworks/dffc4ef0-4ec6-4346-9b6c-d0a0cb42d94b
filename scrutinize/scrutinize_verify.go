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

	ifacetypes "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/scrutinize"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	httpcustomer "github.com/valkyrieworks/rpc/customer/http"
	ordinalermocks "github.com/valkyrieworks/status/ordinaler/simulations"
	statemulators "github.com/valkyrieworks/status/simulations"
	transferindexmocks "github.com/valkyrieworks/status/transordinal/simulations"
	"github.com/valkyrieworks/kinds"
	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func VerifyScrutinizeBuilder(t *testing.T) {
	cfg := verify.RestoreVerifyOrigin("REDACTED")
	t.Cleanup(leaktest.Check(t))
	defer func() { _ = os.RemoveAll(cfg.OriginFolder) }()
	t.Run("REDACTED", func(t *testing.T) {
		d, err := scrutinize.NewFromSettings(cfg)
		require.NoError(t, err)
		require.NotNil(t, d)
	})
}

func VerifyScrutinizeRun(t *testing.T) {
	cfg := verify.RestoreVerifyOrigin("REDACTED")
	t.Cleanup(leaktest.Check(t))
	defer func() { _ = os.RemoveAll(cfg.OriginFolder) }()
	t.Run("REDACTED", func(t *testing.T) {
		d, err := scrutinize.NewFromSettings(cfg)
		require.NoError(t, err)
		ctx, revoke := context.WithCancel(context.Background())
		ceasedGroup := &sync.WaitGroup{}
		ceasedGroup.Add(1)
		go func() {
			require.NoError(t, d.Run(ctx))
			ceasedGroup.Done()
		}()
		revoke()
		ceasedGroup.Wait()
	})
}

func VerifyLedger(t *testing.T) {
	verifyLevel := int64(1)
	verifyLedger := new(kinds.Ledger)
	verifyLedger.Level = verifyLevel
	verifyLedger.FinalEndorseDigest = []byte("REDACTED")
	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)

	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(verifyLevel)
	ledgerDepotEmulate.On("REDACTED").Return(int64(0))
	ledgerDepotEmulate.On("REDACTED", verifyLevel).Return(&kinds.LedgerMeta{})
	ledgerDepotEmulate.On("REDACTED", verifyLevel).Return(verifyLedger)
	ledgerDepotEmulate.On("REDACTED").Return(nil)

	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}

	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)
	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)
	outcomeLedger, err := cli.Ledger(context.Background(), &verifyLevel)
	require.NoError(t, err)
	require.Equal(t, verifyLedger.Level, outcomeLedger.Ledger.Level)
	require.Equal(t, verifyLedger.FinalEndorseDigest, outcomeLedger.Ledger.FinalEndorseDigest)
	revoke()
	wg.Wait()

	ledgerDepotEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
}

func VerifyTransferScan(t *testing.T) {
	verifyDigest := []byte("REDACTED")
	verifyTransfer := []byte("REDACTED")
	verifyInquire := fmt.Sprintf("REDACTED", string(verifyDigest))
	verifyTransferOutcome := &ifacetypes.TransOutcome{
		Level: 1,
		Ordinal:  100,
		Tx:     verifyTransfer,
	}

	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(nil)
	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}
	transferOrdinalerEmulate.On("REDACTED", mock.Anything,
		mock.MatchedBy(func(q *inquire.Inquire) bool {
			return verifyInquire == strings.ReplaceAll(q.String(), "REDACTED", "REDACTED")
		})).
		Return([]*ifacetypes.TransOutcome{verifyTransferOutcome}, nil)

	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)
	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)

	screen := 1
	outcomeTransferScan, err := cli.TransferScan(context.Background(), verifyInquire, false, &screen, &screen, "REDACTED")
	require.NoError(t, err)
	require.Len(t, outcomeTransferScan.Txs, 1)
	require.Equal(t, kinds.Tx(verifyTransfer), outcomeTransferScan.Txs[0].Tx)

	revoke()
	wg.Wait()

	transferOrdinalerEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
	ledgerDepotEmulate.AssertExpectations(t)
}

func VerifyTransfer(t *testing.T) {
	verifyDigest := []byte("REDACTED")
	verifyTransfer := []byte("REDACTED")

	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(nil)
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}
	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	transferOrdinalerEmulate.On("REDACTED", verifyDigest).Return(&ifacetypes.TransOutcome{
		Tx: verifyTransfer,
	}, nil)

	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)
	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)

	res, err := cli.Tx(context.Background(), verifyDigest, false)
	require.NoError(t, err)
	require.Equal(t, kinds.Tx(verifyTransfer), res.Tx)

	revoke()
	wg.Wait()

	transferOrdinalerEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
	ledgerDepotEmulate.AssertExpectations(t)
}

func VerifyAgreementOptions(t *testing.T) {
	verifyLevel := int64(1)
	verifyMaximumFuel := int64(55)
	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate.On("REDACTED").Return(verifyLevel)
	ledgerDepotEmulate.On("REDACTED").Return(int64(0))
	statusDepotEmulate.On("REDACTED", verifyLevel).Return(kinds.AgreementOptions{
		Ledger: kinds.LedgerOptions{
			MaximumFuel: verifyMaximumFuel,
		},
	}, nil)
	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}
	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)

	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)
	options, err := cli.AgreementOptions(context.Background(), &verifyLevel)
	require.NoError(t, err)
	require.Equal(t, options.AgreementOptions.Ledger.MaximumFuel, verifyMaximumFuel)

	revoke()
	wg.Wait()

	ledgerDepotEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
}

func VerifyLedgerOutcomes(t *testing.T) {
	verifyLevel := int64(1)
	verifyFuelApplied := int64(100)
	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)
	//
	statusDepotEmulate.On("REDACTED", verifyLevel).Return(&ifacetypes.ReplyCompleteLedger{
		TransOutcomes: []*ifacetypes.InvokeTransferOutcome{
			{
				FuelApplied: verifyFuelApplied,
			},
		},
	}, nil)
	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate.On("REDACTED").Return(int64(0))
	ledgerDepotEmulate.On("REDACTED").Return(verifyLevel)
	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}
	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)

	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)
	res, err := cli.LedgerOutcomes(context.Background(), &verifyLevel)
	require.NoError(t, err)
	require.Equal(t, res.TransOutcomes[0].FuelApplied, verifyFuelApplied)

	revoke()
	wg.Wait()

	ledgerDepotEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
}

func VerifyEndorse(t *testing.T) {
	verifyLevel := int64(1)
	verifyEpoch := int32(101)
	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate.On("REDACTED").Return(int64(0))
	ledgerDepotEmulate.On("REDACTED").Return(verifyLevel)
	ledgerDepotEmulate.On("REDACTED", verifyLevel).Return(&kinds.LedgerMeta{}, nil)
	ledgerDepotEmulate.On("REDACTED", verifyLevel).Return(&kinds.Endorse{
		Level: verifyLevel,
		Cycle:  verifyEpoch,
	}, nil)
	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}
	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)

	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)
	res, err := cli.Endorse(context.Background(), &verifyLevel)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Endorse.Cycle, verifyEpoch)

	revoke()
	wg.Wait()

	ledgerDepotEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
}

func VerifyLedgerByDigest(t *testing.T) {
	verifyLevel := int64(1)
	verifyDigest := []byte("REDACTED")
	verifyLedger := new(kinds.Ledger)
	verifyLedger.Level = verifyLevel
	verifyLedger.FinalEndorseDigest = verifyDigest
	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate.On("REDACTED", verifyLevel).Return(&kinds.LedgerMeta{
		LedgerUID: kinds.LedgerUID{
			Digest: verifyDigest,
		},
		Heading: kinds.Heading{
			Level: verifyLevel,
		},
	}, nil)
	ledgerDepotEmulate.On("REDACTED", verifyDigest).Return(verifyLedger, nil)
	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}
	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)

	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)
	res, err := cli.LedgerByDigest(context.Background(), verifyDigest)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, []byte(res.LedgerUID.Digest), verifyDigest)

	revoke()
	wg.Wait()

	ledgerDepotEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
}

func VerifyLedgerchain(t *testing.T) {
	verifyLevel := int64(1)
	verifyLedger := new(kinds.Ledger)
	verifyLedgerDigest := []byte("REDACTED")
	verifyLedger.Level = verifyLevel
	verifyLedger.FinalEndorseDigest = verifyLedgerDigest
	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)

	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate.On("REDACTED").Return(verifyLevel)
	ledgerDepotEmulate.On("REDACTED").Return(int64(0))
	ledgerDepotEmulate.On("REDACTED", verifyLevel).Return(&kinds.LedgerMeta{
		LedgerUID: kinds.LedgerUID{
			Digest: verifyLedgerDigest,
		},
	})
	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}
	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)

	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)
	res, err := cli.LedgerchainDetails(context.Background(), 0, 100)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, verifyLedgerDigest, []byte(res.LedgerMetadata[0].LedgerUID.Digest))

	revoke()
	wg.Wait()

	ledgerDepotEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
}

func VerifyRatifiers(t *testing.T) {
	verifyLevel := int64(1)
	verifyPollingEnergy := int64(100)
	verifyRatifiers := kinds.RatifierAssign{
		Ratifiers: []*kinds.Ratifier{
			{
				PollingEnergy: verifyPollingEnergy,
			},
		},
	}
	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)
	statusDepotEmulate.On("REDACTED", verifyLevel).Return(&verifyRatifiers, nil)

	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(nil)
	ledgerDepotEmulate.On("REDACTED").Return(verifyLevel)
	ledgerDepotEmulate.On("REDACTED").Return(int64(0))
	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}
	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)

	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)

	verifyScreen := 1
	verifyEachScreen := 100
	res, err := cli.Ratifiers(context.Background(), &verifyLevel, &verifyScreen, &verifyEachScreen)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, verifyPollingEnergy, res.Ratifiers[0].PollingEnergy)

	revoke()
	wg.Wait()

	ledgerDepotEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
}

func VerifyLedgerScan(t *testing.T) {
	verifyLevel := int64(1)
	verifyLedgerDigest := []byte("REDACTED")
	verifyInquire := "REDACTED"
	statusDepotEmulate := &statemulators.Depot{}
	statusDepotEmulate.On("REDACTED").Return(nil)

	ledgerDepotEmulate := &statemulators.LedgerDepot{}
	ledgerDepotEmulate.On("REDACTED").Return(nil)

	transferOrdinalerEmulate := &transferindexmocks.TransOrdinaler{}
	recordIndexEmulate := &ordinalermocks.LedgerOrdinaler{}
	ledgerDepotEmulate.On("REDACTED", verifyLevel).Return(&kinds.Ledger{
		Heading: kinds.Heading{
			Level: verifyLevel,
		},
	}, nil)
	ledgerDepotEmulate.On("REDACTED", verifyLevel).Return(&kinds.LedgerMeta{
		LedgerUID: kinds.LedgerUID{
			Digest: verifyLedgerDigest,
		},
	})
	recordIndexEmulate.On("REDACTED", mock.Anything,
		mock.MatchedBy(func(q *inquire.Inquire) bool { return verifyInquire == q.String() })).
		Return([]int64{verifyLevel}, nil)
	rpcSettings := settings.VerifyRPCSettings()
	d := scrutinize.New(rpcSettings, ledgerDepotEmulate, statusDepotEmulate, transferOrdinalerEmulate, recordIndexEmulate)

	ctx, revoke := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	launchedGroup := &sync.WaitGroup{}
	launchedGroup.Add(1)
	go func() {
		launchedGroup.Done()
		defer wg.Done()
		require.NoError(t, d.Run(ctx))
	}()
	//
	//
	launchedGroup.Wait()
	demandAttach(t, rpcSettings.AcceptLocation, 20)
	cli, err := httpcustomer.New(rpcSettings.AcceptLocation, "REDACTED")
	require.NoError(t, err)

	verifyScreen := 1
	verifyEachScreen := 100
	verifySequenceBy := "REDACTED"
	res, err := cli.LedgerScan(context.Background(), verifyInquire, &verifyScreen, &verifyEachScreen, verifySequenceBy)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, verifyLedgerDigest, []byte(res.Ledgers[0].LedgerUID.Digest))

	revoke()
	wg.Wait()

	ledgerDepotEmulate.AssertExpectations(t)
	statusDepotEmulate.AssertExpectations(t)
}

func demandAttach(t testing.TB, address string, attempts int) {
	segments := strings.SplitN(address, "REDACTED", 2)
	if len(segments) != 2 {
		t.Fatalf("REDACTED", address)
	}
	var err error
	for i := 0; i < attempts; i++ {
		var link net.Conn
		link, err = net.Dial(segments[0], segments[1])
		if err == nil {
			link.Close()
			return
		}
		//
		time.Sleep(time.Microsecond * 100)
	}
	t.Fatalf("REDACTED", address, attempts, err)
}
