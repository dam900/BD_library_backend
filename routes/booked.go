package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/storage"
	"net/http"
)

func SetUpBookedBooksEndpoint(echoClient *echo.Echo, store *storage.PostgresStorage) {
	echoClient.GET(GET_BOOKED_BOOKS, storage.WithDb(store, getBookedBooks))
	echoClient.POST(POST_BOOKED_BOOKS, storage.WithDb(store, postBookedBook))
	echoClient.PUT(PUT_BOOKED_BOOK, storage.WithDb(store, putBookedBook))
	echoClient.DELETE(DELETE_BOOKED_BOOK, storage.WithDb(store, deleteBookedBook))
}

func getBookedBooks(c echo.Context, store *storage.PostgresStorage) error {
	return c.String(http.StatusOK, "ENDPOINT FOR BOOKED BOOKS")
}

func postBookedBook(c echo.Context, store *storage.PostgresStorage) error {
	return c.String(http.StatusOK, "ENDPOINT FOR POST BOOKED BOOKS")
}

func putBookedBook(c echo.Context, store *storage.PostgresStorage) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteBookedBook(c echo.Context, store *storage.PostgresStorage) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
