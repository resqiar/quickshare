package handlers

import (
	"log"
	"quickshare/inputs"
	"quickshare/services"

	"github.com/gofiber/fiber/v2"
)

type APIHandler interface {
	ParseMD(c *fiber.Ctx) error
	ShareMD(c *fiber.Ctx) error
}

type APIHandlerImpl struct {
	UtilService services.UtilService
	PostService services.PostService
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

func (h *APIHandlerImpl) ShareMD(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var payload inputs.CreatePostInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	result, err := h.PostService.CreatePost(&payload, userID.(string))
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).SendString(result)
}
