package routes

import (
	"short-it/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	app.Get("/", controller.Health)

    app.Post("/shorten", controller.Shorten)
    app.Get("/:short_url/info", controller.Info)
    app.Get("/:short_url", controller.Redirect)
}
