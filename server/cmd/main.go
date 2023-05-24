package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/CAEL0/tic-tac-toe/server"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile("ENV_FILE")
	port := getIntVariable("PORT")
	if serveError := server.New(port, newDbClient()).ListenAndServe(); serveError != nil {
		log.Fatalf("Failed toserver: %v", serveError)
	}
}

func loadEnvFile(envFilename string) {
	envFilePath := os.Getenv(envFilename)
	if envFilePath == "" {
		log.Fatalf("Environment variable (%s) is required.", envFilename)
	}
	if loadError := godotenv.Load(envFilePath); loadError != nil {
		log.Fatalf("Failed to load env file: %v", loadError)
	}
}

func getIntVariable(name string) int {
	value, parseError := strconv.ParseInt(os.Getenv(name), 10, 64)
	if parseError != nil {
		log.Fatalf("Failed to parse to int: %v", parseError)
	}
	return int(value)
}

func getStringVariable(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Failed to load %s", name)
	}
	return value
}

func newDbClient() *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		getStringVariable("MYSQL_USERNAME"),
		getStringVariable("MYSQL_PASSWORD"),
		getStringVariable("MYSQL_URL"),
		getStringVariable("MYSQL_SCHEMA"),
	)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer db.Close()
	return db
}
