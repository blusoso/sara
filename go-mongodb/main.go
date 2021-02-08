package main

import (
	"log"

	"github.com/fahstjlps/sara-mongodb/database"
	"github.com/fahstjlps/sara-mongodb/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
