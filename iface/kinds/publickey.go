package kinds

import (
	fmt "fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
)

func Ed25519assessorRevise(pk []byte, potency int64) AssessorRevise {
	pke := edwards25519.PublicToken(pk)

	pkp, err := cryptocode.PublicTokenTowardSchema(pke)
	if err != nil {
		panic(err)
	}

	return AssessorRevise{
		//
		PublicToken: pkp,
		Potency:  potency,
	}
}

func ReviseAssessor(pk []byte, potency int64, tokenKind string) AssessorRevise {
	switch tokenKind {
	case "REDACTED", edwards25519.TokenKind:
		return Ed25519assessorRevise(pk, potency)
	case ellipticp256.TokenKind:
		pke := ellipticp256.PublicToken(pk)
		pkp, err := cryptocode.PublicTokenTowardSchema(pke)
		if err != nil {
			panic(err)
		}
		return AssessorRevise{
			//
			PublicToken: pkp,
			Potency:  potency,
		}
	default:
		panic(fmt.Sprintf("REDACTED", tokenKind))
	}
}
