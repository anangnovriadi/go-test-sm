package handlers

import (
	"github.com/gofiber/fiber/v2"

	"go-test-sm/database"
	"go-test-sm/models"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Locals("userID").(float64)
	var user models.User
	database.DB.First(&user, id)
	return c.JSON(user)
}
