package crypto_test

import (
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
)

func InstanceHash256() {
	sum := security.Hash256([]byte("REDACTED"))
	fmt.Printf("REDACTED", sum)
	//
	//
}
