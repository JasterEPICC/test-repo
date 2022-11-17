package router

import (
	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Gateway API")
	})

	campaign := app.Group("/campaign")
	Campaign(campaign)

	category := app.Group("/category")
	Category(category)

	voucher := app.Group("/voucher")
	Voucher(voucher)
}
