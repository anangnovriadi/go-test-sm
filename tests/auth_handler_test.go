package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go-test-sm/database"
	"go-test-sm/handlers"
	"go-test-sm/models"
)

func setupAuthApp() *fiber.App {
	app := fiber.New()
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	return app
}

func TestAuthRegister(t *testing.T) {
	database.ConnectTestDB()
	app := setupAuthApp()

	user := models.User{Email: "reguser@gmail.com", Password: "123456"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestAuthLogin(t *testing.T) {
	database.ConnectTestDB()
	app := setupAuthApp()

	registerBody, _ := json.Marshal(models.User{Email: "loginuser@gmail.com", Password: "123456"})
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(registerBody))
	req.Header.Set("Content-Type", "application/json")
	app.Test(req, -1)

	loginBody, _ := json.Marshal(models.User{Email: "loginuser@gmail.com", Password: "123456"})
	req2 := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(loginBody))
	req2.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req2, -1)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}
