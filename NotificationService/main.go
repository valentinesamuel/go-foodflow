package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	dbName := getEnvVar("NOTIFICATION_DB_NAME")

	fmt.Println("Hello Notification! => 🚀", dbName)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading notification .env file📛")
	}
	return os.Getenv(key)
}
