package main

import (
	"fmt"

	"github.com/MeeL1ght/gopher-rest-api/config"
	db "github.com/MeeL1ght/gopher-rest-api/database"
	"github.com/MeeL1ght/gopher-rest-api/middlewares"
	router "github.com/MeeL1ght/gopher-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	config.ViperConfig()
}

func main() {
	// - Fiber App
	app := fiber.New()

	// Environment variables
	host := viper.Get("app.host")
	port := viper.Get("app.port")
	address := fmt.Sprintf("%v:%v", host, port)

	// Middlewares
	middlewares.Load(app)

	// MongoDB Connection
	db.NewMongoClient()

	// Routes
	router.Load(app)

	// Run Server
	app.Listen(address)
}
