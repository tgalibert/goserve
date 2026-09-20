package main

import (
	"goserve/internal/network"
	"log"
)

func main() {
	server := network.NewServer(8000)

	err := server.Start()

	if err != nil {
		log.Fatal(err.Error())
	}
}