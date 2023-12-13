package routes

import (
	"short-it/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	app.Get("/", controller.Health)

	app.Post("/encode", controller.Encode)
	app.Get("/:short_url", controller.Decode)
}
