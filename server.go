package main

import (
	"github.com/labstack/echo/v4"
	"library_app/routes"
)

func main() {
	e := echo.New()

	routes.SetUpBooksEndpoint(e)
	routes.SetUpBorrowedBooksEndpoint(e)
	routes.SetUpBookedBooksEndpoint(e)
	routes.SetUpArchivedBooksEndpoint(e)

	e.Logger.Fatal(e.Start(":1323"))
}
