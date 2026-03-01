package hashmap

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	cmtsecurity "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
)

const AttestationActionCascade = "REDACTED"

//
//
type CascadeAction struct {
	key    string //
	Influx  string
	Emission string
}

func FreshCascadeAction(key, influx, emission string) CascadeAction {
	return CascadeAction{
		key:    key,
		Influx:  influx,
		Emission: emission,
	}
}

func (dop CascadeAction) AttestationAction() cmtsecurity.AttestationAction {
	doschema := cmtsecurity.CascadeAction{
		Key:    dop.key,
		Influx:  dop.Influx,
		Emission: dop.Emission,
	}
	bz, err := doschema.Serialize()
	if err != nil {
		panic(err)
	}

	return cmtsecurity.AttestationAction{
		Kind: AttestationActionCascade,
		Key:  []byte(dop.key),
		Data: bz,
	}
}

func (dop CascadeAction) Run(influx [][]byte) (emission [][]byte, err error) {
	if len(influx) != 1 {
		return nil, errors.New("REDACTED")
	}
	if string(influx[0]) != dop.Influx {
		return nil, fmt.Errorf("REDACTED",
			dop.Influx, string(influx[0]))
	}
	return [][]byte{[]byte(dop.Emission)}, nil
}

func (dop CascadeAction) ObtainToken() []byte {
	return []byte(dop.key)
}

//

func VerifyAttestationHandlers(t *testing.T) {
	var err error

	//
	//

	//
	op1 := FreshCascadeAction("REDACTED", "REDACTED", "REDACTED")
	op2 := FreshCascadeAction("REDACTED", "REDACTED", "REDACTED")
	op3 := FreshCascadeAction("REDACTED", "REDACTED", "REDACTED")
	op4 := FreshCascadeAction("REDACTED", "REDACTED", "REDACTED")

	//
	removals := AttestationHandlers([]AttestationHandler{op1, op2, op3, op4})
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.Nil(t, err)
	err = removals.ValidateDatum(bz("REDACTED"), "REDACTED", bz("REDACTED"))
	assert.Nil(t, err)

	//
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)
	err = removals.ValidateDatum(bz("REDACTED"), "REDACTED", bz("REDACTED"))
	assert.NotNil(t, err)

	//
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	removals = []AttestationHandler{op1, op2, op4}
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	removals = []AttestationHandler{op4, op3, op2, op1}
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	removals = []AttestationHandler{}
	err = removals.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)
}

func bz(s string) []byte {
	return []byte(s)
}

func VerifyAttestationCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias      string
		distortAttestation func(*Attestation)
		faultTxt        string
	}{
		{"REDACTED", func(sp *Attestation) {}, "REDACTED"},
		{"REDACTED", func(sp *Attestation) { sp.Sum = -1 }, "REDACTED"},
		{"REDACTED", func(sp *Attestation) { sp.Ordinal = -1 }, "REDACTED"},
		{
			"REDACTED", func(sp *Attestation) { sp.NodeDigest = make([]byte, 10) },
			"REDACTED",
		},
		{
			"REDACTED", func(sp *Attestation) { sp.Kin = make([][]byte, MaximumKin+1) },
			"REDACTED",
		},
		{
			"REDACTED", func(sp *Attestation) { sp.Kin[0] = make([]byte, 10) },
			"REDACTED",
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.verifyAlias, func(t *testing.T) {
			_, attestations := AttestationsOriginatingOctetSegments([][]byte{
				[]byte("REDACTED"),
				[]byte("REDACTED"),
				[]byte("REDACTED"),
			})
			tc.distortAttestation(attestations[0])
			err := attestations[0].CertifyFundamental()
			if tc.faultTxt != "REDACTED" {
				assert.Contains(t, err.Error(), tc.faultTxt)
			}
		})
	}
}

func VerifyBallotSchemaformat(t *testing.T) {
	_, attestations := AttestationsOriginatingOctetSegments([][]byte{
		[]byte("REDACTED"),
		[]byte("REDACTED"),
		[]byte("REDACTED"),
	})

	verifyScenarios := []struct {
		verifyAlias string
		v1       *Attestation
		expirationPhrase  bool
	}{
		{"REDACTED", &Attestation{}, false},
		{"REDACTED", nil, false},
		{"REDACTED", attestations[0], true},
	}
	for _, tc := range verifyScenarios {
		pb := tc.v1.TowardSchema()

		v, err := AttestationOriginatingSchema(pb)
		if tc.expirationPhrase {
			require.NoError(t, err)
			require.Equal(t, tc.v1, v, tc.verifyAlias)
		} else {
			require.Error(t, err)
		}
	}
}

//
func Verifyvsa2022_hundred(t *testing.T) {
	//
	key := []byte{0x13}
	datum := []byte{0x37}
	datadigest := tenderminthash.Sum(datum)
	bz := new(bytes.Buffer)
	_ = serializeOctetSegment(bz, key)
	_ = serializeOctetSegment(bz, datadigest)
	mapdigest := tenderminthash.Sum(append([]byte{0}, bz.Bytes()...))

	//
	op := FreshDatumAction(
		key,
		&Attestation{NodeDigest: mapdigest},
	)

	//
	var origin []byte

	assert.NotNil(t, AttestationHandlers{op}.Validate(origin, "REDACTED"+string(key), [][]byte{datum}))
}
