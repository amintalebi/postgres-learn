package store

import (
	"context"

	"github.com/amintalebi/db-learn/domain"
)

func (p postgres) CreateOrder(ctx context.Context, order domain.Order) (uint, error) {
	var id uint
	err := p.conn.QueryRow(ctx, `
		INSERT INTO orders (customer_id, date, status, total_amount)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, order.CustomerID, order.Date, order.Status, order.TotalAmount).Scan(&id)

	return id, err
}
