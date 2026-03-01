package kinds

import (
	"math"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

var (
	verifyNomination *Nomination
	pbp          *commitchema.Nomination
)

func initialize() {
	imprint, err := time.Parse(MomentLayout, "REDACTED")
	if err != nil {
		panic(err)
	}
	verifyNomination = &Nomination{
		Altitude: 12345,
		Iteration:  23456,
		LedgerUUID: LedgerUUID{
			Digest:          []byte("REDACTED"),
			FragmentAssignHeading: FragmentAssignHeading{Sum: 111, Digest: []byte("REDACTED")},
		},
		PolicyIteration:  -1,
		Timestamp: imprint,
	}
	pbp = verifyNomination.TowardSchema()
}

func VerifyNominationNotatable(t *testing.T) {
	successionUUID := "REDACTED"
	attestOctets := NominationAttestOctets(successionUUID, pbp)
	pb := NormalizeNomination(successionUUID, pbp)

	anticipated, err := protocolio.SerializeSeparated(&pb)
	require.NoError(t, err)
	require.Equal(t, anticipated, attestOctets, "REDACTED")
}

func VerifyNominationText(t *testing.T) {
	str := verifyNomination.Text()
	anticipated := "REDACTED" //
	if str != anticipated {
		t.Errorf("REDACTED", anticipated, str)
	}
}

func VerifyNominationValidateNotation(t *testing.T) {
	privateItem := FreshSimulatePRV()
	publicToken, err := privateItem.ObtainPublicToken()
	require.NoError(t, err)

	item := FreshNomination(
		4, 2, 2,
		LedgerUUID{commitrand.Octets(tenderminthash.Extent), FragmentAssignHeading{777, commitrand.Octets(tenderminthash.Extent)}})
	p := item.TowardSchema()
	attestOctets := NominationAttestOctets("REDACTED", p)

	//
	err = privateItem.AttestNomination("REDACTED", p)
	require.NoError(t, err)
	item.Notation = p.Notation

	//
	sound := publicToken.ValidateNotation(attestOctets, item.Notation)
	require.True(t, sound)

	//
	freshItem := new(commitchema.Nomination)
	pb := item.TowardSchema()

	bs, err := proto.Marshal(pb)
	require.NoError(t, err)

	err = proto.Unmarshal(bs, freshItem)
	require.NoError(t, err)

	np, err := NominationOriginatingSchema(freshItem)
	require.NoError(t, err)

	//
	freshAttestOctets := NominationAttestOctets("REDACTED", pb)
	require.Equal(t, string(attestOctets), string(freshAttestOctets))
	sound = publicToken.ValidateNotation(freshAttestOctets, np.Notation)
	require.True(t, sound)
}

func AssessmentNominationRecordAttestOctets(b *testing.B) {
	for b.Loop() {
		NominationAttestOctets("REDACTED", pbp)
	}
}

func AssessmentNominationAttest(b *testing.B) {
	privateItem := FreshSimulatePRV()
	for b.Loop() {
		err := privateItem.AttestNomination("REDACTED", pbp)
		if err != nil {
			b.Error(err)
		}
	}
}

func AssessmentNominationValidateNotation(b *testing.B) {
	privateItem := FreshSimulatePRV()
	err := privateItem.AttestNomination("REDACTED", pbp)
	require.NoError(b, err)
	publicToken, err := privateItem.ObtainPublicToken()
	require.NoError(b, err)

	for b.Loop() {
		publicToken.ValidateNotation(NominationAttestOctets("REDACTED", pbp), verifyNomination.Notation)
	}
}

func VerifyNominationCertifyFundamental(t *testing.T) {
	privateItem := FreshSimulatePRV()
	verifyScenarios := []struct {
		verifyAlias         string
		distortNomination func(*Nomination)
		anticipateFault        bool
	}{
		{"REDACTED", func(p *Nomination) {}, false},
		{"REDACTED", func(p *Nomination) { p.Kind = commitchema.PreendorseKind }, true},
		{"REDACTED", func(p *Nomination) { p.Altitude = -1 }, true},
		{"REDACTED", func(p *Nomination) { p.Iteration = -1 }, true},
		{"REDACTED", func(p *Nomination) { p.PolicyIteration = -2 }, true},
		{"REDACTED", func(p *Nomination) {
			p.LedgerUUID = LedgerUUID{[]byte{1, 2, 3}, FragmentAssignHeading{111, []byte("REDACTED")}}
		}, true},
		{"REDACTED", func(p *Nomination) {
			p.Notation = make([]byte, 0)
		}, true},
		{"REDACTED", func(p *Nomination) {
			p.Notation = make([]byte, MaximumNotationExtent+1)
		}, true},
	}
	ledgerUUID := createLedgerUUID(tenderminthash.Sum([]byte("REDACTED")), math.MaxInt32, tenderminthash.Sum([]byte("REDACTED")))

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			item := FreshNomination(
				4, 2, 2,
				ledgerUUID)
			p := item.TowardSchema()
			err := privateItem.AttestNomination("REDACTED", p)
			item.Notation = p.Notation
			require.NoError(t, err)
			tc.distortNomination(item)
			assert.Equal(t, tc.anticipateFault, item.CertifyFundamental() != nil, "REDACTED")
		})
	}
}

func VerifyNominationSchemaArea(t *testing.T) {
	nomination := FreshNomination(1, 2, 3, createLedgerUUID([]byte("REDACTED"), 2, []byte("REDACTED")))
	nomination.Notation = []byte("REDACTED")
	item2 := FreshNomination(1, 2, 3, LedgerUUID{})

	verifyScenarios := []struct {
		msg     string
		p1      *Nomination
		expirationPhrase bool
	}{
		{"REDACTED", nomination, true},
		{"REDACTED", item2, false}, //
		{"REDACTED", &Nomination{}, false},
		{"REDACTED", nil, false},
	}
	for _, tc := range verifyScenarios {
		schemaNomination := tc.p1.TowardSchema()

		p, err := NominationOriginatingSchema(schemaNomination)
		if tc.expirationPhrase {
			require.NoError(t, err)
			require.Equal(t, tc.p1, p, tc.msg)
		} else {
			require.Error(t, err)
		}
	}
}

func VerifyNominationCertifyLedgerExtent(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias     string
		maximumLedgerExtent int64
		nomination     *Nomination
		anticipatePhrase   bool
	}{
		{"REDACTED", int64(10 * LedgerFragmentExtentOctets), FreshNomination(0, 0, 0, LedgerUUID{FragmentAssignHeading: FragmentAssignHeading{Sum: 5}}), true},
		{"REDACTED", int64(10 * LedgerFragmentExtentOctets), FreshNomination(0, 0, 0, LedgerUUID{FragmentAssignHeading: FragmentAssignHeading{Sum: 20}}), false},
		{"REDACTED", int64(10 * LedgerFragmentExtentOctets), FreshNomination(0, 0, 0, LedgerUUID{FragmentAssignHeading: FragmentAssignHeading{Sum: math.MaxUint32}}), false},
		{"REDACTED", -1, FreshNomination(0, 0, 0, LedgerUUID{FragmentAssignHeading: FragmentAssignHeading{Sum: math.MaxUint32}}), false},
		{"REDACTED", -1, FreshNomination(0, 0, 0, LedgerUUID{FragmentAssignHeading: FragmentAssignHeading{Sum: math.MaxUint32}}), false},
		{"REDACTED", -1, FreshNomination(0, 0, 0, LedgerUUID{FragmentAssignHeading: FragmentAssignHeading{Sum: 1600}}), true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			err := tc.nomination.CertifyLedgerExtent(tc.maximumLedgerExtent)
			if tc.anticipatePhrase {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
