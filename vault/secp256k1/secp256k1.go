package secp256k1

import (
	"bytes"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"io"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"golang.org/x/crypto/ripemd160" //

	"github.com/valkyrieworks/vault"
	cometjson "github.com/valkyrieworks/utils/json"
)

//
const (
	PrivateKeyLabel = "REDACTED"
	PublicKeyLabel  = "REDACTED"

	KeyKind     = "REDACTED"
	PrivateKeyVolume = 32
)

func init() {
	cometjson.EnrollKind(PublicKey{}, PublicKeyLabel)
	cometjson.EnrollKind(PrivateKey{}, PrivateKeyLabel)
}

var _ vault.PrivateKey = PrivateKey{}

//
type PrivateKey []byte

//
func (privateKey PrivateKey) Octets() []byte {
	return []byte(privateKey)
}

//
//
func (privateKey PrivateKey) PublicKey() vault.PublicKey {
	secpPrivateKey := secp256k1.PrivKeyFromBytes(privateKey)

	pk := secpPrivateKey.PubKey().SerializeCompressed()

	return PublicKey(pk)
}

//
//
func (privateKey PrivateKey) Matches(another vault.PrivateKey) bool {
	if anotherSecp, ok := another.(PrivateKey); ok {
		return subtle.ConstantTimeCompare(privateKey[:], anotherSecp[:]) == 1
	}
	return false
}

func (privateKey PrivateKey) Kind() string {
	return KeyKind
}

//
//
func GeneratePrivateKey() PrivateKey {
	return generatePrivateKey(vault.CScanner())
}

//
func generatePrivateKey(random io.Reader) PrivateKey {
	var privateKeyOctets [PrivateKeyVolume]byte
	d := new(big.Int)

	for {
		privateKeyOctets = [PrivateKeyVolume]byte{}
		_, err := io.ReadFull(random, privateKeyOctets[:])
		if err != nil {
			panic(err)
		}

		d.SetBytes(privateKeyOctets[:])
		//
		isSoundFieldComponent := 0 < d.Sign() && d.Cmp(secp256k1.S256().N) < 0
		if isSoundFieldComponent {
			break
		}
	}

	return PrivateKey(privateKeyOctets[:])
}

var one = new(big.Int).SetInt64(1)

//
//
//
//
//
//
//
//
//
//
func GeneratePrivateKeySecp256k1(key []byte) PrivateKey {
	securityDigest := sha256.Sum256(key)
	//
	//
	//
	//
	fe := new(big.Int).SetBytes(securityDigest[:])
	n := new(big.Int).Sub(secp256k1.S256().N, one)
	fe.Mod(fe, n)
	fe.Add(fe, one)

	feB := fe.Bytes()
	privateKey32 := make([]byte, PrivateKeyVolume)
	//
	copy(privateKey32[32-len(feB):32], feB)

	return PrivateKey(privateKey32)
}

//
//
func (privateKey PrivateKey) Attest(msg []byte) ([]byte, error) {
	private := secp256k1.PrivKeyFromBytes(privateKey)

	sum := sha256.Sum256(msg)
	sig := ecdsa.SignCompact(private, sum[:], false)

	//
	return sig[1:], nil
}

//

var _ vault.PublicKey = PublicKey{}

//
//
const PublicKeyVolume = 33

//
//
//
//
//
type PublicKey []byte

//
func (publicKey PublicKey) Location() vault.Location {
	if len(publicKey) != PublicKeyVolume {
		panic("REDACTED")
	}
	digesterSha256 := sha256.New()
	_, _ = digesterSha256.Write(publicKey) //
	sha := digesterSha256.Sum(nil)

	digesterRipemd160 := ripemd160.New()
	_, _ = digesterRipemd160.Write(sha) //

	return vault.Location(digesterRipemd160.Sum(nil))
}

//
func (publicKey PublicKey) Octets() []byte {
	return []byte(publicKey)
}

func (publicKey PublicKey) String() string {
	return fmt.Sprintf("REDACTED", []byte(publicKey))
}

func (publicKey PublicKey) Matches(another vault.PublicKey) bool {
	if anotherSecp, ok := another.(PublicKey); ok {
		return bytes.Equal(publicKey[:], anotherSecp[:])
	}
	return false
}

func (publicKey PublicKey) Kind() string {
	return KeyKind
}

//
//
func (publicKey PublicKey) ValidateAutograph(msg []byte, signatureStr []byte) bool {
	if len(signatureStr) != 64 {
		return false
	}

	pub, err := secp256k1.ParsePubKey(publicKey)
	if err != nil {
		return false
	}

	//
	autograph := autographFromOctets(signatureStr)
	//
	//
	//
	//
	alteredAutograph, analyzeErr := ecdsa.ParseDERSignature(autograph.Serialize())
	if analyzeErr != nil {
		return false
	}
	if !autograph.IsEqual(alteredAutograph) {
		return false
	}

	return autograph.Verify(vault.Sha256(msg), pub)
}

//
//
func autographFromOctets(signatureStr []byte) *ecdsa.Signature {
	var r secp256k1.ModNScalar
	r.SetByteSlice(signatureStr[:32])
	var s secp256k1.ModNScalar
	s.SetByteSlice(signatureStr[32:64])
	return ecdsa.NewSignature(&r, &s)
}
