package query

// SELECT //
const (
	SelectBookQuery = `SELECT b.id,
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
						WHERE b.id = $1
						GROUP BY b.id, b2.date_to, b2.user_id, b3.user_id, b3.date_from, b3.date_to`

	SelectBooksQuery = `SELECT b.id,
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
						OFFSET $1 LIMIT 100;`

	SelectAuthorsQuery = `SELECT * FROM authors;`

	SelectAuthorQuery = `SELECT * FROM authors WHERE id = $1;`
)

// CREATE //
const (
	CreateAuthorQuery = `INSERT INTO authors (name, last_name) VALUES ($1, $2);
`

	CreateBookQuery = `WITH new_book AS (
						INSERT
						INTO books (title, genre)
						VALUES ($1, $2) RETURNING id
							)
						SELECT *
						FROM new_book;`

	CreateBorrowedStatusQuery = `INSERT INTO borrowed (book_id, user_id, date_from, date_to)
									VALUES ($1, $2, $3, $4)
									ON CONFLICT (book_id)
									DO UPDATE SET
									  user_id = EXCLUDED.user_id,
									  date_from = EXCLUDED.date_from,
									  date_to = EXCLUDED.date_to;`

	CreateBookedStatusQuery = `INSERT INTO booked (book_id, user_id, date_to)
									VALUES ($1, $2, $3)
									ON CONFLICT (book_id)
									DO UPDATE SET
									  user_id = EXCLUDED.user_id,
									  date_to = EXCLUDED.date_to;`

	CreateAuthorsToBooksQuery = `INSERT INTO books2authors (book_id, author_id) VALUES ($1, $2);`
)

// UPDATE //

const (
	UpdateBooksQuery = `UPDATE books
						SET title = $1, genre = $2
						WHERE id = $3`

	UpdateAuthorQuery = `UPDATE authors SET name = $1, last_name = $2 WHERE id = $3;`
)

// DELETE //

const (
	DeleteBookQuery = `DELETE FROM books WHERE id=$1`

	DeleteAuthorQuery = `DELETE FROM authors WHERE id = $1;`
)
