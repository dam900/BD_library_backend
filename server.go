package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"library_app/routes"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_custom} method=${method}, uri=${uri}, status=${status}\n" +
			"header=${header}\n" +
			"query=${query}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00",
	}))

	routes.SetUpBooksEndpoint(e)
	routes.SetUpBorrowedBooksEndpoint(e)
	routes.SetUpBookedBooksEndpoint(e)
	routes.SetUpArchivedBooksEndpoint(e)

	e.Logger.Fatal(e.Start(":1323"))
}
