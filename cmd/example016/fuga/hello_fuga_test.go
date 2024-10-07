package fuga

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestHelloFuga(t *testing.T) {
	fmt.Printf("TestHelloFuga started pid=%d\n", os.Getpid())
	for range 3 {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println()
	fmt.Println("TestHelloFuga finished")
}
