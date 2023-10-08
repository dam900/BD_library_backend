package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const GET_BOOKS = "/books"
const GET_BOOK_WITH_ID = "/books/:id"
const PUT_BOOK = "/books/:id"
const DELETE_BOOK = "/books/:id"

func SetUpBooksEndpoint(e *echo.Echo) {
	e.GET(GET_BOOKS, getBook)
	e.GET(GET_BOOK_WITH_ID, getBookWithId)
	e.PUT(PUT_BOOK, putBook)
	e.DELETE(DELETE_BOOK, deleteBook)
}

func getBook(c echo.Context) error {
	return c.String(http.StatusOK, "GET ALL BOOKS")
}

func getBookWithId(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func putBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
