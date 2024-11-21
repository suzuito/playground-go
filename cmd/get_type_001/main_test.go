package main

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var anyint = 0

type testcase struct {
	name     string
	input    any
	expected any
	wantErr  bool
	errMsg   string
}

var testcases = []testcase{
	//
	// row1
	//
	// int8 int8
	{
		input:    int8(math.MinInt8),
		expected: int8(math.MinInt8),
		name:     "min",
	},
	{
		input:    int8(math.MaxInt8),
		expected: int8(math.MaxInt8),
		name:     "max",
	},
	// int8 int16
	{
		input:    int8(math.MinInt8),
		expected: int16(math.MinInt8),
		name:     "min",
	},
	{
		input:    int8(math.MaxInt8),
		expected: int16(math.MaxInt8),
		name:     "max",
	},
	// int8 int32
	{
		input:    int8(math.MinInt8),
		expected: int32(math.MinInt8),
		name:     "min",
	},
	{
		input:    int8(math.MaxInt8),
		expected: int32(math.MaxInt8),
		name:     "max",
	},
	// int8 int64
	{
		input:    int8(math.MinInt8),
		expected: int64(math.MinInt8),
		name:     "min",
	},
	{
		input:    int8(math.MaxInt8),
		expected: int64(math.MaxInt8),
		name:     "max",
	},
	// int8 uint8
	{
		input:    int8(-1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int8(0),
		expected: uint8(0),
		name:     "min",
	},
	{
		input:    int8(math.MaxInt8),
		expected: uint8(math.MaxInt8),
		name:     "max",
	},
	// int8 uint16
	{
		input:    int8(-1),
		expected: uint16(anyint),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int8(0),
		expected: uint16(0),
		name:     "min",
	},
	{
		input:    int8(math.MaxInt8),
		expected: uint16(math.MaxInt8),
		name:     "max",
	},
	// int8 uint32
	{
		input:    int8(-1),
		expected: uint32(anyint),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int8(0),
		expected: uint32(0),
		name:     "min",
	},
	{
		input:    int8(math.MaxInt8),
		expected: uint32(math.MaxInt8),
		name:     "max",
	},
	// int8 uint64
	{
		input:    int8(-1),
		expected: uint64(0),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int8(0),
		expected: uint64(0),
		name:     "min",
	},
	{
		input:    int8(math.MaxInt8),
		expected: uint64(math.MaxInt8),
		name:     "max",
	},
	//
	// row2
	//
	// int16 int8
	{
		input:    int16(math.MinInt8 - 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int16(math.MinInt8),
		expected: int8(math.MinInt8),
		name:     "min",
	},
	{
		input:    int16(math.MaxInt8),
		expected: int8(math.MaxInt8),
		name:     "max",
	},
	{
		input:    int16(math.MaxInt8 + 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int16 int16
	{
		input:    int16(math.MinInt16),
		expected: int16(math.MinInt16),
		name:     "min",
	},
	{
		input:    int16(math.MaxInt16),
		expected: int16(math.MaxInt16),
		name:     "max",
	},
	// int16 int32
	{
		input:    int16(math.MinInt16),
		expected: int32(math.MinInt16),
		name:     "min",
	},
	{
		input:    int16(math.MaxInt16),
		expected: int32(math.MaxInt16),
		name:     "max",
	},
	// int16 int64
	{
		input:    int16(math.MinInt16),
		expected: int64(math.MinInt16),
		name:     "min",
	},
	{
		input:    int16(math.MaxInt16),
		expected: int64(math.MaxInt16),
		name:     "max",
	},
	// int16 uint8
	{
		input:    int16(-1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int16(0),
		expected: uint8(0),
		name:     "min",
	},
	{
		input:    int16(math.MaxUint8),
		expected: uint8(math.MaxUint8),
		name:     "max",
	},
	{
		input:    int16(math.MaxUint8 + 1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int16 uint16
	{
		input:    int16(-1),
		expected: uint16(anyint),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int16(0),
		expected: uint16(0),
		name:     "min",
	},
	{
		input:    int16(math.MaxInt16),
		expected: uint16(math.MaxInt16),
		name:     "max",
	},
	// int16 uint32
	{
		input:    int16(-1),
		expected: uint32(anyint),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int16(0),
		expected: uint32(0),
		name:     "min",
	},
	{
		input:    int16(math.MaxInt16),
		expected: uint32(math.MaxInt16),
		name:     "max",
	},
	// int16 uint64
	{
		input:    int16(-1),
		expected: uint64(anyint),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int16(0),
		expected: uint64(0),
		name:     "min",
	},
	{
		input:    int16(math.MaxInt16),
		expected: uint64(math.MaxInt16),
		name:     "max",
	},
	//
	// row3
	//
	// int32 int8
	{
		input:    int32(math.MinInt8 - 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int32(math.MinInt8),
		expected: int8(math.MinInt8),
		name:     "min",
	},
	{
		input:    int32(math.MaxInt8),
		expected: int8(math.MaxInt8),
		name:     "max",
	},
	{
		input:    int32(math.MaxInt8 + 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int32 int16
	{
		input:    int32(math.MinInt16 - 1),
		expected: int16(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int32(math.MinInt16),
		expected: int16(math.MinInt16),
		name:     "min",
	},
	{
		input:    int32(math.MaxInt16),
		expected: int16(math.MaxInt16),
		name:     "max",
	},
	{
		input:    int32(math.MaxInt16 + 1),
		expected: int16(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int32 int32
	{
		input:    int32(math.MinInt32),
		expected: int32(math.MinInt32),
		name:     "min",
	},
	{
		input:    int32(math.MaxInt32),
		expected: int32(math.MaxInt32),
		name:     "max",
	},
	// int32 int64
	{
		input:    int32(math.MinInt32),
		expected: int64(math.MinInt32),
		name:     "min",
	},
	{
		input:    int32(math.MaxInt32),
		expected: int64(math.MaxInt32),
		name:     "max",
	},
	// int32 uint8
	{
		input:    int32(-1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int32(0),
		expected: uint8(0),
		name:     "min",
	},
	{
		input:    int32(math.MaxUint8),
		expected: uint8(math.MaxUint8),
		name:     "max",
	},
	{
		input:    int32(math.MaxUint8 + 1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int32 uint16
	{
		input:    int32(-1),
		expected: uint16(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int32(0),
		expected: uint16(0),
		name:     "min",
	},
	{
		input:    int32(math.MaxUint16),
		expected: uint16(math.MaxUint16),
		name:     "max",
	},
	{
		input:    int32(math.MaxUint16 + 1),
		expected: uint16(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int32 uint32
	{
		input:    int32(-1),
		expected: uint32(anyint),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int32(0),
		expected: uint32(0),
		name:     "min",
	},
	{
		input:    int32(math.MaxInt32),
		expected: uint32(math.MaxInt32),
		name:     "max",
	},
	// int32 uint64
	{
		input:    int32(-1),
		expected: uint64(anyint),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int32(0),
		expected: uint64(0),
		name:     "min",
	},
	{
		input:    int32(math.MaxInt32),
		expected: uint64(math.MaxInt32),
		name:     "max",
	},
	//
	// row4
	//
	// int64 int8
	{
		input:    int64(math.MinInt8 - 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int64(math.MinInt8),
		expected: int8(math.MinInt8),
		name:     "min",
	},
	{
		input:    int64(math.MaxInt8),
		expected: int8(math.MaxInt8),
		name:     "max",
	},
	{
		input:    int64(math.MaxInt8 + 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int64 int16
	{
		input:    int64(math.MinInt16 - 1),
		expected: int16(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int64(math.MinInt16),
		expected: int16(math.MinInt16),
		name:     "min",
	},
	{
		input:    int64(math.MaxInt16),
		expected: int16(math.MaxInt16),
		name:     "max",
	},
	{
		input:    int64(math.MaxInt16 + 1),
		expected: int16(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int64 int32
	{
		input:    int64(math.MinInt32 - 1),
		expected: int32(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int64(math.MinInt32),
		expected: int32(math.MinInt32),
		name:     "min",
	},
	{
		input:    int64(math.MaxInt32),
		expected: int32(math.MaxInt32),
		name:     "max",
	},
	{
		input:    int64(math.MaxInt32 + 1),
		expected: int32(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int64 int64
	{
		input:    int64(math.MinInt64),
		expected: int64(math.MinInt64),
		name:     "min",
	},
	{
		input:    int64(math.MaxInt64),
		expected: int64(math.MaxInt64),
		name:     "max",
	},
	// int64 uint8
	{
		input:    int64(-1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int64(0),
		expected: uint8(0),
		name:     "min",
	},
	{
		input:    int64(math.MaxUint8),
		expected: uint8(math.MaxUint8),
		name:     "max",
	},
	{
		input:    int64(math.MaxUint8 + 1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int64 uint16
	{
		input:    int64(-1),
		expected: uint16(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int64(0),
		expected: uint16(0),
		name:     "min",
	},
	{
		input:    int64(math.MaxUint16),
		expected: uint16(math.MaxUint16),
		name:     "max",
	},
	{
		input:    int64(math.MaxUint16 + 1),
		expected: uint16(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int64 uint32
	{
		input:    int64(-1),
		expected: uint32(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over min",
	},
	{
		input:    int64(0),
		expected: uint32(0),
		name:     "min",
	},
	{
		input:    int64(math.MaxInt32),
		expected: uint32(math.MaxInt32),
		name:     "max",
	},
	{
		input:    int64(math.MaxUint32 + 1),
		expected: uint32(anyint),
		wantErr:  true,
		errMsg:   "p1",
		name:     "over max",
	},
	// int64 uint64
	{
		input:    int64(-1),
		expected: uint64(anyint),
		wantErr:  true,
		errMsg:   "p3",
		name:     "over min",
	},
	{
		input:    int64(0),
		expected: uint64(0),
		name:     "min",
	},
	{
		input:    int64(math.MaxInt64),
		expected: uint64(math.MaxInt64),
		name:     "max",
	},
	//
	// row5
	//
	// uint8 int8
	{
		input:    uint8(0),
		expected: int8(0),
		name:     "min",
	},
	{
		input:    uint8(math.MaxInt8),
		expected: int8(math.MaxInt8),
		name:     "max",
	},
	{
		input:    uint8(math.MaxInt8 + 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint8 int16
	{
		input:    uint8(0),
		expected: int16(0),
		name:     "min",
	},
	{
		input:    uint8(math.MaxUint8),
		expected: int16(math.MaxUint8),
		name:     "max",
	},
	// uint8 int32
	{
		input:    uint8(0),
		expected: int32(0),
		name:     "min",
	},
	{
		input:    uint8(math.MaxUint8),
		expected: int32(math.MaxUint8),
		name:     "max",
	},
	// uint8 int64
	{
		input:    uint8(0),
		expected: int64(0),
		name:     "min",
	},
	{
		input:    uint8(math.MaxUint8),
		expected: int64(math.MaxUint8),
		name:     "max",
	},
	// uint8 uint8
	{
		input:    uint8(0),
		expected: uint8(0),
		name:     "min",
	},
	{
		input:    uint8(math.MaxUint8),
		expected: uint8(math.MaxUint8),
		name:     "max",
	},
	// uint8 uint16
	{
		input:    uint8(0),
		expected: uint16(0),
		name:     "min",
	},
	{
		input:    uint8(math.MaxUint8),
		expected: uint16(math.MaxUint8),
		name:     "max",
	},
	// uint8 uint32
	{
		input:    uint8(0),
		expected: uint32(0),
		name:     "min",
	},
	{
		input:    uint8(math.MaxUint8),
		expected: uint32(math.MaxUint8),
		name:     "max",
	},
	// uint8 uint64
	{
		input:    uint8(0),
		expected: uint64(0),
		name:     "min",
	},
	{
		input:    uint8(math.MaxUint8),
		expected: uint64(math.MaxUint8),
		name:     "max",
	},
	//
	// row6
	//
	// uint16 int8
	{
		input:    uint16(0),
		expected: int8(0),
		name:     "min",
	},
	{
		input:    uint16(math.MaxInt8),
		expected: int8(math.MaxInt8),
		name:     "max",
	},
	{
		input:    uint16(math.MaxInt8 + 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint16 int16
	{
		input:    uint16(0),
		expected: int16(0),
		name:     "min",
	},
	{
		input:    uint16(math.MaxInt16),
		expected: int16(math.MaxInt16),
		name:     "max",
	},
	{
		input:    uint16(math.MaxInt16 + 1),
		expected: int16(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint16 int32
	{
		input:    uint16(0),
		expected: int32(0),
		name:     "min",
	},
	{
		input:    uint16(math.MaxUint16),
		expected: int32(math.MaxUint16),
		name:     "max",
	},
	// uint16 int64
	{
		input:    uint16(0),
		expected: int64(0),
		name:     "min",
	},
	{
		input:    uint16(math.MaxUint16),
		expected: int64(math.MaxUint16),
		name:     "max",
	},
	// uint16 uint8
	{
		input:    uint16(0),
		expected: uint8(0),
		name:     "min",
	},
	{
		input:    uint16(math.MaxUint8),
		expected: uint8(math.MaxUint8),
		name:     "max",
	},
	{
		input:    uint16(math.MaxUint8 + 1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "max",
	},
	// uint16 uint16
	{
		input:    uint16(0),
		expected: uint16(0),
		name:     "min",
	},
	{
		input:    uint16(math.MaxUint16),
		expected: uint16(math.MaxUint16),
		name:     "max",
	},
	// uint16 uint32
	{
		input:    uint16(0),
		expected: uint32(0),
		name:     "min",
	},
	{
		input:    uint16(math.MaxUint16),
		expected: uint32(math.MaxUint16),
		name:     "max",
	},
	// uint16 uint64
	{
		input:    uint16(0),
		expected: uint64(0),
		name:     "min",
	},
	{
		input:    uint16(math.MaxUint16),
		expected: uint64(math.MaxUint16),
		name:     "max",
	},
	//
	// row7
	//
	// uint32 int8
	{
		input:    uint32(0),
		expected: int8(0),
		name:     "min",
	},
	{
		input:    uint32(math.MaxInt8),
		expected: int8(math.MaxInt8),
		name:     "max",
	},
	{
		input:    uint32(math.MaxInt8 + 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint32 int16
	{
		input:    uint32(0),
		expected: int16(0),
		name:     "min",
	},
	{
		input:    uint32(math.MaxInt16),
		expected: int16(math.MaxInt16),
		name:     "max",
	},
	{
		input:    uint32(math.MaxInt16 + 1),
		expected: int16(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint32 int32
	{
		input:    uint32(0),
		expected: int32(0),
		name:     "min",
	},
	{
		input:    uint32(math.MaxInt32),
		expected: int32(math.MaxInt32),
		name:     "max",
	},
	{
		input:    uint32(math.MaxInt32 + 1),
		expected: int32(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint32 int64
	{
		input:    uint32(0),
		expected: int64(0),
		name:     "min",
	},
	{
		input:    uint32(math.MaxUint32),
		expected: int64(math.MaxUint32),
		name:     "max",
	},
	// uint32 uint8
	{
		input:    uint32(0),
		expected: uint8(0),
		name:     "min",
	},
	{
		input:    uint32(math.MaxUint8),
		expected: uint8(math.MaxUint8),
		name:     "max",
	},
	{
		input:    uint32(math.MaxUint8 + 1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "max",
	},
	// uint32 uint16
	{
		input:    uint32(0),
		expected: uint16(0),
		name:     "min",
	},
	{
		input:    uint32(math.MaxUint16),
		expected: uint16(math.MaxUint16),
		name:     "max",
	},
	{
		input:    uint32(math.MaxUint16 + 1),
		expected: uint16(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "max",
	},
	// uint32 uint32
	{
		input:    uint32(0),
		expected: uint32(0),
		name:     "min",
	},
	{
		input:    uint32(math.MaxUint32),
		expected: uint32(math.MaxUint32),
		name:     "max",
	},
	// uint32 uint64
	{
		input:    uint32(0),
		expected: uint64(0),
		name:     "min",
	},
	{
		input:    uint32(math.MaxUint32),
		expected: uint64(math.MaxUint32),
		name:     "max",
	},
	//
	// row8
	//
	// uint64 int8
	{
		input:    uint64(0),
		expected: int8(0),
		name:     "min",
	},
	{
		input:    uint64(math.MaxInt8),
		expected: int8(math.MaxInt8),
		name:     "max",
	},
	{
		input:    uint64(math.MaxInt8 + 1),
		expected: int8(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint64 int16
	{
		input:    uint64(0),
		expected: int16(0),
		name:     "min",
	},
	{
		input:    uint64(math.MaxInt16),
		expected: int16(math.MaxInt16),
		name:     "max",
	},
	{
		input:    uint64(math.MaxInt16 + 1),
		expected: int16(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint64 int32
	{
		input:    uint64(0),
		expected: int32(0),
		name:     "min",
	},
	{
		input:    uint64(math.MaxInt32),
		expected: int32(math.MaxInt32),
		name:     "max",
	},
	{
		input:    uint64(math.MaxInt32 + 1),
		expected: int32(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint64 int64
	{
		input:    uint64(0),
		expected: int64(0),
		name:     "min",
	},
	{
		input:    uint64(math.MaxInt64),
		expected: int64(math.MaxInt64),
		name:     "max",
	},
	{
		input:    uint64(math.MaxInt64 + 1),
		expected: int64(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "over max",
	},
	// uint64 uint8
	{
		input:    uint64(0),
		expected: uint8(0),
		name:     "min",
	},
	{
		input:    uint64(math.MaxUint8),
		expected: uint8(math.MaxUint8),
		name:     "max",
	},
	{
		input:    uint64(math.MaxUint8 + 1),
		expected: uint8(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "max",
	},
	// uint64 uint16
	{
		input:    uint64(0),
		expected: uint16(0),
		name:     "min",
	},
	{
		input:    uint64(math.MaxUint16),
		expected: uint16(math.MaxUint16),
		name:     "max",
	},
	{
		input:    uint64(math.MaxUint16 + 1),
		expected: uint16(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "max",
	},
	// uint64 uint32
	{
		input:    uint64(0),
		expected: uint32(0),
		name:     "min",
	},
	{
		input:    uint64(math.MaxUint32),
		expected: uint32(math.MaxUint32),
		name:     "max",
	},
	{
		input:    uint64(math.MaxUint32 + 1),
		expected: uint32(anyint),
		wantErr:  true,
		errMsg:   "p2",
		name:     "max",
	},
	// uint64 uint64
	{
		input:    uint64(0),
		expected: uint64(0),
		name:     "min",
	},
	{
		input:    uint64(math.MaxUint64),
		expected: uint64(math.MaxUint64),
		name:     "max",
	},
}

type convertIntegerFuncs struct {
	int8int8     convertIntegerFunc[int8, int8]
	int8int16    convertIntegerFunc[int8, int16]
	int8int32    convertIntegerFunc[int8, int32]
	int8int64    convertIntegerFunc[int8, int64]
	int8uint8    convertIntegerFunc[int8, uint8]
	int8uint16   convertIntegerFunc[int8, uint16]
	int8uint32   convertIntegerFunc[int8, uint32]
	int8uint64   convertIntegerFunc[int8, uint64]
	int16int8    convertIntegerFunc[int16, int8]
	int16int16   convertIntegerFunc[int16, int16]
	int16int32   convertIntegerFunc[int16, int32]
	int16int64   convertIntegerFunc[int16, int64]
	int16uint8   convertIntegerFunc[int16, uint8]
	int16uint16  convertIntegerFunc[int16, uint16]
	int16uint32  convertIntegerFunc[int16, uint32]
	int16uint64  convertIntegerFunc[int16, uint64]
	int32int8    convertIntegerFunc[int32, int8]
	int32int16   convertIntegerFunc[int32, int16]
	int32int32   convertIntegerFunc[int32, int32]
	int32int64   convertIntegerFunc[int32, int64]
	int32uint8   convertIntegerFunc[int32, uint8]
	int32uint16  convertIntegerFunc[int32, uint16]
	int32uint32  convertIntegerFunc[int32, uint32]
	int32uint64  convertIntegerFunc[int32, uint64]
	int64int8    convertIntegerFunc[int64, int8]
	int64int16   convertIntegerFunc[int64, int16]
	int64int32   convertIntegerFunc[int64, int32]
	int64int64   convertIntegerFunc[int64, int64]
	int64uint8   convertIntegerFunc[int64, uint8]
	int64uint16  convertIntegerFunc[int64, uint16]
	int64uint32  convertIntegerFunc[int64, uint32]
	int64uint64  convertIntegerFunc[int64, uint64]
	uint8int8    convertIntegerFunc[uint8, int8]
	uint8int16   convertIntegerFunc[uint8, int16]
	uint8int32   convertIntegerFunc[uint8, int32]
	uint8int64   convertIntegerFunc[uint8, int64]
	uint8uint8   convertIntegerFunc[uint8, uint8]
	uint8uint16  convertIntegerFunc[uint8, uint16]
	uint8uint32  convertIntegerFunc[uint8, uint32]
	uint8uint64  convertIntegerFunc[uint8, uint64]
	uint16int8   convertIntegerFunc[uint16, int8]
	uint16int16  convertIntegerFunc[uint16, int16]
	uint16int32  convertIntegerFunc[uint16, int32]
	uint16int64  convertIntegerFunc[uint16, int64]
	uint16uint8  convertIntegerFunc[uint16, uint8]
	uint16uint16 convertIntegerFunc[uint16, uint16]
	uint16uint32 convertIntegerFunc[uint16, uint32]
	uint16uint64 convertIntegerFunc[uint16, uint64]
	uint32int8   convertIntegerFunc[uint32, int8]
	uint32int16  convertIntegerFunc[uint32, int16]
	uint32int32  convertIntegerFunc[uint32, int32]
	uint32int64  convertIntegerFunc[uint32, int64]
	uint32uint8  convertIntegerFunc[uint32, uint8]
	uint32uint16 convertIntegerFunc[uint32, uint16]
	uint32uint32 convertIntegerFunc[uint32, uint32]
	uint32uint64 convertIntegerFunc[uint32, uint64]
	uint64int8   convertIntegerFunc[uint64, int8]
	uint64int16  convertIntegerFunc[uint64, int16]
	uint64int32  convertIntegerFunc[uint64, int32]
	uint64int64  convertIntegerFunc[uint64, int64]
	uint64uint8  convertIntegerFunc[uint64, uint8]
	uint64uint16 convertIntegerFunc[uint64, uint16]
	uint64uint32 convertIntegerFunc[uint64, uint32]
	uint64uint64 convertIntegerFunc[uint64, uint64]
}

func (funcs *convertIntegerFuncs) call(
	input any,
	expected any,
) (actual any, err error) {
	switch inputT := input.(type) {
	case int8:
		switch expected.(type) {
		case int8:
			actual, err = funcs.int8int8(inputT)
		case int16:
			actual, err = funcs.int8int16(inputT)
		case int32:
			actual, err = funcs.int8int32(inputT)
		case int64:
			actual, err = funcs.int8int64(inputT)
		case uint8:
			actual, err = funcs.int8uint8(inputT)
		case uint16:
			actual, err = funcs.int8uint16(inputT)
		case uint32:
			actual, err = funcs.int8uint32(inputT)
		case uint64:
			actual, err = funcs.int8uint64(inputT)
		}
	case int16:
		switch expected.(type) {
		case int8:
			actual, err = funcs.int16int8(inputT)
		case int16:
			actual, err = funcs.int16int16(inputT)
		case int32:
			actual, err = funcs.int16int32(inputT)
		case int64:
			actual, err = funcs.int16int64(inputT)
		case uint8:
			actual, err = funcs.int16uint8(inputT)
		case uint16:
			actual, err = funcs.int16uint16(inputT)
		case uint32:
			actual, err = funcs.int16uint32(inputT)
		case uint64:
			actual, err = funcs.int16uint64(inputT)
		}
	case int32:
		switch expected.(type) {
		case int8:
			actual, err = funcs.int32int8(inputT)
		case int16:
			actual, err = funcs.int32int16(inputT)
		case int32:
			actual, err = funcs.int32int32(inputT)
		case int64:
			actual, err = funcs.int32int64(inputT)
		case uint8:
			actual, err = funcs.int32uint8(inputT)
		case uint16:
			actual, err = funcs.int32uint16(inputT)
		case uint32:
			actual, err = funcs.int32uint32(inputT)
		case uint64:
			actual, err = funcs.int32uint64(inputT)
		}
	case int64:
		switch expected.(type) {
		case int8:
			actual, err = funcs.int64int8(inputT)
		case int16:
			actual, err = funcs.int64int16(inputT)
		case int32:
			actual, err = funcs.int64int32(inputT)
		case int64:
			actual, err = funcs.int64int64(inputT)
		case uint8:
			actual, err = funcs.int64uint8(inputT)
		case uint16:
			actual, err = funcs.int64uint16(inputT)
		case uint32:
			actual, err = funcs.int64uint32(inputT)
		case uint64:
			actual, err = funcs.int64uint64(inputT)
		}
	case uint8:
		switch expected.(type) {
		case int8:
			actual, err = funcs.uint8int8(inputT)
		case int16:
			actual, err = funcs.uint8int16(inputT)
		case int32:
			actual, err = funcs.uint8int32(inputT)
		case int64:
			actual, err = funcs.uint8int64(inputT)
		case uint8:
			actual, err = funcs.uint8uint8(inputT)
		case uint16:
			actual, err = funcs.uint8uint16(inputT)
		case uint32:
			actual, err = funcs.uint8uint32(inputT)
		case uint64:
			actual, err = funcs.uint8uint64(inputT)
		}
	case uint16:
		switch expected.(type) {
		case int8:
			actual, err = funcs.uint16int8(inputT)
		case int16:
			actual, err = funcs.uint16int16(inputT)
		case int32:
			actual, err = funcs.uint16int32(inputT)
		case int64:
			actual, err = funcs.uint16int64(inputT)
		case uint8:
			actual, err = funcs.uint16uint8(inputT)
		case uint16:
			actual, err = funcs.uint16uint16(inputT)
		case uint32:
			actual, err = funcs.uint16uint32(inputT)
		case uint64:
			actual, err = funcs.uint16uint64(inputT)
		}
	case uint32:
		switch expected.(type) {
		case int8:
			actual, err = funcs.uint32int8(inputT)
		case int16:
			actual, err = funcs.uint32int16(inputT)
		case int32:
			actual, err = funcs.uint32int32(inputT)
		case int64:
			actual, err = funcs.uint32int64(inputT)
		case uint8:
			actual, err = funcs.uint32uint8(inputT)
		case uint16:
			actual, err = funcs.uint32uint16(inputT)
		case uint32:
			actual, err = funcs.uint32uint32(inputT)
		case uint64:
			actual, err = funcs.uint32uint64(inputT)
		}
	case uint64:
		switch expected.(type) {
		case int8:
			actual, err = funcs.uint64int8(inputT)
		case int16:
			actual, err = funcs.uint64int16(inputT)
		case int32:
			actual, err = funcs.uint64int32(inputT)
		case int64:
			actual, err = funcs.uint64int64(inputT)
		case uint8:
			actual, err = funcs.uint64uint8(inputT)
		case uint16:
			actual, err = funcs.uint64uint16(inputT)
		case uint32:
			actual, err = funcs.uint64uint32(inputT)
		case uint64:
			actual, err = funcs.uint64uint64(inputT)
		}
	}
	return actual, err
}

func testConvertIntegerRun(
	t *testing.T,
	tc testcase,
	funcs *convertIntegerFuncs,
) {
	actual, err := funcs.call(tc.input, tc.expected)
	if tc.wantErr {
		require.EqualError(t, err, tc.errMsg)
	} else {
		require.NoError(t, err)

		require.Equal(t, tc.expected, actual)
		require.IsType(t, tc.expected, actual)
	}
}

var convertIntegerV1Funcs = convertIntegerFuncs{
	int8int8:     convertIntegerV1[int8, int8],
	int8int16:    convertIntegerV1[int8, int16],
	int8int32:    convertIntegerV1[int8, int32],
	int8int64:    convertIntegerV1[int8, int64],
	int8uint8:    convertIntegerV1[int8, uint8],
	int8uint16:   convertIntegerV1[int8, uint16],
	int8uint32:   convertIntegerV1[int8, uint32],
	int8uint64:   convertIntegerV1[int8, uint64],
	int16int8:    convertIntegerV1[int16, int8],
	int16int16:   convertIntegerV1[int16, int16],
	int16int32:   convertIntegerV1[int16, int32],
	int16int64:   convertIntegerV1[int16, int64],
	int16uint8:   convertIntegerV1[int16, uint8],
	int16uint16:  convertIntegerV1[int16, uint16],
	int16uint32:  convertIntegerV1[int16, uint32],
	int16uint64:  convertIntegerV1[int16, uint64],
	int32int8:    convertIntegerV1[int32, int8],
	int32int16:   convertIntegerV1[int32, int16],
	int32int32:   convertIntegerV1[int32, int32],
	int32int64:   convertIntegerV1[int32, int64],
	int32uint8:   convertIntegerV1[int32, uint8],
	int32uint16:  convertIntegerV1[int32, uint16],
	int32uint32:  convertIntegerV1[int32, uint32],
	int32uint64:  convertIntegerV1[int32, uint64],
	int64int8:    convertIntegerV1[int64, int8],
	int64int16:   convertIntegerV1[int64, int16],
	int64int32:   convertIntegerV1[int64, int32],
	int64int64:   convertIntegerV1[int64, int64],
	int64uint8:   convertIntegerV1[int64, uint8],
	int64uint16:  convertIntegerV1[int64, uint16],
	int64uint32:  convertIntegerV1[int64, uint32],
	int64uint64:  convertIntegerV1[int64, uint64],
	uint8int8:    convertIntegerV1[uint8, int8],
	uint8int16:   convertIntegerV1[uint8, int16],
	uint8int32:   convertIntegerV1[uint8, int32],
	uint8int64:   convertIntegerV1[uint8, int64],
	uint8uint8:   convertIntegerV1[uint8, uint8],
	uint8uint16:  convertIntegerV1[uint8, uint16],
	uint8uint32:  convertIntegerV1[uint8, uint32],
	uint8uint64:  convertIntegerV1[uint8, uint64],
	uint16int8:   convertIntegerV1[uint16, int8],
	uint16int16:  convertIntegerV1[uint16, int16],
	uint16int32:  convertIntegerV1[uint16, int32],
	uint16int64:  convertIntegerV1[uint16, int64],
	uint16uint8:  convertIntegerV1[uint16, uint8],
	uint16uint16: convertIntegerV1[uint16, uint16],
	uint16uint32: convertIntegerV1[uint16, uint32],
	uint16uint64: convertIntegerV1[uint16, uint64],
	uint32int8:   convertIntegerV1[uint32, int8],
	uint32int16:  convertIntegerV1[uint32, int16],
	uint32int32:  convertIntegerV1[uint32, int32],
	uint32int64:  convertIntegerV1[uint32, int64],
	uint32uint8:  convertIntegerV1[uint32, uint8],
	uint32uint16: convertIntegerV1[uint32, uint16],
	uint32uint32: convertIntegerV1[uint32, uint32],
	uint32uint64: convertIntegerV1[uint32, uint64],
	uint64int8:   convertIntegerV1[uint64, int8],
	uint64int16:  convertIntegerV1[uint64, int16],
	uint64int32:  convertIntegerV1[uint64, int32],
	uint64int64:  convertIntegerV1[uint64, int64],
	uint64uint8:  convertIntegerV1[uint64, uint8],
	uint64uint16: convertIntegerV1[uint64, uint16],
	uint64uint32: convertIntegerV1[uint64, uint32],
	uint64uint64: convertIntegerV1[uint64, uint64],
}

func TestConvertIntegerV1(t *testing.T) {
	for _, c := range testcases {
		t.Run(
			reflect.TypeOf(c.input).Name()+reflect.TypeOf(c.expected).Name()+c.name,
			func(t *testing.T) {
				testConvertIntegerRun(t, c, &convertIntegerV1Funcs)
			},
		)
	}
}

func BenchmarkConvertIntegerV1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			convertIntegerV1Funcs.call(c.input, c.expected)
		}
	}
}
