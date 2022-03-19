package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/harrieson/go-fiber/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/app", welcome)

	log.Fatal(app.Listen(":3000"))
}
