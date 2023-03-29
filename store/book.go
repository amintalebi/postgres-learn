package store

import (
	"context"

	"github.com/amintalebi/db-learn/domain"
)

func (p postgres) CreateBook(ctx context.Context, book domain.Book) (uint, error) {
	var id uint
	err := p.conn.QueryRow(ctx, `
		INSERT INTO books (title, author, ISBN, publisher, publication_date, price, stock, category_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`, book.Title, book.Author, book.ISBN, book.Publisher, book.PublicationDate, book.Price, book.Stock, book.CategoryID).Scan(&id)

	return id, err
}
