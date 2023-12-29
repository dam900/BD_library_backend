package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"library_app/auth"
	"library_app/routes"
	"library_app/storage"
	"log"
)

func main() {
	log.Println("Creating echo client")
	e := echo.New()
	e.File(routes.INDEX, "static/index.html")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_custom} method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00",
	}))

	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	e.Use(storage.DbMiddleware(store))

	api := e.Group("/api")
	login := e.Group("/login")
	routes.SetUpLoginEndpoint(login)

	api.Use(middleware.BasicAuth(auth.Authorize))

	log.Println("Setting up routes")

	routes.SetUpBooksEndpoint(api)
	routes.SetUpAuthorsEndpoint(api)
	routes.SetUpBookedEndpoint(api)
	routes.SetUpBorrowedEndpoint(api)

	log.Println("Starting a server")
	log.Println("Server running on: http://localhost:1323/")
	e.Logger.Fatal(e.Start(":1323"))

}
