package controller

import (
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
	repo "github.com/anhht1999/Golang_week4/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllImage(c *fiber.Ctx) error {
	return c.JSON(repo.Images.GetAllImages())
}

func GetAllImageById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	Image, err := repo.Images.FindImageById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(Image)
}

func DeleteImage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	//xóa Image
	err = repo.Images.DeleteImage(int64(id))

	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {

		return c.SendString("delete successfully")
	}
}

func CreateImage(c *fiber.Ctx) error {
	Image := new(model.Image)

	err := c.BodyParser(&Image)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	Product, err := repo.Products.FindProductById(int64(Image.ProductId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	result := repo.Images.CreateNewImage(Image)
	values := []*model.Image{}
    for _, value := range result {
        values = append(values, value)
    }
	Product.Images = values

	return c.SendString(fmt.Sprintf("New Image is created successfully with id = %d",Image.Id))

}

func UpdateImage(c *fiber.Ctx) error {
	updatedImage := new(model.Image)

	err := c.BodyParser(&updatedImage)

	// Image, error := repo.Images.FindImageById(int64(updatedImage.ProductId))
	// if error != nil {
	// 	return c.Status(404).SendString(err.Error())
	// }
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Images.UpdateImage(updatedImage)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	// Product, err := repo.Products.FindProductById(int64(Image.ProductId))
	// if err != nil {
	// 	return c.Status(404).SendString(err.Error())
	// }

	// result := repo.Images.AverageRating(int64(Image.ProductId))
	// Product.Rating = float32(result[int64(Image.ProductId)])

	return c.SendString(fmt.Sprintf("Image with id = %d is successfully updated", updatedImage.Id))

}

// func UpsertImage(c *fiber.Ctx) error {
// 	Image := new(model.Image)

// 	err := c.BodyParser(&Image)
// 	// if error
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cannot parse JSON",
// 			"error":   err,
// 		})
// 	}

// 	id := repo.Images.UpsertImage(Image)
// 	return c.SendString(fmt.Sprintf("Image with id = %d is successfully upserted", id))
// }