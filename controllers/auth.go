package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/models"
	"gnutta.com/repository"
	"gnutta.com/security"
)

func Login(c *fiber.Ctx) error {
	var user = new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	sendPass := user.Password
	// get password in database
	result := db.Database.Find(&user, "username = ?", &user.Username)
	if result.RowsAffected == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid username"})
	}
	//check 
	ok := security.ChaeckPasswordHash(sendPass, user.Password)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Password"})
	}
	result = db.Database.Find(&user, "username = ? AND password = ?", &user.Username, &user.Password)
	if result.RowsAffected == 0 {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Server error."})
	}
	// token genarate
	token, err := security.GenerateToken(string(user.ID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Server error."})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"user": &user, "token": fmt.Sprintf("Bearer %s", token)})
}

func Register(c *fiber.Ctx) error {
	var user = new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	result := db.Database.Find(&user, "username = ?", &user.Username)
	if result.RowsAffected == 1 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "username already exists."})
	}
	hash, err := security.HashPassword(user.Password)
	if err != nil {
		log.Fatal("Has an error")
		return err
	}
	user.Password = hash
	repository.AddUser(user, c)
	return c.Status(http.StatusOK).JSON(user)
}