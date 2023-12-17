package handlers

import (
	"quickshare/entities"
	"quickshare/services"

	"github.com/gofiber/fiber/v2"
)

type WebHandler interface {
	SendDashboard(c *fiber.Ctx) error
	SendLogin(c *fiber.Ctx) error
}

type WebHandlerImpl struct {
	UserService services.UserService
}

func (h *WebHandlerImpl) SendDashboard(c *fiber.Ctx) error {
	userID := c.Locals("userID")

	var user *entities.User

	if userID != nil {
		user, _ = h.UserService.FindUserByID(userID.(string))
	}

	return c.Render("pages/index", fiber.Map{
		"User": user,
	})
}

func (h *WebHandlerImpl) SendLogin(c *fiber.Ctx) error {
	return c.Render("pages/login", nil)
}
