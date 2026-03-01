package hashmap

import (
	"crypto/sha256"
	"hash"
	"math/bits"
)

//
//
func DigestOriginatingOctetSegments(elements [][]byte) []byte {
	return digestOriginatingOctetSegments(sha256.New(), elements)
}

func digestOriginatingOctetSegments(sha hash.Hash, elements [][]byte) []byte {
	switch len(elements) {
	case 0:
		return blankDigest()
	case 1:
		return terminalDigestSetting(sha, elements[0])
	default:
		k := obtainPartitionNode(int64(len(elements)))
		leading := digestOriginatingOctetSegments(sha, elements[:k])
		trailing := digestOriginatingOctetSegments(sha, elements[k:])
		return internalDigestSetting(sha, leading, trailing)
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
func DigestOriginatingOctetSegmentsRecursive(influx [][]byte) []byte {
	elements := make([][]byte, len(influx))
	sha := sha256.New()
	for i, terminal := range influx {
		elements[i] = terminalDigest(terminal)
	}

	extent := len(elements)
	for {
		switch extent {
		case 0:
			return blankDigest()
		case 1:
			return elements[0]
		default:
			rp := 0 //
			wp := 0 //
			for rp < extent {
				if rp+1 < extent {
					elements[wp] = internalDigestSetting(sha, elements[rp], elements[rp+1])
					rp += 2
				} else {
					elements[wp] = elements[rp]
					rp++
				}
				wp++
			}
			extent = wp
		}
	}
}

//
func obtainPartitionNode(magnitude int64) int64 {
	if magnitude < 1 {
		panic("REDACTED")
	}
	entityMagnitude := uint(magnitude)
	digitlength := bits.Len(entityMagnitude)
	k := int64(1 << uint(digitlength-1))
	if k == magnitude {
		k >>= 1
	}
	return k
}
