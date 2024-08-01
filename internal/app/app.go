package app

import (
	"github.com/AeronCyther/leet_tutor/internal/components"
	"github.com/gofiber/fiber/v3"
)

func Init() *fiber.App {
	app := fiber.New()
	component := components.Hello("World")

	app.Get("/", func(c fiber.Ctx) error {
		err := component.Render(c.Context(), c.Response().BodyWriter())
		if err != nil {
			return err
		}
		c.Response().Header.Add("Content-Type", "text/html")
		return nil
	})

	return app
}
