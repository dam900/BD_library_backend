package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/types"
	"net/http"
	"time"
)

func SetUpBooksEndpoint(echoClient *echo.Echo) {
	echoClient.GET(GET_BOOKS, getBooks)
	echoClient.POST(POST_BOOKS, postBooks)
	echoClient.GET(GET_BOOK_WITH_ID, getBookWithId)
	echoClient.PUT(PUT_BOOK, putBook)
	echoClient.DELETE(DELETE_BOOK, deleteBook)
}

func getBooks(c echo.Context) error {

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

func postBooks(c echo.Context) error {
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
