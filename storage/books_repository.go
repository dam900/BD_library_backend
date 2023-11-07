package storage

import (
	"database/sql"
	"encoding/json"
	"library_app/types"
)

type BooksRepository struct {
	db *sql.DB
}

func (b BooksRepository) Create(book types.BookDto, opt QueryOptions) error {
	//ctx := opt.ctx.Request().Context()
	//tx, err := b.db.BeginTx(ctx, &sql.TxOptions{})
	//defer tx.Rollback()
	return nil
}

func (b BooksRepository) Retrieve(id int, opt QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) RetrieveAll(opt QueryOptions) ([]types.BookDto, error) {
	offset := opt.Offset
	query := `SELECT b.id,
       b.title,
       b.genre,
       JSON_BUILD_OBJECT(
               'bookedBy', b2.user_id,
               'to', b2.date_to
           ) AS booked_status,
       JSON_BUILD_OBJECT(
               'borrowedBy', b3.user_id,
               'from', b3.date_from,
               'to', b3.date_to
           ) AS borrowed_status,
       JSON_AGG(
               JSON_BUILD_OBJECT(
                       'id', a.id,
                       'name', a.name,
                       'lastName', a.last_name)
           ) AS authors
FROM books AS b
         JOIN books2authors AS ba ON b.id = ba.book_id
         JOIN authors AS a ON ba.author_id = a.id
         LEFT JOIN booked b2 on b.id = b2.book_id
         LEFT JOIN borrowed b3 on b.id = b3.book_id
GROUP BY b.id, b2.date_to, b2.user_id, b3.user_id, b3.date_from, b3.date_to
OFFSET $1 LIMIT 100`

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
		err := rows.Scan(&book.Id, &book.Title, &book.Genre, &bookedJson, &borrowedJson, &authorsJson)
		if err != nil {
			return books, err
		}
		json.Unmarshal(bookedJson, &bookedStatus)
		json.Unmarshal(borrowedJson, &borrowedStatus)
		json.Unmarshal(authorsJson, &authors)

		book.BookedStatus = &bookedStatus
		book.BorrowedStatus = &borrowedStatus
		book.Authors = authors

		books = append(books, book)
	}

	return books, nil
}

func (b BooksRepository) Delete(id int, opt QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) Update(id int, newItem types.BookDto, opt QueryOptions) error {
	//TODO implement me
	panic("implement me")
}
