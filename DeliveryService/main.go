package main

import (
	"DeliveryService/db"
	"DeliveryService/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-foodflow/utilities/models"
)

func main() {
	db.ConnectDeliveryDB()

	db.DeliveryDB.AutoMigrate(&models.DeliveryTracking{})

	server := gin.Default()

	routes.RegisterRoutes(server)

	dbName := getEnvVar("DELIVERY_DB_NAME")
	port := getEnvVar("PORT")

	fmt.Println("Hello Delivery! 🚀=>", dbName)

	server.Run(":" + port)
}

func getEnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading deliveryservice .env file📛")
	}
	return os.Getenv(key)
}
