package routes

import (
	"quickshare/handlers"
	"quickshare/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitWebRoute(server *fiber.App, handler handlers.WebHandler) {
	server.Get("/", middlewares.LooseRoute, handler.SendDashboard)
	server.Get("/login", handler.SendLogin)
	server.Get("/u/:id", middlewares.LooseRoute, handler.SendPosts)
	server.Get("/:id", middlewares.LooseRoute, handler.SendPost)
}
