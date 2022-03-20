package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrieson/go-fiber/database"
	"github.com/harrieson/go-fiber/models"
)

// Product serializer
type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseproduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, SerialNumber: productModel.SerialNumber, Name: productModel.Name}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseproduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := CreateResponseproduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}
