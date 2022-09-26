package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gnutta.com/db"
	"gnutta.com/migration"
	"gnutta.com/routes"
)


func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Load file .env have error: ", err)
	}
}


func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	err := db.NewConnection()
	if err != nil {
		log.Panic(err)
	}
	migration.AutoMigrate(db.Database)
	routes.Install(app)
	port := ":" + os.Getenv("PORT")
	if port != ":" {
		log.Fatal(app.Listen(port))
	}
	log.Fatal(app.Listen(port))
}