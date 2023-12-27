package storage

import (
	"context"
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
	return createBook(book, opt, booksRepository, ctx)
}

func (booksRepository BooksRepository) Retrieve(id string, opt *QueryOptions) (*types.BookDto, error) {
	return selectBook(id, booksRepository)
}

func (booksRepository BooksRepository) RetrieveAll(opt *QueryOptions) ([]types.BookDto, error) {
	offset := opt.Offset
	return selectBooks(booksRepository, offset)
}

func (booksRepository BooksRepository) Delete(id string, opt *QueryOptions) error {
	//_, err := booksRepository.Db.Exec(Query.DeleteBookQuery, id)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (booksRepository BooksRepository) Update(id string, book types.BookDto, opt *QueryOptions) (*types.BookDto, error) {
	ctx := opt.Ctx.Request().Context()
	log.Println("Opening transaction for /books")
	tx, err := booksRepository.Db.BeginTx(ctx, &sql.TxOptions{})

	//rollback := func() {
	//	if err := tx.Rollback(); err != nil {
	//		log.Printf("Unable to roll back: %v", err)
	//	}
	//}

	if err != nil {
		log.Printf("Cannot open transaction: %v", err)
		return nil, err
	}
	log.Println("Success")

	b, err := selectBook(id, booksRepository)
	if err != nil {
		log.Printf("Book doesn't exist", err)
		return nil, err
	}
	if b.Title != book.Title || b.Genre != book.Genre {
		tx.Exec(Query.UpdateBooksQuery, book.Title, book.Genre)
	}
	//if b.BookedStatus != book.BookedStatus {
	//	//tx.Exec(Query.UpdateBooksQuery, book.Title, book.Genre)
	//}
	if b.BorrowedStatus != book.BorrowedStatus {
		s := book.BorrowedStatus
		tx.Exec(Query.UpdateBorrowedQuery, s.From, s.To, s.BorrowedBy, book.Id)
	}
	return &book, nil
}

func selectBook(id string, booksRepository BooksRepository) (*types.BookDto, error) {

	query := Query.SelectBookQuery

	rows, err := booksRepository.Db.Query(query, id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	rows.Next()
	book, err := unmarshallBook(rows)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func selectBooks(booksRepository BooksRepository, offset int) ([]types.BookDto, error) {
	rows, err := booksRepository.Db.Query(Query.SelectBooksQuery, offset)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, err
		}
	}

	var books []types.BookDto
	for rows.Next() {
		book, err := unmarshallBook(rows)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func unmarshallBook(rows *sql.Rows) (types.BookDto, error) {
	var borrowedJson []byte
	var borrowedStatus types.BorrowedStatus
	var authorsJson []byte
	var authors []types.Author
	var bookedJson []byte
	var bookedStatus types.BookedStatus

	var book types.BookDto
	if err := rows.Scan(&book.Id, &book.Title, &book.Genre, &bookedJson, &borrowedJson, &authorsJson); err != nil {
		return types.BookDto{}, err
	}
	if err := json.Unmarshal(bookedJson, &bookedStatus); err != nil {
		return types.BookDto{}, err
	}
	book.BookedStatus = &bookedStatus
	if err := json.Unmarshal(borrowedJson, &borrowedStatus); err != nil {
		return types.BookDto{}, err
	}
	book.BorrowedStatus = &borrowedStatus
	if err := json.Unmarshal(authorsJson, &authors); err != nil {
		return types.BookDto{}, err
	}
	book.Authors = authors
	return book, nil
}

func createBook(book *types.BookDto, opt *QueryOptions, booksRepository BooksRepository, ctx context.Context) (*types.BookDto, error) {
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
	if err := row.Scan(&book.Id); err != nil {
		log.Printf("Unnable to query row: %v", err)
		return nil, err
	}
	if _, err := tx.Exec(Query.CreateBorrowedStatus, book.Id); err != nil {
		log.Printf("Unnable to create default borrowed status: %v", err)
		return nil, err
	}

	for _, author := range book.Authors {
		_, err := tx.ExecContext(ctx, Query.CreateAuthorsToBooksQuery, book.Id, author.Id)
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
	b, err := booksRepository.Retrieve(book.Id, opt)

	if err != nil {
		log.Printf("Failed retriving the info back: %v", err)
		return nil, err
	}
	return b, nil
}
