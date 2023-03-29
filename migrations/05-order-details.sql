CREATE TABLE IF NOT EXISTS order_details (
 id SERIAL PRIMARY KEY,
 order_id INT NOT NULL,
 book_id INT NOT NULL,
 quantity INT NOT NULL,
 unit_price DECIMAL(10,2) NOT NULL,
 FOREIGN KEY (order_id) REFERENCES orders(id),
 FOREIGN KEY (book_id) REFERENCES books(id)
);
