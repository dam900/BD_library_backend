package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/storage"
	"net/http"
)

func SetUpArchivedBooksEndpoint(echoClient *echo.Echo, store *storage.PostgresStorage) {
	echoClient.GET(GET_ARCHIVED_BOOKS, storage.WithStorage(store, getArchivedBooks))
	echoClient.POST(POST_ARCHIVED_BOOKS, storage.WithStorage(store, postArchivedBook))
	echoClient.PUT(PUT_ARCHIVED_BOOK, storage.WithStorage(store, putArchivedBook))
	echoClient.DELETE(DELETE_ARCHIVED_BOOK, storage.WithStorage(store, deleteArchivedBook))
}

func getArchivedBooks(c echo.Context, store *storage.PostgresStorage) error {
	return c.String(http.StatusOK, "ENDPOINT FOR ARCHIVED BOOKS")
}

func postArchivedBook(c echo.Context, store *storage.PostgresStorage) error {
	return c.String(http.StatusOK, "ENDPOINT FOR POST ARCHIVED BOOKS")
}

func putArchivedBook(c echo.Context, store *storage.PostgresStorage) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteArchivedBook(c echo.Context, store *storage.PostgresStorage) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
