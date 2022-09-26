package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"gnutta.com/controllers"
	"gnutta.com/repository"
	"gnutta.com/security"
)

func Authenticate(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: security.JwtSecretKey,
		SigningMethod: security.JwtSigningMethod,
		TokenLookup: "header:Authorization", 	
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(err)
		},
		})(c)
}


func Install(app *fiber.App) {
	// v1 := app.Group("/api", middleware)
	// login	
	app.Post("/login", controllers.Login)
	// users
	app.Get("/users",Authenticate, repository.GetUsers)
	app.Get("/users/:id", repository.GetUser)
	app.Post("/users", controllers.Register)
	app.Put("/users/:id", repository.UpdateUser)
	app.Delete("/users/:id", repository.DeleteUser)
	// prefixes
	app.Get("/prefixes", repository.GetPrefixes)
	app.Get("/prefixes/:id", repository.GetPrefix)
	app.Post("/prefixes", repository.AddPrefix)
	app.Put("/prefixes/:id", repository.UpdatePrefix)
	app.Delete("/prefixes/:id", repository.DeletePrefix)

	app.Get("/parents", repository.GetParents)
	app.Get("/parents/:id", repository.GetParent)
	app.Post("/parents", repository.AddParent)
	app.Put("/parents/:id", repository.UpdateParent)
	app.Delete("/parents/:id", repository.DeleteParent)

	app.Get("/classrooms", repository.GetClassrooms)
	app.Get("/classrooms/:id", repository.GetClassroom)
	app.Post("/classrooms", repository.AddClassroom)
	app.Put("/classrooms/:id", repository.UpdateClassroom)
	app.Delete("/classrooms/:id", repository.DeleteClassroom)

	app.Get("/categories", repository.GetCategories)
	app.Get("/categories/:id", repository.GetCategory)
	app.Post("/categories", repository.AddCategory)
	app.Put("/categories/:id", repository.UpdateCategory)
	app.Delete("/categories/:id", repository.DeleteCategory)

	app.Get("/finances", repository.GetFinances)
	app.Get("/finances/:id", repository.GetFinance)
	app.Post("/finances", repository.AddFinance)
	app.Put("/finances/:id", repository.UpdateFinance)
	app.Delete("/finances/:id", repository.DeleteFinance)

	app.Get("/transections", repository.GetTransections)
	app.Get("/transections/:id", repository.GetTransection)
	app.Post("/transections", repository.AddTransection)
	app.Put("/transections/:id", repository.UpdateTransection)
	app.Delete("/transections/:id", repository.DeleteTransection)

	app.Get("/students", repository.GetStudents)
	app.Get("/students/:id", repository.GetStudent)
	app.Post("/students", repository.AddStudent)
	app.Put("/students/:id", repository.UpdateStudent)
	app.Delete("/students/:id", repository.DeleteStudent)
}