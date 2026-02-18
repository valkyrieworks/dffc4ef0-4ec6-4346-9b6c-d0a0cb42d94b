package crypt_test

import (
	"fmt"

	"github.com/valkyrieworks/vault"
)

func InstanceSha256() {
	sum := vault.Sha256([]byte("REDACTED"))
	fmt.Printf("REDACTED", sum)
	//
	//
}
