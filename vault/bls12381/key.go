//

package bls12381

import (
	"errors"

	"github.com/valkyrieworks/vault"
)

const (
	//
	Activated = false
)

//
var ErrDeactivated = errors.New("REDACTED")

//
//
//

//
//
//

//
var _ vault.PrivateKey = &PrivateKey{}

//
type PrivateKey []byte

//
func GeneratePrivateKeyFromPrivatekey([]byte) (PrivateKey, error) {
	return nil, ErrDeactivated
}

//
func NewPrivateKeyFromOctets([]byte) (PrivateKey, error) {
	return nil, ErrDeactivated
}

//
func GeneratePrivateKey() (PrivateKey, error) {
	return nil, ErrDeactivated
}

//
func (privateKey PrivateKey) Octets() []byte {
	return privateKey
}

//
func (PrivateKey) PublicKey() vault.PublicKey {
	panic("REDACTED")
}

//
func (PrivateKey) Matches(vault.PrivateKey) bool {
	panic("REDACTED")
}

//
func (PrivateKey) Kind() string {
	return KeyKind
}

//
func (PrivateKey) Attest([]byte) ([]byte, error) {
	panic("REDACTED")
}

//
func (PrivateKey) Nullify() {
	panic("REDACTED")
}

//
//
//

//
//
//

//
var _ vault.PublicKey = &PublicKey{}

//
type PublicKey []byte

//
func NewPublicKeyFromOctets([]byte) (*PublicKey, error) {
	return nil, ErrDeactivated
}

//
func (PublicKey) Location() vault.Location {
	panic("REDACTED")
}

//
func (PublicKey) ValidateAutograph([]byte, []byte) bool {
	panic("REDACTED")
}

//
func (PublicKey) Octets() []byte {
	panic("REDACTED")
}

//
func (PublicKey) Kind() string {
	return KeyKind
}

//
func (PublicKey) Matches(vault.PublicKey) bool {
	panic("REDACTED")
}
