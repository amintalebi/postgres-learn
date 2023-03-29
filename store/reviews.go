package store

import (
	"context"

	"github.com/amintalebi/db-learn/domain"
)

func (p postgres) CreateReview(ctx context.Context, review domain.Review) (uint, error) {
	var id uint
	err := p.conn.QueryRow(ctx, `
		INSERT INTO book_reviews (book_id, rate)
		VALUES ($1, $2)
		RETURNING id
	`, review.BookID, review.Rate).Scan(&id)

	return id, err
}
