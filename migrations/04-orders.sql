CREATE TABLE IF NOT EXISTS orders (
 id SERIAL PRIMARY KEY,
 customer_id INT NOT NULL,
 date DATE NOT NULL,
 total_amount DECIMAL(10,2) NOT NULL,
 status VARCHAR(255) NOT NULL,
 FOREIGN KEY (customer_id) REFERENCES customers(id)
);
