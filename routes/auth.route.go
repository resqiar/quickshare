package routes

import (
	"quickshare/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoute(server *fiber.App, handler handlers.AuthHandler) {
	api := server.Group("auth")

	api.Get("/google", handler.SendGoogle)
	api.Get("/google/callback", handler.SendGoogleCallback)
}
