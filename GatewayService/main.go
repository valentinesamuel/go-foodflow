package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	dbName := getEnvVar("GATEWAY_DB_NAME")

	fmt.Println("Hello Gateway! 🚀 => ", dbName)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading gateway service .env file📛")
	}
	return os.Getenv(key)
}
