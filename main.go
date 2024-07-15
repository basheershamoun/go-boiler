package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	env := os.Getenv("FOO_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env

}
func main() {

	app := fiber.New()

	app.Get("/", home)

	err := app.Listen(":3333")
	if err != nil {
		return
	}

}

func home(c *fiber.Ctx) error {
	c.Status(200)

	// set content type
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString("hello world")
}
