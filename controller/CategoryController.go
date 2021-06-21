package controller

import (
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
	repo "github.com/anhht1999/Golang_week4/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategory(c *fiber.Ctx) error {
	return c.JSON(repo.Categorys.GetAllCategorys())
}

func GetCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	Category, err := repo.Categorys.FindCategoryById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(Category)
}

func DeleteCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Categorys.DeleteCategoryById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateCategory(c *fiber.Ctx) error {
	Category := new(model.Category)

	err := c.BodyParser(&Category)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	CategoryId := repo.Categorys.CreateNewCategory(Category)
	return c.SendString(fmt.Sprintf("New Category is created successfully with id = %d", CategoryId))

}

func UpdateCategory(c *fiber.Ctx) error {
	updatedCategory := new(model.Category)

	err := c.BodyParser(&updatedCategory)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Categorys.UpdateCategory(updatedCategory)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Category with id = %d is successfully updated", updatedCategory.Id))

}

func UpsertCategory(c *fiber.Ctx) error {
	Category := new(model.Category)

	err := c.BodyParser(&Category)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Categorys.Upsert(Category)
	return c.SendString(fmt.Sprintf("Category with id = %d is successfully upserted", id))
}
