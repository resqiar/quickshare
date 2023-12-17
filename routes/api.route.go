package routes

import (
	"quickshare/handlers"
	"quickshare/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitAPIRoute(server *fiber.App, handler handlers.APIHandler) {
	api := server.Group("api")

	api.Post("/parse", handler.ParseMD)
	api.Post("/share", middlewares.ProtectedRoute, handler.ShareMD)
}
