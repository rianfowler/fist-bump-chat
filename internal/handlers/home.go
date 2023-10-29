package handlers

import "github.com/gofiber/fiber/v2"

func GetRootPage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Render("web/templates/pages/root", nil, "web/templates/layouts/main")
}
