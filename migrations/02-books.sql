CREATE TABLE IF NOT EXISTS books (
 id SERIAL PRIMARY KEY,
 title VARCHAR(255) NOT NULL,
 author VARCHAR(255) NOT NULL,
 ISBN VARCHAR(255),
 publisher VARCHAR(255) NOT NULL,
 publication_date DATE NOT NULL,
 price DECIMAL(10,2) NOT NULL,
 stock INT NOT NULL,
 category_id INT NOT NULL,
 FOREIGN KEY (category_id) REFERENCES categories(id)
);
