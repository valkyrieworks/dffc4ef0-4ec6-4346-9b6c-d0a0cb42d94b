package directives

import (
	"context"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	ifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongmindsettings "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	blockmocks "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
	transfermocks "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	altitude int64 = 10
	foundation   int64 = 2
)

func configureAgainOrdinalIncidentDirective() *cobra.Command {
	againOrdinalIncidentDirective := &cobra.Command{
		Use: AgainOrdinalIncidentDirective.Use,
		Run: func(cmd *cobra.Command, arguments []string) {},
	}

	_ = againOrdinalIncidentDirective.ExecuteContext(context.Background())

	return againOrdinalIncidentDirective
}

func VerifyAgainOrdinalIncidentInspectAltitude(t *testing.T) {
	simulateLedgerDepot := &simulations.LedgerDepot{}
	simulateLedgerDepot.
		On("REDACTED").Return(foundation).
		On("REDACTED").Return(altitude)

	verifyScenarios := []struct {
		initiateAltitude int64
		terminateAltitude   int64
		soundAltitude bool
	}{
		{0, 0, true},
		{0, foundation, true},
		{0, foundation - 1, false},
		{0, altitude, true},
		{0, altitude + 1, true},
		{0, 0, true},
		{foundation - 1, 0, false},
		{foundation, 0, true},
		{foundation, foundation, true},
		{foundation, foundation - 1, false},
		{foundation, altitude, true},
		{foundation, altitude + 1, true},
		{altitude, 0, true},
		{altitude, foundation, false},
		{altitude, altitude - 1, false},
		{altitude, altitude, true},
		{altitude, altitude + 1, true},
		{altitude + 1, 0, false},
	}

	for _, tc := range verifyScenarios {
		initiateAltitude = tc.initiateAltitude
		terminateAltitude = tc.terminateAltitude

		err := inspectSoundAltitude(simulateLedgerDepot)
		if tc.soundAltitude {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}

func VerifyFetchIncidentReceiver(t *testing.T) {
	verifyScenarios := []struct {
		receivers   string
		linkWebroute string
		fetchFault bool
	}{
		{"REDACTED", "REDACTED", true},
		{"REDACTED", "REDACTED", true},
		{"REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", true}, //
		//
		{"REDACTED", "REDACTED", true},
	}

	for idx, tc := range verifyScenarios {
		cfg := strongmindsettings.VerifySettings()
		cfg.TransferOrdinal.Ordinalizer = tc.receivers
		cfg.TransferOrdinal.SqlsLink = tc.linkWebroute
		_, _, err := fetchIncidentReceivers(cfg, verify.FallbackVerifySuccessionUUID)
		if tc.fetchFault {
			require.Error(t, err, idx)
		} else {
			require.NoError(t, err, idx)
		}
	}
}

func VerifyFetchLedgerDepot(t *testing.T) {
	cfg := strongmindsettings.VerifySettings()
	cfg.DatastoreRoute = t.TempDir()
	_, _, err := fetchStatusAlsoLedgerDepot(cfg)
	require.Error(t, err)

	_, err = dbm.FreshDatastore("REDACTED", dbm.ProceedStratumDatastoreRepository, cfg.DatastorePath())
	require.NoError(t, err)

	//
	_, err = dbm.FreshDatastore("REDACTED", dbm.ProceedStratumDatastoreRepository, cfg.DatastorePath())
	require.NoError(t, err)

	bs, ss, err := fetchStatusAlsoLedgerDepot(cfg)
	require.NoError(t, err)
	require.NotNil(t, bs)
	require.NotNil(t, ss)
}

func VerifyAgainOrdinalIncident(t *testing.T) {
	simulateLedgerDepot := &simulations.LedgerDepot{}
	simulateStatusDepot := &simulations.Depot{}
	simulateLedgerOrdinalizer := &blockmocks.LedgerOrdinalizer{}
	simulateTransferOrdinalizer := &transfermocks.TransferOrdinalizer{}

	simulateLedgerDepot.
		On("REDACTED").Return(foundation).
		On("REDACTED").Return(altitude).
		On("REDACTED", foundation).Return(nil).Once().
		On("REDACTED", foundation).Return(&kinds.Ledger{Data: kinds.Data{Txs: kinds.Txs{make(kinds.Tx, 1)}}}).
		On("REDACTED", altitude).Return(&kinds.Ledger{Data: kinds.Data{Txs: kinds.Txs{make(kinds.Tx, 1)}}})

	ifaceReply := &ifacetypes.ReplyCulminateLedger{
		TransferOutcomes: []*ifacetypes.InvokeTransferOutcome{
			{Cipher: 1},
		},
	}

	simulateLedgerOrdinalizer.
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(errors.New("REDACTED")).Once().
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)

	simulateTransferOrdinalizer.
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(errors.New("REDACTED")).Once().
		On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)

	simulateStatusDepot.
		On("REDACTED", foundation).Return(nil, errors.New("REDACTED")).Once().
		On("REDACTED", foundation).Return(ifaceReply, nil).
		On("REDACTED", altitude).Return(ifaceReply, nil)

	verifyScenarios := []struct {
		initiateAltitude int64
		terminateAltitude   int64
		againOrdinalFault  bool
	}{
		{foundation, altitude, true}, //
		{foundation, altitude, true}, //
		{foundation, altitude, true}, //
		{foundation, altitude, true}, //
		{foundation, foundation, false},
		{altitude, altitude, false},
	}

	for _, tc := range verifyScenarios {
		arguments := incidentAgainOrdinalArguments{
			initiateAltitude:  tc.initiateAltitude,
			terminateAltitude:    tc.terminateAltitude,
			ledgerOrdinalizer: simulateLedgerOrdinalizer,
			transferOrdinalizer:    simulateTransferOrdinalizer,
			ledgerDepot:   simulateLedgerDepot,
			statusDepot:   simulateStatusDepot,
		}

		err := incidentAgainOrdinal(configureAgainOrdinalIncidentDirective(), arguments)
		if tc.againOrdinalFault {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
