package storage

import (
	"database/sql"
	"library_app/types"
)

type BooksRepository struct {
	db *sql.DB
}

func (b BooksRepository) Create(book types.Book, opt QueryOptions) error {
	ctx := opt.ctx.Request().Context()
	tx, err := b.db.BeginTx(ctx, &sql.TxOptions{})
	defer tx.Rollback()

	if err != nil {
		return err
	}
	insertBookQuery := "INSERT INTO books (id, title, genre) VALUES (?,?,?);"
	tx.ExecContext(ctx, insertBookQuery, book.Id, book.Name, book.Genre)

	connectToAuthorQuery := "INSERT INTO books2authors (book_id, author_id) VALUES (?, ?);"
	tx.ExecContext(ctx, connectToAuthorQuery, book.Id, book.Author.Id)

	return nil
}

func (b BooksRepository) Retrieve(id int, opt QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) RetrieveAll(opt QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) Delete(id int, opt QueryOptions) error {
	//TODO implement me
	panic("implement me")
}

func (b BooksRepository) Update(id int, newItem types.Book, opt QueryOptions) error {
	//TODO implement me
	panic("implement me")
}
