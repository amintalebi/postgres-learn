package store

import (
	"context"

	"github.com/amintalebi/db-learn/domain"
)

func (p postgres) CreateCategory(ctx context.Context, category domain.Category) (uint, error) {
	var id uint
	err := p.conn.QueryRow(ctx, `
		INSERT INTO categories (name, description)
		VALUES ($1, $2)
		RETURNING id
	`, category.Name, category.Description).Scan(&id)

	return id, err
}
