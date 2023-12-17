package routes

import (
	"quickshare/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitAPIRoute(server *fiber.App, handler handlers.APIHandler) {
	api := server.Group("api")

	api.Post("/parse", handler.ParseMD)
}
