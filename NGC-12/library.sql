CREATE DATABASE IF NOT EXISTS library;

CREATE TABLE IF NOT EXISTS authors (
    author_id INTEGER AUTO_INCREMENT,
    author_name VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    
    PRIMARY KEY(author_id)
);

CREATE TABLE IF NOT EXISTS books (
    book_id INTEGER AUTO_INCREMENT,
    author_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    publication_year YEAR NOT NULL,
    available_quantity INTEGER NOT NULL,
    
    PRIMARY KEY(book_id),
    FOREIGN KEY(author_id) REFERENCES authors(author_id)
);

CREATE TABLE IF NOT EXISTS book_loans (
    loan_id INTEGER AUTO_INCREMENT,
    book_id INTEGER NOT NULL,
    borrower_name VARCHAR(255) NOT NULL,
    loan_date DATE NOT NULL,
    return_date DATE NOT NULL,

    PRIMARY KEY(loan_id),
    FOREIGN KEY(book_id) REFERENCES books(book_id)
);

-- Add appropriate constraints to ensure that book titles and author names are unique.
ALTER TABLE books ADD UNIQUE (title);
ALTER TABLE authors ADD UNIQUE (author_name);


-- Create indexes to optimize the retrieval of books based on the author's ID and to find overdue book loans efficiently.
CREATE INDEX book_by_author_id ON books (author_id);
CREATE INDEX overdue_loans ON book_loans(return_date) WHERE return_date < CURRENT_DATE();