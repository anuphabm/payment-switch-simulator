package transport

import (
	"bufio"
	"log"
	"net"
	"payment-switch-simulator/internal/handler"
	"time"
)

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

		response := handler.HandleMessage(msg)
		conn.Write([]byte(response + "\n"))
	}
}
