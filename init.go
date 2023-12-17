package main

import (
	"quickshare/handlers"
	"quickshare/repositories"
	"quickshare/routes"
	"quickshare/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(server *fiber.App, db *pgxpool.Pool) {
	tileRepo := repositories.InitTileRepo(db)

	tileService := services.TileServiceImpl{
		Repository: tileRepo,
	}

	_ = handlers.TileHandlerImpl{TileService: &tileService}
	webHandler := handlers.WebHandlerImpl{}

	routes.InitWebRoute(server, &webHandler)
}
