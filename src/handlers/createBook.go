package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	struct_test "github.com/neyaadeez/go-fiber/src/struct_type"
)

func CreateBook(c *fiber.Ctx) error {
	book := new(struct_test.Books)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).SendString(fmt.Sprintf("the data is not parsable: %v", err.Error()))
	}

	struct_test.BooksLibrary = append(struct_test.BooksLibrary, *book)

	return c.JSON(struct_test.BooksLibrary)
}
