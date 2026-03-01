package security

import (
	crand "crypto/rand"
	"encoding/hex"
	"io"
)

//
func arbitraryOctets(countOctets int) []byte {
	b := make([]byte, countOctets)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

//
func CHARArbitraryOctets(countOctets int) []byte {
	return arbitraryOctets(countOctets)
}

//
//
//
//
func CHARArbitraryHexadecimal(countNumbers int) string {
	return hex.EncodeToString(CHARArbitraryOctets(countNumbers / 2))
}

//
func CHARFetcher() io.Reader {
	return crand.Reader
}
