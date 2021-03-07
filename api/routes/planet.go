package routes

import (
	"planets-api/api/responses"
	"planets-api/pkg/planet"

	"github.com/gofiber/fiber/v2"
)

// Planet defines the URL routing for Planet.
func Planet(app fiber.Router, service planet.Service) {
	app.Get("/planet", getAllPlanets(service))
	app.Get("/planet/:id", getOnePlanet(service))
	app.Post("/planet", addPlanet(service))
	app.Put("/planet", updatePlanet(service))
	app.Delete("/planet/:id", removePlanet(service))
}

// addPlanet through API service request.
func addPlanet(service planet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody planet.Planet
		if err := c.BodyParser(&reqBody); err != nil {
			return responses.BadRequest(c, err)
		}
		result, err := service.Insert(&reqBody)
		if err != nil {
			return responses.UnprocessableEntity(c, err)
		}
		return responses.Created(c, result)
	}
}

// updatePlanet through API client request.
func updatePlanet(service planet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody planet.Planet
		if err := c.BodyParser(&reqBody); err != nil {
			return responses.BadRequest(c, err)
		}
		result, err := service.Update(&reqBody)
		if err != nil {
			return responses.UnprocessableEntity(c, err)
		}
		return responses.ResetContent(c, result)
	}
}

// removePlanet through API client request.
func removePlanet(service planet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := service.Remove(id); err != nil {
			return responses.UnprocessableEntity(c, err)
		}
		return responses.ResetContent(c, id)
	}
}

// getAllPlanets through API client request.
func getAllPlanets(service planet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.FindAll()
		if err != nil {
			return responses.UnprocessableEntity(c, err)
		}
		return responses.OK(c, result)
	}
}

// getOnePlanet through API client request.
func getOnePlanet(service planet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.FindOne(id)
		if err != nil {
			return responses.UnprocessableEntity(c, err)
		}
		return responses.OK(c, result)
	}
}
