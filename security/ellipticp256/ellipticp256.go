package ellipticp256

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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
)

//
const (
	PrivateTokenAlias = "REDACTED"
	PublicTokenAlias  = "REDACTED"

	TokenKind     = "REDACTED"
	PrivateTokenExtent = 32
)

func initialize() {
	strongmindjson.EnrollKind(PublicToken{}, PublicTokenAlias)
	strongmindjson.EnrollKind(PrivateToken{}, PrivateTokenAlias)
}

var _ security.PrivateToken = PrivateToken{}

//
type PrivateToken []byte

//
func (privateToken PrivateToken) Octets() []byte {
	return []byte(privateToken)
}

//
//
func (privateToken PrivateToken) PublicToken() security.PublicToken {
	ellipticPrivateToken := secp256k1.PrivKeyFromBytes(privateToken)

	pk := ellipticPrivateToken.PubKey().SerializeCompressed()

	return PublicToken(pk)
}

//
//
func (privateToken PrivateToken) Matches(another security.PrivateToken) bool {
	if anotherElliptic, ok := another.(PrivateToken); ok {
		return subtle.ConstantTimeCompare(privateToken[:], anotherElliptic[:]) == 1
	}
	return false
}

func (privateToken PrivateToken) Kind() string {
	return TokenKind
}

//
//
func ProducePrivateToken() PrivateToken {
	return producePrivateToken(security.CHARFetcher())
}

//
func producePrivateToken(arbitrary io.Reader) PrivateToken {
	var privateTokenOctets [PrivateTokenExtent]byte
	d := new(big.Int)

	for {
		privateTokenOctets = [PrivateTokenExtent]byte{}
		_, err := io.ReadFull(arbitrary, privateTokenOctets[:])
		if err != nil {
			panic(err)
		}

		d.SetBytes(privateTokenOctets[:])
		//
		equalsSoundAttributeComponent := 0 < d.Sign() && d.Cmp(secp256k1.S256().N) < 0
		if equalsSoundAttributeComponent {
			break
		}
	}

	return PrivateToken(privateTokenOctets[:])
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
func ProducePrivateTokenEllipticp256(credential []byte) PrivateToken {
	secondDigest := sha256.Sum256(credential)
	//
	//
	//
	//
	fe := new(big.Int).SetBytes(secondDigest[:])
	n := new(big.Int).Sub(secp256k1.S256().N, one)
	fe.Mod(fe, n)
	fe.Add(fe, one)

	feB := fe.Bytes()
	privateKeypair32 := make([]byte, PrivateTokenExtent)
	//
	copy(privateKeypair32[32-len(feB):32], feB)

	return PrivateToken(privateKeypair32)
}

//
//
func (privateToken PrivateToken) Attest(msg []byte) ([]byte, error) {
	private := secp256k1.PrivKeyFromBytes(privateToken)

	sum := sha256.Sum256(msg)
	sig := ecdsa.SignCompact(private, sum[:], false)

	//
	return sig[1:], nil
}

//

var _ security.PublicToken = PublicToken{}

//
//
const PublicTokenExtent = 33

//
//
//
//
//
type PublicToken []byte

//
func (publicToken PublicToken) Location() security.Location {
	if len(publicToken) != PublicTokenExtent {
		panic("REDACTED")
	}
	digesterHash256 := sha256.New()
	_, _ = digesterHash256.Write(publicToken) //
	sha := digesterHash256.Sum(nil)

	digesterHashmd160 := ripemd160.New()
	_, _ = digesterHashmd160.Write(sha) //

	return security.Location(digesterHashmd160.Sum(nil))
}

//
func (publicToken PublicToken) Octets() []byte {
	return []byte(publicToken)
}

func (publicToken PublicToken) Text() string {
	return fmt.Sprintf("REDACTED", []byte(publicToken))
}

func (publicToken PublicToken) Matches(another security.PublicToken) bool {
	if anotherElliptic, ok := another.(PublicToken); ok {
		return bytes.Equal(publicToken[:], anotherElliptic[:])
	}
	return false
}

func (publicToken PublicToken) Kind() string {
	return TokenKind
}

//
//
func (publicToken PublicToken) ValidateNotation(msg []byte, signatureTxt []byte) bool {
	if len(signatureTxt) != 64 {
		return false
	}

	pub, err := secp256k1.ParsePubKey(publicToken)
	if err != nil {
		return false
	}

	//
	signing := signingOriginatingOctets(signatureTxt)
	//
	//
	//
	//
	adjustedSigning, analyzeFault := ecdsa.ParseDERSignature(signing.Serialize())
	if analyzeFault != nil {
		return false
	}
	if !signing.IsEqual(adjustedSigning) {
		return false
	}

	return signing.Verify(security.Hash256(msg), pub)
}

//
//
func signingOriginatingOctets(signatureTxt []byte) *ecdsa.Signature {
	var r secp256k1.ModNScalar
	r.SetByteSlice(signatureTxt[:32])
	var s secp256k1.ModNScalar
	s.SetByteSlice(signatureTxt[32:64])
	return ecdsa.NewSignature(&r, &s)
}
