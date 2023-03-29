CREATE TABLE IF NOT EXISTS book_reviews (
 id SERIAL PRIMARY KEY,
 book_id INT NOT NULL,
 rate INT NOT NULL,
 FOREIGN KEY (book_id) REFERENCES books(id)
);
