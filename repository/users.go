package repository

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/middlewares"
	"gnutta.com/models"
)

func GetUsers(ctx *fiber.Ctx) error {
	_, err := middlewares.AuthRequestWithId(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(err)
	}
	var users []models.User
	db.Database.Find(&users)
	return ctx.Status(http.StatusOK).JSON(users)
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User

	result := db.Database.Find(&user, id)
	if result.RowsAffected == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Not found user"})
	}
	return ctx.Status(http.StatusOK).JSON(&user)
}


func AddUser(user *models.User, ctx *fiber.Ctx) {
	db.Database.Create(&user)
	ctx.Next()
}

func UpdateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	id := ctx.Params("id")

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	db.Database.Where("id = ?", id).Updates(&user)
	return ctx.Status(http.StatusOK).JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User

	result := db.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	return ctx.SendStatus(http.StatusOK)
}