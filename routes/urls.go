package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scitax/testGo/handlers"
)

func NewRouter() *fiber.App {

	app := fiber.New()

	// Use middlewares for each route
	// RespTimeMW should be first middleware to concider other middlewares performance
	app.Use(handlers.RespTimeMW, handlers.PreprocessingMW)
	app.Get("/", handlers.HwFiber)

	return app
}
