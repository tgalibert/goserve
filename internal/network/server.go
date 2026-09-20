package network

import (
	"fmt"
	"io"
	"net"
)

type Server struct {
	Port int
}

func NewServer(port int) *Server {

	return &Server{
		Port: port,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))

	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	raw := io.Reader(conn)
	request := NewRequest(raw)
	request.Parse()
}