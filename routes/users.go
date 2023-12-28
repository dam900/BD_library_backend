package routes

import (
	"github.com/labstack/echo/v4"
	"library_app/auth"
	"library_app/storage"
	"library_app/types"
	"net/http"
)

func SetUpLoginEndpoint(echoClient *echo.Group) {
	echoClient.POST(PostLoginUsers, login)
}

func login(c echo.Context) error {
	var creds types.Credentials
	err := c.Bind(&creds)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Wrong body")
	}
	if creds.Login == "" || creds.Password == "" {
		return c.JSON(http.StatusBadRequest, "Missing credentials")
	}
	db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
	user, exists := db.UsersRepository.DoesExist(creds.Login)
	if exists == false {
		return c.JSON(http.StatusNotFound, "User does not exist")
	}
	good, _ := auth.GoodCredentials(creds.Login, creds.Password, user.Login, user.Password)
	if good == false || err != nil {
		return c.JSON(http.StatusUnauthorized, "Wrong credentials")
	}
	return c.JSON(http.StatusOK, "")
}
