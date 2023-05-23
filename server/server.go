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
	http.HandleFunc("/", index)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
