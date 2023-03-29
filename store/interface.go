package store

import (
	"context"

	"github.com/amintalebi/db-learn/domain"
)

type IStore interface {
	Migrateable
	IBook
	ICustomer
	IOrder
	IOrderDetail
	ICatergory
	IReview
}

type Migrateable interface {
	Migrate(path string) error
}

type IBook interface {
	CreateBook(ctx context.Context, book domain.Book) (uint, error)
}

type ICustomer interface {
	CreateCustomer(ctx context.Context, customer domain.Customer) (uint, error)
}

type IOrder interface {
	CreateOrder(ctx context.Context, order domain.Order) (uint, error)
}

type IOrderDetail interface {
	CreateOrderDetail(ctx context.Context, orderDetail domain.OrderDetail) (uint, error)
}

type ICatergory interface {
	CreateCategory(ctx context.Context, category domain.Category) (uint, error)
}

type IReview interface {
	CreateReview(ctx context.Context, review domain.Review) (uint, error)
}
