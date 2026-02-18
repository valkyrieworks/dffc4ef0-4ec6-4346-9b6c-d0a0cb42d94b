package vault

import (
	crand "crypto/rand"
	"encoding/hex"
	"io"
)

//
func randomOctets(countOctets int) []byte {
	b := make([]byte, countOctets)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

//
func CRandomOctets(countOctets int) []byte {
	return randomOctets(countOctets)
}

//
//
//
//
func CRandomHex(countNumerals int) string {
	return hex.EncodeToString(CRandomOctets(countNumerals / 2))
}

//
func CScanner() io.Reader {
	return crand.Reader
}
