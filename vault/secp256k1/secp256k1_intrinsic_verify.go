package secp256k1

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func Verify_generateprivatekey(t *testing.T) {
	empty := make([]byte, 32)
	oneBYTE := big.NewInt(1).Bytes()
	oneCushioned := make([]byte, 32)
	copy(oneCushioned[32-len(oneBYTE):32], oneBYTE)
	t.Logf("REDACTED", oneCushioned, len(oneCushioned))

	soundOne := bytes.Join([][]byte{empty, oneCushioned}, nil)
	verifies := []struct {
		label        string
		noSoRandom   []byte
		mustAlarm bool
	}{
		{"REDACTED", empty, true},
		{"REDACTED", secp256k1.S256().N.Bytes(), true},
		{"REDACTED", soundOne, false},
	}
	for _, tt := range verifies {

		t.Run(tt.label, func(t *testing.T) {
			if tt.mustAlarm {
				require.Panics(t, func() {
					generatePrivateKey(bytes.NewReader(tt.noSoRandom))
				})
				return
			}
			got := generatePrivateKey(bytes.NewReader(tt.noSoRandom))
			fe := new(big.Int).SetBytes(got[:])
			require.True(t, fe.Cmp(secp256k1.S256().N) < 0)
			require.True(t, fe.Sign() > 0)
		})
	}
}

//
//
//
func VerifyAutographValidationAndDeclineUpperS(t *testing.T) {
	msg := []byte("REDACTED")
	for i := 0; i < 500; i++ {
		private := GeneratePrivateKey()
		signatureStr, err := private.Attest(msg)
		require.NoError(t, err)
		var r secp256k1.ModNScalar
		r.SetByteSlice(signatureStr[:32])
		var s secp256k1.ModNScalar
		s.SetByteSlice(signatureStr[32:64])
		require.False(t, s.IsOverHalfOrder())

		pub := private.PublicKey()
		require.True(t, pub.ValidateAutograph(msg, signatureStr))

		//
		var S256 secp256k1.ModNScalar
		S256.SetByteSlice(secp256k1.S256().N.Bytes())
		s.Negate().Add(&S256)
		require.True(t, s.IsOverHalfOrder())

		readerOctets := r.Bytes()
		sOctets := s.Bytes()
		maliciousSignatureStr := make([]byte, 64)
		copy(maliciousSignatureStr[32-len(readerOctets):32], readerOctets[:])
		copy(maliciousSignatureStr[64-len(sOctets):64], sOctets[:])

		require.False(t, pub.ValidateAutograph(msg, maliciousSignatureStr),
			"REDACTED",
			maliciousSignatureStr,
			private,
		)
	}
}
