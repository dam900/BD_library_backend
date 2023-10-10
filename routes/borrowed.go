package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const GET_BORROWED_BOOKS = "/books/borrowed"
const POST_BORROWED_BOOKS = "/books/borrowed"
const PUT_BORROWED_BOOK = "/books/borrowed/:id"
const DELETE_BORROWED_BOOK = "/books/borrowed/:id"

func SetUpBorrowedBooksEndpoint(e *echo.Echo) {
	e.GET(GET_BORROWED_BOOKS, getBorrowedBooks)
	e.POST(POST_BORROWED_BOOKS, postBorrowedBook)
	e.PUT(PUT_BORROWED_BOOK, putBorrowedBook)
	e.DELETE(DELETE_BORROWED_BOOK, deleteBorrowedBook)
}

func getBorrowedBooks(c echo.Context) error {
	return c.String(http.StatusOK, "ENDPOINT FOR BORROWED BOOKS")
}

func postBorrowedBook(c echo.Context) error {
	return c.String(http.StatusOK, "ENDPOINT FOR POST BORROWED BOOKS")
}

func putBorrowedBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteBorrowedBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
