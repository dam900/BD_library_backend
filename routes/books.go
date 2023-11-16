package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"library_app/storage"
	"library_app/types"
	"net/http"
)

func SetUpBooksEndpoint(echoClient *echo.Echo) {
	echoClient.GET(GetBooks, getBooks)
	echoClient.POST(PostBook, postBooks)
	echoClient.GET(GetBookWithId, getBookWithId)
	echoClient.PUT(PutBook, putBook)
	echoClient.DELETE(DeleteBook, deleteBook)
}

func getBooks(c echo.Context) error {
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	result, err := db.BooksRepository.RetrieveAll(&storage.QueryOptions{c, 0})
	if err != nil {
		return c.String(http.StatusInternalServerError, "There was an error")
	}
	return c.JSONPretty(http.StatusOK, result, "	")
}

func getBookWithId(c echo.Context) error {
	id := c.Param("id")
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	result, err := db.BooksRepository.Retrieve(id, &storage.QueryOptions{c, 0})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSONPretty(http.StatusOK, result, "	")
}

func postBooks(c echo.Context) error {
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	b := &types.BookDto{}
	if err := c.Bind(b); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	b, err := db.BooksRepository.Create(b, &storage.QueryOptions{Ctx: c})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSONPretty(http.StatusOK, b, "	")
}

func putBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteBook(c echo.Context) error {
	id := c.Param("id")
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	err := db.BooksRepository.Delete(id, &storage.QueryOptions{
		Ctx: c,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("Deleted book with id: %s", id))
}
