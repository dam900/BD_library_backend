package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/storage"
	"library_app/types"
	"net/http"
)

func SetUpAuthorsEndpoint(echoClient *echo.Group) {
	echoClient.GET(GetAuthors, getAuthors)
	echoClient.POST(PostAuthor, postAuthors)
	echoClient.GET(GetAuthorWithId, getAuthorsWithId)
}

func getAuthors(c echo.Context) error {
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	result, err := db.AuthorsRepository.RetrieveAll(&storage.QueryOptions{c, 0})
	if err != nil {
		return c.String(http.StatusInternalServerError, "There was an error")
	}
	return c.JSONPretty(http.StatusOK, result, "	")
}

func getAuthorsWithId(c echo.Context) error {
	id := c.Param("id")
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	if !db.AuthorsRepository.DoesExist(id) {
		return c.JSON(http.StatusNotFound, "Author doesn't exist")
	}
	result, err := db.AuthorsRepository.Retrieve(id, &storage.QueryOptions{c, 0})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSONPretty(http.StatusOK, result, "	")
}

func postAuthors(c echo.Context) error {
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	a := &types.Author{}
	if err := c.Bind(a); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	b, err := db.AuthorsRepository.Create(a, &storage.QueryOptions{Ctx: c})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSONPretty(http.StatusOK, b, "	")
}
