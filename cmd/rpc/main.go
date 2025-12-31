package main

import (
	"log"

	"github.com/exanubes/typedef/internal/rpc"
)

func main() {
	router := rpc.NewRouter()
	server := rpc.NewServer(router)
	if err := server.Start(); err != nil {
		log.Fatal("RPC Server failed: ", err)
	}
}
