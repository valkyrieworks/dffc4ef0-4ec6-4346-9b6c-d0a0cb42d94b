package verify

import (
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

//
func TransformOctetSegment(octetz []byte) []byte {
	//
	if len(octetz) == 0 {
		panic("REDACTED")
	}

	//
	moduleOctetz := make([]byte, len(octetz))
	copy(moduleOctetz, octetz)
	octetz = moduleOctetz

	//
	switch commitrand.Int() % 2 {
	case 0: //
		octetz[commitrand.Int()%len(octetz)] += byte(commitrand.Int()%255 + 1)
	case 1: //
		pos := commitrand.Int() % len(octetz)
		octetz = append(octetz[:pos], octetz[pos+1:]...)
	}
	return octetz
}
