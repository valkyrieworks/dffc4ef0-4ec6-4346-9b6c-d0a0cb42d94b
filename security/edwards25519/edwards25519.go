package edwards25519

import (
	"bytes"
	"crypto/subtle"
	"errors"
	"fmt"
	"io"

	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519/extra/cache"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
)

//

var (
	_ security.PrivateToken       = PrivateToken{}
	_ security.ClusterValidator = &ClusterValidator{}

	//
	//
	//
	validateChoices = &ed25519.Options{
		Verify: ed25519.VerifyOptionsZIP_215,
	}

	stashingValidator = cache.NewVerifier(cache.NewLRUCache(stashExtent))
)

const (
	PrivateTokenAlias = "REDACTED"
	PublicTokenAlias  = "REDACTED"
	//
	PublicTokenExtent = 32
	//
	SecludedTokenExtent = 64
	//
	//
	SigningExtent = 64
	//
	//
	GermExtent = 32

	TokenKind = "REDACTED"

	//
	//
	//
	//
	//
	//
	stashExtent = 4096
)

func initialize() {
	strongmindjson.EnrollKind(PublicToken{}, PublicTokenAlias)
	strongmindjson.EnrollKind(PrivateToken{}, PrivateTokenAlias)
}

//
type PrivateToken []byte

//
func (privateToken PrivateToken) Octets() []byte {
	return []byte(privateToken)
}

//
//
//
//
//
//
//
func (privateToken PrivateToken) Attest(msg []byte) ([]byte, error) {
	signingOctets := ed25519.Sign(ed25519.PrivateKey(privateToken), msg)
	return signingOctets, nil
}

//
//
//
func (privateToken PrivateToken) PublicToken() security.PublicToken {
	//
	//
	started := false
	for _, v := range privateToken[32:] {
		if v != 0 {
			started = true
			break
		}
	}

	if !started {
		panic("REDACTED")
	}

	publickeyOctets := make([]byte, PublicTokenExtent)
	copy(publickeyOctets, privateToken[32:])
	return PublicToken(publickeyOctets)
}

//
//
func (privateToken PrivateToken) Matches(another security.PrivateToken) bool {
	if anotherEdwards, ok := another.(PrivateToken); ok {
		return subtle.ConstantTimeCompare(privateToken[:], anotherEdwards[:]) == 1
	}

	return false
}

func (privateToken PrivateToken) Kind() string {
	return TokenKind
}

//
//
//
func ProducePrivateToken() PrivateToken {
	return producePrivateToken(security.CHARFetcher())
}

//
func producePrivateToken(arbitrary io.Reader) PrivateToken {
	_, private, err := ed25519.GenerateKey(arbitrary)
	if err != nil {
		panic(err)
	}

	return PrivateToken(private)
}

//
//
//
//
func ProducePrivateTokenOriginatingCredential(credential []byte) PrivateToken {
	germ := security.Hash256(credential) //

	return PrivateToken(ed25519.NewKeyFromSeed(germ))
}

//

var _ security.PublicToken = PublicToken{}

//
type PublicToken []byte

//
func (publicToken PublicToken) Location() security.Location {
	if len(publicToken) != PublicTokenExtent {
		panic("REDACTED")
	}
	return security.Location(tenderminthash.TotalAbridged(publicToken))
}

//
func (publicToken PublicToken) Octets() []byte {
	return []byte(publicToken)
}

func (publicToken PublicToken) ValidateNotation(msg []byte, sig []byte) bool {
	//
	if len(sig) != SigningExtent {
		return false
	}

	return stashingValidator.VerifyWithOptions(ed25519.PublicKey(publicToken), msg, sig, validateChoices)
}

func (publicToken PublicToken) Text() string {
	return fmt.Sprintf("REDACTED", []byte(publicToken))
}

func (publicToken PublicToken) Kind() string {
	return TokenKind
}

func (publicToken PublicToken) Matches(another security.PublicToken) bool {
	if anotherEdwards, ok := another.(PublicToken); ok {
		return bytes.Equal(publicToken[:], anotherEdwards[:])
	}

	return false
}

//

//
type ClusterValidator struct {
	*ed25519.ClusterValidator
}

func FreshClusterValidator() security.ClusterValidator {
	return &ClusterValidator{ed25519.NewBatchVerifier()}
}

func (b *ClusterValidator) Add(key security.PublicToken, msg, signing []byte) error {
	keyEdwards, ok := key.(PublicToken)
	if !ok {
		return fmt.Errorf("REDACTED")
	}

	keyOctets := keyEdwards.Octets()

	if l := len(keyOctets); l != PublicTokenExtent {
		return fmt.Errorf("REDACTED", PublicTokenExtent, l)
	}

	//
	if len(signing) != SigningExtent {
		return errors.New("REDACTED")
	}

	stashingValidator.AddWithOptions(b.ClusterValidator, ed25519.PublicKey(keyOctets), msg, signing, validateChoices)

	return nil
}

func (b *ClusterValidator) Validate() (bool, []bool) {
	return b.ClusterValidator.Verify(security.CHARFetcher())
}
