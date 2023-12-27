package storage

import (
	"database/sql"
)

type UsersRepository struct {
	db *sql.DB
}

func (u UsersRepository) Create(item *UsersRepository, opt *QueryOptions) (*UsersRepository, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) Retrieve(id string, opt *QueryOptions) (*UsersRepository, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) RetrieveAll(opt *QueryOptions) ([]UsersRepository, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) Delete(id string, opt *QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) Update(id string, newItem UsersRepository, opt *QueryOptions) error {
	//TODO implement me
	panic("implement me")
}
