package storage

import (
	"database/sql"
	"library_app/types"
)

type BooksRepository struct {
	db *sql.DB
}

func (b BooksRepository) Create(item types.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) Retrieve(id string) error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) RetrieveAll() error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) Update(id string, newItem types.Book) error {
	//TODO implement me
	panic("implement me")
}
