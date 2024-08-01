package app

import (
	"github.com/AeronCyther/leet_tutor/internal/views"
	"github.com/gofiber/fiber/v3"
)

func Init() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return RenderComponent(c, views.Home())
	})

	app.Get("/fragment/home", func(c fiber.Ctx) error {
		return RenderComponent(c, views.HomeFragment())
	})

	app.Get("/about", func(c fiber.Ctx) error {
		return RenderComponent(c, views.About())
	})

	app.Get("/fragment/about", func(c fiber.Ctx) error {
		return RenderComponent(c, views.AboutFragment())
	})

	return app
}
