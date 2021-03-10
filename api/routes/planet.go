package routes

import (
	"fmt"
	"planets-api/pkg/planet"

	"github.com/gofiber/fiber/v2"
)

// Planet defines the URL routing for Planet.
func Planet(app fiber.Router, service planet.Service) {
	app.Post("/planet", addPlanet(service))
	app.Delete("/planet/:id", removePlanet(service))
	app.Get("/planet", getAllPlanets(service))
	app.Get("/planet/:find", getOnePlanet(service))
}

// addPlanet through API service request.
func addPlanet(service planet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody planet.Planet
		if err := c.BodyParser(&reqBody); err != nil {
			return c.Status(400).SendString(fmt.Sprint(err, reqBody))
		}
		if find, err := service.FindOneWithName(reqBody.Name); find != nil {
			if err != nil {
				return c.Status(400).SendString(fmt.Sprint(err, reqBody))
			}
			return c.Status(200).JSON(find)
		}
		res, err := service.Insert(&reqBody)
		if err != nil {
			return c.Status(422).SendString(fmt.Sprint(err))
		}

		return c.Status(201).JSON(res)
	}
}

// removePlanet through API client request.
func removePlanet(service planet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := service.Remove(id); err != nil {
			return c.Status(422).SendString(fmt.Sprint(err))
		}

		return c.Status(200).SendString(id)
	}
}

// getAllPlanets through API client request.
func getAllPlanets(service planet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.FindAll()
		if err != nil {
			return c.Status(422).SendString(fmt.Sprint(err))
		}

		return c.Status(200).JSON(result)
	}
}

// getOnePlanetWithID through API client request.
func getOnePlanet(service planet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			err    error
			result *planet.Planet
			params = c.Params("find")
		)
		if len(params) == 24 { // id
			result, err = service.FindOneWithID(params)
		} else { // name
			result, err = service.FindOneWithName(params)
		}
		if err != nil {
			return c.Status(422).SendString(fmt.Sprint(err))
		}

		return c.Status(200).JSON(result)
	}
}
