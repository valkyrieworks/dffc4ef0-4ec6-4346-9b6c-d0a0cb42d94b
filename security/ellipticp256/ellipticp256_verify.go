package elliptic2_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/btcsuite/btcd/btcutil/base58"
	underlyingsecp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
)

type tokenData struct {
	private string
	pub  string
	location string
}

var ellipticDataRegistry = []tokenData{
	{
		private: "REDACTED",
		pub:  "REDACTED",
		location: "REDACTED",
	},
}

func VerifyPublicTokenElliptic256locator(t *testing.T) {
	for _, d := range ellipticDataRegistry {
		privateBYTE, _ := hex.DecodeString(d.private)
		publicBYTE, _ := hex.DecodeString(d.pub)
		locationBftzero, _, _ := base58.CheckDecode(d.location)
		locationBYTE := security.Location(locationBftzero)

		private := ellipticp256.PrivateToken(privateBYTE)

		publicToken := private.PublicToken()
		publicTYP, _ := publicToken.(ellipticp256.PublicToken)
		pub := publicTYP
		location := publicToken.Location()

		assert.Equal(t, pub, ellipticp256.PublicToken(publicBYTE), "REDACTED")
		assert.Equal(t, location, locationBYTE, "REDACTED")
	}
}

func VerifyAttestAlsoCertifyEllipticp256(t *testing.T) {
	privateToken := ellipticp256.ProducePrivateToken()
	publicToken := privateToken.PublicToken()

	msg := security.CHARArbitraryOctets(128)
	sig, err := privateToken.Attest(msg)
	require.Nil(t, err)

	assert.True(t, publicToken.ValidateNotation(msg, sig))

	//
	sig[3] ^= byte(0x01)

	assert.False(t, publicToken.ValidateNotation(msg, sig))
}

//
//
func VerifyElliptic256loadSecludedkeyAlsoMarshalEqualsCredential(t *testing.T) {
	numeralBelongingVerifies := 256
	for i := 0; i < numeralBelongingVerifies; i++ {
		//
		privateTokenOctets := [32]byte{}
		copy(privateTokenOctets[:], security.CHARArbitraryOctets(32))

		//
		//
		private := underlyingsecp256k1.PrivKeyFromBytes(privateTokenOctets[:])
		//
		//
		//
		//
		marshaledOctets := private.Serialize()
		require.Equal(t, privateTokenOctets[:], marshaledOctets)
	}
}

func VerifyProducePrivateTokenEllipticp256(t *testing.T) {
	//
	N := underlyingsecp256k1.S256().N
	verifies := []struct {
		alias   string
		credential []byte
	}{
		{"REDACTED", []byte{}},
		{
			"REDACTED",
			[]byte("REDACTED" +
				"REDACTED"),
		},
		{"REDACTED", []byte{0}},
		{"REDACTED", []byte("REDACTED")},
		{"REDACTED", []byte("REDACTED")},
	}
	for _, tt := range verifies {

		t.Run(tt.alias, func(t *testing.T) {
			attainedPrivateToken := ellipticp256.ProducePrivateTokenEllipticp256(tt.credential)
			require.NotNil(t, attainedPrivateToken)
			//
			fe := new(big.Int).SetBytes(attainedPrivateToken[:])
			require.True(t, fe.Cmp(N) < 0)
			require.True(t, fe.Sign() > 0)
		})
	}
}
