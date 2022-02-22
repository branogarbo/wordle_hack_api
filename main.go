package main

import (
	"time"

	"github.com/branogarbo/wordle_hack"
	"github.com/gofiber/fiber/v2"
)

func main() {
	s := fiber.New(fiber.Config{
		GETOnly: true,
	})

	api := s.Group("api")

	api.Get("/today", func(c *fiber.Ctx) error {
		word, err := wordle_hack.GetWordByDate(time.Now())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"word": word})
	})

	api.Get("/day", func(c *fiber.Ctx) error {
		type Date struct {
			Year  int
			Month int
			Day   int
		}

		date := new(Date)

		err := c.QueryParser(date)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		word, err := wordle_hack.GetWordByInteger(date.Year, date.Month, date.Day)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"word": word})
	})

	api.Get("/day/:dateString", func(c *fiber.Ctx) error {
		date := c.Params("dateString")

		word, err := wordle_hack.GetWordByString(date)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"word": word})
	})

	s.Get("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	s.Listen(":3000")
}
