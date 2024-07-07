package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	dbName := getEnvVar("MENU_DB_NAME")

	fmt.Println("Hello Menu!🚀 => ", dbName)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading menuservice .env file📛")
	}
	return os.Getenv(key)
}
