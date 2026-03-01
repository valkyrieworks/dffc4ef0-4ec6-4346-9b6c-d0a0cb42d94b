package ellipticp256

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func Verify_produceprivatekey(t *testing.T) {
	blank := make([]byte, 32)
	singleBYTE := big.NewInt(1).Bytes()
	singleStuffed := make([]byte, 32)
	copy(singleStuffed[32-len(singleBYTE):32], singleBYTE)
	t.Logf("REDACTED", singleStuffed, len(singleStuffed))

	soundSingle := bytes.Join([][]byte{blank, singleStuffed}, nil)
	verifies := []struct {
		alias        string
		negationThereforeArbitrary   []byte
		mustAlarm bool
	}{
		{"REDACTED", blank, true},
		{"REDACTED", secp256k1.S256().N.Bytes(), true},
		{"REDACTED", soundSingle, false},
	}
	for _, tt := range verifies {

		t.Run(tt.alias, func(t *testing.T) {
			if tt.mustAlarm {
				require.Panics(t, func() {
					producePrivateToken(bytes.NewReader(tt.negationThereforeArbitrary))
				})
				return
			}
			got := producePrivateToken(bytes.NewReader(tt.negationThereforeArbitrary))
			fe := new(big.Int).SetBytes(got[:])
			require.True(t, fe.Cmp(secp256k1.S256().N) < 0)
			require.True(t, fe.Sign() > 0)
		})
	}
}

//
//
//
func VerifySigningValidationAlsoDeclineHigherSTR(t *testing.T) {
	msg := []byte("REDACTED")
	for i := 0; i < 500; i++ {
		private := ProducePrivateToken()
		signatureTxt, err := private.Attest(msg)
		require.NoError(t, err)
		var r secp256k1.ModNScalar
		r.SetByteSlice(signatureTxt[:32])
		var s secp256k1.ModNScalar
		s.SetByteSlice(signatureTxt[32:64])
		require.False(t, s.IsOverHalfOrder())

		pub := private.PublicToken()
		require.True(t, pub.ValidateNotation(msg, signatureTxt))

		//
		var Curve256 secp256k1.ModNScalar
		Curve256.SetByteSlice(secp256k1.S256().N.Bytes())
		s.Negate().Add(&Curve256)
		require.True(t, s.IsOverHalfOrder())

		readerOctets := r.Bytes()
		strOctets := s.Bytes()
		harmSignatureTxt := make([]byte, 64)
		copy(harmSignatureTxt[32-len(readerOctets):32], readerOctets[:])
		copy(harmSignatureTxt[64-len(strOctets):64], strOctets[:])

		require.False(t, pub.ValidateNotation(msg, harmSignatureTxt),
			"REDACTED",
			harmSignatureTxt,
			private,
		)
	}
}
