package depot

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

//
//
//
func AssessmentIteratedFetchObservedEndorseIdenticalLedger(b *testing.B) {
	status, bs, sanitize := createStatusAlsoLedgerDepot()
	defer sanitize()
	h := bs.Altitude() + 1
	ledger, err := status.CreateLedger(h, verify.CreateNTHTrans(h, 10), new(kinds.Endorse), nil, status.Assessors.ObtainNominator().Location)
	require.NoError(b, err)
	observedEndorse := createVerifyAddnEndorseUsingCountSignatures(ledger.Heading.Altitude, committime.Now(), 100).TowardEndorse()
	ps, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(b, err)
	bs.PersistLedger(ledger, ps, observedEndorse)

	//
	res := bs.FetchObservedEndorse(ledger.Altitude)
	require.Equal(b, observedEndorse, res)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := bs.FetchObservedEndorse(ledger.Altitude)
		require.NotNil(b, res)
	}
}
