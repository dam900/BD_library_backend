package storage

import (
	"database/sql"
	"encoding/json"
	Query "library_app/other/query"
	"library_app/types"
	"log"
)

type BooksRepository struct {
	Db *sql.DB
}

func (booksRepository BooksRepository) Create(book *types.BookDto, opt *QueryOptions) (*types.BookDto, error) {
	ctx := opt.Ctx.Request().Context()
	cp := book
	log.Println("Opening transaction for /books")
	tx, err := booksRepository.Db.BeginTx(ctx, &sql.TxOptions{})

	rollback := func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("Unable to roll back: %v", err)
		}
	}

	if err != nil {
		log.Printf("Cannot open transaction: %v", err)
		return nil, err
	}
	log.Println("Success")
	row := tx.QueryRowContext(ctx, Query.CreateBookQuery, book.Title, book.Genre)
	if err := row.Scan(&cp.Id); err != nil {
		log.Printf("Unnable to query row: %v", err)
		return nil, err
	}
	for _, author := range book.Authors {
		_, err := tx.ExecContext(ctx, Query.ConnectAuthorsToBooksQuery, cp.Id, author.Id)
		if err != nil {
			log.Printf("Unable to connect books with authors: %v", err)
			rollback()
		}
	}

	log.Println("Committing transaction")
	if err := tx.Commit(); err != nil {
		log.Printf("Failed committing: %v", err)
		rollback()
		return nil, err
	}
	log.Println("Success")
	b, err := booksRepository.Retrieve(cp.Id, opt)
	if err != nil {
		log.Printf("Failed retriving the info back: %v", err)
		return nil, err
	}
	return b, nil
}

func (booksRepository BooksRepository) Retrieve(id string, opt *QueryOptions) (*types.BookDto, error) {
	query := Query.SelectBookQuery

	rows, err := booksRepository.Db.Query(query, id)
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

func (booksRepository BooksRepository) RetrieveAll(opt *QueryOptions) ([]types.BookDto, error) {
	offset := opt.Offset
	query := Query.SelectBooksQuery

	rows, err := booksRepository.Db.Query(query, offset)
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

func (booksRepository BooksRepository) Delete(id string, opt *QueryOptions) error {
	//_, err := booksRepository.Db.Exec(Query.DeleteBookQuery, id)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (booksRepository BooksRepository) Update(id string, newItem types.BookDto, opt *QueryOptions) error {
	//TODO implement me
	panic("implement me")
}
