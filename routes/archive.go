package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/storage"
	"net/http"
)

func SetUpArchiveEndpoint(echoClient *echo.Group) {
	echoClient.GET(GetArchive, getArchive)
}

func getArchive(c echo.Context) error {
	userID := c.Param("user")
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	_, exists := db.UsersRepository.DoesExist(userID)
	if !exists {
		return c.JSON(http.StatusNotFound, "User doesn't exist")
	}
	books, err := db.BooksRepository.RetrieveArchive(userID, &storage.QueryOptions{Ctx: c, Offset: 0})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, books, "	")
}
