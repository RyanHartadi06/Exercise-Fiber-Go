package main

import (
	"fiberv2/database"
	"fiberv2/database/migration"
	"fiberv2/route"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.DatabaseInit()

	migration.RunMingration()

	app := fiber.New()

	route.Init(app)

	app.Listen(":3000")
}
