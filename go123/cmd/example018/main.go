package main

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	// status.Errorf関数は%w修飾子を使用できないためエラーをWrapできない
	errTarget := fmt.Errorf("hoge")
	err1 := status.Errorf(codes.Internal, "fuga : %w", errTarget)
	fmt.Println(errors.Is(err1, errTarget))
	err2 := fmt.Errorf("fuga : %w", errTarget)
	fmt.Println(errors.Is(err2, errTarget))
}
