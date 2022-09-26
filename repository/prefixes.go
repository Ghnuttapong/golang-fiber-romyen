package repository

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/models"
)

func GetPrefixes(ctx *fiber.Ctx) error{
	var prefixes []models.Prefix
	db.Database.Find(&prefixes)
	return ctx.Status(http.StatusOK).JSON(&prefixes)
}

func GetPrefix(ctx *fiber.Ctx) error {
	id := ctx.Params("id")	
	var prefix models.Prefix
	result := db.Database.First(&prefix, id)
	if result.RowsAffected == 0 {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	return ctx.Status(http.StatusOK).JSON(&prefix)
}


func AddPrefix(ctx *fiber.Ctx) error {
	prefix := new(models.Prefix)

	if err := ctx.BodyParser(prefix); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	db.Database.Create(&prefix)
	return ctx.Status(http.StatusOK).JSON(prefix)
}

func UpdatePrefix(ctx *fiber.Ctx) error {
	prefix := new(models.Prefix)
	id := ctx.Params("id")
	if err := ctx.BodyParser(prefix); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	db.Database.Where("id = ?", id).Updates(&prefix)
	return ctx.Status(http.StatusOK).JSON(prefix)
}

func DeletePrefix(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var prefix models.Prefix
	result := db.Database.Delete(&prefix, id)
	if result.RowsAffected == 0 {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	return ctx.SendStatus(http.StatusOK)
}