package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	/*return c.Render("layouts/public/index", fiber.Map{
		"Title": os.Getenv("APP_NAME"),
		"csrf_": c.Cookies("csrf_"),
	}, "layouts/main")*/
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "ok",
		"cookie":  c.Cookies("csrf_"),
	})
}

func AboutPage(c *fiber.Ctx) error {
	fmt.Println(c.Cookies("csrf_"))
	fmt.Println(c.Locals("token"))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func ContactPage(c *fiber.Ctx) error {
	return c.SendString("Contact Page")
}
