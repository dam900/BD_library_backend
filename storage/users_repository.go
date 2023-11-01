package storage

import (
	"database/sql"
	"library_app/types"
)

type UsersRepository struct {
	db *sql.DB
}

func (u UsersRepository) Create(item types.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) Retrieve(id string) error {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) RetrieveAll() error {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) Update(id string, newItem types.User) error {
	//TODO implement me
	panic("implement me")
}
