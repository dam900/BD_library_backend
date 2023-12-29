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

func (booksRepository BooksRepository) DoesExist(id string) bool {
	_, err := booksRepository.Db.Query(Query.SelectBookQuery, id)
	if err != nil {
		return false
	}
	return true
}

func (booksRepository BooksRepository) Create(book *types.BookDto, opt *QueryOptions) (*types.BookDto, error) {
	ctx := opt.Ctx.Request().Context()
	return createBook(book, booksRepository, ctx)
}

func (booksRepository BooksRepository) Retrieve(id string, opt *QueryOptions) (*types.BookDto, error) {
	return selectBook(id, booksRepository)
}

func (booksRepository BooksRepository) RetrieveAll(opt *QueryOptions) ([]types.BookDto, error) {
	offset := opt.Offset
	return selectBooks(booksRepository, offset)
}

func (booksRepository BooksRepository) Delete(id string, opt *QueryOptions) error {
	return deleteBook(id, booksRepository)
}

func (booksRepository BooksRepository) Update(id string, newBook *types.BookDto, opt *QueryOptions) (*types.BookDto, error) {
	ctx := opt.Ctx.Request().Context()
	return updateBook(id, newBook, booksRepository, ctx)
}

func (booksRepository BooksRepository) BookBook(id string, status *types.BookedStatus, opt *QueryOptions) error {
	book, err := booksRepository.Retrieve(id, opt)
	if err != nil {
		return err
	}
	_, err = booksRepository.Db.Exec(Query.CreateBookedStatusQuery, book.Id, status.BookedBy, status.To.String())
	if err != nil {
		return err
	}
	return nil
}

func (booksRepository BooksRepository) UnBookBook(id, userId string) error {
	_, err := booksRepository.Db.Exec("DELETE FROM booked WHERE book_id=$1 AND user_id=$2", id, userId)
	if err != nil {
		return err
	}
	return nil
}

func (booksRepository BooksRepository) BorrowBook(id string, status *types.BorrowedStatus, opt *QueryOptions) error {
	book, err := booksRepository.Retrieve(id, opt)
	ctx := opt.Ctx.Request().Context()
	if err != nil {
		return err
	}
	tx, err := booksRepository.Db.BeginTx(ctx, &sql.TxOptions{})
	if book.IsBooked() {
		_, err = tx.Exec("DELETE FROM booked WHERE book_id = $1 AND user_id = $2", book.Id, status.BorrowedBy)
		if err != nil {
			return err
		}
	}
	_, err = tx.Exec(Query.CreateBorrowedStatusQuery, book.Id, status.BorrowedBy, status.From.String(), status.To.String())
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (booksRepository BooksRepository) UnBorrowBook(id, userId string, opt *QueryOptions) error {
	ctx := opt.Ctx.Request().Context()
	tx, err := booksRepository.Db.BeginTx(ctx, &sql.TxOptions{})
	_, err = tx.Exec("DELETE FROM borrowed WHERE book_id = $1 AND user_id = $2", id, userId)
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO archive (book_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING", id, userId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func deleteBook(id string, booksRepository BooksRepository) error {
	_, err := booksRepository.Db.Exec(Query.DeleteBookQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func selectBook(id string, booksRepository BooksRepository) (*types.BookDto, error) {
	rows, err := booksRepository.Db.Query(Query.SelectBookQuery, id)

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

func createBook(book *types.BookDto, booksRepository BooksRepository, ctx context.Context) (*types.BookDto, error) {
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
	var genreId int
	err = tx.QueryRow("SELECT id FROM genres WHERE genre = $1", book.Genre).Scan(&genreId)
	if err != nil {
		return nil, err
	}
	row := tx.QueryRowContext(ctx, Query.CreateBookQuery, book.Title, genreId)
	if err := row.Scan(&book.Id); err != nil {
		log.Printf("Unnable to query row: %v", err)
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
	b, err := selectBook(book.Id, booksRepository)
	if err != nil {
		log.Printf("Failed retriving the info back: %v", err)
		return nil, err
	}
	return b, nil
}

func updateBook(id string, newBook *types.BookDto, booksRepository BooksRepository, ctx context.Context) (*types.BookDto, error) {

	b, err := selectBook(id, booksRepository)
	if err != nil {
		log.Printf("Book doesn't exist: %v", err)
		return nil, err
	}

	log.Println("Opening transaction for /books")
	tx, err := booksRepository.Db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		log.Printf("Cannot open transaction: %v", err)
		return nil, err
	}
	log.Println("Success")

	if b.Title != newBook.Title || b.Genre != newBook.Genre {
		if newBook.Title == "" {
			newBook.Title = b.Title
		}
		if newBook.Genre == "" {
			newBook.Genre = b.Genre
		}
		if _, err := tx.Exec(Query.UpdateBooksQuery, newBook.Title, newBook.Genre, id); err != nil {
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return newBook, nil
}
