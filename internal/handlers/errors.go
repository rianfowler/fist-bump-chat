package handlers

import "github.com/gofiber/fiber/v2"

func Get404Page(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).Render("web/templates/404", nil, "web/templates/layouts/main")
}
