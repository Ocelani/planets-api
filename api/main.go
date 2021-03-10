package main

import (
	"log"
	"planets-api/api/database"
	"planets-api/api/routes"
	"planets-api/pkg/planet"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	defer app.Shutdown()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome!"))
	})

	routes.Planet(app, *planet.NewService(
		database.Connection().Collection("planets"),
	))

	log.Fatal(app.Listen(":8080"))
}
