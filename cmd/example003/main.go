package main

import (
	"fmt"
	"time"
)

func seq(yield func() bool) {
	for {
		// yield関数がfalseを返した
		// = for文を抜けた
		// yield関数がfalseを返した後
		// イテレーターを直ちに終了させるべき
		if !yield() {
			return
		}
	}
}

func main() {
	start := time.Now()
	for range seq {
		fmt.Printf("+")
		// 100マイクロ秒後にbreak
		if time.Now().Sub(start) > time.Duration(100)*time.Microsecond {
			break
		}
	}
}
