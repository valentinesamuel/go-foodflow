package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	dbName := getEnvVar("RESTAURANT_DB_NAME")

	fmt.Println("Hello Restaurant! 🚀 =>", dbName)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading restaurant service .env file📛")
	}
	return os.Getenv(key)
}
