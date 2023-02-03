package accounting

import (
	"github.com/gofiber/fiber/v2"
)

func Account(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Account Page",
	})
}

func AddAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Add Account Page",
	})
}

func EditAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Edit Account Page",
		"id":      id,
	})
}

func DetailAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Detail Account Page",
		"id":      id,
	})
}
