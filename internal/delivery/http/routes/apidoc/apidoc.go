package apidoc

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/gofiber/swagger"
)

func RegisterAPIDoc(router fiber.Router) {
	router.Get("/swagger", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	router.Get("/swagger/*", fiberSwagger.HandlerDefault)
}
