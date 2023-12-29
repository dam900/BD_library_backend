package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"library_app/storage"
	"library_app/types"
	"net/http"
)

func SetUpBookedEndpoint(echoClient *echo.Group) {
	echoClient.POST(PostBookedBooks, postBooked)
	echoClient.DELETE(DeleteBookedBook, deleteBooked)
}

func postBooked(c echo.Context) error {
	id := c.Param("id")
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	var bookedStatus types.BookedStatus
	if err := c.Bind(&bookedStatus); err != nil {
		return c.JSON(http.StatusBadRequest, "Wrong body")
	}
	book, err := db.BooksRepository.Retrieve(id, &storage.QueryOptions{Ctx: c, Offset: 0})
	if err != nil {
		return c.JSON(http.StatusNotFound, "Book doesn't exist")
	}
	if book.IsBooked() {
		return c.JSON(http.StatusConflict, "Already booked")
	}
	if book.IsBorrowed() {
		return c.JSON(http.StatusConflict, "Cannot book borrowed book")
	}
	err = db.BooksRepository.BookBook(id, &bookedStatus, &storage.QueryOptions{Ctx: c, Offset: 0})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
	}
	return c.JSON(http.StatusOK, "")
}

func deleteBooked(c echo.Context) error {
	id := c.Param("id")
	userId := c.Param("user")
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	book, err := db.BooksRepository.Retrieve(id, &storage.QueryOptions{Ctx: c, Offset: 0})
	if err != nil {
		return c.JSON(http.StatusNotFound, "Book doesn't exist")
	}
	if !book.IsBooked() {
		return c.JSON(http.StatusConflict, "Book was never booked")
	}
	err = db.BooksRepository.UnBookBook(id, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "")
}
