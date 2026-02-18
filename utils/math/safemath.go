package math

import (
	"errors"
	"math"
)

var (
	ErrOverloadInt32 = errors.New("REDACTED")
	ErrOverloadUint8 = errors.New("REDACTED")
	ErrOverloadInt8  = errors.New("REDACTED")
)

//
//
func SecureAppendInt32(a, b int32) int32 {
	if b > 0 && (a > math.MaxInt32-b) {
		panic(ErrOverloadInt32)
	} else if b < 0 && (a < math.MinInt32-b) {
		panic(ErrOverloadInt32)
	}
	return a + b
}

//
//
func SecureSubtractInt32(a, b int32) int32 {
	if b > 0 && (a < math.MinInt32+b) {
		panic(ErrOverloadInt32)
	} else if b < 0 && (a > math.MaxInt32+b) {
		panic(ErrOverloadInt32)
	}
	return a - b
}

//
//
func SecureTransformInt32(a int64) int32 {
	if a > math.MaxInt32 {
		panic(ErrOverloadInt32)
	} else if a < math.MinInt32 {
		panic(ErrOverloadInt32)
	}
	return int32(a)
}

//
//
func SecureTransformUint8(a int64) (uint8, error) {
	if a > math.MaxUint8 {
		return 0, ErrOverloadUint8
	} else if a < 0 {
		return 0, ErrOverloadUint8
	}
	return uint8(a), nil
}

//
//
func SecureTransformInt8(a int64) (int8, error) {
	if a > math.MaxInt8 {
		return 0, ErrOverloadInt8
	} else if a < math.MinInt8 {
		return 0, ErrOverloadInt8
	}
	return int8(a), nil
}
