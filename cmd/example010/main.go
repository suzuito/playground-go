package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(">>> ", string(line))
	}
	// bufioパッケージのScannerというAPIを用いることで
	// かなり簡潔に記述できるようになる。
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { // 条件式
		fmt.Println(">>> ", scanner.Text())
	}
}
