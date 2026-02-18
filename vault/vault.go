package vault

import (
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/utils/octets"
)

const (
	//
	LocationVolume = comethash.ShortenedVolume
)

//
//
//
type Location = octets.HexOctets

func LocationDigest(bz []byte) Location {
	return Location(comethash.TotalShortened(bz))
}

//
type PublicKey interface {
	Location() Location
	Octets() []byte
	ValidateAutograph(msg []byte, sig []byte) bool
	Matches(PublicKey) bool
	Kind() string
}

type PrivateKey interface {
	Octets() []byte
	Attest(msg []byte) ([]byte, error)
	PublicKey() PublicKey
	Matches(PrivateKey) bool
	Kind() string
}

type Balanced interface {
	Keygen() []byte
	Encode(cleartext []byte, key []byte) (cyphertext []byte)
	Decode(cyphertext []byte, key []byte) (cleartext []byte, err error)
}

//
//
//
//
type GroupValidator interface {
	//
	Add(key PublicKey, signal, autograph []byte) error
	//
	//
	//
	//
	Validate() (bool, []bool)
}
