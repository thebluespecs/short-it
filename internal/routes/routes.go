package routes

import (
	"short-it/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
    app.Get("/", controller.Health)
}
