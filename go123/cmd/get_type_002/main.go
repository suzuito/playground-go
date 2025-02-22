package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type intOverflowCheckersMap map[reflect.Kind]map[reflect.Kind]intOverflowChecker

var intOverflowCheckers intOverflowCheckersMap = map[reflect.Kind]map[reflect.Kind]intOverflowChecker{
	reflect.Int8: {
		reflect.Int8:  nil,
		reflect.Int16: nil,
		reflect.Int32: nil,
		reflect.Int64: nil,
		reflect.Uint8: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
		reflect.Uint16: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
		reflect.Uint32: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
		reflect.Uint64: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
	},
	reflect.Int16: {
		reflect.Int8: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: math.MinInt8, MaxI: math.MaxInt8,
		},
		reflect.Int16: nil,
		reflect.Int32: nil,
		reflect.Int64: nil,
		reflect.Uint8: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: 0, MaxI: math.MaxUint8,
		},
		reflect.Uint16: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
		reflect.Uint32: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
		reflect.Uint64: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
	},
	reflect.Int32: {
		reflect.Int8: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: math.MinInt8, MaxI: math.MaxInt8,
		},
		reflect.Int16: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: math.MinInt16, MaxI: math.MaxInt16,
		},
		reflect.Int32: nil,
		reflect.Int64: nil,
		reflect.Uint8: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: 0, MaxI: math.MaxUint8,
		},
		reflect.Uint16: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: 0, MaxI: math.MaxUint16,
		},
		reflect.Uint32: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
		reflect.Uint64: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
	},
	reflect.Int64: {
		reflect.Int8: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: math.MinInt8, MaxI: math.MaxInt8,
		},
		reflect.Int16: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: math.MinInt16, MaxI: math.MaxInt16,
		},
		reflect.Int32: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: math.MinInt32, MaxI: math.MaxInt32,
		},
		reflect.Int64: nil,
		reflect.Uint8: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: 0, MaxI: math.MaxUint8,
		},
		reflect.Uint16: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: 0, MaxI: math.MaxUint16,
		},
		reflect.Uint32: &intOverflowCheckerImpl{
			Type: MAX_AND_MIN, MinI: 0, MaxI: math.MaxUint32,
		},
		reflect.Uint64: &intOverflowCheckerImpl{
			Type: ONLY_MIN, MinI: 0,
		},
	},
	reflect.Uint8: {
		reflect.Int8: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt8,
		},
		reflect.Int16:  nil,
		reflect.Int32:  nil,
		reflect.Int64:  nil,
		reflect.Uint8:  nil,
		reflect.Uint16: nil,
		reflect.Uint32: nil,
		reflect.Uint64: nil,
	},
	reflect.Uint16: {
		reflect.Int8: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt8,
		},
		reflect.Int16: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt16,
		},
		reflect.Int32: nil,
		reflect.Int64: nil,
		reflect.Uint8: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxUint8,
		},
		reflect.Uint16: nil,
		reflect.Uint32: nil,
		reflect.Uint64: nil,
	},
	reflect.Uint32: {
		reflect.Int8: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt8,
		},
		reflect.Int16: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt16,
		},
		reflect.Int32: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt32,
		},
		reflect.Int64: nil,
		reflect.Uint8: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxUint8,
		},
		reflect.Uint16: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxUint16,
		},
		reflect.Uint32: nil,
		reflect.Uint64: nil,
	},
	reflect.Uint64: {
		reflect.Int8: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt8,
		},
		reflect.Int16: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt16,
		},
		reflect.Int32: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt32,
		},
		reflect.Int64: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxInt64,
		},
		reflect.Uint8: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxUint8,
		},
		reflect.Uint16: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxUint16,
		},
		reflect.Uint32: &intOverflowCheckerImpl{
			Type: ONLY_MAX, MaxUI: math.MaxUint32,
		},
		reflect.Uint64: nil,
	},
}

type intOverflowChecker interface {
	Check(v any) error
}

type intOverflowCheckerImpl struct {
	Type  intOverflowCheckerType
	MaxI  int64
	MinI  int64
	MaxUI uint64
	MinUI uint64
}

var errIntOverflow error = errors.New("integer overflow")

func (t *intOverflowCheckerImpl) Check(
	src any,
) error {
	srcValue := reflect.ValueOf(src)
	if srcValue.CanInt() {
		srcInt64 := srcValue.Int()
		switch t.Type {
		case MAX_AND_MIN:
			if srcInt64 < t.MinI || srcInt64 > t.MaxI {
				return errIntOverflow
			}
		case ONLY_MAX:
			if srcInt64 > t.MaxI {
				return errIntOverflow
			}
		case ONLY_MIN:
			if srcInt64 < t.MinI {
				return errIntOverflow
			}
		}
		return nil
	}

	if srcValue.CanUint() {
		srcUint64 := srcValue.Uint()
		switch t.Type {
		case MAX_AND_MIN:
			if srcUint64 < t.MinUI || srcUint64 > t.MaxUI {
				return errIntOverflow
			}
		case ONLY_MAX:
			if srcUint64 > t.MaxUI {
				return errIntOverflow
			}
		case ONLY_MIN:
			if srcUint64 < t.MinUI {
				return errIntOverflow
			}
		}
		return nil
	}

	return fmt.Errorf("not integer value")
}

type intOverflowCheckerType uint8

const (
	MAX_AND_MIN intOverflowCheckerType = iota
	ONLY_MAX
	ONLY_MIN
)

func checkIntOverflow(
	srcValue any,
	dstType reflect.Kind,
) error {
	srcType := reflect.TypeOf(srcValue).Kind()
	var checker intOverflowChecker
	if checkers, ok := intOverflowCheckers[srcType]; !ok {
		return fmt.Errorf("srcType is not intN type")
	} else {
		if checker, ok = checkers[dstType]; !ok {
			return fmt.Errorf("dstType is not intN type")
		}
	}

	if checker == nil {
		return nil
	}

	return checker.Check(srcValue)
}

type convertIntegerFunc[T, U integer] func(input T) (U, error)

func convertIntegerV1[T, U integer](input T) (U, error) {
	var zero U
	if err := checkIntOverflow(
		input,
		reflect.TypeOf(zero).Kind(),
	); err != nil {
		return zero, err
	}

	return U(input), nil
}
