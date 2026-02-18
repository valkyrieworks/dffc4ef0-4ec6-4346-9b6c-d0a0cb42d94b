package keykinds

import (
	"fmt"
	"strings"

	"github.com/valkyrieworks/vault"

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/secp256k1"
)

var keyKinds map[string]func() (vault.PrivateKey, error)

func init() {
	keyKinds = map[string]func() (vault.PrivateKey, error){
		ed25519.KeyKind: func() (vault.PrivateKey, error) { //
			return ed25519.GeneratePrivateKey(), nil
		},
		secp256k1.KeyKind: func() (vault.PrivateKey, error) { //
			return secp256k1.GeneratePrivateKey(), nil
		},
	}
}

func GeneratePrivateKey(keyKind string) (vault.PrivateKey, error) {
	generateF, ok := keyKinds[keyKind]
	if !ok {
		return nil, fmt.Errorf("REDACTED", keyKind)
	}
	return generateF()
}

func UpheldKeyKindsStr() string {
	keyKindsSection := make([]string, 0, len(keyKinds))
	for k := range keyKinds {
		keyKindsSection = append(keyKindsSection, fmt.Sprintf("REDACTED", k))
	}
	return strings.Join(keyKindsSection, "REDACTED")
}

func CatalogUpheldKeyKinds() []string {
	keyKindsSection := make([]string, 0, len(keyKinds))
	for k := range keyKinds {
		keyKindsSection = append(keyKindsSection, k)
	}
	return keyKindsSection
}
