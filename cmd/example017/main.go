package main

import (
	"log"
	"net/http"
)

func main() {
	cli := http.DefaultClient
	_, err := cli.Get("http://localhost:1111")
	if err != nil {
		log.Fatal(err)
	}
}
