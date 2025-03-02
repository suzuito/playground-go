package main

import (
	"context"
	"fmt"
	"testing/synctest"
)

func main() {
	synctest.Run(func() {
		// Create a context.Context which can be canceled.
		ctx, cancel := context.WithCancel(context.Background())

		// context.AfterFunc registers a function to be called
		// when a context is canceled.
		afterFuncCalled := false
		context.AfterFunc(ctx, func() {
			afterFuncCalled = true
		})

		// The context has not been canceled, so the AfterFunc is not called.
		synctest.Wait()
		fmt.Printf("before context is canceled: afterFuncCalled=%v\n", afterFuncCalled)

		// Cancel the context and wait for the AfterFunc to finish executing.
		// Verify that the AfterFunc ran.
		cancel()
		synctest.Wait()
		fmt.Printf("after context is canceled:  afterFuncCalled=%v\n", afterFuncCalled)

	})
}
