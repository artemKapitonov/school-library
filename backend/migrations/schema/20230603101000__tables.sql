-- +goose Up
CREATE TABLE students
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    class TEXT NOT NULL
);

CREATE TABLE books
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author TEXT NOT NULL,
    title TEXT NOT NULL,
    year INTEGER,
    count INTEGER
);

CREATE TABLE students_books
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    student_id INTEGER  REFERENCES students (id),
    book_id INTEGER REFERENCES books (id)
);

-- +goose Down

DROP TABLE students;

DROP TABLE books;

DROP TABLE students_books;