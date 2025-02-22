package main

import "fmt"

func seq(yield func(k string, v int) bool) {
	// yield関数の引数として値を渡す
	yield("hoge", 10)
	yield("fuga", 30)
	yield("foo", 20)
}

func main() {
	for k, v := range seq {
		// yield関数の引数として渡された値は
		// kとvに格納されている
		fmt.Println(k, v)
	}
}
