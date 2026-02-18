package directives

import (
	"context"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	ifacetypes "github.com/valkyrieworks/iface/kinds"
	cometsettings "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/intrinsic/verify"
	ledgermocks "github.com/valkyrieworks/status/ordinaler/simulations"
	"github.com/valkyrieworks/status/simulations"
	txmocks "github.com/valkyrieworks/status/transordinal/simulations"
	"github.com/valkyrieworks/kinds"
)

const (
	level int64 = 10
	root   int64 = 2
)

func configureReOrdinalEventCommand() *cobra.Command {
	reOrdinalEventCommand := &cobra.Command{
		Use: ReOrdinalEventCommand.Use,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	_ = reOrdinalEventCommand.ExecuteContext(context.Background())

	return reOrdinalEventCommand
}

func VerifyReOrdinalEventInspectLevel(t *testing.T) {
	emulateLedgerDepot := &simulations.LedgerDepot{}
	emulateLedgerDepot.
		On("REDACTED").Return(root).
		On("REDACTED").Return(level)

	verifyScenarios := []struct {
		beginLevel int64
		terminateLevel   int64
		soundLevel bool
	}{
		{0, 0, true},
		{0, root, true},
		{0, root - 1, false},
		{0, level, true},
		{0, level + 1, true},
		{0, 0, true},
		{root - 1, 0, false},
		{root, 0, true},
		{root, root, true},
		{root, root - 1, false},
		{root, level, true},
		{root, level + 1, true},
		{level, 0, true},
		{level, root, false},
		{level, level - 1, false},
		{level, level, true},
		{level, level + 1, true},
		{level + 1, 0, false},
	}

	for _, tc := range verifyScenarios {
		beginLevel = tc.beginLevel
		terminateLevel = tc.terminateLevel

		err := inspectSoundLevel(emulateLedgerDepot)
		if tc.soundLevel {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}

func VerifyImportEventDrain(t *testing.T) {
	verifyScenarios := []struct {
		drains   string
		linkURL string
		importErr bool
	}{
		{"REDACTED", "REDACTED", true},
		{"REDACTED", "REDACTED", true},
		{"REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", true}, //
		//
		{"REDACTED", "REDACTED", true},
	}

	for idx, tc := range verifyScenarios {
		cfg := cometsettings.VerifySettings()
		cfg.TransOrdinal.Ordinaler = tc.drains
		cfg.TransOrdinal.PsqlLink = tc.linkURL
		_, _, err := importEventDrains(cfg, verify.StandardVerifyLedgerUID)
		if tc.importErr {
			require.Error(t, err, idx)
		} else {
			require.NoError(t, err, idx)
		}
	}
}

func VerifyImportLedgerDepot(t *testing.T) {
	cfg := cometsettings.VerifySettings()
	cfg.StoreRoute = t.TempDir()
	_, _, err := importStatusAndLedgerDepot(cfg)
	require.Error(t, err)

	_, err = dbm.NewStore("REDACTED", dbm.GoLayerStoreServer, cfg.StoreFolder())
	require.NoError(t, err)

	//
	_, err = dbm.NewStore("REDACTED", dbm.GoLayerStoreServer, cfg.StoreFolder())
	require.NoError(t, err)

	bs, ss, err := importStatusAndLedgerDepot(cfg)
	require.NoError(t, err)
	require.NotNil(t, bs)
	require.NotNil(t, ss)
}

func VerifyReOrdinalEvent(t *testing.T) {
	emulateLedgerDepot := &simulations.LedgerDepot{}
	emulateStatusDepot := &simulations.Depot{}
	emulateLedgerOrdinaler := &ledgermocks.LedgerOrdinaler{}
	emulateTransferOrdinaler := &txmocks.TransOrdinaler{}

	emulateLedgerDepot.
		On("REDACTED").Return(root).
		On("REDACTED").Return(level).
		On("REDACTED", root).Return(nil).Once().
		On("REDACTED", root).Return(&kinds.Ledger{Data: kinds.Data{Txs: kinds.Txs{make(kinds.Tx, 1)}}}).
		On("REDACTED", level).Return(&kinds.Ledger{Data: kinds.Data{Txs: kinds.Txs{make(kinds.Tx, 1)}}})

	ifaceReply := &ifacetypes.ReplyCompleteLedger{
		TransOutcomes: []*ifacetypes.InvokeTransferOutcome{
			{Code: 1},
		},
	}

	emulateLedgerOrdinaler.
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(errors.New("REDACTED")).Once().
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)

	emulateTransferOrdinaler.
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(errors.New("REDACTED")).Once().
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)

	emulateStatusDepot.
		On("REDACTED", root).Return(nil, errors.New("REDACTED")).Once().
		On("REDACTED", root).Return(ifaceReply, nil).
		On("REDACTED", level).Return(ifaceReply, nil)

	verifyScenarios := []struct {
		beginLevel int64
		terminateLevel   int64
		reOrdinalErr  bool
	}{
		{root, level, true}, //
		{root, level, true}, //
		{root, level, true}, //
		{root, level, true}, //
		{root, root, false},
		{level, level, false},
	}

	for _, tc := range verifyScenarios {
		args := eventReOrdinalArgs{
			beginLevel:  tc.beginLevel,
			terminateLevel:    tc.terminateLevel,
			ledgerOrdinaler: emulateLedgerOrdinaler,
			transOrdinaler:    emulateTransferOrdinaler,
			ledgerDepot:   emulateLedgerDepot,
			statusDepot:   emulateStatusDepot,
		}

		err := eventReOrdinal(configureReOrdinalEventCommand(), args)
		if tc.reOrdinalErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
