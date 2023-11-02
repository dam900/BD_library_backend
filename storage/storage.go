package storage

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
)

type ApiHandlerFunc func(c echo.Context, store *PostgresStorage) error

type Repository[T any] interface {
	Create(item T) error
	Retrieve(id int) error
	RetrieveAll() error
	Delete(id int) error
	Update(id int, newItem T) error
}

type PostgresStorage struct {
	BooksRepository *BooksRepository
	UsersRepository *UsersRepository
}

func NewPostgresStore() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=hello_world sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}
	log.Println("Connection to database opened correctly")
	log.Println("Pinging database")
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Ping succeeded")

	return &PostgresStorage{
		BooksRepository: &BooksRepository{db: db},
		UsersRepository: &UsersRepository{db: db},
	}, nil

}

func WithStorage(store *PostgresStorage, f ApiHandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return f(c, store)
	}
}
