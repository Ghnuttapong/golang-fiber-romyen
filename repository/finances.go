package repository

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/models"
)

func GetFinances(ctx *fiber.Ctx) error {
	var finances []models.Finance

	db.Database.Find(&finances)
	return ctx.Status(http.StatusOK).JSON(finances)
}

func GetFinance(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var finance models.Finance

	result := db.Database.Find(&finance, id)

	if result.RowsAffected == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Not found user"})
	}
	return ctx.Status(http.StatusOK).JSON(&finance)
}


func AddFinance(ctx *fiber.Ctx) error {
	finance := new(models.Finance)

	if err := ctx.BodyParser(finance); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	result := db.Database.Find(&finance, "name = ?", &finance.Name)
	if result.RowsAffected == 0 {
		db.Database.Create(&finance)
		return ctx.Status(http.StatusOK).JSON(finance)
	}
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"There is 1 item already."})
}

func UpdateFinance(ctx *fiber.Ctx) error {
	finance := new(models.Finance)
	id := ctx.Params("id")

	if err := ctx.BodyParser(finance); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	db.Database.Where("id = ?", id).Updates(&finance)
	return ctx.Status(http.StatusOK).JSON(finance)
}

func DeleteFinance(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var finance models.Finance

	result := db.Database.Delete(&finance, id)

	if result.RowsAffected == 0 {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	return ctx.SendStatus(http.StatusOK)
}