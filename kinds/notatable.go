package kinds

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/signature381"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
)

//
//
//
var MaximumSigningExtent = strongarithmetic.MaximumInteger(
	edwards25519.SigningExtent,
	signature381.SigningMagnitude,
)

//
//
//
//
//
//
type Notatable interface {
	AttestOctets(successionUUID string) []byte
}
