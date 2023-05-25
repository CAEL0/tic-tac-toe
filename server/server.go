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
}

func New(port int, db *sql.DB) *Server {
	return &Server{
		port: port,
		db:   db,
	}
}

func (s *Server) ListenAndServe() error {
	hb := hub.New()
	go hb.Run()

	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWebsocket(hb, w, r)
	})
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}
