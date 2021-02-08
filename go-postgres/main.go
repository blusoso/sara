package main

import (
	"log"

    "github.com/gofiber/fiber/v2"

    "github.com/fahstjlps/go-sara/router"
    "github.com/fahstjlps/go-sara/database"
)

func main()  {
    app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}