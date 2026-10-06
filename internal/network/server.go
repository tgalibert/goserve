package network

import (
	"errors"
	"fmt"
	"goserve/internal/constant"
	"goserve/internal/lib"
	"io"
	"net"
	"time"
)

type Server struct {
	Port int
	corsEnable bool
	origins []string
	Router *Router
	ReadTimeout	time.Duration
	WriteTimeout time.Duration
}

func NewServer(port int) *Server {

	return &Server{
		Port: port,
		corsEnable: false,
		origins: make([]string, 0),
		Router: NewRouter(),
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
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

func (s *Server) SetReadTimeout(d time.Duration) {
	s.ReadTimeout = d
}

func (s *Server) SetWriteTimeout(d time.Duration) {
	s.WriteTimeout = d
}

func (s *Server) SetTimeout(d time.Duration) {
	s.ReadTimeout = d
	s.WriteTimeout = d
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	raw := io.Reader(conn)
	parser := lib.NewIoParser(raw)
	
	req := NewRequest(parser)
	res := NewResponse()
	if s.ReadTimeout > 0 {
		conn.SetReadDeadline(time.Now().Add(s.ReadTimeout))
	}
	err := req.Parse()

	if s.WriteTimeout > 0 {
		conn.SetWriteDeadline(time.Now().Add(s.WriteTimeout))
	}
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			res.StatusCode = constant.StatusRequestTimeout
			res.SetBody([]byte("408 Request Timeout"))
		} else {
			res.StatusCode = constant.StatusBadRequest
			res.SetBody([]byte("400 Bad Request"))
		}
		res.WriteTo(conn)
		return
	}
	if s.corsEnable {
		clientOrigin := req.Headers["origin"]
		matched := s.getAllowOrigin(clientOrigin)
		if matched != "" {
			res.AddCorsHeaders(matched)
		}
	}
	if req.Method == OPTIONS {
		res.Preflight()
		err = res.WriteTo(conn)
		return
	}
	s.Router.Serve(req, res)

	err = res.WriteTo(conn)
}