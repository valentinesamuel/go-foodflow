package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DeliveryDB *gorm.DB

func ConnectDeliveryDB() {
	host := getEnvVar("DELIVERY_HOST")
	user := getEnvVar("DELIVERY_DB_USER")
	password := getEnvVar("DELIVERY_DB_PASSWORD")
	dbname := getEnvVar("DELIVERY_DB_NAME")
	port, _ := strconv.ParseInt(getEnvVar("DELIVERY_DB_PORT"), 10, 64)
	sslmode := getEnvVar("DELIVERY_SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, dbname, port, sslmode)

	fmt.Println("===>>>>>>>>>>>>>",dsn)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("ERR: Error connecting to the delivery database📛📛📛")
	}
	DeliveryDB = connection

	fmt.Println("Delivery Database connected🚀🚀🚀")
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading delivery .env file")
	}
	return os.Getenv(key)
}