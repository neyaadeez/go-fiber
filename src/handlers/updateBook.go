package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	struct_test "github.com/neyaadeez/go-fiber/src/struct_type"
)

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).SendString("id is required to update the book")
	}

	for i, book := range struct_test.BooksLibrary {
		if book.Id == id {

			temp := new(struct_test.Books)
			err := c.BodyParser(temp)
			if err != nil {
				return c.Status(400).SendString(fmt.Sprintf("error parsing the data: %s", err.Error()))
			}

			if temp.Author != "" {
				struct_test.BooksLibrary[i].Author = temp.Author
			}

			if temp.Name != "" {
				struct_test.BooksLibrary[i].Name = temp.Name
			}

			if temp.ReleaseDate != "" {
				struct_test.BooksLibrary[i].ReleaseDate = temp.ReleaseDate
			}

			return c.Status(200).SendString(fmt.Sprintf("%s data is successfully updated!", id))
		}
	}

	return c.Status(404).SendString(fmt.Sprintf("%s this id doesn't exits", id))
}
