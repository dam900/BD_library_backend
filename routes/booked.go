package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetUpBookedBooksEndpoint(echoClient *echo.Echo) {
	echoClient.POST(PostBookedBooks, setBookBookedStatus)
	echoClient.DELETE(DeleteBookedBook, deleteBookedBook)
}

func setBookBookedStatus(c echo.Context) error {
	return c.String(http.StatusOK, "ENDPOINT FOR POST BOOKED BOOKS")
}

func deleteBookedBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
