package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/harrieson/go-fiber/database"
	"github.com/harrieson/go-fiber/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App) {
	//Welcome
	app.Get("/api", welcome)

	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
