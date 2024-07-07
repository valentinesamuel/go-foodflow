package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	dbName := getEnvVar("USER_DB_NAME")

	fmt.Println("Hello User! 🚀 =>", dbName)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading userservice .env file📛")
	}
	return os.Getenv(key)
}
