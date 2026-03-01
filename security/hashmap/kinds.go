package hashmap

import (
	"encoding/binary"
	"io"
)

//
type Graph interface {
	Extent() (extent int)
	Altitude() (altitude int8)
	Has(key []byte) (has bool)
	Attestation(key []byte) (datum []byte, attestation []byte, present bool) //
	Get(key []byte) (ordinal int, datum []byte, present bool)
	ObtainViaOrdinal(ordinal int) (key []byte, datum []byte)
	Set(key []byte, datum []byte) (modified bool)
	Discard(key []byte) (datum []byte, discarded bool)
	DigestUsingTally() (digest []byte, tally int)
	Digest() (digest []byte)
	Persist() (digest []byte)
	Fetch(digest []byte)
	Duplicate() Graph
	Traverse(func(key []byte, datum []byte) (halt bool)) (halted bool)
	TraverseScope(initiate []byte, end []byte, increasing bool, fx func(key []byte, datum []byte) (halt bool)) (halted bool)
}

//

//
func serializeOctetSegment(w io.Writer, bz []byte) (err error) {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(buf[:], uint64(len(bz)))
	_, err = w.Write(buf[0:n])
	if err != nil {
		return
	}
	_, err = w.Write(bz)
	return
}
