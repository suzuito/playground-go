package main

import (
	"fmt"
)

func main() {
	items := []string{"hoge", "fuga"}
	for _, v := range items {
		fmt.Println(v)
	}
}
