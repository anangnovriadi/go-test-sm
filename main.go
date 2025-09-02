package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"go-test-sm/database"
	"go-test-sm/models"
	"go-test-sm/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system env")
	}

	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})

	app := fiber.New()
	routes.SetupRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3010"
	}
	log.Fatal(app.Listen(":" + port))
}
