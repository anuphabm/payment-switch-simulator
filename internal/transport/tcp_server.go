package transport

import (
	"bufio"
	"log"
	"net"
	"payment-switch-simulator/internal/async"
	"payment-switch-simulator/internal/handler"
	"payment-switch-simulator/internal/resilience"
	"time"
)

var limiter = resilience.NewRateLimiter(10, 5, time.Second)

var globalQueue *async.Queue

func SetQueue(q *async.Queue) {
	globalQueue = q
}

type TCPServer struct {
	addr string
}

func NewTCPServer(addr string) *TCPServer {
	return &TCPServer{addr: addr}
}

func (s *TCPServer) Start() {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Println("read error:", err)
			return
		}

		// 🔥 rate limiting
		if !limiter.Allow() {
			conn.Write([]byte(`{"error":"rate limit exceeded"}` + "\n"))
			continue
		}

		response := handler.HandleMessage(msg, globalQueue)
		conn.Write([]byte(response + "\n"))
	}
}
