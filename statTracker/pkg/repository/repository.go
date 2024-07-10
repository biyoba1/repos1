package repository

import (
	"github.com/biyoba1/statistic_service"
	"github.com/jmoiron/sqlx"
)

type OrderBookService interface {
	GetOrderBook(exchange_name, pair string) ([]*statTracker.OrderBook, error)
	SaveOrderBook(exchangeName, pair string, orderBook []*statTracker.DepthOrder) error
	GetOrderHistory(client *statTracker.Client) ([]*statTracker.HistoryOrder, error)
	SaveOrder(client *statTracker.Client, order *statTracker.HistoryOrder) error
}

type Repository struct {
	OrderBookService
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		OrderBookService: NewTrackerOrderPostgres(db),
	}
}
