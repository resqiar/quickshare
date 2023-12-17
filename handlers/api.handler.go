package handlers

import (
	"quickshare/services"

	"github.com/gofiber/fiber/v2"
)

type APIHandler interface {
	ParseMD(c *fiber.Ctx) error
}

type APIHandlerImpl struct {
	UtilService services.UtilService
}

func (h *APIHandlerImpl) ParseMD(c *fiber.Ctx) error {
	payload := c.Body()

	parsed := h.UtilService.ParseMD(payload)
	if parsed == nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	c.Type("application/octet-stream")
	return c.Status(fiber.StatusOK).Send(parsed)
}
