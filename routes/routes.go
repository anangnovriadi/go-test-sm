package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-test-sm/handlers"
	"go-test-sm/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)

	api.Get("/user", middleware.AuthRequired, handlers.GetUser)
}
