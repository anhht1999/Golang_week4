package controller

import (
	"fmt"
	"time"

	"github.com/anhht1999/Golang_week4/model"
	repo "github.com/anhht1999/Golang_week4/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	return c.JSON(repo.Products.GetAllProducts())
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	Product, err := repo.Products.FindProductById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(Product)
}

func DeleteProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Products.DeleteProductById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateProduct(c *fiber.Ctx) error {
	Product := new(model.Product)

	err := c.BodyParser(&Product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	ProductId := repo.Products.CreateNewProduct(Product)
	return c.SendString(fmt.Sprintf("New Product is created successfully with id = %d", ProductId))

}

func UpdateProduct(c *fiber.Ctx) error {
	updatedProduct := new(model.Product)

	err := c.BodyParser(&updatedProduct)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}


	ProductPrice := repo.Products.GetPriceProductById(updatedProduct.Id)
	if updatedProduct.Price != ProductPrice {
		t := time.Now().Second()
		HistoryPrice := new(model.HistoryPrice)
		HistoryPrice.ProductId = updatedProduct.Id
		HistoryPrice.OldPrice = float64(updatedProduct.Price)
		HistoryPrice.ModifiedAt = int64(t)
		err := c.BodyParser(&HistoryPrice)
		// if error
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot parse JSON",
				"error":   err,
			})
		}
		errr := repo.HistoryPrice.CreateNewHistoryPrice(HistoryPrice)
		if errr != nil {
			return c.Status(404).SendString(err.Error())
		}
	}

	err = repo.Products.UpdateProduct(updatedProduct)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Product with id = %d is successfully updated", updatedProduct.Id))

}

func UpsertProduct(c *fiber.Ctx) error {
	Product := new(model.Product)

	err := c.BodyParser(&Product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Products.Upsert(Product)
	return c.SendString(fmt.Sprintf("Product with id = %d is successfully upserted", id))
}

func GetAllHistoryPrices(c *fiber.Ctx) error {
	return c.JSON(repo.HistoryPrice.GetAllHistoryPrices())
}