DROP TABLE IF EXISTS books2authors;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS borrowed;
DROP TABLE IF EXISTS booked;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS archive;


CREATE TABLE authors
(
    id        UUID DEFAULT gen_random_uuid(),
    name      TEXT NOT NULL,
    last_name TEXT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE genres
(
    id SERIAL,
    genre TEXT,
    PRIMARY KEY (id)
);


CREATE TABLE books
(
    id    UUID DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    genre_id INT NOT NULL,
    CONSTRAINT fk_genre FOREIGN KEY (genre_id) REFERENCES genres (id) ON DELETE CASCADE,
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
    user_id   TEXT NOT NULL,
    date_from DATE,
    date_to   DATE,
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (login_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, book_id)
);

CREATE TABLE booked
(
    book_id UUID NOT NULL,
    user_id TEXT NOT NULL,
    date_to DATE,
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (login_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, book_id)
);

CREATE TABLE archive
(
    book_id UUID NOT NULL,
    user_id TEXT NOT NULL,
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (login_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, book_id)
);
