package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/storage"
	"net/http"
)

func SetUpBorrowedBooksEndpoint(echoClient *echo.Echo, store *storage.PostgresStorage) {
	echoClient.GET(GET_BORROWED_BOOKS, storage.WithDb(store, getBorrowedBooks))
	echoClient.POST(POST_BORROWED_BOOKS, storage.WithDb(store, postBorrowedBook))
	echoClient.PUT(PUT_BORROWED_BOOK, storage.WithDb(store, putBorrowedBook))
	echoClient.DELETE(DELETE_BORROWED_BOOK, storage.WithDb(store, deleteBorrowedBook))
}

func getBorrowedBooks(c echo.Context, store *storage.PostgresStorage) error {
	return c.String(http.StatusOK, "ENDPOINT FOR BORROWED BOOKS")
}

func postBorrowedBook(c echo.Context, store *storage.PostgresStorage) error {
	return c.String(http.StatusOK, "ENDPOINT FOR POST BORROWED BOOKS")
}

func putBorrowedBook(c echo.Context, store *storage.PostgresStorage) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteBorrowedBook(c echo.Context, store *storage.PostgresStorage) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
