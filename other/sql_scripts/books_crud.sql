-- CREATE
With new_book as (
INSERT
INTO books (title, genre)
VALUES ('test', 'test') RETURNING id
    )
SELECT *
FROM new_book;
-- READ ALL

SELECT b.id,
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
OFFSET 0 LIMIT 100;

-- READ CONCRETE
SELECT b.id,
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
WHERE b.id = 'uuid'
GROUP BY b.id, b2.date_to, b2.user_id, b3.user_id, b3.date_from, b3.date_to;
-- UPDATE add later

UPDATE books
SET title = 'title1', genre = 'genre1'
WHERE id = 'id'

-- DELETE add later

DELETE FROM books WHERE books.id = 'uuid';

/* BOOKS 2 AUTHORS CRUD */
-- CREATE
INSERT INTO books2authors (book_id, author_id)
VALUES (1, 1);

-- READ unused

-- UPDATE unused

-- DELETE
DELETE
FROM books2authors
WHERE book_id = 1
  AND author_id = 1;
