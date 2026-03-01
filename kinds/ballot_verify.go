package kinds

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

func instancePreballot() *Ballot {
	return instanceBallot(byte(commitchema.PreballotKind))
}

func instancePreendorse() *Ballot {
	ballot := instanceBallot(byte(commitchema.PreendorseKind))
	ballot.Addition = []byte("REDACTED")
	ballot.AdditionNotation = []byte("REDACTED")
	return ballot
}

func instanceBallot(t byte) *Ballot {
	imprint, err := time.Parse(MomentLayout, "REDACTED")
	if err != nil {
		panic(err)
	}

	return &Ballot{
		Kind:      commitchema.AttestedSignalKind(t),
		Altitude:    12345,
		Iteration:     2,
		Timestamp: imprint,
		LedgerUUID: LedgerUUID{
			Digest: tenderminthash.Sum([]byte("REDACTED")),
			FragmentAssignHeading: FragmentAssignHeading{
				Sum: 1000000,
				Digest:  tenderminthash.Sum([]byte("REDACTED")),
			},
		},
		AssessorLocation: security.LocatorDigest([]byte("REDACTED")),
		AssessorOrdinal:   56789,
	}
}

func VerifyBallotNotatable(t *testing.T) {
	ballot := instancePreendorse()
	v := ballot.TowardSchema()
	attestOctets := BallotAttestOctets("REDACTED", v)
	pb := NormalizeBallot("REDACTED", v)
	anticipated, err := protocolio.SerializeSeparated(&pb)
	require.NoError(t, err)

	require.Equal(t, anticipated, attestOctets, "REDACTED")
}

func VerifyBallotAttestOctetsVerifyArrays(t *testing.T) {
	verifies := []struct {
		successionUUID string
		ballot    *Ballot
		desire    []byte
	}{
		0: {
			"REDACTED", &Ballot{},
			//
			[]byte{0xd, 0x2a, 0xb, 0x8, 0x80, 0x92, 0xb8, 0xc3, 0x98, 0xfe, 0xff, 0xff, 0xff, 0x1},
		},
		//
		1: {
			"REDACTED", &Ballot{Altitude: 1, Iteration: 1, Kind: commitchema.PreendorseKind},
			[]byte{
				0x21,                                   //
				0x8,                                    //
				0x2,                                    //
				0x11,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				0x19,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				0x2a, //
				//
				0xb, 0x8, 0x80, 0x92, 0xb8, 0xc3, 0x98, 0xfe, 0xff, 0xff, 0xff, 0x1,
			},
		},
		//
		2: {
			"REDACTED", &Ballot{Altitude: 1, Iteration: 1, Kind: commitchema.PreballotKind},
			[]byte{
				0x21,                                   //
				0x8,                                    //
				0x1,                                    //
				0x11,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				0x19,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				0x2a, //
				//
				0xb, 0x8, 0x80, 0x92, 0xb8, 0xc3, 0x98, 0xfe, 0xff, 0xff, 0xff, 0x1,
			},
		},
		3: {
			"REDACTED", &Ballot{Altitude: 1, Iteration: 1},
			[]byte{
				0x1f,                                   //
				0x11,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				0x19,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				//
				0x2a,
				0xb, 0x8, 0x80, 0x92, 0xb8, 0xc3, 0x98, 0xfe, 0xff, 0xff, 0xff, 0x1,
			},
		},
		//
		4: {
			"REDACTED", &Ballot{Altitude: 1, Iteration: 1},
			[]byte{
				0x2e,                                   //
				0x11,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				0x19,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				//
				0x2a,                                                                //
				0xb, 0x8, 0x80, 0x92, 0xb8, 0xc3, 0x98, 0xfe, 0xff, 0xff, 0xff, 0x1, //
				//
				0x32,
				0xd, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
			}, //
		},
		//
		5: {
			"REDACTED", &Ballot{
				Altitude:    1,
				Iteration:     1,
				Addition: []byte("REDACTED"),
			},
			[]byte{
				0x2e,                                   //
				0x11,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				0x19,                                   //
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, //
				//
				0x2a,                                                                //
				0xb, 0x8, 0x80, 0x92, 0xb8, 0xc3, 0x98, 0xfe, 0xff, 0xff, 0xff, 0x1, //
				//
				0x32,
				0xd, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, //
			}, //
		},
	}
	for i, tc := range verifies {
		v := tc.ballot.TowardSchema()
		got := BallotAttestOctets(tc.successionUUID, v)
		assert.Equal(t, len(tc.desire), len(got), "REDACTED", i)
		assert.Equal(t, tc.desire, got, "REDACTED", i)
	}
}

func VerifyBallotNominationNegationEquals(t *testing.T) {
	cv := NormalizeBallot("REDACTED", &commitchema.Ballot{Altitude: 1, Iteration: 1})
	p := NormalizeNomination("REDACTED", &commitchema.Nomination{Altitude: 1, Iteration: 1})
	vb, err := proto.Marshal(&cv)
	require.NoError(t, err)
	pb, err := proto.Marshal(&p)
	require.NoError(t, err)
	require.NotEqual(t, vb, pb)
}

func VerifyBallotValidateNotation(t *testing.T) {
	privateItem := FreshSimulatePRV()
	publickey, err := privateItem.ObtainPublicToken()
	require.NoError(t, err)

	ballot := instancePreendorse()
	v := ballot.TowardSchema()
	attestOctets := BallotAttestOctets("REDACTED", v)

	//
	err = privateItem.AttestBallot("REDACTED", v)
	require.NoError(t, err)

	//
	sound := publickey.ValidateNotation(BallotAttestOctets("REDACTED", v), v.Notation)
	require.True(t, sound)

	//
	preendorse := new(commitchema.Ballot)
	bs, err := proto.Marshal(v)
	require.NoError(t, err)
	err = proto.Unmarshal(bs, preendorse)
	require.NoError(t, err)

	//
	freshAttestOctets := BallotAttestOctets("REDACTED", preendorse)
	require.Equal(t, string(attestOctets), string(freshAttestOctets))
	sound = publickey.ValidateNotation(freshAttestOctets, preendorse.Notation)
	require.True(t, sound)
}

//
//
func VerifyBallotAddition(t *testing.T) {
	verifyScenarios := []struct {
		alias             string
		addition        []byte
		encompassNotation bool
		anticipateFailure      bool
	}{
		{
			alias:             "REDACTED",
			addition:        []byte("REDACTED"),
			encompassNotation: true,
			anticipateFailure:      false,
		},
		{
			alias:             "REDACTED",
			addition:        []byte("REDACTED"),
			encompassNotation: false,
			anticipateFailure:      true,
		},
		{
			alias:             "REDACTED",
			encompassNotation: true,
			anticipateFailure:      false,
		},
		{
			alias:             "REDACTED",
			encompassNotation: false,
			anticipateFailure:      true,
		},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.alias, func(t *testing.T) {
			altitude, iteration := int64(1), int32(0)
			privateItem := FreshSimulatePRV()
			pk, err := privateItem.ObtainPublicToken()
			require.NoError(t, err)
			ballot := &Ballot{
				AssessorLocation: pk.Location(),
				AssessorOrdinal:   0,
				Altitude:           altitude,
				Iteration:            iteration,
				Timestamp:        committime.Now(),
				Kind:             commitchema.PreendorseKind,
				LedgerUUID:          createLedgerUUIDUnpredictable(),
			}

			v := ballot.TowardSchema()
			err = privateItem.AttestBallot("REDACTED", v)
			require.NoError(t, err)
			ballot.Notation = v.Notation
			if tc.encompassNotation {
				ballot.AdditionNotation = v.AdditionNotation
			}
			err = ballot.ValidateAddition("REDACTED", pk)
			if tc.anticipateFailure {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func VerifyEqualsBallotKindSound(t *testing.T) {
	tc := []struct {
		alias string
		in   commitchema.AttestedSignalKind
		out  bool
	}{
		{"REDACTED", commitchema.PreballotKind, true},
		{"REDACTED", commitchema.PreendorseKind, true},
		{"REDACTED", commitchema.AttestedSignalKind(0x3), false},
	}

	for _, tt := range tc {
		t.Run(tt.alias, func(st *testing.T) {
			if rs := EqualsBallotKindSound(tt.in); rs != tt.out {
				t.Errorf("REDACTED", rs, tt.out)
			}
		})
	}
}

func VerifyBallotValidate(t *testing.T) {
	privateItem := FreshSimulatePRV()
	publickey, err := privateItem.ObtainPublicToken()
	require.NoError(t, err)

	ballot := instancePreballot()
	ballot.AssessorLocation = publickey.Location()

	err = ballot.Validate("REDACTED", edwards25519.ProducePrivateToken().PublicToken())
	if assert.Error(t, err) {
		assert.Equal(t, FaultBallotUnfitAssessorLocator, err)
	}

	err = ballot.Validate("REDACTED", publickey)
	if assert.Error(t, err) {
		assert.Equal(t, FaultBallotUnfitSigning, err)
	}
}

func VerifyBallotText(t *testing.T) {
	str := instancePreendorse().Text()
	anticipated := "REDACTED" //
	if str != anticipated {
		t.Errorf("REDACTED", anticipated, str)
	}

	string2 := instancePreballot().Text()
	anticipated = "REDACTED" //
	if string2 != anticipated {
		t.Errorf("REDACTED", anticipated, string2)
	}
}

func attestBallot(t *testing.T, pv PrivateAssessor, successionUUID string, ballot *Ballot) {
	t.Helper()

	v := ballot.TowardSchema()
	require.NoError(t, pv.AttestBallot(successionUUID, v))
	ballot.Notation = v.Notation
	ballot.AdditionNotation = v.AdditionNotation
}

func VerifySoundBallots(t *testing.T) {
	privateItem := FreshSimulatePRV()

	verifyScenarios := []struct {
		alias         string
		ballot         *Ballot
		distortBallot func(*Ballot)
	}{
		{"REDACTED", instancePreballot(), func(v *Ballot) {}},
		{"REDACTED", instancePreendorse(), func(v *Ballot) { v.Addition = nil }},
		{"REDACTED", instancePreendorse(), func(v *Ballot) { v.Addition = []byte("REDACTED") }},
	}
	for _, tc := range verifyScenarios {
		attestBallot(t, privateItem, "REDACTED", tc.ballot)
		tc.distortBallot(tc.ballot)
		require.NoError(t, tc.ballot.CertifyFundamental(), "REDACTED", tc.alias)
		require.NoError(t, tc.ballot.AssureAddition(), "REDACTED", tc.alias)
	}
}

func VerifyUnfitBallots(t *testing.T) {
	privateItem := FreshSimulatePRV()

	verifyScenarios := []struct {
		alias         string
		distortBallot func(*Ballot)
	}{
		{"REDACTED", func(v *Ballot) { v.Altitude = -1 }},
		{"REDACTED", func(v *Ballot) { v.Iteration = -1 }},
		{"REDACTED", func(v *Ballot) { v.Altitude = 0 }},
		{"REDACTED", func(v *Ballot) { v.LedgerUUID = LedgerUUID{[]byte{1, 2, 3}, FragmentAssignHeading{111, []byte("REDACTED")}} }},
		{"REDACTED", func(v *Ballot) { v.AssessorLocation = make([]byte, 1) }},
		{"REDACTED", func(v *Ballot) { v.AssessorOrdinal = -1 }},
		{"REDACTED", func(v *Ballot) { v.Notation = nil }},
		{"REDACTED", func(v *Ballot) { v.Notation = make([]byte, MaximumSigningExtent+1) }},
	}
	for _, tc := range verifyScenarios {
		preballot := instancePreballot()
		attestBallot(t, privateItem, "REDACTED", preballot)
		tc.distortBallot(preballot)
		require.Error(t, preballot.CertifyFundamental(), "REDACTED", tc.alias)
		require.NoError(t, preballot.AssureAddition(), "REDACTED", tc.alias)

		preendorse := instancePreendorse()
		attestBallot(t, privateItem, "REDACTED", preendorse)
		tc.distortBallot(preendorse)
		require.Error(t, preendorse.CertifyFundamental(), "REDACTED", tc.alias)
		require.NoError(t, preendorse.AssureAddition(), "REDACTED", tc.alias)
	}
}

func VerifyUnfitPreballots(t *testing.T) {
	privateItem := FreshSimulatePRV()

	verifyScenarios := []struct {
		alias         string
		distortBallot func(*Ballot)
	}{
		{"REDACTED", func(v *Ballot) { v.Addition = []byte("REDACTED") }},
		{"REDACTED", func(v *Ballot) { v.AdditionNotation = []byte("REDACTED") }},
	}
	for _, tc := range verifyScenarios {
		preballot := instancePreballot()
		attestBallot(t, privateItem, "REDACTED", preballot)
		tc.distortBallot(preballot)
		require.Error(t, preballot.CertifyFundamental(), "REDACTED", tc.alias)
		require.NoError(t, preballot.AssureAddition(), "REDACTED", tc.alias)
	}
}

func VerifyUnfitPreendorseAdditions(t *testing.T) {
	privateItem := FreshSimulatePRV()

	verifyScenarios := []struct {
		alias         string
		distortBallot func(*Ballot)
	}{
		{"REDACTED", func(v *Ballot) {
			v.Addition = []byte("REDACTED")
			v.AdditionNotation = nil
		}},
		{"REDACTED", func(v *Ballot) { v.AdditionNotation = make([]byte, MaximumSigningExtent+1) }},
	}
	for _, tc := range verifyScenarios {
		preendorse := instancePreendorse()
		attestBallot(t, privateItem, "REDACTED", preendorse)
		tc.distortBallot(preendorse)
		//
		require.Error(t, preendorse.CertifyFundamental(), "REDACTED", tc.alias)
	}
}

func VerifyAssureBallotAddition(t *testing.T) {
	privateItem := FreshSimulatePRV()

	verifyScenarios := []struct {
		alias         string
		distortBallot func(*Ballot)
		anticipateFailure  bool
	}{
		{"REDACTED", func(v *Ballot) {
			v.Addition = nil
			v.AdditionNotation = nil
		}, true},
		{"REDACTED", func(v *Ballot) {
			v.AdditionNotation = []byte("REDACTED")
		}, false},
	}
	for _, tc := range verifyScenarios {
		preendorse := instancePreendorse()
		attestBallot(t, privateItem, "REDACTED", preendorse)
		tc.distortBallot(preendorse)
		if tc.anticipateFailure {
			require.Error(t, preendorse.AssureAddition(), "REDACTED", tc.alias)
		} else {
			require.NoError(t, preendorse.AssureAddition(), "REDACTED", tc.alias)
		}
	}
}

func VerifyBallotSchemaformat(t *testing.T) {
	privateItem := FreshSimulatePRV()
	ballot := instancePreendorse()
	v := ballot.TowardSchema()
	err := privateItem.AttestBallot("REDACTED", v)
	ballot.Notation = v.Notation
	require.NoError(t, err)

	verifyScenarios := []struct {
		msg                 string
		ballot                *Ballot
		transformsOkay          bool
		phasesCertifyFundamental bool
	}{
		{"REDACTED", ballot, true, true},
		{"REDACTED", &Ballot{}, true, false},
	}
	for _, tc := range verifyScenarios {
		schemaNomination := tc.ballot.TowardSchema()

		v, err := BallotOriginatingSchema(schemaNomination)
		if tc.transformsOkay {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}

		err = v.CertifyFundamental()
		if tc.phasesCertifyFundamental {
			require.NoError(t, err)
			require.Equal(t, tc.ballot, v, tc.msg)
		} else {
			require.Error(t, err)
		}
	}
}

func VerifyAttestAlsoInspectBallot(t *testing.T) {
	privateItem := FreshSimulatePRV()

	verifyScenarios := []struct {
		alias              string
		additionsActivated bool
		ballot              *Ballot
		anticipateFailure       bool
	}{
		{
			alias:              "REDACTED",
			additionsActivated: true,
			ballot:              instancePreendorse(),
			anticipateFailure:       false,
		},
		{
			alias:              "REDACTED",
			additionsActivated: false,
			ballot:              instancePreendorse(),
			anticipateFailure:       false,
		},
		{
			alias:              "REDACTED",
			additionsActivated: true,
			ballot: func() *Ballot {
				v := instancePreendorse()
				v.LedgerUUID = LedgerUUID{make([]byte, 0), FragmentAssignHeading{0, make([]byte, 0)}}
				return v
			}(),
			anticipateFailure: true,
		},
		{
			alias:              "REDACTED",
			additionsActivated: false,
			ballot: func() *Ballot {
				v := instancePreendorse()
				v.LedgerUUID = LedgerUUID{make([]byte, 0), FragmentAssignHeading{0, make([]byte, 0)}}
				return v
			}(),
			anticipateFailure: true,
		},
		{
			alias:              "REDACTED",
			additionsActivated: true,
			ballot: func() *Ballot {
				v := instancePreendorse()
				v.Addition = make([]byte, 0)
				return v
			}(),
			anticipateFailure: false,
		},
		{
			alias:              "REDACTED",
			additionsActivated: false,
			ballot: func() *Ballot {
				v := instancePreendorse()
				v.Addition = make([]byte, 0)
				return v
			}(),
			anticipateFailure: false,
		},
		{
			alias:              "REDACTED",
			additionsActivated: true,
			ballot:              instancePreballot(),
			anticipateFailure:       true,
		},
		{
			alias:              "REDACTED",
			additionsActivated: false,
			ballot:              instancePreballot(),
			anticipateFailure:       false,
		},
		{
			alias:              "REDACTED",
			additionsActivated: true,
			ballot: func() *Ballot {
				v := instancePreballot()
				v.Addition = []byte("REDACTED")
				return v
			}(),
			anticipateFailure: true,
		},
		{
			alias:              "REDACTED",
			additionsActivated: false,
			ballot: func() *Ballot {
				v := instancePreballot()
				v.Addition = []byte("REDACTED")
				return v
			}(),
			anticipateFailure: true,
		},
	}

	for _, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", tc.alias, tc.additionsActivated), func(t *testing.T) {
			_, err := AttestAlsoInspectBallot(tc.ballot, privateItem, "REDACTED", tc.additionsActivated)
			if tc.anticipateFailure {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
