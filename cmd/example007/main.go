package main

import (
	"bufio"
	"fmt"
	"iter"
	"log"
	"os"
	"strings"
)

func ReadLineFromFile(
	filePath string,
) (iter.Seq[string], error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return func(yield func(string) bool) {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}, nil
}

func Filter[V any](f func(V) bool, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

func main() {
	seq, err := ReadLineFromFile("cmd/example007/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	seq = Filter(
		func(line string) bool {
			return strings.HasPrefix(line, "5")
		},
		seq,
	)
	for line := range seq {
		fmt.Println(line)
	}
}
