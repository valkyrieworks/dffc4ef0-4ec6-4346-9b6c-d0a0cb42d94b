package merkle

import (
	"crypto/sha256"
	"hash"
	"math/bits"
)

//
//
func DigestFromOctetSegments(items [][]byte) []byte {
	return digestFromOctetSegments(sha256.New(), items)
}

func digestFromOctetSegments(sha hash.Hash, items [][]byte) []byte {
	switch len(items) {
	case 0:
		return emptyDigest()
	case 1:
		return elementDigestOption(sha, items[0])
	default:
		k := fetchDivideSpot(int64(len(items)))
		left := digestFromOctetSegments(sha, items[:k])
		correct := digestFromOctetSegments(sha, items[k:])
		return deeperDigestOption(sha, left, correct)
	}
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
func DigestFromOctetSegmentsRecursive(influx [][]byte) []byte {
	items := make([][]byte, len(influx))
	sha := sha256.New()
	for i, element := range influx {
		items[i] = elementDigest(element)
	}

	volume := len(items)
	for {
		switch volume {
		case 0:
			return emptyDigest()
		case 1:
			return items[0]
		default:
			rp := 0 //
			wp := 0 //
			for rp < volume {
				if rp+1 < volume {
					items[wp] = deeperDigestOption(sha, items[rp], items[rp+1])
					rp += 2
				} else {
					items[wp] = items[rp]
					rp++
				}
				wp++
			}
			volume = wp
		}
	}
}

//
func fetchDivideSpot(extent int64) int64 {
	if extent < 1 {
		panic("REDACTED")
	}
	uExtent := uint(extent)
	bitextent := bits.Len(uExtent)
	k := int64(1 << uint(bitextent-1))
	if k == extent {
		k >>= 1
	}
	return k
}
