package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port int
}

func New(port int) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) ListenAndServe() error {
	http.Handle("/", http.FileServer(http.Dir("client")))
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}
