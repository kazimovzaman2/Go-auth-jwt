package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kazimovzaman2/auth-jwt/handler"
	"github.com/kazimovzaman2/auth-jwt/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	product := api.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
