package storage

import (
	"database/sql"
	"encoding/json"
	Query "library_app/other/query"
	"library_app/types"
)

type BooksRepository struct {
	db *sql.DB
}

func (b BooksRepository) Create(book *types.BookDto, opt *QueryOptions) (*types.BookDto, error) {
	ctx := opt.Ctx.Request().Context()
	tx, err := b.db.BeginTx(ctx, &sql.TxOptions{})
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	row := tx.QueryRow(Query.CreateBookQuery, book.Title, book.Genre)
	if err := row.Scan(&book.Id); err != nil {
		return nil, err
	}
	return book, nil
}

func (b BooksRepository) Retrieve(id string, opt *QueryOptions) (*types.BookDto, error) {
	query := Query.SelectBookQuery

	rows, err := b.db.Query(query, id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	rows.Next()
	var bookedJson []byte
	var bookedStatus types.BookedStatus
	var borrowedJson []byte
	var borrowedStatus types.BorrowedStatus
	var authorsJson []byte
	var authors []types.Author

	var book types.BookDto
	if err = rows.Scan(&book.Id, &book.Title, &book.Genre, &bookedJson, &borrowedJson, &authorsJson); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bookedJson, &bookedStatus); err != nil {
		return nil, err
	}
	book.BookedStatus = &bookedStatus
	if err := json.Unmarshal(borrowedJson, &borrowedStatus); err != nil {
		return nil, err
	}
	book.BorrowedStatus = &borrowedStatus
	if err := json.Unmarshal(authorsJson, &authors); err != nil {
		return nil, err
	}
	book.Authors = authors

	return &book, nil
}

func (b BooksRepository) RetrieveAll(opt *QueryOptions) ([]types.BookDto, error) {
	offset := opt.Offset
	query := Query.SelectBooksQuery

	rows, err := b.db.Query(query, offset)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, err
		}
	}

	var books []types.BookDto
	for rows.Next() {
		var bookedJson []byte
		var bookedStatus types.BookedStatus
		var borrowedJson []byte
		var borrowedStatus types.BorrowedStatus
		var authorsJson []byte
		var authors []types.Author

		var book types.BookDto
		if err = rows.Scan(&book.Id, &book.Title, &book.Genre, &bookedJson, &borrowedJson, &authorsJson); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(bookedJson, &bookedStatus); err != nil {
			return nil, err
		}
		book.BookedStatus = &bookedStatus
		if err := json.Unmarshal(borrowedJson, &borrowedStatus); err != nil {
			return nil, err
		}
		book.BorrowedStatus = &borrowedStatus
		if err := json.Unmarshal(authorsJson, &authors); err != nil {
			return nil, err
		}
		book.Authors = authors
		books = append(books, book)
	}

	return books, nil
}

func (b BooksRepository) Delete(id string, opt *QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) Update(id string, newItem types.BookDto, opt *QueryOptions) error {
	//TODO implement me
	panic("implement me")
}
