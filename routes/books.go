package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetUpBooksEndpoint(echoClient *echo.Echo) {
	echoClient.GET(GET_BOOKS, getBooks)
	echoClient.POST(POST_BOOKS, postBooks)
	echoClient.GET(GET_BOOK_WITH_ID, getBookWithId)
	echoClient.PUT(PUT_BOOK, putBook)
	echoClient.DELETE(DELETE_BOOK, deleteBook)
}

func getBooks(c echo.Context) error {
	return c.String(http.StatusOK, "GetBook")
}

func postBooks(c echo.Context) error {
	//storage := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	return c.String(http.StatusOK, "POST BOOKS")
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
