package handlers

import (
	"quickshare/services"

	"github.com/gofiber/fiber/v2"
)

type TileHandler interface {
	SendAllTiles(c *fiber.Ctx) error
}

type TileHandlerImpl struct {
	TileService services.TileService
}

func (h *TileHandlerImpl) SendAllTiles(c *fiber.Ctx) error {
	msg, error := h.TileService.GetAllTiles()
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": error,
		})
	}

	return c.JSON(msg)
}
