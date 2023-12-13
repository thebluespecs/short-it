package controller

import (
	"short-it/internal/db/mongo"
	"short-it/internal/logger"
	short "short-it/internal/services"

	"github.com/gofiber/fiber/v2"
)

type encodeInput struct {
	Url string `json:"url"`
}

func Encode(c *fiber.Ctx) error {
	var input encodeInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "cannot parse JSON",
		})
	}
	shortened_url, err := short.Encode(input.Url, mongo.GetInstance())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "cannot encode url",
		})
	}
	var response struct {
		OriginalUrl  string `json:"original_url"`
		ShortenedUrl string `json:"shortened_url"`
	}
	response.OriginalUrl = input.Url
	response.ShortenedUrl = shortened_url
	return c.JSON(fiber.Map{
		"status": "ok",
		"data":   response,
	})
}

func Decode(c *fiber.Ctx) error {
	// get short url from params
	short_url := c.Params("short_url")
	logger.Info("short url: " + short_url)
	// decode the short url
	url, err := short.Decode(short_url, mongo.GetInstance())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "cannot decode url",
		})
	}
	// redirect to the original url
	return c.JSON(fiber.Map{
		"status": "ok",
		"data":   url,
	})
}
