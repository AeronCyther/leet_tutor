package app

import (
	"log"

	"github.com/AeronCyther/leet_tutor/internal/config"
	"github.com/AeronCyther/leet_tutor/internal/views"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func Init() *fiber.App {
	app := fiber.New()

	conf := config.GetOrInitConfig()

	if conf.Env == "dev" {
		sess_id, err := uuid.NewUUID()
		if err != nil {
			log.Fatal("Error generating session ID")
		}
		app.Head("*", func(c fiber.Ctx) error {
			c.Response().Header.Add("ETag", sess_id.String())
			c.Response().Header.Add("Content-Type", "text/html")
			return nil
		})
	}

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
