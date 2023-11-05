-- CREATE
INSERT INTO books (title, genre) VALUES ('Władca Pierścieni', 'Fantasy');

-- READ ALL

SELECT b.id,
       b.title,
       b.genre,
       JSON_BUILD_OBJECT(
               'booked_by', b2.user_id,
               'booked_to', b2.date_to
           ) AS booked_status,
       JSON_BUILD_OBJECT(
               'borrowed_by', b3.user_id,
               'borrowed_from', b3.date_from,
               'borrowed_to', b3.date_to
           ) AS borrowed_status,
       JSON_AGG(
               JSON_BUILD_OBJECT(
                       'author_id', a.id,
                       'author_name', a.name,
                       'author_lastname', a.last_name)
           ) AS authors
FROM books AS b
         JOIN books2authors AS ba ON b.id = ba.book_id
         JOIN authors AS a ON ba.author_id = a.id
         LEFT JOIN booked b2 on b.id = b2.book_id
         LEFT JOIN borrowed b3 on b.id = b3.book_id
GROUP BY b.id, b2.date_to, b2.user_id, b3.user_id, b3.date_from, b3.date_to
OFFSET 0 LIMIT 2;

-- READ CONCRETE

SELECT b.id,
       b.title,
       b.genre,
       JSON_BUILD_OBJECT(
               'booked_by', b2.user_id,
               'booked_to', b2.date_to
           ) AS booked_status,
       JSON_BUILD_OBJECT(
               'borrowed_by', b3.user_id,
               'borrowed_from', b3.date_from,
               'borrowed_to', b3.date_to
           ) AS borrowed_status,
       JSON_AGG(
               JSON_BUILD_OBJECT(
                       'author_id', a.id,
                       'author_name', a.name,
                       'author_lastname', a.last_name)
           ) AS authors
FROM books AS b
         JOIN books2authors AS ba ON b.id = ba.book_id
         JOIN authors AS a ON ba.author_id = a.id
         LEFT JOIN booked b2 on b.id = b2.book_id
         LEFT JOIN borrowed b3 on b.id = b3.book_id
WHERE b.id = 1
GROUP BY b.id, b2.date_to, b2.user_id, b3.user_id, b3.date_from, b3.date_to;

-- UPDATE add later

-- DELETE add later

/* BOOKS 2 AUTHORS CRUD */
-- CREATE
INSERT INTO books2authors (book_id, author_id) VALUES (1,1);

-- READ unused

-- UPDATE unused

-- DELETE
DELETE FROM books2authors WHERE book_id=1 AND author_id = 1;
