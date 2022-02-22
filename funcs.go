package main

import "github.com/gofiber/fiber/v2"

func sendError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{"error": err.Error()})
}
