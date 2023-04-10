package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

const originalIcsUrlKey = "ORIGINAL_ICS_URL"

func main() {
	originalIcsUrl := os.Getenv(originalIcsUrlKey)
	if originalIcsUrl == "" {
		log.Fatalf("%s is not set", originalIcsUrlKey)
	}

	startWebServer(originalIcsUrl)
}

func startWebServer(originalIcsUrl string) {
	app := fiber.New()

	app.Use(cache.New(cache.Config{
		Expiration: 15 * time.Minute,
	}))

	logRequests := os.Getenv("LOG_REQUESTS") == "true"

	app.Get("/", func(c *fiber.Ctx) error {
		if logRequests {
			log.Println("Request received")
		}

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
