package main

import (
	"github.com/anhht1999/Golang_week4/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	UserRouter := app.Group("/api/user")
	routes.ConfigUserRouter(&UserRouter) //http://localhost:3000/api/User

	CategoryRouter := app.Group("/api/category")
	routes.ConfigCategoryRouter(&CategoryRouter) //http://localhost:3000/api/Category

	ProductRouter := app.Group("/api/product")
	routes.ConfigProductRouter(&ProductRouter) //http://localhost:3000/api/product

	ReviewRouter := app.Group("/api/review")
	routes.ConfigReviewRouter(&ReviewRouter) //http://localhost:3000/api/review

	ImageRouter := app.Group("/api/image")
	routes.ConfigImageRouter(&ImageRouter) //http://localhost:3000/api/image

	app.Listen(":3000")
}
