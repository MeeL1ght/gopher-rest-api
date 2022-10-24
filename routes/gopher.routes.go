package routes

import (
	gopherController "github.com/MeeL1ght/gopher-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

// Gopher routes
func gopherRoutes(app *fiber.App) {
	gopherRoute := app.Group("/gopher")

	// Create
	gopherRoute.Post("/", gopherController.CreateGopher)

	// Read All
	gopherRoute.Get("/", gopherController.ReadGophers)

	// Read One
	gopherRoute.Get("/:id", gopherController.ReadGopher)

	// Update
	gopherRoute.Put("/:id", gopherController.UpdateGopher)

	// Delete
	gopherRoute.Delete("/:id", gopherController.DeleteGopher)
}
