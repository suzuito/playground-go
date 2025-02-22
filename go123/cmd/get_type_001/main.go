package main

import (
	"errors"
	"math"
)

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type convertIntegerFunc[T, U integer] func(input T) (U, error)

var errp1 error = errors.New("p1")
var errp2 error = errors.New("p2")
var errp3 error = errors.New("p3")

func convertIntegerV1[T, U integer](input T) (U, error) {
	var zero U
	switch inputT := any(input).(type) {
	case int8:
		switch any(zero).(type) {
		// case int8, int16, int32, int64: // 可読性向上のためにコメントを残す
		case uint8, uint16, uint32, uint64:
			if input < 0 {
				return zero, errp3
			}
		}
	case int16:
		switch any(zero).(type) {
		case int8:
			if inputT < math.MinInt8 || inputT > math.MaxInt8 {
				return zero, errp1
			}
		// case int16, int32, int64: // 可読性向上のためにコメントを残す
		case uint8:
			if inputT < 0 || inputT > math.MaxUint8 {
				return zero, errp1
			}
		case uint16, uint32, uint64:
			if input < 0 {
				return zero, errp3
			}
		}
	case int32:
		switch any(zero).(type) {
		case int8:
			if inputT < math.MinInt8 || inputT > math.MaxInt8 {
				return zero, errp1
			}
		case int16:
			if inputT < math.MinInt16 || inputT > math.MaxInt16 {
				return zero, errp1
			}
		// case int32, int64: // 可読性向上のためにコメントを残す
		case uint8:
			if inputT < 0 || inputT > math.MaxUint8 {
				return zero, errp1
			}
		case uint16:
			if inputT < 0 || inputT > math.MaxUint16 {
				return zero, errp1
			}
		case uint32, uint64:
			if inputT < 0 {
				return zero, errp3
			}
		}
	case int64:
		switch any(zero).(type) {
		case int8:
			if inputT < math.MinInt8 || inputT > math.MaxInt8 {
				return zero, errp1
			}
		case int16:
			if inputT < math.MinInt16 || inputT > math.MaxInt16 {
				return zero, errp1
			}
		case int32:
			if inputT < math.MinInt32 || inputT > math.MaxInt32 {
				return zero, errp1
			}
		// case int64: // 可読性向上のためにコメントを残す
		case uint8:
			if inputT < 0 || inputT > math.MaxUint8 {
				return zero, errp1
			}
		case uint16:
			if inputT < 0 || inputT > math.MaxUint16 {
				return zero, errp1
			}
		case uint32:
			if inputT < 0 || inputT > math.MaxUint32 {
				return zero, errp1
			}
		case uint64:
			if inputT < 0 {
				return zero, errp3
			}
		}
	case uint8:
		switch any(zero).(type) {
		case int8:
			if inputT > math.MaxInt8 {
				return zero, errp2
			}
			// case int16, int32, int64: // 可読性向上のためにコメントを残す
			// case uint8, uint16, uint32, uint64: // 可読性向上のためにコメントを残す
		}
	case uint16:
		switch any(zero).(type) {
		case int8:
			if inputT > math.MaxInt8 {
				return zero, errp2
			}
		case int16:
			if inputT > math.MaxInt16 {
				return zero, errp2
			}
		// case int32, int64: // 可読性向上のためにコメントを残す
		case uint8:
			if inputT > math.MaxUint8 {
				return zero, errp2
			}
		case uint16, uint32, uint64:
		}
	case uint32:
		switch any(zero).(type) {
		case int8:
			if inputT > math.MaxInt8 {
				return zero, errp2
			}
		case int16:
			if inputT > math.MaxInt16 {
				return zero, errp2
			}
		case int32:
			if inputT > math.MaxInt32 {
				return zero, errp2
			}
		// case int64: // 可読性向上のためにコメントを残す
		case uint8:
			if inputT > math.MaxUint8 {
				return zero, errp2
			}
		case uint16:
			if inputT > math.MaxUint16 {
				return zero, errp2
			}
			// case uint32, uint64: // 可読性向上のためにコメントを残す
		}
	case uint64:
		switch any(zero).(type) {
		case int8:
			if inputT > math.MaxInt8 {
				return zero, errp2
			}
		case int16:
			if inputT > math.MaxInt16 {
				return zero, errp2
			}
		case int32:
			if inputT > math.MaxInt32 {
				return zero, errp2
			}
		case int64:
			if inputT > math.MaxInt64 {
				return zero, errp2
			}
		case uint8:
			if inputT > math.MaxUint8 {
				return zero, errp2
			}
		case uint16:
			if inputT > math.MaxUint16 {
				return zero, errp2
			}
		case uint32:
			if inputT > math.MaxUint32 {
				return zero, errp2
			}
		case uint64:
		}
	}
	return U(input), nil
}

func main() {}
