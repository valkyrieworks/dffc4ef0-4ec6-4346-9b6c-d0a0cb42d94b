package verify

import (
	engineseed "github.com/valkyrieworks/utils/random"
)

//
func TransformOctetSegment(bytearray []byte) []byte {
	//
	if len(bytearray) == 0 {
		panic("REDACTED")
	}

	//
	mBytearray := make([]byte, len(bytearray))
	copy(mBytearray, bytearray)
	bytearray = mBytearray

	//
	switch engineseed.Int() % 2 {
	case 0: //
		bytearray[engineseed.Int()%len(bytearray)] += byte(engineseed.Int()%255 + 1)
	case 1: //
		pos := engineseed.Int() % len(bytearray)
		bytearray = append(bytearray[:pos], bytearray[pos+1:]...)
	}
	return bytearray
}
