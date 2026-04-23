package main

import (
	"log"
	"payment-switch-simulator/internal/transport"
)

func main() {
	server := transport.NewTCPServer(":9090")
	log.Println("Payment Switch Simulator running on :9090")
	server.Start()
}
