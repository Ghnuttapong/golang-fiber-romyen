package repository

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/models"
)

func GetCategories(ctx *fiber.Ctx) error{
	var categories []models.Category
	db.Database.Find(&categories)
	return ctx.Status(http.StatusOK).JSON(&categories)
}

func GetCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")	
	var category models.Category
	result := db.Database.First(&category, id)
	if result.RowsAffected == 0 {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	return ctx.Status(http.StatusOK).JSON(&category)
}


func AddCategory(ctx *fiber.Ctx) error {
	category := new(models.Category)

	if err := ctx.BodyParser(category); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	result := db.Database.First(&category, "name = ?", &category.Name)
	if result.RowsAffected == 0 {
		db.Database.Create(&category)
		return ctx.Status(http.StatusOK).JSON(category)
	}
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"There is already a name for this category."})
}

func UpdateCategory(ctx *fiber.Ctx) error {
	category := new(models.Category)
	id := ctx.Params("id")
	if err := ctx.BodyParser(category); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	result := db.Database.First(&category, "name = ?", &category.Name)
	if result.RowsAffected == 0 {
		db.Database.Where("id = ?", id).Updates(&category)
		return ctx.Status(http.StatusOK).JSON(category)
	}
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"There is already a name for this category."})
}

func DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var category models.Category
	result := db.Database.Delete(&category, id)
	if result.RowsAffected == 0 {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	return ctx.SendStatus(http.StatusOK)
}