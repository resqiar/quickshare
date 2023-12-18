package main

import (
	"fmt"
	"log"
	"os"
	"quickshare/config"
	"quickshare/database"
	"quickshare/modules"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var (
	engine = html.New("./views", ".html")
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env")
	}

	server := fiber.New(fiber.Config{
		Views: engine,
	})

	server.Static("/static", "./views/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		MaxAge:    3600,
	})

	conn := database.InitDatabase()
	defer conn.Close()

	// Init redis
	database.InitRedis()

	// Initialize sessions
	config.InitSession()
	config.InitStateSession()

	modules.InitModule(server, conn)

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := server.Listen(PORT); err != nil {
		log.Fatal(err)
	}
}
