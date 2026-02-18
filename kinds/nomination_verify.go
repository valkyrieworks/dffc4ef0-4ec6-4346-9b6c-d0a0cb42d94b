package kinds

import (
	"math"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/utils/protoio"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

var (
	verifyNomination *Nomination
	pbp          *engineproto.Nomination
)

func init() {
	imprint, err := time.Parse(TimeLayout, "REDACTED")
	if err != nil {
		panic(err)
	}
	verifyNomination = &Nomination{
		Level: 12345,
		Cycle:  23456,
		LedgerUID: LedgerUID{
			Digest:          []byte("REDACTED"),
			SegmentAssignHeading: SegmentAssignHeading{Sum: 111, Digest: []byte("REDACTED")},
		},
		POLDuration:  -1,
		Timestamp: imprint,
	}
	pbp = verifyNomination.ToSchema()
}

func VerifyNominationSignable(t *testing.T) {
	ledgerUID := "REDACTED"
	attestOctets := NominationAttestOctets(ledgerUID, pbp)
	pb := StandardizeNomination(ledgerUID, pbp)

	anticipated, err := protoio.SerializeSeparated(&pb)
	require.NoError(t, err)
	require.Equal(t, anticipated, attestOctets, "REDACTED")
}

func VerifyNominationString(t *testing.T) {
	str := verifyNomination.String()
	anticipated := "REDACTED" //
	if str != anticipated {
		t.Errorf("REDACTED", anticipated, str)
	}
}

func VerifyNominationValidateAutograph(t *testing.T) {
	privateValue := NewEmulatePV()
	publicKey, err := privateValue.FetchPublicKey()
	require.NoError(t, err)

	nomination := NewNomination(
		4, 2, 2,
		LedgerUID{engineseed.Octets(comethash.Volume), SegmentAssignHeading{777, engineseed.Octets(comethash.Volume)}})
	p := nomination.ToSchema()
	attestOctets := NominationAttestOctets("REDACTED", p)

	//
	err = privateValue.AttestNomination("REDACTED", p)
	require.NoError(t, err)
	nomination.Autograph = p.Autograph

	//
	sound := publicKey.ValidateAutograph(attestOctets, nomination.Autograph)
	require.True(t, sound)

	//
	newNomination := new(engineproto.Nomination)
	pb := nomination.ToSchema()

	bs, err := proto.Marshal(pb)
	require.NoError(t, err)

	err = proto.Unmarshal(bs, newNomination)
	require.NoError(t, err)

	np, err := NominationFromSchema(newNomination)
	require.NoError(t, err)

	//
	newAttestOctets := NominationAttestOctets("REDACTED", pb)
	require.Equal(t, string(attestOctets), string(newAttestOctets))
	sound = publicKey.ValidateAutograph(newAttestOctets, np.Autograph)
	require.True(t, sound)
}

func CriterionNominationRecordAttestOctets(b *testing.B) {
	for b.Loop() {
		NominationAttestOctets("REDACTED", pbp)
	}
}

func CriterionNominationAttest(b *testing.B) {
	privateValue := NewEmulatePV()
	for b.Loop() {
		err := privateValue.AttestNomination("REDACTED", pbp)
		if err != nil {
			b.Error(err)
		}
	}
}

func CriterionNominationValidateAutograph(b *testing.B) {
	privateValue := NewEmulatePV()
	err := privateValue.AttestNomination("REDACTED", pbp)
	require.NoError(b, err)
	publicKey, err := privateValue.FetchPublicKey()
	require.NoError(b, err)

	for b.Loop() {
		publicKey.ValidateAutograph(NominationAttestOctets("REDACTED", pbp), verifyNomination.Autograph)
	}
}

func VerifyNominationCertifySimple(t *testing.T) {
	privateValue := NewEmulatePV()
	verifyScenarios := []struct {
		verifyLabel         string
		distortNomination func(*Nomination)
		anticipateErr        bool
	}{
		{"REDACTED", func(p *Nomination) {}, false},
		{"REDACTED", func(p *Nomination) { p.Kind = engineproto.PreendorseKind }, true},
		{"REDACTED", func(p *Nomination) { p.Level = -1 }, true},
		{"REDACTED", func(p *Nomination) { p.Cycle = -1 }, true},
		{"REDACTED", func(p *Nomination) { p.POLDuration = -2 }, true},
		{"REDACTED", func(p *Nomination) {
			p.LedgerUID = LedgerUID{[]byte{1, 2, 3}, SegmentAssignHeading{111, []byte("REDACTED")}}
		}, true},
		{"REDACTED", func(p *Nomination) {
			p.Autograph = make([]byte, 0)
		}, true},
		{"REDACTED", func(p *Nomination) {
			p.Autograph = make([]byte, MaximumAutographVolume+1)
		}, true},
	}
	ledgerUID := createLedgerUID(comethash.Sum([]byte("REDACTED")), math.MaxInt32, comethash.Sum([]byte("REDACTED")))

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			nomination := NewNomination(
				4, 2, 2,
				ledgerUID)
			p := nomination.ToSchema()
			err := privateValue.AttestNomination("REDACTED", p)
			nomination.Autograph = p.Autograph
			require.NoError(t, err)
			tc.distortNomination(nomination)
			assert.Equal(t, tc.anticipateErr, nomination.CertifySimple() != nil, "REDACTED")
		})
	}
}

func VerifyNominationSchemaBuffer(t *testing.T) {
	nomination := NewNomination(1, 2, 3, createLedgerUID([]byte("REDACTED"), 2, []byte("REDACTED")))
	nomination.Autograph = []byte("REDACTED")
	nomination2 := NewNomination(1, 2, 3, LedgerUID{})

	verifyScenarios := []struct {
		msg     string
		p1      *Nomination
		expirationPass bool
	}{
		{"REDACTED", nomination, true},
		{"REDACTED", nomination2, false}, //
		{"REDACTED", &Nomination{}, false},
		{"REDACTED", nil, false},
	}
	for _, tc := range verifyScenarios {
		schemaNomination := tc.p1.ToSchema()

		p, err := NominationFromSchema(schemaNomination)
		if tc.expirationPass {
			require.NoError(t, err)
			require.Equal(t, tc.p1, p, tc.msg)
		} else {
			require.Error(t, err)
		}
	}
}

func VerifyNominationCertifyLedgerVolume(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel     string
		maximumLedgerVolume int64
		nomination     *Nomination
		anticipatePass   bool
	}{
		{"REDACTED", int64(10 * LedgerSegmentVolumeOctets), NewNomination(0, 0, 0, LedgerUID{SegmentAssignHeading: SegmentAssignHeading{Sum: 5}}), true},
		{"REDACTED", int64(10 * LedgerSegmentVolumeOctets), NewNomination(0, 0, 0, LedgerUID{SegmentAssignHeading: SegmentAssignHeading{Sum: 20}}), false},
		{"REDACTED", int64(10 * LedgerSegmentVolumeOctets), NewNomination(0, 0, 0, LedgerUID{SegmentAssignHeading: SegmentAssignHeading{Sum: math.MaxUint32}}), false},
		{"REDACTED", -1, NewNomination(0, 0, 0, LedgerUID{SegmentAssignHeading: SegmentAssignHeading{Sum: math.MaxUint32}}), false},
		{"REDACTED", -1, NewNomination(0, 0, 0, LedgerUID{SegmentAssignHeading: SegmentAssignHeading{Sum: math.MaxUint32}}), false},
		{"REDACTED", -1, NewNomination(0, 0, 0, LedgerUID{SegmentAssignHeading: SegmentAssignHeading{Sum: 1600}}), true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			err := tc.nomination.CertifyLedgerVolume(tc.maximumLedgerVolume)
			if tc.anticipatePass {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
