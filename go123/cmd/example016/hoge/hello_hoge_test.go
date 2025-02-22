package hoge

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestHelloHoge001(t *testing.T) {
	fmt.Printf("TestHelloHoge001 started pid=%d\n", os.Getpid())
	for range 3 {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println()
	fmt.Println("TestHelloHoge001 finished")
}

func TestHelloHoge002(t *testing.T) {
	fmt.Printf("TestHelloHoge002 started pid=%d\n", os.Getpid())
	for range 3 {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println()
	fmt.Println("TestHelloHoge002 finished")
}
