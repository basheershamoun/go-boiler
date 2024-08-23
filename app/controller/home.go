package controller

import "github.com/gofiber/fiber/v2"

func HomeGet(c *fiber.Ctx) error {
	c.Status(200)

	// set content type
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString("Home Get request path:/ ")
}

func HomePost(c *fiber.Ctx) error {
	c.Status(200)

	// set content type
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString("This page will show on post  ")
}
