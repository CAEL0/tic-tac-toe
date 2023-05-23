package main

import (
	"log"

	"github.com/CAEL0/tic-tac-toe/server"
)

func main() {
	port := 8080
	if serveError := server.New(port).ListenAndServe(); serveError != nil {
		log.Fatalf("Failed toserver: %v", serveError)
	}
}
