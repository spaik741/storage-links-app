package web

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"time"
)

func StartListening() {
	configs := fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	app := fiber.New(configs)
	app.Get("/links/:chatId", GetLinks)
	app.Post("/links", Save)
	app.Delete("/links", Clear)
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Error start web app", err)
	}
}
