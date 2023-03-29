CREATE TABLE IF NOT EXISTS customers (
 id SERIAL PRIMARY KEY,
 first_name VARCHAR(255) NOT NULL,
 last_name VARCHAR(255) NOT NULL,
 email VARCHAR(255) NOT NULL,
 address VARCHAR(255),
 phone_number VARCHAR(255) NOT NULL,
 attributes json 
);
