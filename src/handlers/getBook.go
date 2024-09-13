package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	struct_test "github.com/neyaadeez/go-fiber/src/struct_type"
)

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, book := range struct_test.BooksLibrary {
		if book.Id == id {
			return c.JSON(book)
		}
	}

	return c.Status(404).SendString(fmt.Sprintf("%v this book doesn't exists", id))
}
