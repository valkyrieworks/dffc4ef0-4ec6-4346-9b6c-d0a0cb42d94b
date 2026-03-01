package arithmetic

import (
	"errors"
	"math"
)

var (
	FaultOverrunInteger32 = errors.New("REDACTED")
	FaultOverrunOctet = errors.New("REDACTED")
	FaultOverrunInt8n  = errors.New("REDACTED")
)

//
//
func SecureAppendInteger32(a, b int32) int32 {
	if b > 0 && (a > math.MaxInt32-b) {
		panic(FaultOverrunInteger32)
	} else if b < 0 && (a < math.MinInt32-b) {
		panic(FaultOverrunInteger32)
	}
	return a + b
}

//
//
func SecureUnderInteger32(a, b int32) int32 {
	if b > 0 && (a < math.MinInt32+b) {
		panic(FaultOverrunInteger32)
	} else if b < 0 && (a > math.MaxInt32+b) {
		panic(FaultOverrunInteger32)
	}
	return a - b
}

//
//
func SecureAdaptInteger32(a int64) int32 {
	if a > math.MaxInt32 {
		panic(FaultOverrunInteger32)
	} else if a < math.MinInt32 {
		panic(FaultOverrunInteger32)
	}
	return int32(a)
}

//
//
func SecureTransformOctet(a int64) (uint8, error) {
	if a > math.MaxUint8 {
		return 0, FaultOverrunOctet
	} else if a < 0 {
		return 0, FaultOverrunOctet
	}
	return uint8(a), nil
}

//
//
func SecureAdaptInt8n(a int64) (int8, error) {
	if a > math.MaxInt8 {
		return 0, FaultOverrunInt8n
	} else if a < math.MinInt8 {
		return 0, FaultOverrunInt8n
	}
	return int8(a), nil
}
