DROP TABLE IF EXISTS books2authors;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS borrowed;
DROP TABLE IF EXISTS booked;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS users;


CREATE TABLE authors
(
    id        UUID DEFAULT gen_random_uuid(),
    name      TEXT NOT NULL,
    last_name TEXT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE books
(
    id    UUID DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    genre TEXT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE books2authors
(
    book_id   UUID NOT NULL,
    author_id UUID NOT NULL,
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, author_id)
);

CREATE TABLE users
(
    name     TEXT NOT NULL,
    lastname TEXT NOT NULL,
    login_id TEXT NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY (login_id)
);

CREATE TABLE borrowed
(
    book_id   UUID NOT NULL,
    user_id   TEXT DEFAULT NULL,
    date_from DATE DEFAULT NULL,
    date_to   DATE DEFAULT NULL,
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (login_id) ON DELETE CASCADE,
    PRIMARY KEY (book_id)
);

CREATE TABLE booked
(
    book_id UUID NOT NULL,
    user_id TEXT DEFAULT NULL,
    date_to DATE DEFAULT NULL,
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (login_id) ON DELETE CASCADE,
    PRIMARY KEY (book_id)
);
