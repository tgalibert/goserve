package network

import (
	"fmt"
	"goserve/internal/lib"
	"io"
	"net"
)

type Server struct {
	Port int
	corsEnable bool
	origins []string
}

func NewServer(port int) *Server {

	return &Server{
		Port: port,
		corsEnable: false,
		origins: make([]string, 0),
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
			fmt.Println(err.Error())
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) AllowCors(origins ...string) {
	s.corsEnable = true
	if len(origins) > 0 {
		s.origins = append(s.origins, origins...)
	} else {
		s.origins = append(s.origins, "*")
	}
}

func (s *Server) getAllowOrigin(clientOrigin string) string {
	for _, origin := range(s.origins) {
		if origin == "*" || origin == clientOrigin {
			if origin == "*" {
				return "*"
			} else {
				return clientOrigin
			}
		}
	}
	return ""
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	raw := io.Reader(conn)
	parser := lib.NewIoParser(raw)
	
	request := NewRequest(parser)
	err := request.Parse()
	if err != nil {
		fmt.Printf(err.Error())
	}

	res := NewResponse()
	if s.corsEnable {
		clientOrigin := request.Headers["origin"]
		matched := s.getAllowOrigin(clientOrigin)
		if matched != "" {
			res.AddCorsHeaders(matched)
		}
	}
	if request.Method == OPTIONS {
		res.Preflight()
		err = res.WriteTo(conn)
		return
	}
	err = res.WriteTo(conn)
}