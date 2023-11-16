package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetUpBorrowedBooksEndpoint(echoClient *echo.Echo) {
	echoClient.POST(PostBorrowedBooks, postBorrowedBook)
	echoClient.DELETE(DeleteBorrowedBook, deleteBorrowedBook)
}

func postBorrowedBook(c echo.Context) error {
	return c.String(http.StatusOK, "ENDPOINT FOR POST BORROWED BOOKS")
}

func deleteBorrowedBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
