//

package signature381

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"

	blst "github.com/supranational/blst/bindings/go"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
)

const (
	//
	Enabled = true
)

var (
	//
	ErrDeserialization = errors.New("REDACTED")
	//
	//
	ErrInfinitePubKey = errors.New("REDACTED")

	dstMinPk = []byte("REDACTED")
)

//
//
//
//
type (
	blstPublicKey          = blst.P1Affine
	blstSignature          = blst.P2Affine
	blstAggregateSignature = blst.P1Aggregate
	blstAggregatePublicKey = blst.P2Aggregate
)

//

func init() {
	strongmindjson.RegisterType(PubKey{}, PubKeyName)
	strongmindjson.RegisterType(PrivKey{}, PrivKeyName)
}

//
//
//

//
//
//

var _ security.PrivKey = &PrivKey{}

type PrivKey struct {
	sk *blst.SecretKey
}

//
func GenPrivKeyFromSecret(secret []byte) (*PrivKey, error) {
	if len(secret) != 32 {
		seed := sha256.Sum256(secret) //
		secret = seed[:]
	}

	sk := blst.KeyGen(secret)
	return &PrivKey{sk: sk}, nil
}

//
func NewPrivateKeyFromBytes(bz []byte) (*PrivKey, error) {
	sk := new(blst.SecretKey).Deserialize(bz)
	if sk == nil {
		return nil, ErrDeserialization
	}
	return &PrivKey{sk: sk}, nil
}

//
func GenPrivKey() (*PrivKey, error) {
	var ikm [32]byte
	_, err := rand.Read(ikm[:])
	if err != nil {
		return nil, err
	}
	return GenPrivKeyFromSecret(ikm[:])
}

//
func (privKey PrivKey) Bytes() []byte {
	return privKey.sk.Serialize()
}

//
//
func (privKey PrivKey) PubKey() security.PubKey {
	return PubKey{pk: new(blstPublicKey).From(privKey.sk)}
}

//
func (privKey PrivKey) Equals(other security.PrivKey) bool {
	return privKey.Type() == other.Type() && bytes.Equal(privKey.Bytes(), other.Bytes())
}

//
func (PrivKey) Type() string {
	return KeyType
}

//
func (privKey PrivKey) Sign(msg []byte) ([]byte, error) {
	signature := new(blstSignature).Sign(privKey.sk, msg, dstMinPk)
	return signature.Compress(), nil
}

//
func (privKey *PrivKey) Zeroize() {
	privKey.sk.Zeroize()
}

//
//
//
//
func (privKey PrivKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(privKey.Bytes())
}

//
func (privKey *PrivKey) UnmarshalJSON(bz []byte) error {
	var rawBytes []byte
	if err := json.Unmarshal(bz, &rawBytes); err != nil {
		return err
	}
	pk, err := NewPrivateKeyFromBytes(rawBytes)
	if err != nil {
		return err
	}
	privKey.sk = pk.sk
	return nil
}

//
//
//

//
//
//

var _ security.PubKey = &PubKey{}

type PubKey struct {
	pk *blstPublicKey
}

//
func NewPublicKeyFromBytes(bz []byte) (*PubKey, error) {
	pk := new(blstPublicKey).Deserialize(bz)
	if pk == nil {
		return nil, ErrDeserialization
	}
	//
	if !pk.KeyValidate() {
		return nil, ErrInfinitePubKey
	}
	return &PubKey{pk: pk}, nil
}

//
//
//
func (pubKey PubKey) Address() security.Address {
	return security.Address(tenderminthash.SumTruncated(pubKey.pk.Serialize()))
}

//
func (pubKey PubKey) VerifySignature(msg, sig []byte) bool {
	signature := new(blstSignature).Uncompress(sig)
	if signature == nil {
		return false
	}

	//
	//
	if !signature.SigValidate(false) {
		return false
	}

	return signature.Verify(false, pubKey.pk, false, msg, dstMinPk)
}

//
func (pubKey PubKey) Bytes() []byte {
	return pubKey.pk.Serialize()
}

//
func (PubKey) Type() string {
	return KeyType
}

//
func (pubKey PubKey) Equals(other security.PubKey) bool {
	return pubKey.Type() == other.Type() && bytes.Equal(pubKey.Bytes(), other.Bytes())
}

//
//
//
//
func (pubkey PubKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(pubkey.Bytes())
}

//
func (pubkey *PubKey) UnmarshalJSON(bz []byte) error {
	var rawBytes []byte
	if err := json.Unmarshal(bz, &rawBytes); err != nil {
		return err
	}
	pk, err := NewPublicKeyFromBytes(rawBytes)
	if err != nil {
		return err
	}
	pubkey.pk = pk.pk
	return nil
}
