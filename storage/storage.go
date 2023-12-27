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
		Create(item *T, opt *QueryOptions) (*T, error)
		Retrieve(id string, opt *QueryOptions) (*T, error)
		RetrieveAll(opt *QueryOptions) ([]T, error)
		Delete(id string, opt *QueryOptions) error
		Update(id string, newItem T, opt *QueryOptions) (*T, error)
	}
	PostgresStorage struct {
		db                *sql.DB
		BooksRepository   BooksRepository
		UsersRepository   UsersRepository
		AuthorsRepository AuthorsRepository
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
		db:                db,
		BooksRepository:   BooksRepository{Db: db},
		UsersRepository:   UsersRepository{Db: db},
		AuthorsRepository: AuthorsRepository{Db: db},
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
