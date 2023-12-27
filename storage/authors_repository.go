package storage

import (
	"database/sql"
	"library_app/types"
)

type AuthorsRepository struct {
	Db *sql.DB
}

func (a AuthorsRepository) Create(item *types.Author, opt *QueryOptions) (*types.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorsRepository) Retrieve(id string, opt *QueryOptions) (*types.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorsRepository) RetrieveAll(opt *QueryOptions) ([]types.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorsRepository) Delete(id string, opt *QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthorsRepository) Update(id string, newItem types.Author, opt *QueryOptions) error {
	//TODO implement me
	panic("implement me")
}
