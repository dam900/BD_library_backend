package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetUpBorrowedBooksEndpoint(echoClient *echo.Echo) {
	echoClient.GET(GET_BORROWED_BOOKS, getBorrowedBooks)
	echoClient.POST(POST_BORROWED_BOOKS, postBorrowedBook)
	echoClient.PUT(PUT_BORROWED_BOOK, putBorrowedBook)
	echoClient.DELETE(DELETE_BORROWED_BOOK, deleteBorrowedBook)
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
