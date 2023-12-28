package auth

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"library_app/storage"
)

func needsAuth(c echo.Context) bool {
	if c.Request().Method == "" || c.Request().Method == "GET" {
		return false
	}
	return true
}

func Authorize(username string, password string, c echo.Context) (bool, error) {
	if needsAuth(c) {
		db := c.Get(storage.DbContextKey).(*storage.PostgresStorage)
		user, exists := db.UsersRepository.DoesExist(username)
		if exists == false {
			return false, nil
		}
		good, err := GoodCredentials(username, password, user.Login, user.Password)
		if good == true && err == nil {
			return true, nil
		}
		return false, nil
	}
	return true, nil
}

func GoodCredentials(username string, password, providedUserName, providedPassword string) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte(providedUserName)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(providedPassword)) == 1 {
		return true, nil
	}
	return false, nil
}
