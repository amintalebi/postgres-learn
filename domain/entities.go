package domain

import "time"

type Book struct {
	ID              uint
	Title           string
	Author          string
	ISBN            *string
	Publisher       string
	PublicationDate time.Time
	Price           int
	Stock           int
	CategoryID      int
}

type Customer struct {
	ID          uint
	FirstName   string
	LastName    string
	Email       string
	Address     *string
	PhoneNumber string
	Attributes  map[string]interface{}
}

type Order struct {
	ID          int
	CustomerID  int
	Date        time.Time
	TotalAmount int
	Status      string
}

type OrderDetail struct {
	ID        int
	OrderID   uint
	BookID    int
	Quantity  int
	UnitPrice float64
}

type Category struct {
	ID          int
	Name        string
	Description string
}

type Review struct {
	ID     int
	BookID int
	Rate   int
}
