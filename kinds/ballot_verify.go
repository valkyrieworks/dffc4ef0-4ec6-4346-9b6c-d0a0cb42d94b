package kinds

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/utils/protoio"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

func instancePreballot() *Ballot {
	return instanceBallot(byte(engineproto.PreballotKind))
}

func instancePreendorse() *Ballot {
	ballot := instanceBallot(byte(engineproto.PreendorseKind))
	ballot.Addition = []byte("REDACTED")
	ballot.AdditionAutograph = []byte("REDACTED")
	return ballot
}

func instanceBallot(t byte) *Ballot {
	imprint, err := time.Parse(TimeLayout, "REDACTED")
	if err != nil {
		panic(err)
	}

	return &Ballot{
		Kind:      engineproto.AttestedMessageKind(t),
		Level:    12345,
		Cycle:     2,
		Timestamp: imprint,
		LedgerUID: LedgerUID{
			Digest: comethash.Sum([]byte("REDACTED")),
			SegmentAssignHeading: SegmentAssignHeading{
				Sum: 1000000,
				Digest:  comethash.Sum([]byte("REDACTED")),
			},
		},
		RatifierLocation: vault.LocationDigest([]byte("REDACTED")),
		RatifierOrdinal:   56789,
	}
}

func VerifyBallotSignable(t *testing.T) {
	ballot := instancePreendorse()
	v := ballot.ToSchema()
	attestOctets := BallotAttestOctets("REDACTED", v)
	pb := StandardizeBallot("REDACTED", v)
	anticipated, err := protoio.SerializeSeparated(&pb)
	require.NoError(t, err)

	require.Equal(t, anticipated, attestOctets, "REDACTED")
}

func VerifyBallotAttestOctetsVerifyArrays(t *testing.T) {
	verifies := []struct {
		ledgerUID string
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
			"REDACTED", &Ballot{Level: 1, Cycle: 1, Kind: engineproto.PreendorseKind},
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
			"REDACTED", &Ballot{Level: 1, Cycle: 1, Kind: engineproto.PreballotKind},
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
			"REDACTED", &Ballot{Level: 1, Cycle: 1},
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
			"REDACTED", &Ballot{Level: 1, Cycle: 1},
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
				Level:    1,
				Cycle:     1,
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
		v := tc.ballot.ToSchema()
		got := BallotAttestOctets(tc.ledgerUID, v)
		assert.Equal(t, len(tc.desire), len(got), "REDACTED", i)
		assert.Equal(t, tc.desire, got, "REDACTED", i)
	}
}

func VerifyBallotNominationNoEqual(t *testing.T) {
	cv := StandardizeBallot("REDACTED", &engineproto.Ballot{Level: 1, Cycle: 1})
	p := StandardizeNomination("REDACTED", &engineproto.Nomination{Level: 1, Cycle: 1})
	vb, err := proto.Marshal(&cv)
	require.NoError(t, err)
	pb, err := proto.Marshal(&p)
	require.NoError(t, err)
	require.NotEqual(t, vb, pb)
}

func VerifyBallotValidateAutograph(t *testing.T) {
	privateValue := NewEmulatePV()
	publickey, err := privateValue.FetchPublicKey()
	require.NoError(t, err)

	ballot := instancePreendorse()
	v := ballot.ToSchema()
	attestOctets := BallotAttestOctets("REDACTED", v)

	//
	err = privateValue.AttestBallot("REDACTED", v)
	require.NoError(t, err)

	//
	sound := publickey.ValidateAutograph(BallotAttestOctets("REDACTED", v), v.Autograph)
	require.True(t, sound)

	//
	preendorse := new(engineproto.Ballot)
	bs, err := proto.Marshal(v)
	require.NoError(t, err)
	err = proto.Unmarshal(bs, preendorse)
	require.NoError(t, err)

	//
	newAttestOctets := BallotAttestOctets("REDACTED", preendorse)
	require.Equal(t, string(attestOctets), string(newAttestOctets))
	sound = publickey.ValidateAutograph(newAttestOctets, preendorse.Autograph)
	require.True(t, sound)
}

//
//
func VerifyBallotAddition(t *testing.T) {
	verifyScenarios := []struct {
		label             string
		addition        []byte
		encompassAutograph bool
		anticipateFault      bool
	}{
		{
			label:             "REDACTED",
			addition:        []byte("REDACTED"),
			encompassAutograph: true,
			anticipateFault:      false,
		},
		{
			label:             "REDACTED",
			addition:        []byte("REDACTED"),
			encompassAutograph: false,
			anticipateFault:      true,
		},
		{
			label:             "REDACTED",
			encompassAutograph: true,
			anticipateFault:      false,
		},
		{
			label:             "REDACTED",
			encompassAutograph: false,
			anticipateFault:      true,
		},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.label, func(t *testing.T) {
			level, epoch := int64(1), int32(0)
			privateValue := NewEmulatePV()
			pk, err := privateValue.FetchPublicKey()
			require.NoError(t, err)
			ballot := &Ballot{
				RatifierLocation: pk.Location(),
				RatifierOrdinal:   0,
				Level:           level,
				Cycle:            epoch,
				Timestamp:        engineclock.Now(),
				Kind:             engineproto.PreendorseKind,
				LedgerUID:          createLedgerUIDArbitrary(),
			}

			v := ballot.ToSchema()
			err = privateValue.AttestBallot("REDACTED", v)
			require.NoError(t, err)
			ballot.Autograph = v.Autograph
			if tc.encompassAutograph {
				ballot.AdditionAutograph = v.AdditionAutograph
			}
			err = ballot.ValidateAddition("REDACTED", pk)
			if tc.anticipateFault {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func VerifyIsBallotKindSound(t *testing.T) {
	tc := []struct {
		label string
		in   engineproto.AttestedMessageKind
		out  bool
	}{
		{"REDACTED", engineproto.PreballotKind, true},
		{"REDACTED", engineproto.PreendorseKind, true},
		{"REDACTED", engineproto.AttestedMessageKind(0x3), false},
	}

	for _, tt := range tc {
		t.Run(tt.label, func(st *testing.T) {
			if rs := IsBallotKindSound(tt.in); rs != tt.out {
				t.Errorf("REDACTED", rs, tt.out)
			}
		})
	}
}

func VerifyBallotValidate(t *testing.T) {
	privateValue := NewEmulatePV()
	publickey, err := privateValue.FetchPublicKey()
	require.NoError(t, err)

	ballot := instancePreballot()
	ballot.RatifierLocation = publickey.Location()

	err = ballot.Validate("REDACTED", ed25519.GeneratePrivateKey().PublicKey())
	if assert.Error(t, err) {
		assert.Equal(t, ErrBallotCorruptRatifierLocation, err)
	}

	err = ballot.Validate("REDACTED", publickey)
	if assert.Error(t, err) {
		assert.Equal(t, ErrBallotCorruptAutograph, err)
	}
}

func VerifyBallotString(t *testing.T) {
	str := instancePreendorse().String()
	anticipated := "REDACTED" //
	if str != anticipated {
		t.Errorf("REDACTED", anticipated, str)
	}

	str2 := instancePreballot().String()
	anticipated = "REDACTED" //
	if str2 != anticipated {
		t.Errorf("REDACTED", anticipated, str2)
	}
}

func attestBallot(t *testing.T, pv PrivateRatifier, ledgerUID string, ballot *Ballot) {
	t.Helper()

	v := ballot.ToSchema()
	require.NoError(t, pv.AttestBallot(ledgerUID, v))
	ballot.Autograph = v.Autograph
	ballot.AdditionAutograph = v.AdditionAutograph
}

func VerifySoundBallots(t *testing.T) {
	privateValue := NewEmulatePV()

	verifyScenarios := []struct {
		label         string
		ballot         *Ballot
		distortBallot func(*Ballot)
	}{
		{"REDACTED", instancePreballot(), func(v *Ballot) {}},
		{"REDACTED", instancePreendorse(), func(v *Ballot) { v.Addition = nil }},
		{"REDACTED", instancePreendorse(), func(v *Ballot) { v.Addition = []byte("REDACTED") }},
	}
	for _, tc := range verifyScenarios {
		attestBallot(t, privateValue, "REDACTED", tc.ballot)
		tc.distortBallot(tc.ballot)
		require.NoError(t, tc.ballot.CertifySimple(), "REDACTED", tc.label)
		require.NoError(t, tc.ballot.AssureAddition(), "REDACTED", tc.label)
	}
}

func VerifyCorruptBallots(t *testing.T) {
	privateValue := NewEmulatePV()

	verifyScenarios := []struct {
		label         string
		distortBallot func(*Ballot)
	}{
		{"REDACTED", func(v *Ballot) { v.Level = -1 }},
		{"REDACTED", func(v *Ballot) { v.Cycle = -1 }},
		{"REDACTED", func(v *Ballot) { v.Level = 0 }},
		{"REDACTED", func(v *Ballot) { v.LedgerUID = LedgerUID{[]byte{1, 2, 3}, SegmentAssignHeading{111, []byte("REDACTED")}} }},
		{"REDACTED", func(v *Ballot) { v.RatifierLocation = make([]byte, 1) }},
		{"REDACTED", func(v *Ballot) { v.RatifierOrdinal = -1 }},
		{"REDACTED", func(v *Ballot) { v.Autograph = nil }},
		{"REDACTED", func(v *Ballot) { v.Autograph = make([]byte, MaximumAutographVolume+1) }},
	}
	for _, tc := range verifyScenarios {
		preballot := instancePreballot()
		attestBallot(t, privateValue, "REDACTED", preballot)
		tc.distortBallot(preballot)
		require.Error(t, preballot.CertifySimple(), "REDACTED", tc.label)
		require.NoError(t, preballot.AssureAddition(), "REDACTED", tc.label)

		preendorse := instancePreendorse()
		attestBallot(t, privateValue, "REDACTED", preendorse)
		tc.distortBallot(preendorse)
		require.Error(t, preendorse.CertifySimple(), "REDACTED", tc.label)
		require.NoError(t, preendorse.AssureAddition(), "REDACTED", tc.label)
	}
}

func VerifyCorruptPreballots(t *testing.T) {
	privateValue := NewEmulatePV()

	verifyScenarios := []struct {
		label         string
		distortBallot func(*Ballot)
	}{
		{"REDACTED", func(v *Ballot) { v.Addition = []byte("REDACTED") }},
		{"REDACTED", func(v *Ballot) { v.AdditionAutograph = []byte("REDACTED") }},
	}
	for _, tc := range verifyScenarios {
		preballot := instancePreballot()
		attestBallot(t, privateValue, "REDACTED", preballot)
		tc.distortBallot(preballot)
		require.Error(t, preballot.CertifySimple(), "REDACTED", tc.label)
		require.NoError(t, preballot.AssureAddition(), "REDACTED", tc.label)
	}
}

func VerifyCorruptPreendorsePlugins(t *testing.T) {
	privateValue := NewEmulatePV()

	verifyScenarios := []struct {
		label         string
		distortBallot func(*Ballot)
	}{
		{"REDACTED", func(v *Ballot) {
			v.Addition = []byte("REDACTED")
			v.AdditionAutograph = nil
		}},
		{"REDACTED", func(v *Ballot) { v.AdditionAutograph = make([]byte, MaximumAutographVolume+1) }},
	}
	for _, tc := range verifyScenarios {
		preendorse := instancePreendorse()
		attestBallot(t, privateValue, "REDACTED", preendorse)
		tc.distortBallot(preendorse)
		//
		require.Error(t, preendorse.CertifySimple(), "REDACTED", tc.label)
	}
}

func VerifyAssureBallotAddition(t *testing.T) {
	privateValue := NewEmulatePV()

	verifyScenarios := []struct {
		label         string
		distortBallot func(*Ballot)
		anticipateFault  bool
	}{
		{"REDACTED", func(v *Ballot) {
			v.Addition = nil
			v.AdditionAutograph = nil
		}, true},
		{"REDACTED", func(v *Ballot) {
			v.AdditionAutograph = []byte("REDACTED")
		}, false},
	}
	for _, tc := range verifyScenarios {
		preendorse := instancePreendorse()
		attestBallot(t, privateValue, "REDACTED", preendorse)
		tc.distortBallot(preendorse)
		if tc.anticipateFault {
			require.Error(t, preendorse.AssureAddition(), "REDACTED", tc.label)
		} else {
			require.NoError(t, preendorse.AssureAddition(), "REDACTED", tc.label)
		}
	}
}

func VerifyBallotProtobuf(t *testing.T) {
	privateValue := NewEmulatePV()
	ballot := instancePreendorse()
	v := ballot.ToSchema()
	err := privateValue.AttestBallot("REDACTED", v)
	ballot.Autograph = v.Autograph
	require.NoError(t, err)

	verifyScenarios := []struct {
		msg                 string
		ballot                *Ballot
		transformsOk          bool
		succeedsCertifySimple bool
	}{
		{"REDACTED", ballot, true, true},
		{"REDACTED", &Ballot{}, true, false},
	}
	for _, tc := range verifyScenarios {
		schemaNomination := tc.ballot.ToSchema()

		v, err := BallotFromSchema(schemaNomination)
		if tc.transformsOk {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}

		err = v.CertifySimple()
		if tc.succeedsCertifySimple {
			require.NoError(t, err)
			require.Equal(t, tc.ballot, v, tc.msg)
		} else {
			require.Error(t, err)
		}
	}
}

func VerifyAttestAndInspectBallot(t *testing.T) {
	privateValue := NewEmulatePV()

	verifyScenarios := []struct {
		label              string
		pluginsActivated bool
		ballot              *Ballot
		anticipateFault       bool
	}{
		{
			label:              "REDACTED",
			pluginsActivated: true,
			ballot:              instancePreendorse(),
			anticipateFault:       false,
		},
		{
			label:              "REDACTED",
			pluginsActivated: false,
			ballot:              instancePreendorse(),
			anticipateFault:       false,
		},
		{
			label:              "REDACTED",
			pluginsActivated: true,
			ballot: func() *Ballot {
				v := instancePreendorse()
				v.LedgerUID = LedgerUID{make([]byte, 0), SegmentAssignHeading{0, make([]byte, 0)}}
				return v
			}(),
			anticipateFault: true,
		},
		{
			label:              "REDACTED",
			pluginsActivated: false,
			ballot: func() *Ballot {
				v := instancePreendorse()
				v.LedgerUID = LedgerUID{make([]byte, 0), SegmentAssignHeading{0, make([]byte, 0)}}
				return v
			}(),
			anticipateFault: true,
		},
		{
			label:              "REDACTED",
			pluginsActivated: true,
			ballot: func() *Ballot {
				v := instancePreendorse()
				v.Addition = make([]byte, 0)
				return v
			}(),
			anticipateFault: false,
		},
		{
			label:              "REDACTED",
			pluginsActivated: false,
			ballot: func() *Ballot {
				v := instancePreendorse()
				v.Addition = make([]byte, 0)
				return v
			}(),
			anticipateFault: false,
		},
		{
			label:              "REDACTED",
			pluginsActivated: true,
			ballot:              instancePreballot(),
			anticipateFault:       true,
		},
		{
			label:              "REDACTED",
			pluginsActivated: false,
			ballot:              instancePreballot(),
			anticipateFault:       false,
		},
		{
			label:              "REDACTED",
			pluginsActivated: true,
			ballot: func() *Ballot {
				v := instancePreballot()
				v.Addition = []byte("REDACTED")
				return v
			}(),
			anticipateFault: true,
		},
		{
			label:              "REDACTED",
			pluginsActivated: false,
			ballot: func() *Ballot {
				v := instancePreballot()
				v.Addition = []byte("REDACTED")
				return v
			}(),
			anticipateFault: true,
		},
	}

	for _, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", tc.label, tc.pluginsActivated), func(t *testing.T) {
			_, err := AttestAndInspectBallot(tc.ballot, privateValue, "REDACTED", tc.pluginsActivated)
			if tc.anticipateFault {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
