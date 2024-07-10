package service

import (
	"github.com/biyoba1/statistic_service"
	"github.com/biyoba1/statistic_service/pkg/repository"
)

type OrderBookService interface {
	GetOrderBook(exchange_name, pair string) ([]*statTracker.OrderBook, error)
	SaveOrderBook(exchange_name, pair string, orderBook []*statTracker.DepthOrder) error
	GetOrderHistory(client *statTracker.Client) ([]*statTracker.HistoryOrder, error)
	SaveOrder(client *statTracker.Client, order *statTracker.HistoryOrder) error
}

type Service struct {
	OrderBookService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		OrderBookService: NewTrackerOrderService(repos.OrderBookService),
	}
}
