package main

import (
	"os"
	"time"
)

func main() {
	for range 10 {
		go func() {
			time.Sleep(5 * time.Second)
		}()
	}
	go func() {
		os.Exit(10)
	}()
	time.Sleep(5 * time.Second)
}
