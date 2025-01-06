package main

import (
	"errors"
	"fmt"
	"math"
)

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

//nolint:cyclop,revive
func isConvertible[T, U integer](value T) bool {
	var uMin, uMax int64

	if isSigned[U]() {
		uMin, uMax = minSigned[U](), maxSigned[U]()
	} else {
		uMin, uMax = 0, int64(maxUnsigned[U]())
	}

	switch {
	// int to int
	case isSigned[T]() && isSigned[U]():
		return int64(value) >= uMin && int64(value) <= uMax
	// uint to int
	case !isSigned[T]() && isSigned[U]():
		return uint64(value) <= uint64(uMax)
	// int to uint
	case isSigned[T]() && !isSigned[U]():
		return int64(value) >= 0 && uint64(value) <= uint64(uMax)
	// uint to uint
	case !isSigned[T]() && !isSigned[U]():
		return uint64(value) <= uint64(uMax)
	}

	return false
}

func isSigned[T integer]() bool {
	var t T
	return t-1 < t
}

func minSigned[T integer]() int64 {
	switch any(T(0)).(type) {
	case int8:
		return math.MinInt8
	case int16:
		return math.MinInt16
	case int32:
		return math.MinInt32
	case int64:
		return math.MinInt64
	}

	return 0
}

func maxSigned[T integer]() int64 {
	switch any(T(0)).(type) {
	case int8:
		return math.MaxInt8
	case int16:
		return math.MaxInt16
	case int32:
		return math.MaxInt32
	case int64:
		return math.MaxInt64
	}

	return 0
}

func maxUnsigned[T integer]() uint64 {
	switch any(T(0)).(type) {
	case uint8:
		return math.MaxUint8
	case uint16:
		return math.MaxUint16
	case uint32:
		return math.MaxUint32
	case uint64:
		return math.MaxUint64
	}

	return 0
}

var errIntOverflow = errors.New("integer overflow")

func ConvertInteger[T, U integer](input T) (U, error) {
	var output U

	if ok := isConvertible[T, U](input); !ok {
		return output, errIntOverflow
	}

	return U(input), nil
}

func main() {
	fmt.Println(isConvertible[int8, int16](math.MinInt8))
	fmt.Println(isConvertible[int8, int16](0))
	fmt.Println(isConvertible[int8, int16](math.MaxInt8))
	//fmt.Println(isConvertible[int8, int64](math.MinInt8))
	//fmt.Println(isConvertible[int8, int64](0))
	//fmt.Println(isConvertible[int8, int64](math.MaxInt8))
	//fmt.Println(isConvertible[int8, uint16](math.MinInt8))
	//fmt.Println(isConvertible[int8, uint16](0))
	//fmt.Println(isConvertible[int8, uint16](math.MaxInt8))
}
