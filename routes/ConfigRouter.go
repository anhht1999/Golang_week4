package routes

import (
	"github.com/anhht1999/Golang_week4/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigUserRouter(router *fiber.Router) {

	(*router).Get("/", controller.GetAllUser)

	(*router).Get("/:id", controller.GetUserById) 

	(*router).Delete("/:id", controller.DeleteUserById) 

	(*router).Post("", controller.CreateUser) 

	(*router).Patch("", controller.UpdateUser)

	(*router).Put("", controller.UpsertUser) 

}

func ConfigCategoryRouter(router *fiber.Router) {

	// (*router).Get("/", controller.GetAllCategory) //Liệt kê
	(*router).Delete("/:id", controller.DeleteCategoryById) 

	(*router).Post("", controller.CreateCategory) 

	(*router).Patch("", controller.UpdateCategory) 

	(*router).Put("", controller.UpsertCategory) 

}

func ConfigProductRouter(router *fiber.Router) {

	(*router).Get("/", controller.GetAllProduct)

	(*router).Delete("/:id", controller.DeleteProductById) 

	(*router).Post("", controller.CreateProduct) 

	(*router).Patch("", controller.UpdateProduct) 

	(*router).Put("", controller.UpsertProduct) 

}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Delete("/:id", controller.DeleteReview) 

	(*router).Post("", controller.CreateReview) 

	(*router).Patch("", controller.UpdateReview) 

	(*router).Put("", controller.UpsertReview) 
}

func ConfigImageRouter(router *fiber.Router) {
	(*router).Delete("/:id", controller.DeleteReview) 

	(*router).Post("", controller.CreateImage) 

	// (*router).Patch("", controller.UpdateReview) 

	// (*router).Put("", controller.UpsertReview) 
}