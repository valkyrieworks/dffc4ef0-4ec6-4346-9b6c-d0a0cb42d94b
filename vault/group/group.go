package group

import (
	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
)

//
//
func InstantiateGroupValidator(pk vault.PublicKey) (vault.GroupValidator, bool) {
	switch pk.Kind() {
	case ed25519.KeyKind:
		return ed25519.NewGroupValidator(), true
	default:
		return nil, false
	}
}

//
//
func SustainsGroupValidator(pk vault.PublicKey) bool {
	if pk == nil {
		return false
	}

	switch pk.Kind() {
	case ed25519.KeyKind:
		return true
	default:
		return false
	}
}
