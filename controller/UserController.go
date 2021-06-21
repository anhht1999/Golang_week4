package controller

import (
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
	repo "github.com/anhht1999/Golang_week4/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(repo.Users.GetAllUsers())
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	User, err := repo.Users.FindUserById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(User)
}

func DeleteUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Users.DeleteUserById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateUser(c *fiber.Ctx) error {
	User := new(model.User)

	err := c.BodyParser(&User)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	UserId := repo.Users.CreateNewUser(User)
	return c.SendString(fmt.Sprintf("New User is created successfully with id = %d", UserId))

}

func UpdateUser(c *fiber.Ctx) error {
	updatedUser := new(model.User)

	err := c.BodyParser(&updatedUser)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Users.UpdateUser(updatedUser)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("User with id = %d is successfully updated", updatedUser.Id))

}

func UpsertUser(c *fiber.Ctx) error {
	User := new(model.User)

	err := c.BodyParser(&User)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Users.Upsert(User)
	return c.SendString(fmt.Sprintf("User with id = %d is successfully upserted", id))
}
