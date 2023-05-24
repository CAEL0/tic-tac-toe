package main

import (
	"log"
	"os"
	"strconv"

	"github.com/CAEL0/tic-tac-toe/server"
	"github.com/joho/godotenv"
)

func main() {
	envFile := "ENV_FILE"
	envFilePath := os.Getenv(envFile)
	if envFilePath == "" {
		log.Fatalf("Environment variable (%s) is required.", envFile)
	}
	if loadErr := godotenv.Load(envFilePath); loadErr != nil {
		log.Fatalf("Failed to load env file: %v", loadErr)
	}
	port := getIntVariable("PORT")
	if serveError := server.New(port).ListenAndServe(); serveError != nil {
		log.Fatalf("Failed toserver: %v", serveError)
	}
}

func getIntVariable(name string) int {
	value, err := strconv.ParseInt(os.Getenv(name), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse to int: %v", err)
	}
	return int(value)
}
