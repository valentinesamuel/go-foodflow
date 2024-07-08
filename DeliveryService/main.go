package main

import (
	"DeliveryService/db"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	db.ConnectDeliveryDB()
	dbName := getEnvVar("DELIVERY_DB_NAME")

	fmt.Println("Hello Delivery! 🚀=>", dbName)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading deliveryservice .env file📛")
	}
	return os.Getenv(key)
}
