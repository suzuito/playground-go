package main

import (
	"fmt"
	"math/big"
)

func fib(n int) func(yield func(v *big.Int) bool) {
	return func(yield func(v *big.Int) bool) {
		a, b := big.NewInt(0), big.NewInt(1)
		yield(a)
		yield(b)
		for range n {
			a.Add(a, b)
			yield(a)
			a, b = b, a
		}
	}
}

func main() {
	for v := range fib(1000) {
		fmt.Println(v)
	}
	fmt.Println()
}
