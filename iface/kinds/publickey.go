package kinds

import (
	fmt "fmt"

	"github.com/valkyrieworks/vault/ed25519"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/vault/secp256k1"
)

func Ed25519ratifierModify(pk []byte, energy int64) RatifierModify {
	pke := ed25519.PublicKey(pk)

	pkp, err := cryptocode.PublicKeyToSchema(pke)
	if err != nil {
		panic(err)
	}

	return RatifierModify{
		//
		PublicKey: pkp,
		Energy:  energy,
	}
}

func ModifyRatifier(pk []byte, energy int64, keyKind string) RatifierModify {
	switch keyKind {
	case "REDACTED", ed25519.KeyKind:
		return Ed25519ratifierModify(pk, energy)
	case secp256k1.KeyKind:
		pke := secp256k1.PublicKey(pk)
		pkp, err := cryptocode.PublicKeyToSchema(pke)
		if err != nil {
			panic(err)
		}
		return RatifierModify{
			//
			PublicKey: pkp,
			Energy:  energy,
		}
	default:
		panic(fmt.Sprintf("REDACTED", keyKind))
	}
}
