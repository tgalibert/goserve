package main

import (
	"goserve/internal/network"
	"log"
)

func main() {
	server := network.NewServer(8002)

	err := server.Start()

	if err != nil {
		log.Fatal(err.Error())
	}
}