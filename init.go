package main

import (
	"quickshare/handlers"
	"quickshare/routes"
	"quickshare/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(server *fiber.App, db *pgxpool.Pool) {
	utilService := services.UtilServiceImpl{}

	apiHandler := handlers.APIHandlerImpl{UtilService: &utilService}
	webHandler := handlers.WebHandlerImpl{}

	routes.InitAPIRoute(server, &apiHandler)
	routes.InitWebRoute(server, &webHandler)
}
