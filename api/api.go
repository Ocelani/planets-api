package api

import (
	"log"
	"planets-api/api/database"
	"planets-api/api/routes"

	"github.com/gofiber/fiber/v2"
)

func main(api *fiber.App) {
	api := fiber.New()

	routes.Planet(api, database.Planets())

	log.Fatal(api.Listen(":3000"))
}
