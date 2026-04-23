package main

import (
	"log"
	"payment-switch-simulator/internal/async"
	"payment-switch-simulator/internal/store"
	"payment-switch-simulator/internal/transport"
	"time"
)

func main() {
	// init async queue
	queue := async.NewQueue(100)
	worker := async.NewWorker(queue)
	worker.Start()

	// inject queue (simple version)
	transport.SetQueue(queue)
	// start TCP server
	store := store.NewMemoryStore()
	go store.Cleanup(5 * time.Minute)
	server := transport.NewTCPServer(":9090")
	log.Println("Payment Switch Simulator running on :9090")
	server.Start()
}
