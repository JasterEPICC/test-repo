package main

import (
	"fmt"
	"privilege-api-myais/lib"
	"privilege-api-myais/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	lib.InitConfig()
	lib.InitTimeZone()
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Register Router
	router.New(app)

	err := app.Listen(fmt.Sprintf(":%s", viper.GetString("app.port")))

	if err != nil {
		lib.LogError(fmt.Sprintf("Cannot start server: %v", err), true)
	}
}
