package secp2_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/btcsuite/btcd/btcutil/base58"
	underlyingsecp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/secp256k1"
)

type keyData struct {
	private string
	pub  string
	address string
}

var secpDataSheet = []keyData{
	{
		private: "REDACTED",
		pub:  "REDACTED",
		address: "REDACTED",
	},
}

func VerifyPublicKeySecp256k1location(t *testing.T) {
	for _, d := range secpDataSheet {
		privateBYTE, _ := hex.DecodeString(d.private)
		publicBYTE, _ := hex.DecodeString(d.pub)
		addressBbz, _, _ := base58.CheckDecode(d.address)
		addressBYTE := vault.Location(addressBbz)

		private := secp256k1.PrivateKey(privateBYTE)

		publicKey := private.PublicKey()
		publicT, _ := publicKey.(secp256k1.PublicKey)
		pub := publicT
		address := publicKey.Location()

		assert.Equal(t, pub, secp256k1.PublicKey(publicBYTE), "REDACTED")
		assert.Equal(t, address, addressBYTE, "REDACTED")
	}
}

func VerifyAttestAndCertifySecp256k1(t *testing.T) {
	privateKey := secp256k1.GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	msg := vault.CRandomOctets(128)
	sig, err := privateKey.Attest(msg)
	require.Nil(t, err)

	assert.True(t, publicKey.ValidateAutograph(msg, sig))

	//
	sig[3] ^= byte(0x01)

	assert.False(t, publicKey.ValidateAutograph(msg, sig))
}

//
//
func VerifySecp256k1loadPrivatekeyAndMarshalIsCredential(t *testing.T) {
	countOfVerifies := 256
	for i := 0; i < countOfVerifies; i++ {
		//
		privateKeyOctets := [32]byte{}
		copy(privateKeyOctets[:], vault.CRandomOctets(32))

		//
		//
		private := underlyingsecp256k1.PrivKeyFromBytes(privateKeyOctets[:])
		//
		//
		//
		//
		marshaledOctets := private.Serialize()
		require.Equal(t, privateKeyOctets[:], marshaledOctets)
	}
}

func VerifyGeneratePrivateKeySecp256k1(t *testing.T) {
	//
	N := underlyingsecp256k1.S256().N
	verifies := []struct {
		label   string
		key []byte
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

		t.Run(tt.label, func(t *testing.T) {
			acquiredPrivateKey := secp256k1.GeneratePrivateKeySecp256k1(tt.key)
			require.NotNil(t, acquiredPrivateKey)
			//
			fe := new(big.Int).SetBytes(acquiredPrivateKey[:])
			require.True(t, fe.Cmp(N) < 0)
			require.True(t, fe.Sign() > 0)
		})
	}
}
