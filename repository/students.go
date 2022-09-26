package repository

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/models"
)

func GetStudents(ctx *fiber.Ctx) error {
	var students []models.Student
	db.Database.Find(&students)
	return ctx.Status(http.StatusOK).JSON(students)
}

func GetStudent(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var student models.Student

	result := db.Database.Find(&student, id)

	if result.RowsAffected == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Not found student"})
	}
	return ctx.Status(http.StatusOK).JSON(&student)
}


func AddStudent(ctx *fiber.Ctx) error {
	student := new(models.Student)

	if err := ctx.BodyParser(student); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	result := db.Database.Find(&student, "firstname = ? AND lastname = ?", &student.Firstname, &student.Lastname)
	if result.RowsAffected == 0 {
		db.Database.Create(&student)
		return ctx.Status(http.StatusOK).JSON(student)
	}
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"student already exists."})
}

func UpdateStudent(ctx *fiber.Ctx) error {
	student := new(models.Student)
	id := ctx.Params("id")

	if err := ctx.BodyParser(student); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	db.Database.Where("id = ?", id).Updates(&student)
	return ctx.Status(http.StatusOK).JSON(student)
}

func DeleteStudent(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var student models.Student

	result := db.Database.Delete(&student, id)

	if result.RowsAffected == 0 {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	return ctx.SendStatus(http.StatusOK)
}