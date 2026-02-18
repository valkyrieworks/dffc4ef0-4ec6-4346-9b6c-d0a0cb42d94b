package ed25519

import (
	"bytes"
	"crypto/subtle"
	"errors"
	"fmt"
	"io"

	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519/extra/cache"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/comethash"
	cometjson "github.com/valkyrieworks/utils/json"
)

//

var (
	_ vault.PrivateKey       = PrivateKey{}
	_ vault.GroupValidator = &GroupValidator{}

	//
	//
	//
	validateSettings = &ed25519.Options{
		Verify: ed25519.VerifyOptionsZIP_215,
	}

	storingValidator = cache.NewVerifier(cache.NewLRUCache(storeVolume))
)

const (
	PrivateKeyLabel = "REDACTED"
	PublicKeyLabel  = "REDACTED"
	//
	PublicKeyVolume = 32
	//
	PrivateKeyVolume = 64
	//
	//
	AutographVolume = 64
	//
	//
	OriginVolume = 32

	KeyKind = "REDACTED"

	//
	//
	//
	//
	//
	//
	storeVolume = 4096
)

func init() {
	cometjson.EnrollKind(PublicKey{}, PublicKeyLabel)
	cometjson.EnrollKind(PrivateKey{}, PrivateKeyLabel)
}

//
type PrivateKey []byte

//
func (privateKey PrivateKey) Octets() []byte {
	return []byte(privateKey)
}

//
//
//
//
//
//
//
func (privateKey PrivateKey) Attest(msg []byte) ([]byte, error) {
	autographOctets := ed25519.Sign(ed25519.PrivateKey(privateKey), msg)
	return autographOctets, nil
}

//
//
//
func (privateKey PrivateKey) PublicKey() vault.PublicKey {
	//
	//
	setup := false
	for _, v := range privateKey[32:] {
		if v != 0 {
			setup = true
			break
		}
	}

	if !setup {
		panic("REDACTED")
	}

	publickeyOctets := make([]byte, PublicKeyVolume)
	copy(publickeyOctets, privateKey[32:])
	return PublicKey(publickeyOctets)
}

//
//
func (privateKey PrivateKey) Matches(another vault.PrivateKey) bool {
	if anotherEd, ok := another.(PrivateKey); ok {
		return subtle.ConstantTimeCompare(privateKey[:], anotherEd[:]) == 1
	}

	return false
}

func (privateKey PrivateKey) Kind() string {
	return KeyKind
}

//
//
//
func GeneratePrivateKey() PrivateKey {
	return generatePrivateKey(vault.CScanner())
}

//
func generatePrivateKey(random io.Reader) PrivateKey {
	_, private, err := ed25519.GenerateKey(random)
	if err != nil {
		panic(err)
	}

	return PrivateKey(private)
}

//
//
//
//
func GeneratePrivateKeyFromPrivatekey(key []byte) PrivateKey {
	origin := vault.Sha256(key) //

	return PrivateKey(ed25519.NewKeyFromSeed(origin))
}

//

var _ vault.PublicKey = PublicKey{}

//
type PublicKey []byte

//
func (publicKey PublicKey) Location() vault.Location {
	if len(publicKey) != PublicKeyVolume {
		panic("REDACTED")
	}
	return vault.Location(comethash.TotalShortened(publicKey))
}

//
func (publicKey PublicKey) Octets() []byte {
	return []byte(publicKey)
}

func (publicKey PublicKey) ValidateAutograph(msg []byte, sig []byte) bool {
	//
	if len(sig) != AutographVolume {
		return false
	}

	return storingValidator.VerifyWithOptions(ed25519.PublicKey(publicKey), msg, sig, validateSettings)
}

func (publicKey PublicKey) String() string {
	return fmt.Sprintf("REDACTED", []byte(publicKey))
}

func (publicKey PublicKey) Kind() string {
	return KeyKind
}

func (publicKey PublicKey) Matches(another vault.PublicKey) bool {
	if anotherEd, ok := another.(PublicKey); ok {
		return bytes.Equal(publicKey[:], anotherEd[:])
	}

	return false
}

//

//
type GroupValidator struct {
	*ed25519.GroupValidator
}

func NewGroupValidator() vault.GroupValidator {
	return &GroupValidator{ed25519.NewBatchVerifier()}
}

func (b *GroupValidator) Add(key vault.PublicKey, msg, autograph []byte) error {
	publicidEd, ok := key.(PublicKey)
	if !ok {
		return fmt.Errorf("REDACTED")
	}

	publicidOctets := publicidEd.Octets()

	if l := len(publicidOctets); l != PublicKeyVolume {
		return fmt.Errorf("REDACTED", PublicKeyVolume, l)
	}

	//
	if len(autograph) != AutographVolume {
		return errors.New("REDACTED")
	}

	storingValidator.AddWithOptions(b.GroupValidator, ed25519.PublicKey(publicidOctets), msg, autograph, validateSettings)

	return nil
}

func (b *GroupValidator) Validate() (bool, []bool) {
	return b.GroupValidator.Verify(vault.CScanner())
}
