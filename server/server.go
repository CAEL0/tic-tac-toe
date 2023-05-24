package server

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	port int
	db   *sql.DB
}

func New(port int, db *sql.DB) *Server {
	return &Server{
		port: port,
		db:   db,
	}
}

func (s *Server) ListenAndServe() error {
	http.Handle("/", http.FileServer(http.Dir("client")))
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}
