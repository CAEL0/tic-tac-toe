package main

import (
	"log"
	"os"

	"github.com/CAEL0/tic-tac-toe/server"
)

func main() {
	envFile := "ENV_FILE"
	envFilePath := os.Getenv(envFile)
	if envFilePath == "" {
		log.Fatalf("Environment variable (%s) is required.", envFile)
	}
	port := 8080
	if serveError := server.New(port).ListenAndServe(); serveError != nil {
		log.Fatalf("Failed toserver: %v", serveError)
	}
}
