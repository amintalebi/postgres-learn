package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/amintalebi/db-learn/domain"
	"github.com/amintalebi/db-learn/store"
)

// make some columns nullable
// make partial index
// measure performace
// views: live view: number of books per author
// views: counters tables, gets updated with each insertion into books, customers, ... table
// fucntions
// views: metrics, number of orders per day
// create materialized view
// create reviews table for books
// top 5 authors based on reviews including number of books, book sales
// ranking: number of books, book sales, average rating based review
// learn triggers, put triggers on every updates

func main() {
	db := store.NewPostgres("localhost:5434", "postgres", "postgres", "postgres")
	err := db.Migrate("migrations")
	if err != nil {
		log.Fatal(err)
	}

	// populateCatergories(db)
	// populateBooks(db)
	// populateCustomers(db)
	populateOrders(db)
	// populateBookReviews(db)
}

func populateCatergories(db store.ICatergory) {
	for name, desc := range bookCategories {
		_, err := db.CreateCategory(context.Background(), domain.Category{Name: name, Description: desc})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func populateBooks(db store.IBook) {
	for i := 0; i < 10000; i++ {
		f, l := GenerateFakeName()
		var isbn *string
		if i%20 != 0 {
			s := GenerateFakeISBN()
			isbn = &s
		}
		book := domain.Book{
			Title:           GenerateFakeBookTitle(),
			Author:          f + " " + l,
			ISBN:            isbn,
			Publisher:       GenerateFakePublisher(),
			PublicationDate: GenerateFakeDate(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			Price:           RandomInt(10, 40),
			Stock:           RandomInt(100, 500),
			CategoryID:      RandomInt(1, 10),
		}
		_, err := db.CreateBook(context.Background(), book)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func populateCustomers(db store.ICustomer) {
	for i := 0; i < 4000; i++ {
		f, l := GenerateFakeName()
		var address *string
		if i%20 != 0 {
			s := GenerateFakeAddress()
			address = &s
		}
		customer := domain.Customer{
			FirstName:   f,
			LastName:    l,
			Email:       f + "." + l + "@gmail.com",
			Address:     address,
			PhoneNumber: GenerateFakePhoneNumber(),
			Attributes: map[string]interface{}{
				"age":            RandomInt(10, 80),
				"height":         RandomInt(150, 200),
				"marital_status": []string{"Single", "Married", "Divorsed"}[rand.Intn(3)],
				"is_student":     []bool{true, false}[rand.Intn(2)], // make it null sometimes
			},
		}
		_, err := db.CreateCustomer(context.Background(), customer)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func populateOrders(db store.IStore) {
	statuses := []string{"SUCCESS", "FAIL", "CANCEL"}
	for i := 0; i < 500_000; i++ {
		order := domain.Order{
			CustomerID:  RandomInt(10, 4000),
			Date:        GenerateFakeDate(time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			TotalAmount: RandomInt(10, 10000),
			Status:      statuses[rand.Intn(len(statuses))],
		}
		id, err := db.CreateOrder(context.Background(), order)
		if err != nil {
			log.Fatal(err)
		}
		od := domain.OrderDetail{
			OrderID:   id,
			BookID:    RandomInt(1, 10000),
			Quantity:  RandomInt(1, 10),
			UnitPrice: float64(RandomInt(10, 100)),
		}
		_, err = db.CreateOrderDetail(context.Background(), od)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func populateBookReviews(db store.IReview) {
	for i := 0; i < 100_000; i++ {
		review := domain.Review{
			BookID: RandomInt(1, 10000),
			Rate:   RandomInt(1, 10),
		}
		_, err := db.CreateReview(context.Background(), review)
		if err != nil {
			log.Fatal(err)
		}
	}
}
