package main

import (
	"context"
	"fmt"
	"testing/synctest"
	"time"
)

func main() {
	main008()
}

func main001() {
	synctest.Run(func() {
		// Create a context.Context which is canceled after a timeout.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Wait just less than the timeout.
		time.Sleep(timeout - time.Nanosecond)
		synctest.Wait()
		fmt.Printf("before timeout: ctx.Err() = %v\n", ctx.Err())

		// Wait the rest of the way until the timeout.
		time.Sleep(time.Nanosecond)
		synctest.Wait()
		fmt.Printf("after timeout:  ctx.Err() = %v\n", ctx.Err())

	})
}

func main002() {
	synctest.Run(func() {
		fmt.Printf("before sleep: %d\n", time.Now().Unix())
		time.Sleep(5 * time.Second)
		fmt.Printf("after sleep1: %d\n", time.Now().Unix())
		fmt.Printf("after sleep2: %d\n", time.Now().Unix())
	})
}

func main003() {
	synctest.Run(func() {
		// バブル(Run関数の中)は、その中で起動している全てのゴルーチンが終わるのを待ってから返る。
		// バブルの中では偽の時間(synthetic time)が流れている。
		fmt.Printf("synthetic time=%s\n", time.Now().Format(time.RFC3339Nano))
	})
}

func main004() {
	synctest.Run(func() {
		// バブルが「永続的にブロックされた(durably blocked)」という概念を理解する。
		// 「バブルは永続的にブロックされた」とは、「バブル中の全てのゴルーチンが永続的にブロックされた」を指す。

		go func() {
			// チャンネルへデータを送信したが受信されない場合、ゴルーチンはブロックされます。
			var ch chan int = make(chan int)
			fmt.Println("ゴルーチン1はブロックされました")
			ch <- 1
		}()

		go func() {
			// time.Sleep関数を実行した場合、ゴルーチンはブロックされます。
			fmt.Println("ゴルーチン2はブロックされました")
			time.Sleep(10 * time.Second)
			fmt.Println("ゴルーチン2はis done")
		}()

		synctest.Wait()
		fmt.Println("hoge")
	})
}

func main005() {
	synctest.Run(func() {
		fmt.Printf("synthetic time=%s\n", time.Now().Format(time.RFC3339Nano)) // (1)
		go func() {
			time.Sleep(1 * time.Second)                                            // (2)
			fmt.Printf("synthetic time=%s\n", time.Now().Format(time.RFC3339Nano)) // (3)
		}()

		synctest.Wait() // (10)
		fmt.Println("synctest.Wait is called")

		// 実行過程の解説
		// (1) 偽の現在時刻の初期値(UTC 2000-01-01 00:00:00)を表示する。
		// (2) このゴルーチンはブロックされる。
		//     すると、バブル中の全てのゴルーチンがブロックされたため時間が進む。
		// (3) 偽の時間が1秒だけ進んだ後の時刻が表示される。
		// (10) バブル中の全てのゴルーチンがブロックされたらsynctest.Waitは返る。
		//      つまり、（2)に到達し、バブルがブロックされるとsynctest.Waitは返る。
	})
}

func main006() {
	synctest.Run(func() { // ゴルーチン1
		ch := make(chan int)
		go func() { // ゴルーチン2
			fmt.Println("a")
			ch <- 1 // (1)
			fmt.Println("b")
		}()

		synctest.Wait() // (2)
		fmt.Println("c")

		<-ch            // (3)
		synctest.Wait() // (4)
		fmt.Println("d")
	})
}

func main007() {
	synctest.Run(func() {
		ch := make(chan int)

		go func() { // 1つ目のゴルーチン
			ch <- 1 // (1)
			fmt.Println("a")
		}()

		go func() { // 2つ目のゴルーチン
			time.Sleep(1 * time.Second) // (2)
			fmt.Println("b")
		}()

		synctest.Wait() // (3)
		fmt.Println("c")

		<-ch // (4)
		fmt.Println("d")

		// (5)

		// 実行過程の解説
		// (1) このゴルーチンはブロックされる。
		// (2) このゴルーチンはブロックされる。
		//     すると、バブル中の全てのゴルーチンがブロックされた。
		// (3) バブル中の全てのゴルーチンがブロックされたらsynctest.Waitは返る。
		//     つまり、(1)と(2)の両方に到達し、バブルがブロックされるとsynctest.Waitは返る。
		// (4) (1)でブロックされたゴルーチンが、ブロック解除される。
		// (5) 1つ目のゴルーチンの実行が完了すると、再びバブルはブロックされる
		//     (2)でtime.Sleepにより2つ目のゴルーチンがブロックされており、バブル中の全てのゴルーチンがブロックされた状態であるため。
		//     そうすると時間が進む。
	})
}

func main008() {
	synctest.Run(func() { // ゴルーチン1
		ch := make(chan int)

		go func() { // ゴルーチン2
			ch <- 1 // (2)
			fmt.Println("a")
		}()

		go func() { // ゴルーチン3
			time.Sleep(1 * time.Second) // (3)
			fmt.Println("b")
		}()

		synctest.Wait() // (1)
		fmt.Println("c")

		synctest.Wait() // (4)
		fmt.Println("d")

		<-ch            // (5)
		synctest.Wait() // (6)
		fmt.Println("e")
	})
}

func main009() {
	synctest.Run(func() {
		fmt.Println("foo1")
		synctest.Wait()
		fmt.Println("foo2")
	})
}
