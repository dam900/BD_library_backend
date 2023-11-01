package storage

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"library_app/types"
)

type Repository[T any] interface {
	Create(item T) error
	Retrieve(id string) error
	RetrieveAll() error
	Delete(id string) error
	Update(id string, newItem T) error
}

type PostgresStorage struct {
	booksRepository *BooksRepository
	usersRepository *UsersRepository
}

func NewPostgresStore() (*PostgresStorage, error) {
	connStr := " "

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		booksRepository: &BooksRepository{db: db},
		usersRepository: &UsersRepository{db: db},
	}, nil

}

func WithStorage(store *PostgresStorage, f types.ApiHandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return f(c, store)
	}
}
