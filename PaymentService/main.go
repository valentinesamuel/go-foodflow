package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	dbName := getEnvVar("PAYMENT_DB_NAME")

	fmt.Println("Hello Payment! 🚀 => ", dbName)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading paymentsevice .env file📛")
	}
	return os.Getenv(key)
}
