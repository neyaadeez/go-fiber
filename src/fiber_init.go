package src

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neyaadeez/go-fiber/src/handlers"
)

func InitServer() {
	app := fiber.New()

	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBook)
	app.Post("/books", handlers.CreateBook)
	app.Post("/books/update/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	app.Listen(":8080")

}
