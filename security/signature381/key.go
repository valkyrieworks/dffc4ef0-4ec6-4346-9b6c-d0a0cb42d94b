//

package signature381

import (
	"errors"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
)

const (
	//
	Activated = false
)

//
var FaultDeactivated = errors.New("REDACTED")

//
//
//

//
//
//

//
var _ security.PrivateToken = &PrivateToken{}

//
type PrivateToken []byte

//
func ProducePrivateTokenOriginatingCredential([]byte) (PrivateToken, error) {
	return nil, FaultDeactivated
}

//
func FreshSecludedTokenOriginatingOctets([]byte) (PrivateToken, error) {
	return nil, FaultDeactivated
}

//
func ProducePrivateToken() (PrivateToken, error) {
	return nil, FaultDeactivated
}

//
func (privateToken PrivateToken) Octets() []byte {
	return privateToken
}

//
func (PrivateToken) PublicToken() security.PublicToken {
	panic("REDACTED")
}

//
func (PrivateToken) Matches(security.PrivateToken) bool {
	panic("REDACTED")
}

//
func (PrivateToken) Kind() string {
	return TokenKind
}

//
func (PrivateToken) Attest([]byte) ([]byte, error) {
	panic("REDACTED")
}

//
func (PrivateToken) Nullify() {
	panic("REDACTED")
}

//
//
//

//
//
//

//
var _ security.PublicToken = &PublicToken{}

//
type PublicToken []byte

//
func FreshCommonTokenOriginatingOctets([]byte) (*PublicToken, error) {
	return nil, FaultDeactivated
}

//
func (PublicToken) Location() security.Location {
	panic("REDACTED")
}

//
func (PublicToken) ValidateNotation([]byte, []byte) bool {
	panic("REDACTED")
}

//
func (PublicToken) Octets() []byte {
	panic("REDACTED")
}

//
func (PublicToken) Kind() string {
	return TokenKind
}

//
func (PublicToken) Matches(security.PublicToken) bool {
	panic("REDACTED")
}
