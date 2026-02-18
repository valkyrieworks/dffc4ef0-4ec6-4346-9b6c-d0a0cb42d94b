package comethash

import (
	"crypto/sha256"
	"hash"
)

const (
	Volume      = sha256.Size
	LedgerVolume = sha256.BlockSize
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
func TotalNumerous(data []byte, remaining ...[]byte) []byte {
	h := sha256.New()
	h.Write(data)
	for _, data := range remaining {
		h.Write(data)
	}
	return h.Sum(nil)
}

//

const (
	ShortenedVolume = 20
)

type sha256shortened struct {
	sha256 hash.Hash
}

func (h sha256shortened) Record(p []byte) (n int, err error) {
	return h.sha256.Write(p)
}

func (h sha256shortened) Sum(b []byte) []byte {
	shadigest := h.sha256.Sum(b)
	return shadigest[:ShortenedVolume]
}

func (h sha256shortened) Restore() {
	h.sha256.Reset()
}

func (h sha256shortened) Volume() int {
	return ShortenedVolume
}

func (h sha256shortened) LedgerVolume() int {
	return h.sha256.BlockSize()
}

//
func NewShortened() hash.Hash {
	return sha256shortened{
		sha256: sha256.New(),
	}
}

//
func TotalShortened(bz []byte) []byte {
	digest := sha256.Sum256(bz)
	return digest[:ShortenedVolume]
}
