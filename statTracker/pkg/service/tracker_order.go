package service

import (
	statTracker "github.com/biyoba1/statistic_service"
	"github.com/biyoba1/statistic_service/pkg/repository"
)

type TrackerOrderService struct {
	repo repository.OrderBookService
}

func NewTrackerOrderService(repo repository.OrderBookService) *TrackerOrderService {
	return &TrackerOrderService{repo: repo}
}

func (s *TrackerOrderService) GetOrderBook(exchange_name, pair string) ([]*statTracker.OrderBook, error) {
	return s.repo.GetOrderBook(exchange_name, pair)
}

func (s *TrackerOrderService) SaveOrderBook(exchange_name, pair string, orderBook []*statTracker.DepthOrder) error {
	return s.repo.SaveOrderBook(exchange_name, pair, orderBook)
}

func (s *TrackerOrderService) GetOrderHistory(client *statTracker.Client) ([]*statTracker.HistoryOrder, error) {
	return s.repo.GetOrderHistory(client)
}

func (s *TrackerOrderService) SaveOrder(client *statTracker.Client, order *statTracker.HistoryOrder) error {
	return s.repo.SaveOrder(client, order)
}
