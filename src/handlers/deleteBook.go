package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	struct_test "github.com/neyaadeez/go-fiber/src/struct_type"
)

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, book := range struct_test.BooksLibrary {
		if book.Id == id {
			struct_test.BooksLibrary = append(struct_test.BooksLibrary[:i], struct_test.BooksLibrary[i+1:]...)
			return c.SendString(fmt.Sprintf("%v book deleted successfully!", id))
		}
	}

	return c.Status(404).SendString(fmt.Sprintf("%v this book doens't exists in library", id))
}
