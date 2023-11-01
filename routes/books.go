package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/storage"
	"library_app/types"
	"net/http"
	"time"
)

func SetUpBooksEndpoint(echoClient *echo.Echo, store *storage.PostgresStorage) {
	echoClient.GET(GET_BOOKS, storage.WithStorage(store, getBooks))
	echoClient.POST(POST_BOOKS, storage.WithStorage(store, postBooks))
	echoClient.GET(GET_BOOK_WITH_ID, storage.WithStorage(store, getBookWithId))
	echoClient.PUT(PUT_BOOK, storage.WithStorage(store, putBook))
	echoClient.DELETE(DELETE_BOOK, storage.WithStorage(store, deleteBook))
}

func getBooks(c echo.Context, store *storage.PostgresStorage) error {

	b := types.Book{
		Id:     0,
		Name:   "Diune",
		Author: "Frank Herbert",
		BorrowedStatus: &types.BorrowedStatus{
			From: time.Now(),
			To:   time.Now(),
		},
		BookedStatus: &types.BookedStatus{To: time.Now()},
	}

	return c.JSONPretty(http.StatusOK, b, "	")
}

func postBooks(c echo.Context, store *storage.PostgresStorage) error {
	return c.String(http.StatusOK, "POST BOOKS")
}

func getBookWithId(c echo.Context, store *storage.PostgresStorage) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func putBook(c echo.Context, store *storage.PostgresStorage) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteBook(c echo.Context, store *storage.PostgresStorage) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
