package main

func seq(yield func() bool) {
	yield()
	// yield関数がfalseを返した後で
	// yield関数を呼ぶとエラーとなる。
	yield()
}

func main() {
	for range seq {
		break
	}
}
