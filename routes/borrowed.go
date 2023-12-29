package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/storage"
	"library_app/types"
	"net/http"
)

func SetUpBorrowedEndpoint(echoClient *echo.Group) {
	echoClient.POST(PostBorrowedBook, postBorrowed)
	echoClient.DELETE(DeleteBorrowedBook, deleteBorrowed)
}

func postBorrowed(c echo.Context) error {
	id := c.Param("id")
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	var borrowedStatus types.BorrowedStatus
	if err := c.Bind(&borrowedStatus); err != nil {
		return c.JSON(http.StatusBadRequest, "Wrong body")
	}
	book, err := db.BooksRepository.Retrieve(id, &storage.QueryOptions{Ctx: c, Offset: 0})
	if err != nil {
		return c.JSON(http.StatusNotFound, "Book doesn't exist")
	}
	if book.IsBorrowed() {
		return c.JSON(http.StatusConflict, "Book is already borrowed")
	}
	if book.IsBooked() && *book.BookedStatus.BookedBy != *borrowedStatus.BorrowedBy {
		return c.JSON(http.StatusConflict, "Book is booked by another user")
	}
	err = db.BooksRepository.BorrowBook(id, &borrowedStatus, &storage.QueryOptions{Ctx: c})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "")
}

func deleteBorrowed(c echo.Context) error {
	id := c.Param("id")
	userId := c.Param("user")
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	book, err := db.BooksRepository.Retrieve(id, &storage.QueryOptions{Ctx: c, Offset: 0})
	if err != nil {
		return c.JSON(http.StatusNotFound, "Book doesn't exist")
	}
	if !book.IsBorrowed() {
		return c.JSON(http.StatusConflict, "Book was never borrowed")
	}
	err = db.BooksRepository.UnBorrowBook(id, userId, &storage.QueryOptions{Ctx: c})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "")
}
