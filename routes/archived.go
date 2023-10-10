package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetUpArchivedBooksEndpoint(echoClient *echo.Echo) {
	echoClient.GET(GET_ARCHIVED_BOOKS, getArchivedBooks)
	echoClient.POST(POST_ARCHIVED_BOOKS, postArchivedBook)
	echoClient.PUT(PUT_ARCHIVED_BOOK, putArchivedBook)
	echoClient.DELETE(DELETE_ARCHIVED_BOOK, deleteArchivedBook)
}

func getArchivedBooks(c echo.Context) error {
	return c.String(http.StatusOK, "ENDPOINT FOR ARCHIVED BOOKS")
}

func postArchivedBook(c echo.Context) error {
	return c.String(http.StatusOK, "ENDPOINT FOR POST ARCHIVED BOOKS")
}

func putArchivedBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteArchivedBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
