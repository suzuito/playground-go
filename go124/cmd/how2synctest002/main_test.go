package main

import (
	"context"
	"testing"
	"testing/synctest"
)

func TestAfterFunc(t *testing.T) {
	synctest.Run(func() {
		ctx, cancel := context.WithCancel(context.Background())

		funcCalled := false
		context.AfterFunc(ctx, func() { // ゴルーチン1
			funcCalled = true
		})

		synctest.Wait() // (1)
		if funcCalled { // (2)
			t.Fatalf("AfterFunc function called before context is canceled")
		}

		cancel() // (3)

		synctest.Wait() // (4)
		if !funcCalled {
			t.Fatalf("AfterFunc function not called after context is canceled")
		}
	})

	// 実行過程の解説
	// (1) この行が実行されている時点で、他に実行されているゴルーチンはない。
	//     要は、バブルは永続的にブロックされている状態であるため、synctest.Waitはすぐに制御を返す。
	// (2) ゴルーチン1はまだ実行されていないためfuncCalledはfalseである。
	// (3) cancelする。するとゴルーチン1が実行される。
	// (4) ゴルーチン1が実行中は、バブルは永続的にブロックされている状態ではないため、synctest.Waitは制御を返さない。
	//     ゴルーチン1が実行完了すると、バブルは永続的にブロックされている状態となるため、synctest.Waitは制御を返す。
}
