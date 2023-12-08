package main

import (
	"short-it/config"
	"short-it/internal/logger"
	"short-it/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.Info("Starting server...")
	app := fiber.New()
	routes.AddRoutes(app)
	app.Listen(":" + config.Get("PORT"))
}
