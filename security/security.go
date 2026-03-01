package security

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
)

const (
	//
	LocatorExtent = tenderminthash.AbridgedExtent
)

//
//
//
type Location = octets.HexadecimalOctets

func LocatorDigest(bz []byte) Location {
	return Location(tenderminthash.TotalAbridged(bz))
}

//
type PublicToken interface {
	Location() Location
	Octets() []byte
	ValidateSigning(msg []byte, sig []byte) bool
	Matches(PublicToken) bool
	Kind() string
}

type PrivateToken interface {
	Octets() []byte
	Attest(msg []byte) ([]byte, error)
	PublicToken() PublicToken
	Matches(PrivateToken) bool
	Kind() string
}

type Balanced interface {
	Tokengen() []byte
	Seal(cleartext []byte, credential []byte) (sealedtext []byte)
	Unseal(sealedtext []byte, credential []byte) (cleartext []byte, err error)
}

//
//
//
//
type ClusterValidator interface {
	//
	Add(key PublicToken, artifact, signing []byte) error
	//
	//
	//
	//
	Validate() (bool, []bool)
}
