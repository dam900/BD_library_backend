package storage

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
)

const (
	DbContextKey = "__db"
)

type (
	QueryOptions struct {
		Ctx    echo.Context
		Offset int
	}
	Repository[T any] interface {
		Create(item T, opt QueryOptions) error
		Retrieve(id int, opt QueryOptions) error
		RetrieveAll(opt QueryOptions) error
		Delete(id int, opt QueryOptions) error
		Update(id int, newItem T, opt QueryOptions) error
	}
	PostgresStorage struct {
		db              *sql.DB
		BooksRepository BooksRepository
		UsersRepository UsersRepository
	}
)

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
		db:              db,
		BooksRepository: BooksRepository{db: db},
		UsersRepository: UsersRepository{db: db},
	}, nil

}

func DbMiddleware(db *PostgresStorage) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DbContextKey, db)
			return next(c)
		}
	}
}
