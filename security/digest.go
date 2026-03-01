package security

import (
	"crypto/sha256"
)

func Hash256(octets []byte) []byte {
	digester := sha256.New()
	digester.Write(octets)
	return digester.Sum(nil)
}
