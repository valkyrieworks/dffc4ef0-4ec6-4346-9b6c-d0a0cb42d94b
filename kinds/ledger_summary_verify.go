package kinds

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

func Testblocksummary_Toschemaformat(t *testing.T) {
	h := createArbitraryHeading()
	bi := LedgerUUID{Digest: h.Digest(), FragmentAssignHeading: FragmentAssignHeading{Sum: 123, Digest: commitrand.Octets(tenderminthash.Extent)}}

	bm := &LedgerSummary{
		LedgerUUID:   bi,
		LedgerExtent: 200,
		Heading:    h,
		CountTrans:    0,
	}

	verifies := []struct {
		verifyAlias string
		bm       *LedgerSummary
		expirationFault   bool
	}{
		{"REDACTED", bm, false},
		{"REDACTED", nil, true},
	}

	for _, tt := range verifies {
		t.Run(tt.verifyAlias, func(t *testing.T) {
			pb := tt.bm.TowardSchema()

			bm, err := LedgerSummaryOriginatingReliableSchema(pb)

			if !tt.expirationFault {
				require.NoError(t, err, tt.verifyAlias)
				require.Equal(t, tt.bm, bm, tt.verifyAlias)
			} else {
				require.Error(t, err, tt.verifyAlias)
			}
		})
	}
}

func Testblocksummary_Certifyfundamental(t *testing.T) {
	h := createArbitraryHeading()
	bi := LedgerUUID{Digest: h.Digest(), FragmentAssignHeading: FragmentAssignHeading{Sum: 123, Digest: commitrand.Octets(tenderminthash.Extent)}}
	bi2 := LedgerUUID{
		Digest:          commitrand.Octets(tenderminthash.Extent),
		FragmentAssignHeading: FragmentAssignHeading{Sum: 123, Digest: commitrand.Octets(tenderminthash.Extent)},
	}
	bi3 := LedgerUUID{
		Digest:          []byte("REDACTED"),
		FragmentAssignHeading: FragmentAssignHeading{Sum: 123, Digest: []byte("REDACTED")},
	}

	bm := &LedgerSummary{
		LedgerUUID:   bi,
		LedgerExtent: 200,
		Heading:    h,
		CountTrans:    0,
	}

	bm2 := &LedgerSummary{
		LedgerUUID:   bi2,
		LedgerExtent: 200,
		Heading:    h,
		CountTrans:    0,
	}

	bm3 := &LedgerSummary{
		LedgerUUID:   bi3,
		LedgerExtent: 200,
		Heading:    h,
		CountTrans:    0,
	}

	verifies := []struct {
		alias    string
		bm      *LedgerSummary
		desireFault bool
	}{
		{"REDACTED", bm, false},
		{"REDACTED", bm2, true},
		{"REDACTED", bm3, true},
	}
	for _, tt := range verifies {
		t.Run(tt.alias, func(t *testing.T) {
			if err := tt.bm.CertifyFundamental(); (err != nil) != tt.desireFault {
				t.Errorf("REDACTED", err, tt.desireFault)
			}
		})
	}
}
