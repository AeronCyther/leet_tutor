package app

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
)

func RenderComponent(c fiber.Ctx, component templ.Component) error {
	err := component.Render(c.Context(), c.Response().BodyWriter())
	if err != nil {
		return err
	}
	c.Response().Header.Add("Content-Type", "text/html")
	return nil
}
