package kinds

import (
	"reflect"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

func VerifyNormalizeLedgerUUID(t *testing.T) {
	unpredictablehash := commitrand.Octets(tenderminthash.Extent)
	ledger1 := commitchema.LedgerUUID{
		Digest:          unpredictablehash,
		FragmentAssignHeading: commitchema.FragmentAssignHeading{Sum: 5, Digest: unpredictablehash},
	}
	ledger2 := commitchema.LedgerUUID{
		Digest:          unpredictablehash,
		FragmentAssignHeading: commitchema.FragmentAssignHeading{Sum: 10, Digest: unpredictablehash},
	}
	ledger1 := commitchema.StandardLedgerUUID{
		Digest:          unpredictablehash,
		FragmentAssignHeading: commitchema.StandardFragmentAssignHeading{Sum: 5, Digest: unpredictablehash},
	}
	ledger2 := commitchema.StandardLedgerUUID{
		Digest:          unpredictablehash,
		FragmentAssignHeading: commitchema.StandardFragmentAssignHeading{Sum: 10, Digest: unpredictablehash},
	}

	verifies := []struct {
		alias string
		arguments commitchema.LedgerUUID
		desire *commitchema.StandardLedgerUUID
	}{
		{"REDACTED", ledger1, &ledger1},
		{"REDACTED", ledger2, &ledger2},
	}
	for _, tt := range verifies {
		t.Run(tt.alias, func(t *testing.T) {
			if got := NormalizeLedgerUUID(tt.arguments); !reflect.DeepEqual(got, tt.desire) {
				t.Errorf("REDACTED", got, tt.desire)
			}
		})
	}
}
