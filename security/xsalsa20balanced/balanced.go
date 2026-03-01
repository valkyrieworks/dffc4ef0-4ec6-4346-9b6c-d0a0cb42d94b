package xsalsa20balanced

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/nacl/secretbox"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
)

//

const (
	numberSize  = 24
	credentialSize = 32
)

//
//
func SealBalanced(cleartext []byte, credential []byte) (sealedtext []byte) {
	if len(credential) != credentialSize {
		panic(fmt.Sprintf("REDACTED", len(credential)))
	}
	number := security.CHARArbitraryOctets(numberSize)
	numberList := [numberSize]byte{}
	copy(numberList[:], number)
	credentialList := [credentialSize]byte{}
	copy(credentialList[:], credential)
	sealedtext = make([]byte, numberSize+secretbox.Overhead+len(cleartext))
	copy(sealedtext, number)
	secretbox.Seal(sealedtext[numberSize:numberSize], cleartext, &numberList, &credentialList)
	return sealedtext
}

//
//
func UnsealBalanced(sealedtext []byte, credential []byte) (cleartext []byte, err error) {
	if len(credential) != credentialSize {
		panic(fmt.Sprintf("REDACTED", len(credential)))
	}
	if len(sealedtext) <= secretbox.Overhead+numberSize {
		return nil, errors.New("REDACTED")
	}
	number := sealedtext[:numberSize]
	numberList := [numberSize]byte{}
	copy(numberList[:], number)
	credentialList := [credentialSize]byte{}
	copy(credentialList[:], credential)
	cleartext = make([]byte, len(sealedtext)-numberSize-secretbox.Overhead)
	_, ok := secretbox.Open(cleartext[:0], sealedtext[numberSize:], &numberList, &credentialList)
	if !ok {
		return nil, errors.New("REDACTED")
	}
	return cleartext, nil
}
