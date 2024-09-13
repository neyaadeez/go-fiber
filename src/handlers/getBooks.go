package handlers

import (
	"github.com/gofiber/fiber/v2"
	struct_test "github.com/neyaadeez/go-fiber/src/struct_type"
)

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(struct_test.BooksLibrary)
}
