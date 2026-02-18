package merkle

import (
	"encoding/binary"
	"io"
)

//
type Graph interface {
	Volume() (volume int)
	Level() (level int8)
	Has(key []byte) (has bool)
	Attestation(key []byte) (item []byte, evidence []byte, present bool) //
	Get(key []byte) (ordinal int, item []byte, present bool)
	FetchByOrdinal(ordinal int) (key []byte, item []byte)
	Set(key []byte, item []byte) (refreshed bool)
	Delete(key []byte) (item []byte, deleted bool)
	DigestWithTotal() (digest []byte, tally int)
	Digest() (digest []byte)
	Persist() (digest []byte)
	Import(digest []byte)
	Clone() Graph
	Recurse(func(key []byte, item []byte) (halt bool)) (ceased bool)
	RecurseScope(begin []byte, end []byte, increasing bool, fx func(key []byte, item []byte) (halt bool)) (ceased bool)
}

//

//
func encodeOctetSegment(w io.Writer, bz []byte) (err error) {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(buf[:], uint64(len(bz)))
	_, err = w.Write(buf[0:n])
	if err != nil {
		return
	}
	_, err = w.Write(bz)
	return
}
