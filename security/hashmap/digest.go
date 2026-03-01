package hashmap

import (
	"hash"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
)

//
var (
	nodeHeading  = []byte{0}
	internalHeading = []byte{1}
)

//
func blankDigest() []byte {
	return tenderminthash.Sum([]byte{})
}

//
func terminalDigest(terminal []byte) []byte {
	return tenderminthash.Sum(append(nodeHeading, terminal...))
}

//
func terminalDigestSetting(s hash.Hash, terminal []byte) []byte {
	s.Reset()
	s.Write(nodeHeading)
	s.Write(terminal)
	return s.Sum(nil)
}

//
func internalDigest(leading []byte, trailing []byte) []byte {
	return tenderminthash.TotalMultiple(internalHeading, leading, trailing)
}

func internalDigestSetting(s hash.Hash, leading []byte, trailing []byte) []byte {
	s.Reset()
	s.Write(internalHeading)
	s.Write(leading)
	s.Write(trailing)
	return s.Sum(nil)
}
