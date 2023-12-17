package routes

import (
	"quickshare/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitWebRoute(server *fiber.App, handler handlers.WebHandler) {
	server.Get("/", handler.SendDashboard)
	server.Get("/login", handler.SendLogin)
}
