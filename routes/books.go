package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/storage"
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
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	result, err := db.BooksRepository.RetrieveAll(&storage.QueryOptions{c, 0})
	if err != nil {
		return c.String(http.StatusInternalServerError, "There was an error")
	}
	return c.JSONPretty(http.StatusOK, result, "	")
}

func postBooks(c echo.Context) error {
	return c.String(http.StatusOK, "POST BOOKS")
}

func getBookWithId(c echo.Context) error {
	id := c.Param("id")
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	result, err := db.BooksRepository.Retrieve(id, &storage.QueryOptions{c, 0})
	if err != nil {
		return c.String(http.StatusInternalServerError, "There was an error")
	}
	return c.JSONPretty(http.StatusOK, result, "	")
}

func putBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
