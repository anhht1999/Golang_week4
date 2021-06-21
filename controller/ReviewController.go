package controller

import (
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
	repo "github.com/anhht1999/Golang_week4/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReviews())
}

func GetAllReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	review, err := repo.Reviews.FindReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(review)
}

func DeleteReview(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	//tìm review chứa productid cần tìm
	review, error := repo.Reviews.FindReviewById(int64(id))
	if error != nil {
		return c.Status(404).SendString(err.Error())
	}

	//xóa review
	err = repo.Reviews.DeleteReview(int64(id))

	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		//cập nhật rating và thông báo xóa
		Product, err := repo.Products.FindProductById(int64(review.ProductId))
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}

		result := repo.Reviews.AverageRating(int64(review.ProductId))
		Product.Rating = float32(result[int64(review.ProductId)])

		return c.SendString("delete successfully")
	}
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	Product, err := repo.Products.FindProductById(int64(review.ProductId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	Id := repo.Reviews.CreateNewReview(review)

	result := repo.Reviews.AverageRating(int64(review.ProductId))
	Product.Rating = float32(result[int64(review.ProductId)])

	return c.SendString(fmt.Sprintf("New review is created successfully with id = %d", Id))

}

func UpdateReview(c *fiber.Ctx) error {
	updatedReview := new(model.Review)

	err := c.BodyParser(&updatedReview)

	review, error := repo.Reviews.FindReviewById(int64(updatedReview.ProductId))
	if error != nil {
		return c.Status(404).SendString(err.Error())
	}
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Reviews.UpdateReview(updatedReview)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	Product, err := repo.Products.FindProductById(int64(review.ProductId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	result := repo.Reviews.AverageRating(int64(review.ProductId))
	Product.Rating = float32(result[int64(review.ProductId)])

	return c.SendString(fmt.Sprintf("Review with id = %d is successfully updated", updatedReview.Id))

}

func UpsertReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)

	review, error := repo.Reviews.FindReviewById(int64(review.ProductId))
	if error != nil {
		return c.Status(404).SendString(err.Error())
	}
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Reviews.UpsertReview(review)

	Product, err := repo.Products.FindProductById(int64(review.ProductId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	result := repo.Reviews.AverageRating(int64(review.ProductId))
	Product.Rating = float32(result[int64(review.ProductId)])

	return c.SendString(fmt.Sprintf("Review with id = %d is successfully upserted", id))
}
