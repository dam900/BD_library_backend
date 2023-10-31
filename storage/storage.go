package storage

import (
	"database/sql"
)

type storage interface {
	Open() (sql.DB, error)
}

type Repository[T any] interface {
	Create(item T) error
	Retrieve(id string) error
	RetrieveAll() error
	Delete(id string) error
	Update(id string, newItem T) error
}
