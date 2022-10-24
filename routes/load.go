package routes

import "github.com/gofiber/fiber/v2"

// Load routes
func Load(app *fiber.App) {
	gopherRoutes(app)
}
