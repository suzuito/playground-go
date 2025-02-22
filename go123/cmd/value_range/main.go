package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("uint32", 0, math.MaxUint32)
	fmt.Println("int32", math.MinInt32, math.MaxInt32)
	fmt.Println("int64", math.MinInt64, math.MaxInt64)

	a := uint32(math.MaxUint32)
	b := int32(a)
	c := int64(a)
	fmt.Printf("uint32=%d int32=%d int64=%d\n", a, b, c)
}
