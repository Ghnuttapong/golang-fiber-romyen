package repository

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/models"
)

func GetParents(ctx *fiber.Ctx) error {
	var parents []models.Parent

	db.Database.Find(&parents)
	return ctx.Status(http.StatusOK).JSON(&parents)
}

func GetParent(ctx *fiber.Ctx) error {
	id := ctx.Params("id")	
	var parent models.Parent
	result := db.Database.First(&parent, id)
	if result.RowsAffected == 0 {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	return ctx.Status(http.StatusOK).JSON(&parent)
}


func AddParent(ctx *fiber.Ctx) error {
	parent := new(models.Parent)

	if err := ctx.BodyParser(parent); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	db.Database.Create(&parent)
	return ctx.Status(http.StatusOK).JSON(parent)
}

func UpdateParent(ctx *fiber.Ctx) error {
	parent := new(models.Parent)
	id := ctx.Params("id")
	if err := ctx.BodyParser(parent); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	db.Database.Where("id = ?", id).Updates(&parent)
	return ctx.Status(http.StatusOK).JSON(parent)
}

func DeleteParent(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var parent models.Parent
	result := db.Database.Delete(&parent, id)
	if result.RowsAffected == 0 {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	return ctx.SendStatus(http.StatusOK)
}