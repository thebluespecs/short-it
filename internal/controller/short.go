package controller

import (
    "time"
	"short-it/internal/db/mongo"
	"short-it/internal/logger"
	short "short-it/internal/services"

	"github.com/gofiber/fiber/v2"
)

type encodeInput struct {
	Url string `json:"url"`
    // ExpireAt is the duration after which the url will expire
    // if not provided, it will expire after 7 days
    ExpireAt time.Duration `json:"expireAt"`
}

func Shorten(c *fiber.Ctx) error {
	var input encodeInput
	if err := c.BodyParser(&input); err != nil {
        logger.Error(err.Error())
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "cannot parse JSON",
		})
	}
    if input.ExpireAt == 0 {
        input.ExpireAt = 7 * 24 * time.Hour
    }
	shortened_url, err := short.Encode(input.Url, input.ExpireAt, mongo.GetInstance())
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

func Info(c *fiber.Ctx) error {
	// get short url from params
	short_url := c.Params("short_url")
	logger.Info("short url: " + short_url)
	// decode the short url
	url, err := short.Decode(short_url, mongo.GetInstance())
	if err != nil {
        logger.Error("cannot decode url in info")
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

func Redirect(c *fiber.Ctx) error {
    // get short url from params
    short_url := c.Params("short_url")
    logger.Info("short url: " + short_url)
    // decode the short url
    url, err := short.Decode(short_url, mongo.GetInstance())
    if err != nil {
        logger.Error("cannot decode url in redirect")
        return c.Status(500).JSON(fiber.Map{
            "status":  "error",
            "message": "cannot decode url",
        })
    }
    // redirect to the original url
    logger.Info("redirecting to: " + url)
    return c.Redirect(url, 301)
}
