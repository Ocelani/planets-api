package responses

import "github.com/gofiber/fiber/v2"

// UnprocessableEntity (422) response status code indicates that the server
// understands, the content type of the request entity, and the syntax of the
// request entity is correct, but it was unable to process the contained instructions.
func UnprocessableEntity(c *fiber.Ctx, err error) error {
	return c.JSON(&fiber.Map{
		"status":  422,
		"success": false,
		"error":   err,
	})
}

// BadRequest (400) response status code indicates that the server cannot or
// will not process the request due to something that is perceived to be a client error
// (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
func BadRequest(c *fiber.Ctx, err error) error {
	return c.JSON(&fiber.Map{
		"status":  400,
		"success": false,
		"error":   err,
	})
}

// ResetContent (205) response status tells the client to reset the document view,
// for example, to clear the content of a form, reset a canvas state, or to refresh the UI.
func ResetContent(c *fiber.Ctx, result interface{}) error {
	return c.JSON(&fiber.Map{
		"status":  205,
		"success": true,
		"result":  result,
	})
}

// Created (201) success status response code indicates that the request has succeeded and has led to the creation of a resource.
func Created(c *fiber.Ctx, result interface{}) error {
	return c.JSON(&fiber.Map{
		"status":  201,
		"success": true,
		"result":  result,
	})
}

// OK (200) success status response code indicates that the request has succeeded.
func OK(c *fiber.Ctx, result interface{}) error {
	return c.JSON(&fiber.Map{
		"status":  200,
		"success": true,
		"result":  result,
	})
}
