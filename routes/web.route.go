package routes

import (
	"quickshare/handlers"
	"quickshare/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitWebRoute(server *fiber.App, handler handlers.WebHandler) {
	server.Get("/", middlewares.LooseRoute, handler.SendDashboard)
	server.Get("/login", handler.SendLogin)
	server.Get("/:id", handler.SendPost)
}
