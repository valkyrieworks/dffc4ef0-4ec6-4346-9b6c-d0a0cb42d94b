package kinds

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/comethash"
	engineseed "github.com/valkyrieworks/utils/random"
)

func Verifyblockmeta_Toschema(t *testing.T) {
	h := createRandomHeading()
	bi := LedgerUID{Digest: h.Digest(), SegmentAssignHeading: SegmentAssignHeading{Sum: 123, Digest: engineseed.Octets(comethash.Volume)}}

	bm := &LedgerMeta{
		LedgerUID:   bi,
		LedgerVolume: 200,
		Heading:    h,
		CountTrans:    0,
	}

	verifies := []struct {
		verifyLabel string
		bm       *LedgerMeta
		expirationErr   bool
	}{
		{"REDACTED", bm, false},
		{"REDACTED", nil, true},
	}

	for _, tt := range verifies {
		t.Run(tt.verifyLabel, func(t *testing.T) {
			pb := tt.bm.ToSchema()

			bm, err := LedgerMetaFromValidatedSchema(pb)

			if !tt.expirationErr {
				require.NoError(t, err, tt.verifyLabel)
				require.Equal(t, tt.bm, bm, tt.verifyLabel)
			} else {
				require.Error(t, err, tt.verifyLabel)
			}
		})
	}
}

func Verifyblockmeta_Verifybasic(t *testing.T) {
	h := createRandomHeading()
	bi := LedgerUID{Digest: h.Digest(), SegmentAssignHeading: SegmentAssignHeading{Sum: 123, Digest: engineseed.Octets(comethash.Volume)}}
	bi2 := LedgerUID{
		Digest:          engineseed.Octets(comethash.Volume),
		SegmentAssignHeading: SegmentAssignHeading{Sum: 123, Digest: engineseed.Octets(comethash.Volume)},
	}
	bi3 := LedgerUID{
		Digest:          []byte("REDACTED"),
		SegmentAssignHeading: SegmentAssignHeading{Sum: 123, Digest: []byte("REDACTED")},
	}

	bm := &LedgerMeta{
		LedgerUID:   bi,
		LedgerVolume: 200,
		Heading:    h,
		CountTrans:    0,
	}

	bm2 := &LedgerMeta{
		LedgerUID:   bi2,
		LedgerVolume: 200,
		Heading:    h,
		CountTrans:    0,
	}

	bm3 := &LedgerMeta{
		LedgerUID:   bi3,
		LedgerVolume: 200,
		Heading:    h,
		CountTrans:    0,
	}

	verifies := []struct {
		label    string
		bm      *LedgerMeta
		desireErr bool
	}{
		{"REDACTED", bm, false},
		{"REDACTED", bm2, true},
		{"REDACTED", bm3, true},
	}
	for _, tt := range verifies {
		t.Run(tt.label, func(t *testing.T) {
			if err := tt.bm.CertifySimple(); (err != nil) != tt.desireErr {
				t.Errorf("REDACTED", err, tt.desireErr)
			}
		})
	}
}
