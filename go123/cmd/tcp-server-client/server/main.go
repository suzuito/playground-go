package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/kelseyhightower/envconfig"
)

type Spec struct {
	Port int
}

// TCPサーバープログラム
// 受信したメッセージをEchoするだけのプログラム
func main() {
	spec := Spec{}
	err := envconfig.Process("", &spec)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", spec.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		// コネクションを待つ
		fmt.Println("Listen", spec.Port)
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// コネクションを処理する
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("%s handle connection", conn.RemoteAddr().String())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("wait")
		incommingMessage := scanner.Text()
		fmt.Println("recv", incommingMessage)
		if _, err := fmt.Fprintf(conn, "You say -> %s\n", incommingMessage); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("%s close connection", conn.RemoteAddr().String())
}
