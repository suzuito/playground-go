package main

import "fmt"

func main() {
	// for文
	for range seq { // seqはイテレーター
		// for文のブロック
		// yield関数の実行後、プログラムの制御はここへ移る
		fmt.Printf("+")
		// for文のブロックの実行後、プログラムの制御はyield関数が呼ばれた直後へ戻る。
	}
	fmt.Println("")
}

// イテレーター
func seq(yield func() bool) {
	// yield関数の実行後、プログラムの制御はfor文の1番上へ移る
	yield()
	// for文のブロックの実行後、プログラムの制御はここへ戻る。
	// 以下同様
	yield()
	yield()
}
