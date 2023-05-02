package route

import (
	"fiberv2/handler"
	"github.com/gofiber/fiber/v2"
)

func Init(r *fiber.App) {
	r.Get("/", handler.UserHandlerGetAll)
	r.Get("/:id", handler.UserHandlerFindById)
	r.Post("/", handler.UserHandlerCreate)
}
