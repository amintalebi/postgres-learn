package store

import (
	"context"

	"github.com/amintalebi/db-learn/domain"
)

func (p postgres) CreateLog(ctx context.Context, log domain.Log) error {
	_, err := p.conn.Exec(ctx, `
		INSERT INTO logs (date, content)
		VALUES ($1, $2)
	`, log.Date, log.Content)

	return err
}
