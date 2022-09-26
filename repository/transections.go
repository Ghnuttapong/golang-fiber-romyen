package repository

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/models"
)

func GetTransections(ctx *fiber.Ctx) error {
	var transections []models.Transaction

	db.Database.Find(&transections)
	return ctx.Status(http.StatusOK).JSON(transections)
}

func GetTransection(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var transection models.Transaction

	result := db.Database.Find(&transection, id)

	if result.RowsAffected == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Not found Transection"})
	}
	return ctx.Status(http.StatusOK).JSON(&transection)
}


func AddTransection(ctx *fiber.Ctx) error {
	transection := new(models.Transaction)

	if err := ctx.BodyParser(transection); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	result := db.Database.Find(&transection, "student_id = ? AND finance_id", &transection.StudentId, &transection.FinanceId)
	if result.RowsAffected == 0 {
		db.Database.Create(&transection)
		return ctx.Status(http.StatusOK).JSON(transection)
	}
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"There is 1 item already."})
}

func UpdateTransection(ctx *fiber.Ctx) error {
	transection := new(models.Transaction)
	id := ctx.Params("id")

	if err := ctx.BodyParser(transection); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	db.Database.Where("id = ?", id).Updates(&transection)
	return ctx.Status(http.StatusOK).JSON(transection)
}

func DeleteTransection(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var transection models.Transaction

	result := db.Database.Delete(&transection, id)

	if result.RowsAffected == 0 {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	return ctx.SendStatus(http.StatusOK)
}
