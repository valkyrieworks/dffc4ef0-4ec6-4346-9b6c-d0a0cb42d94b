package depot

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

//
//
//
func CriterionIteratedImportViewedEndorseIdenticalLedger(b *testing.B) {
	status, bs, sanitize := createStatusAndLedgerDepot()
	defer sanitize()
	h := bs.Level() + 1
	ledger, err := status.CreateLedger(h, verify.CreateNTrans(h, 10), new(kinds.Endorse), nil, status.Ratifiers.FetchRecommender().Location)
	require.NoError(b, err)
	viewedEndorse := createVerifyExtensionEndorseWithCountAutographs(ledger.Heading.Level, engineclock.Now(), 100).ToEndorse()
	ps, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(b, err)
	bs.PersistLedger(ledger, ps, viewedEndorse)

	//
	res := bs.ImportViewedEndorse(ledger.Level)
	require.Equal(b, viewedEndorse, res)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := bs.ImportViewedEndorse(ledger.Level)
		require.NotNil(b, res)
	}
}
