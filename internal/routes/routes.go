package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ksploitx/user-age-api/internal/handler"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	api := app.Group("/users")

	api.Post("/", h.Create)
	api.Get("/", h.List)
	api.Get("/:id", h.GetByID)
	api.Put("/:id", h.Update)
	api.Delete("/:id", h.Delete)
}
