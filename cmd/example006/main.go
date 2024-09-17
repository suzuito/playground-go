package main

import (
	"fmt"
)

func main() {
	i := 0
	for range seq1 {
		fmt.Printf("main関数のfor文はじめ(i=%d)\n", i)
		if i == 2 {
			break
		}
		fmt.Printf("main関数のfor文おわり(i=%d)\n", i)
		i++
	}
}

func seq1(yield func() bool) {
	i := 0
	for range seq2 {
		fmt.Printf(" seq1関数のfor文はじめ(i=%d)\n", i)
		yield()
		fmt.Printf(" seq1関数のfor文おわり(i=%d)\n", i)
		i++
	}
}

func seq2(yield func() bool) {
	i := 0
	for range seq3 {
		fmt.Printf("  seq2関数のfor文はじめ(i=%d)\n", i)
		yield()
		fmt.Printf("  seq2関数のfor文おわり(i=%d)\n", i)
		i++
	}
}

func seq3(yield func() bool) {
	for i := range 10 {
		fmt.Printf("   seq3関数のfor文はじめ(i=%d)\n", i)
		yield()
		fmt.Printf("   seq3関数のfor文おわり(i=%d)\n", i)
		i++
	}
}
