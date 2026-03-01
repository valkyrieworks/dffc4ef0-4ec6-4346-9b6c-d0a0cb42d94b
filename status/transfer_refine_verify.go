package status_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyTransferRefine(t *testing.T) {
	producePaper := unpredictableInaugurationPaper()
	producePaper.AgreementSettings.Ledger.MaximumOctets = 3000
	producePaper.AgreementSettings.Proof.MaximumOctets = 1500

	//
	//
	verifyScenarios := []struct {
		tx    kinds.Tx
		equalsFault bool
	}{
		{kinds.Tx(commitrand.Octets(2122)), false},
		{kinds.Tx(commitrand.Octets(2123)), true},
		{kinds.Tx(commitrand.Octets(3000)), true},
	}

	for i, tc := range verifyScenarios {
		statusDatastore, err := dbm.FreshDatastore("REDACTED", "REDACTED", os.TempDir())
		require.NoError(t, err)
		statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
			EjectIfaceReplies: false,
		})
		status, err := statusDepot.FetchOriginatingDatastoreEitherOriginPaper(producePaper)
		require.NoError(t, err)

		f := sm.TransferPriorInspect(status)
		if tc.equalsFault {
			assert.NotNil(t, f(tc.tx), "REDACTED", i)
		} else {
			assert.Nil(t, f(tc.tx), "REDACTED", i)
		}
	}
}
