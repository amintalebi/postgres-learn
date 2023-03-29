package store

import (
	"context"

	"github.com/amintalebi/db-learn/domain"
)

func (p postgres) CreateCustomer(ctx context.Context, customer domain.Customer) (uint, error) {
	var id uint
	err := p.conn.QueryRow(ctx, `
		INSERT INTO customers (first_name, last_name, email, address, phone_number, attributes)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, customer.FirstName, customer.LastName, customer.Email, customer.Address, customer.PhoneNumber, customer.Attributes).Scan(&id)

	return id, err
}
