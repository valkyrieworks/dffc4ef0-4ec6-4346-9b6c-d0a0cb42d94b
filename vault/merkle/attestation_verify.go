package merkle

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/comethash"
	cmtcrypto "github.com/valkyrieworks/schema/consensuscore/vault"
)

const EvidenceActDomino = "REDACTED"

//
//
type DominoAct struct {
	key    string //
	Influx  string
	Result string
}

func NewDominoAct(key, influx, result string) DominoAct {
	return DominoAct{
		key:    key,
		Influx:  influx,
		Result: result,
	}
}

func (dop DominoAct) EvidenceAct() cmtcrypto.EvidenceAct {
	dopb := cmtcrypto.DominoAct{
		Key:    dop.key,
		Influx:  dop.Influx,
		Result: dop.Result,
	}
	bz, err := dopb.Serialize()
	if err != nil {
		panic(err)
	}

	return cmtcrypto.EvidenceAct{
		Kind: EvidenceActDomino,
		Key:  []byte(dop.key),
		Data: bz,
	}
}

func (dop DominoAct) Run(influx [][]byte) (result [][]byte, err error) {
	if len(influx) != 1 {
		return nil, errors.New("REDACTED")
	}
	if string(influx[0]) != dop.Influx {
		return nil, fmt.Errorf("REDACTED",
			dop.Influx, string(influx[0]))
	}
	return [][]byte{[]byte(dop.Result)}, nil
}

func (dop DominoAct) FetchKey() []byte {
	return []byte(dop.key)
}

//

func VerifyEvidenceHandlers(t *testing.T) {
	var err error

	//
	//

	//
	op1 := NewDominoAct("REDACTED", "REDACTED", "REDACTED")
	op2 := NewDominoAct("REDACTED", "REDACTED", "REDACTED")
	op3 := NewDominoAct("REDACTED", "REDACTED", "REDACTED")
	op4 := NewDominoAct("REDACTED", "REDACTED", "REDACTED")

	//
	popz := EvidenceHandlers([]EvidenceHandler{op1, op2, op3, op4})
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.Nil(t, err)
	err = popz.ValidateItem(bz("REDACTED"), "REDACTED", bz("REDACTED"))
	assert.Nil(t, err)

	//
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)
	err = popz.ValidateItem(bz("REDACTED"), "REDACTED", bz("REDACTED"))
	assert.NotNil(t, err)

	//
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	popz = []EvidenceHandler{op1, op2, op4}
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	popz = []EvidenceHandler{op4, op3, op2, op1}
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)

	//
	popz = []EvidenceHandler{}
	err = popz.Validate(bz("REDACTED"), "REDACTED", [][]byte{bz("REDACTED")})
	assert.NotNil(t, err)
}

func bz(s string) []byte {
	return []byte(s)
}

func VerifyEvidenceCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel      string
		distortEvidence func(*Attestation)
		errStr        string
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

		t.Run(tc.verifyLabel, func(t *testing.T) {
			_, evidences := EvidencesFromOctetSegments([][]byte{
				[]byte("REDACTED"),
				[]byte("REDACTED"),
				[]byte("REDACTED"),
			})
			tc.distortEvidence(evidences[0])
			err := evidences[0].CertifySimple()
			if tc.errStr != "REDACTED" {
				assert.Contains(t, err.Error(), tc.errStr)
			}
		})
	}
}

func VerifyBallotProtobuf(t *testing.T) {
	_, evidences := EvidencesFromOctetSegments([][]byte{
		[]byte("REDACTED"),
		[]byte("REDACTED"),
		[]byte("REDACTED"),
	})

	verifyScenarios := []struct {
		verifyLabel string
		v1       *Attestation
		expirationPass  bool
	}{
		{"REDACTED", &Attestation{}, false},
		{"REDACTED", nil, false},
		{"REDACTED", evidences[0], true},
	}
	for _, tc := range verifyScenarios {
		pb := tc.v1.ToSchema()

		v, err := EvidenceFromSchema(pb)
		if tc.expirationPass {
			require.NoError(t, err)
			require.Equal(t, tc.v1, v, tc.verifyLabel)
		} else {
			require.Error(t, err)
		}
	}
}

//
func Verifyvsa2022_100(t *testing.T) {
	//
	key := []byte{0x13}
	item := []byte{0x37}
	vdigest := comethash.Sum(item)
	bz := new(bytes.Buffer)
	_ = encodeOctetSegment(bz, key)
	_ = encodeOctetSegment(bz, vdigest)
	kvdigest := comethash.Sum(append([]byte{0}, bz.Bytes()...))

	//
	op := NewItemAct(
		key,
		&Attestation{NodeDigest: kvdigest},
	)

	//
	var origin []byte

	assert.NotNil(t, EvidenceHandlers{op}.Validate(origin, "REDACTED"+string(key), [][]byte{item}))
}
