package server

import (
	"io"
	"log/slog"
	"net"
)

type Server struct {
	host string
	port string
}

func New(host string, port string) *Server {
	return &Server{host: host, port: port}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.host+":"+s.port)
	if err != nil {
		panic(err)
	}

	slog.Info("server started", "host", s.host, "port", s.port)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("failed to accept connection", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	slog.Info("reading from connection", "host", conn.RemoteAddr().Network(), "port", conn.RemoteAddr().String())
	defer conn.Close()

	buffer := make([]byte, 1)
	var message []byte

	for {
		_, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				slog.Error("failed to read from connection", err)
			}
			break
		}

		message = append(message, buffer[0])

		if buffer[0] == '\n' {
			conn.Write([]byte(message))
			message = nil
		}
	}
}
