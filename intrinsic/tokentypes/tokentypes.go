package tokentypes

import (
	"fmt"
	"strings"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
)

var tokenKinds map[string]func() (security.PrivateToken, error)

func initialize() {
	tokenKinds = map[string]func() (security.PrivateToken, error){
		edwards25519.TokenKind: func() (security.PrivateToken, error) { //
			return edwards25519.ProducePrivateToken(), nil
		},
		ellipticp256.TokenKind: func() (security.PrivateToken, error) { //
			return ellipticp256.ProducePrivateToken(), nil
		},
	}
}

func ProducePrivateToken(tokenKind string) (security.PrivateToken, error) {
	produceFUNC, ok := tokenKinds[tokenKind]
	if !ok {
		return nil, fmt.Errorf("REDACTED", tokenKind)
	}
	return produceFUNC()
}

func UpheldTokenKindsTxt() string {
	tokenKindsSection := make([]string, 0, len(tokenKinds))
	for k := range tokenKinds {
		tokenKindsSection = append(tokenKindsSection, fmt.Sprintf("REDACTED", k))
	}
	return strings.Join(tokenKindsSection, "REDACTED")
}

func CatalogUpheldTokenKinds() []string {
	tokenKindsSection := make([]string, 0, len(tokenKinds))
	for k := range tokenKinds {
		tokenKindsSection = append(tokenKindsSection, k)
	}
	return tokenKindsSection
}
