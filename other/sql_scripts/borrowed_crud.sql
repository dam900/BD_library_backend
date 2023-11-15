-- CREATE
INSERT INTO borrowed (book_id, user_id, date_from, date_to)
VALUES (2, 'Jowal', '2023-11-05','2023-11-05');

-- READ unused

-- UPDATE unused

-- DELETE
DELETE FROM borrowed WHERE user_id = 'Jowal' AND book_id = 1;