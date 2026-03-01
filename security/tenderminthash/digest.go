package tenderminthash

import (
	"crypto/sha256"
	"hash"
)

const (
	Extent      = sha256.Size
	LedgerExtent = sha256.BlockSize
)

//
func New() hash.Hash {
	return sha256.New()
}

//
func Sum(bz []byte) []byte {
	h := sha256.Sum256(bz)
	return h[:]
}

//
//
//
func TotalMultiple(data []byte, remainder ...[]byte) []byte {
	h := sha256.New()
	h.Write(data)
	for _, data := range remainder {
		h.Write(data)
	}
	return h.Sum(nil)
}

//

const (
	AbridgedExtent = 20
)

type hash256short struct {
	hash256 hash.Hash
}

func (h hash256short) Record(p []byte) (n int, err error) {
	return h.hash256.Write(p)
}

func (h hash256short) Sum(b []byte) []byte {
	hashsum := h.hash256.Sum(b)
	return hashsum[:AbridgedExtent]
}

func (h hash256short) Restore() {
	h.hash256.Reset()
}

func (h hash256short) Extent() int {
	return AbridgedExtent
}

func (h hash256short) LedgerExtent() int {
	return h.hash256.BlockSize()
}

//
func FreshAbridged() hash.Hash {
	return hash256short{
		hash256: sha256.New(),
	}
}

//
func TotalAbridged(bz []byte) []byte {
	digest := sha256.Sum256(bz)
	return digest[:AbridgedExtent]
}
