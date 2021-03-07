package main

import (
	"log"
	"planets-api/api/routes"
	"planets-api/pkg/planet"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})
	routes.Planet(app, planet.NewService())
	log.Fatal(app.Listen(":8080"))
}
