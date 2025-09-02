package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go-test-sm/database"
	"go-test-sm/handlers"
	"go-test-sm/models"
)

func setupUserApp() *fiber.App {
	app := fiber.New()
	app.Get("/user", func(c *fiber.Ctx) error {
		c.Locals("userID", float64(1))
		return handlers.GetUser(c)
	})
	return app
}

func TestUserGetUser(t *testing.T) {
	database.ConnectTestDB()
	database.DB.Create(&models.User{Email: "profile@gmail.com", Password: "123456"})

	app := setupUserApp()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	resp, _ := app.Test(req, -1)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}
