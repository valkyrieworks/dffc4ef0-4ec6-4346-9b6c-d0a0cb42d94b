package cluster

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
)

//
//
func GenerateClusterValidator(pk security.PublicToken) (security.ClusterValidator, bool) {
	switch pk.Kind() {
	case edwards25519.TokenKind:
		return edwards25519.FreshClusterValidator(), true
	default:
		return nil, false
	}
}

//
//
func SustainsClusterValidator(pk security.PublicToken) bool {
	if pk == nil {
		return false
	}

	switch pk.Kind() {
	case edwards25519.TokenKind:
		return true
	default:
		return false
	}
}
