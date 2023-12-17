package handlers

import "github.com/gofiber/fiber/v2"

type WebHandler interface {
	SendDashboard(c *fiber.Ctx) error
}

type WebHandlerImpl struct{}

func (h *WebHandlerImpl) SendDashboard(c *fiber.Ctx) error {
	return c.Render("pages/index", nil)
}
