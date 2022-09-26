package repository


import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/models"
)

func GetClassrooms(ctx *fiber.Ctx) error{
	var classrooms []models.ClassRoom
	db.Database.Find(&classrooms)
	return ctx.Status(http.StatusOK).JSON(&classrooms)
}

func GetClassroom(ctx *fiber.Ctx) error {
	id := ctx.Params("id")	
	var class models.ClassRoom
	result := db.Database.First(&class, id)
	if result.RowsAffected == 0 {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	return ctx.Status(http.StatusOK).JSON(&class)
}


func AddClassroom(ctx *fiber.Ctx) error {
	class := new(models.ClassRoom)

	if err := ctx.BodyParser(class); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	result := db.Database.First(&class, "class = ?", &class.Class)
	if result.RowsAffected == 0 {
		db.Database.Create(&class)
		return ctx.Status(http.StatusOK).JSON(class)
	}
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"class already exists."})
}

func UpdateClassroom(ctx *fiber.Ctx) error {
	class := new(models.ClassRoom)
	id := ctx.Params("id")
	if err := ctx.BodyParser(class); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	result := db.Database.First(&class, "class = ?", &class.Class)
	if result.RowsAffected == 0 {
		db.Database.Where("id = ?", id).Updates(&class)
		return ctx.Status(http.StatusOK).JSON(class)
	}
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"class already exists."})
}

func DeleteClassroom(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var class models.ClassRoom
	result := db.Database.Delete(&class, id)
	if result.RowsAffected == 0 {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	return ctx.SendStatus(http.StatusOK)
}