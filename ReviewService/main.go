package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	dbName := getEnvVar("REVIEW_DB_NAME")

	fmt.Println("Hello Review! 🚀 => ", dbName)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading review service .env file📛")
	}
	return os.Getenv(key)
}
