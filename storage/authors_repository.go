package storage

import (
	"database/sql"
	Query "library_app/other/query"
	"library_app/types"
)

type AuthorsRepository struct {
	Db *sql.DB
}

func (a AuthorsRepository) Create(author *types.Author, opt *QueryOptions) (*types.Author, error) {
	_, err := a.Db.Exec(Query.CreateAuthorQuery, author.Name, author.LastName)
	if err != nil {
		return nil, err
	}

	row := a.Db.QueryRow(Query.CreateBookQuery, author.Name, author.LastName)
	if err := row.Scan(&author.Id); err != nil {
		return nil, err
	}
	return author, nil
}

func (a AuthorsRepository) Retrieve(id string, opt *QueryOptions) (*types.Author, error) {
	var author types.Author
	err := a.Db.QueryRow(Query.SelectAuthorQuery, id).Scan(author)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (a AuthorsRepository) RetrieveAll(opt *QueryOptions) ([]types.Author, error) {
	var authors []types.Author
	rows, err := a.Db.Query(Query.SelectAuthorsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var author types.Author
		if err := rows.Scan(author); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, err
}

func (a AuthorsRepository) Delete(id string, opt *QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthorsRepository) Update(id string, newItem types.Author, opt *QueryOptions) (*types.Author, error) {
	//TODO implement me
	panic("implement me")
}
