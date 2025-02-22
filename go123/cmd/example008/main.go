package main

import "fmt"

func main() {
	for range seq {
	}
}

func seq(yield func() bool) {
	i := 0
	for yield() {
		fmt.Printf(" seq関数のfor文はじめ(i=%d)\n", i)
		fmt.Printf(" seq関数のfor文おわり(i=%d)\n", i)
		i++
	}
}
