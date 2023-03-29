package store

import (
	"context"

	"github.com/amintalebi/db-learn/domain"
)

func (p postgres) CreateOrderDetail(ctx context.Context, orderDetail domain.OrderDetail) (uint, error) {
	var id uint
	err := p.conn.QueryRow(ctx, `
		INSERT INTO order_details (order_id, book_id, quantity, unit_price)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, orderDetail.OrderID, orderDetail.BookID, orderDetail.Quantity, orderDetail.UnitPrice).Scan(&id)

	return id, err
}
