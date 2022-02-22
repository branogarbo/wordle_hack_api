package main

import (
	"github.com/branogarbo/wordle_hack"
	"github.com/gofiber/fiber/v2"
)

func main() {
	s := fiber.New(fiber.Config{
		GETOnly: true,
	})

	api := s.Group("api")
	day := api.Group("day")

	// day.Get("/today", func(c *fiber.Ctx) error {
	// 	word, err := wordle_hack.GetWordByDate(time.Now())
	// 	if err != nil {
	// 		return sendError(c, fiber.StatusInternalServerError, err)
	// 	}

	// 	return c.JSON(fiber.Map{"word": word})
	// })

	// day.Get("/yesterday", func(c *fiber.Ctx) error {
	// 	word, err := wordle_hack.GetWordByDate(time.Now().Add(time.Hour * -24))
	// 	if err != nil {
	// 		return sendError(c, fiber.StatusInternalServerError, err)
	// 	}

	// 	return c.JSON(fiber.Map{"word": word})
	// })

	// day.Get("/tomorrow", func(c *fiber.Ctx) error {
	// 	word, err := wordle_hack.GetWordByDate(time.Now().Add(time.Hour * 24))
	// 	if err != nil {
	// 		return sendError(c, fiber.StatusInternalServerError, err)
	// 	}

	// 	return c.JSON(fiber.Map{"word": word})
	// })

	day.Get("/", func(c *fiber.Ctx) error {
		type Date struct {
			Year  int
			Month int
			Day   int
		}

		date := new(Date)

		err := c.QueryParser(date)
		if err != nil {
			return sendError(c, fiber.StatusBadRequest, err)
		}

		word, err := wordle_hack.GetWordByInteger(date.Year, date.Month, date.Day)
		if err != nil {
			return sendError(c, fiber.StatusBadRequest, err)
		}

		return c.JSON(fiber.Map{"word": word})
	})

	day.Get("/:dateString", func(c *fiber.Ctx) error {
		date := c.Params("dateString")

		word, err := wordle_hack.GetWordByString(date)
		if err != nil {
			return sendError(c, fiber.StatusBadRequest, err)
		}

		return c.JSON(fiber.Map{"word": word})
	})

	s.Get("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	s.Listen(":3000")
}
