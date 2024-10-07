package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Spec struct {
	Port int
}

func main() {
	spec := Spec{}
	err := envconfig.Process("", &spec)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", spec.Port))
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ReceiveMessage(conn)
	go SendMessage(conn)
	wg.Wait()
}

func ReceiveMessage(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("wait")
		incommingMessage := scanner.Text()
		fmt.Println("recv", incommingMessage)
	}
}

func SendMessage(conn net.Conn) {
	for {
		fmt.Println("send")
		time.Sleep(time.Second)
		conn.Write([]byte("foo1\n"))
	}
}
