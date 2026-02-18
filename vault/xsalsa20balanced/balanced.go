package xsalsa20balanced

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/nacl/secretbox"

	"github.com/valkyrieworks/vault"
)

//

const (
	nonceSize  = 24
	privateSize = 32
)

//
//
func EncodeBalanced(cleartext []byte, key []byte) (cyphertext []byte) {
	if len(key) != privateSize {
		panic(fmt.Sprintf("REDACTED", len(key)))
	}
	nonce := vault.CRandomOctets(nonceSize)
	nonceArr := [nonceSize]byte{}
	copy(nonceArr[:], nonce)
	privateArr := [privateSize]byte{}
	copy(privateArr[:], key)
	cyphertext = make([]byte, nonceSize+secretbox.Overhead+len(cleartext))
	copy(cyphertext, nonce)
	secretbox.Seal(cyphertext[nonceSize:nonceSize], cleartext, &nonceArr, &privateArr)
	return cyphertext
}

//
//
func DecodeBalanced(cyphertext []byte, key []byte) (cleartext []byte, err error) {
	if len(key) != privateSize {
		panic(fmt.Sprintf("REDACTED", len(key)))
	}
	if len(cyphertext) <= secretbox.Overhead+nonceSize {
		return nil, errors.New("REDACTED")
	}
	nonce := cyphertext[:nonceSize]
	nonceArr := [nonceSize]byte{}
	copy(nonceArr[:], nonce)
	privateArr := [privateSize]byte{}
	copy(privateArr[:], key)
	cleartext = make([]byte, len(cyphertext)-nonceSize-secretbox.Overhead)
	_, ok := secretbox.Open(cleartext[:0], cyphertext[nonceSize:], &nonceArr, &privateArr)
	if !ok {
		return nil, errors.New("REDACTED")
	}
	return cleartext, nil
}
