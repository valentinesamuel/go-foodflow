package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	dbName := getEnvVar("ORDER_DB_NAME")

	fmt.Println("Hello Order! 🚀 => ", dbName)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading orderservice .env file📛")
	}
	return os.Getenv(key)
}
