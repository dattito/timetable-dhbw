package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func main() {
	originalIcsUrl := os.Getenv("ORIGINAL_ICS_URL")
	if originalIcsUrl == "" {
		log.Fatal("ORIGINAL_ICS_URL is not set")
	}

	startWebServer(originalIcsUrl)
}

func startWebServer(originalIcsUrl string) {
	app := fiber.New()

	app.Use(cache.New(cache.Config{
		Expiration: 15 * time.Minute,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		ics, err := getNewIcsFile(originalIcsUrl)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString("an error occured")
		}

		c.Set("Content-Type", "text/calendar")
		return c.SendString(ics.Serialize())
	})

	log.Println("Starting web server")
	app.Listen(":3000")
}
