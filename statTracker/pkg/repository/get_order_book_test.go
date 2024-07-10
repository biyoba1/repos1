package repository

import (
	statTracker "github.com/biyoba1/statistic_service"
	_ "github.com/lib/pq"
	"testing"
)

func TestGetOrderBook(t *testing.T) {
	db, err := NewPostgresDB(Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewTrackerOrderPostgres(db)

	exchangeName := "BTC"
	pair := "USDT"

	orderBooks, err := repo.GetOrderBook(exchangeName, pair)
	if err != nil {
		t.Errorf("GetOrderBook returned error: %v", err)
	}

	if len(orderBooks) == 0 {
		t.Errorf("GetOrderBook returned no order books")
	}

	for _, orderBook := range orderBooks {
		if orderBook.Exchange != exchangeName {
			t.Errorf("Order book exchange does not match: %s != %s", orderBook.Exchange, exchangeName)
		}
		if orderBook.Pair != pair {
			t.Errorf("Order book pair does not match: %s != %s", orderBook.Pair, pair)
		}
	}
}

func TestGetOrderHistory(t *testing.T) {
	db, err := NewPostgresDB(Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewTrackerOrderPostgres(db)

	client := &statTracker.Client{
		Client_name: "John",
	}

	orders, err := repo.GetOrderHistory(client)
	if err != nil {
		t.Errorf("GetOrderHistory returned error: %v", err)
	}

	if len(orders) == 0 {
		t.Errorf("GetOrderHistory returned no orders")
	}

	for _, order := range orders {
		if order.Client_name != client.Client_name {
			t.Errorf("Order client name does not match: %s != %s", order.Client_name, client.Client_name)
		}
	}
}
