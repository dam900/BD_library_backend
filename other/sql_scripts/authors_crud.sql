-- CREATE
WITH new_author as (
    INSERT INTO authors (name, last_name) VALUES ('J.R.R', 'Tolkien')
        RETURNING id)
SELECT id FROM new_author;

-- READ ALL
SELECT * FROM authors;

-- READ CONCRETE
SELECT * FROM authors WHERE id = 1;

-- UPDATE not important

-- DELETE not important