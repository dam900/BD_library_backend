package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetUpBookedBooksEndpoint(echoClient *echo.Echo) {
	echoClient.GET(GET_BOOKED_BOOKS, getBookedBooks)
	echoClient.POST(POST_BOOKED_BOOKS, postBookedBook)
	echoClient.PUT(PUT_BOOKED_BOOK, putBookedBook)
	echoClient.DELETE(DELETE_BOOKED_BOOK, deleteBookedBook)
}

func getBookedBooks(c echo.Context) error {
	return c.String(http.StatusOK, "ENDPOINT FOR BOOKED BOOKS")
}

func postBookedBook(c echo.Context) error {
	return c.String(http.StatusOK, "ENDPOINT FOR POST BOOKED BOOKS")
}

func putBookedBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteBookedBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
