package status_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	engineseed "github.com/valkyrieworks/utils/random"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

func VerifyTransferRefine(t *testing.T) {
	generatePaper := arbitraryOriginPaper()
	generatePaper.AgreementOptions.Ledger.MaximumOctets = 3000
	generatePaper.AgreementOptions.Proof.MaximumOctets = 1500

	//
	//
	verifyScenarios := []struct {
		tx    kinds.Tx
		isErr bool
	}{
		{kinds.Tx(engineseed.Octets(2122)), false},
		{kinds.Tx(engineseed.Octets(2123)), true},
		{kinds.Tx(engineseed.Octets(3000)), true},
	}

	for i, tc := range verifyScenarios {
		statusStore, err := dbm.NewStore("REDACTED", "REDACTED", os.TempDir())
		require.NoError(t, err)
		statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
			DropIfaceReplies: false,
		})
		status, err := statusDepot.ImportFromStoreOrOriginPaper(generatePaper)
		require.NoError(t, err)

		f := sm.TransferPreInspect(status)
		if tc.isErr {
			assert.NotNil(t, f(tc.tx), "REDACTED", i)
		} else {
			assert.Nil(t, f(tc.tx), "REDACTED", i)
		}
	}
}
