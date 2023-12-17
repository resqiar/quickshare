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
	userRepo := repositories.InitUserRepo(db)

	utilService := services.UtilServiceImpl{}
	userService := services.UserServiceImpl{
		UtilService: &utilService,
		Repository:  userRepo,
	}

	apiHandler := handlers.APIHandlerImpl{UtilService: &utilService}
	authHandler := handlers.AuthHandlerImpl{
		UtilService: &utilService,
		UserService: &userService,
	}
	webHandler := handlers.WebHandlerImpl{
		UserService: &userService,
	}

	routes.InitAPIRoute(server, &apiHandler)
	routes.InitAuthRoute(server, &authHandler)
	routes.InitWebRoute(server, &webHandler)
}
