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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading payment service .env file📛")
	}
	return os.Getenv(key)

}
