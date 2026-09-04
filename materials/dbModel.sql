CREATE TABLE authors
(
    author_id  SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name  VARCHAR(50) NOT NULL
);

CREATE TABLE books
(
    book_id      SERIAL PRIMARY KEY,
    title        VARCHAR(150) NOT NULL,
    author_id    INT REFERENCES authors (author_id),
    publish_year INT
);

CREATE TABLE readers
(
    reader_id SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    phone     VARCHAR(20),
    reg_date  DATE DEFAULT CURRENT_DATE
);

CREATE TABLE loans
(
    loan_id     SERIAL PRIMARY KEY,
    book_id     INT REFERENCES books (book_id),
    reader_id   INT REFERENCES readers (reader_id),
    loan_date   DATE DEFAULT CURRENT_DATE,
    return_date DATE
);