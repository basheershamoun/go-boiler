package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	repository "main/database/repository"
)

func Main() {

	fmt.Println("Running web server...")

	err := repository.ConnectDB()
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}
	defer repository.DB.Close()

	app := fiber.New()

	app.Get("/", home)

	err = app.Listen(":3333")
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
