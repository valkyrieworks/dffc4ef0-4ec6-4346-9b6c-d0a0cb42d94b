package kinds

import (
	"reflect"
	"testing"

	"github.com/valkyrieworks/vault/comethash"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

func VerifyStandardizeLedgerUID(t *testing.T) {
	arbhash := engineseed.Octets(comethash.Volume)
	ledger1 := engineproto.LedgerUID{
		Digest:          arbhash,
		SegmentAssignHeading: engineproto.SegmentAssignHeading{Sum: 5, Digest: arbhash},
	}
	ledger2 := engineproto.LedgerUID{
		Digest:          arbhash,
		SegmentAssignHeading: engineproto.SegmentAssignHeading{Sum: 10, Digest: arbhash},
	}
	vledger1 := engineproto.StandardLedgerUID{
		Digest:          arbhash,
		SegmentAssignHeading: engineproto.StandardSectionCollectionHeading{Sum: 5, Digest: arbhash},
	}
	vledger2 := engineproto.StandardLedgerUID{
		Digest:          arbhash,
		SegmentAssignHeading: engineproto.StandardSectionCollectionHeading{Sum: 10, Digest: arbhash},
	}

	verifies := []struct {
		label string
		args engineproto.LedgerUID
		desire *engineproto.StandardLedgerUID
	}{
		{"REDACTED", ledger1, &vledger1},
		{"REDACTED", ledger2, &vledger2},
	}
	for _, tt := range verifies {
		t.Run(tt.label, func(t *testing.T) {
			if got := StandardizeLedgerUID(tt.args); !reflect.DeepEqual(got, tt.desire) {
				t.Errorf("REDACTED", got, tt.desire)
			}
		})
	}
}
