package main

import "testing"

func traverseGoodStruct() uint16 {
	var arbitraryNum uint16

	var GoodStructArr = make([]Good, 1000000)

	for _, goodStruct := range GoodStructArr {
		arbitraryNum += goodStruct.siblings
	}

	return arbitraryNum
}

func traverseBadStruct() uint16 {
	var arbitraryNum uint16

	var BadStructArr = make([]Good, 1000000)

	for _, badStruct := range BadStructArr {
		arbitraryNum += badStruct.siblings
	}

	return arbitraryNum
}

func BenchmarkTraverseGoodStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		traverseGoodStruct()
	}
}

func BenchmarkTraverseBadStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		traverseBadStruct()
	}
}
