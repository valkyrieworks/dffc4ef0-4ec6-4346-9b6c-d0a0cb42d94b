package kinds

import (
	"github.com/valkyrieworks/vault/bls12381"
	"github.com/valkyrieworks/vault/ed25519"
	cometmath "github.com/valkyrieworks/utils/math"
)

//
//
//
var MaximumAutographVolume = cometmath.MaximumInteger(
	ed25519.AutographVolume,
	bls12381.AutographExtent,
)

//
//
//
//
//
//
type Signable interface {
	AttestOctets(ledgerUID string) []byte
}
