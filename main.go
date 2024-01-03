package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kazimovzaman2/auth-jwt/database"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()
}
