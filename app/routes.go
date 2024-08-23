package app

import (
	"github.com/gofiber/fiber/v2"
	"main/app/controller"
)

func SetupRoutes(app *fiber.App) {

	pages := app.Group("/")

	pages.Get("/", controller.HomeGet)
	pages.Post("/", controller.HomePost)

	// Author
	//api.Get("/authors", controller.GetAuthors)
	//api.Get("/authors/:id", controller.GetAuthor)
	//api.Post("/authors", controller.NewAuthor)
	//api.Delete("/authors/:id", controller.DeleteAuthor)
	//api.Put("/authors/:id", controller.UpdateAuthor)
}
