package main

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
)

func Filter[V any](f func(V) bool, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

func Map[In, Out any](f func(In) Out, seq iter.Seq[In]) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for in := range seq {
			if !yield(f(in)) {
				return
			}
		}
	}
}

func main() {
	data := []int{1, 10, 31, 50}
	seq := slices.Values(data)
	seq = Filter(
		func(v int) bool {
			return v%2 == 0
		},
		seq,
	)
	seq1 := Map(
		func(v int) string {
			return ">>> " + strconv.Itoa(v)
		},
		seq,
	)
	for v := range seq1 {
		fmt.Println(v)
	}
}
