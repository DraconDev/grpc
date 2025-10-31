package main

import (
	"log"
	"os"

	"grpc-example/client"
	"grpc-example/server"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "client" {
		log.Println("Starting gRPC client...")
		client.StartClient()
	} else {
		log.Println("Starting gRPC server...")
		server.StartServer()
	}
}