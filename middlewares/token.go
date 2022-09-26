package middlewares

import (
	"errors"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gnutta.com/db"
	"gnutta.com/security"
)

func AuthRequestWithId(ctx *fiber.Ctx) (*jwt.StandardClaims, error) {
	id := ctx.Params("id")
	user := db.Database.First(id)
	if user == nil {
		return nil, errors.New("Unauthorized")
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		return nil , err
	}
	if payload.Id != id || payload.Issuer != id {
		return nil , errors.New("Unauthorized")
	}
	return payload, nil
}