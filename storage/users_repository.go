package storage

import (
	"database/sql"
	"library_app/types"
)

type UsersRepository struct {
	Db *sql.DB
}

func (u UsersRepository) Create(item *types.User, opt *QueryOptions) (*types.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) DoesExist(id string) (*types.User, bool) {
	var user types.User
	row := u.Db.QueryRow("SELECT * FROM users WHERE login_id = $1", id)
	if err := row.Scan(&user.Name, &user.Lastname, &user.Login, &user.Password); err != nil {
		return nil, false
	}
	return &user, true
}

func (u UsersRepository) Retrieve(id string, opt *QueryOptions) (*types.User, error) {
	var user types.User
	row := u.Db.QueryRow("SELECT * FROM users WHERE login_id = $1", id)
	if err := row.Scan(&user.Name, &user.Lastname, &user.Login, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UsersRepository) RetrieveAll(opt *QueryOptions) ([]types.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) Delete(id string, opt *QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (u UsersRepository) Update(id string, newItem types.User, opt *QueryOptions) (*types.User, error) {
	//TODO implement me
	panic("implement me")
}
