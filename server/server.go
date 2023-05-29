package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/CAEL0/tic-tac-toe/server/hub"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	port int
	db   *sql.DB
	hub  *hub.Hub
}

func New(port int, db *sql.DB) *Server {
	return &Server{
		port: port,
		db:   db,
		hub:  hub.New(),
	}
}

func (s *Server) ListenAndServe() error {
	go s.hub.Run()

	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWebsocket(s.hub, w, r)
	})
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}
