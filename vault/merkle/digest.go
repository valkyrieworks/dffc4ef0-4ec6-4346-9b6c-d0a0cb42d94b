package merkle

import (
	"hash"

	"github.com/valkyrieworks/vault/comethash"
)

//
var (
	elementPrefix  = []byte{0}
	deeperPrefix = []byte{1}
)

//
func emptyDigest() []byte {
	return comethash.Sum([]byte{})
}

//
func elementDigest(element []byte) []byte {
	return comethash.Sum(append(elementPrefix, element...))
}

//
func elementDigestOption(s hash.Hash, element []byte) []byte {
	s.Reset()
	s.Write(elementPrefix)
	s.Write(element)
	return s.Sum(nil)
}

//
func deeperDigest(left []byte, correct []byte) []byte {
	return comethash.TotalNumerous(deeperPrefix, left, correct)
}

func deeperDigestOption(s hash.Hash, left []byte, correct []byte) []byte {
	s.Reset()
	s.Write(deeperPrefix)
	s.Write(left)
	s.Write(correct)
	return s.Sum(nil)
}
