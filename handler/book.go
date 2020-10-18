package handler

import (
	"github.com/gofiber/fiber"
	"github.com/halil-burak/fiber-rest-api/database"
	"github.com/halil-burak/fiber-rest-api/model"
)

// GetBooks returns all books
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	c.JSON(books)
}

// GetBook returns single book
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.Find(&book, id)

	if book.ID == 0 || book.Title == "" {
		c.Status(404).Send("No book found with ID")
		return
	}

	c.JSON(book)
}

// NewBook adds a new book
func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

// DeleteBook removes a book
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.First(&book, id)

	if book.ID == 0 || book.Title == "" {
		c.Status(404).Send("No book found with ID")
		return
	}

	db.Delete(&book)
	c.Send("Book successfully deleted.")
}

// UpdateBook updates an existing book
func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var oldBook model.Book
	db.First(&oldBook, id)

	if oldBook.ID == 0 || oldBook.Title == "" {
		c.Status(404).Send("No book found with ID")
		return
	}

	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	oldBook.Author = book.Author
	oldBook.Title = book.Title
	oldBook.Rating = book.Rating
	db.Save(&oldBook)
	c.Send("Updated book.")
}

// GetAuthorOfBook returns author of a book
func GetAuthorOfBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.First(&book, id)

	if book.ID == 0 || book.Title == "" {
		c.Status(404).Send("No book found with ID")
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Author", "data": book.Author})
}

// GetCategoriesOfBook returns categories of a book
func GetCategoriesOfBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.First(&book, id)

	if book.ID == 0 || book.Title == "" {
		c.Status(404).Send("No book found with ID")
		return
	}

	categories := make([]model.Category, 0)
	db.Preload("Book").Find(&categories)

	c.JSON(fiber.Map{"status": "success", "message": "Categories", "data": categories})
}
