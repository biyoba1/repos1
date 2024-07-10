package repository

import (
	"encoding/json"
	"fmt"
	statTracker "github.com/biyoba1/statistic_service"
	"github.com/jmoiron/sqlx"
	"log"
)

type TrackerOrderPostgres struct {
	db *sqlx.DB
}

/* todo Реализовываем эндпоинты, делает миграцию через goose (https://habr.com/ru/amp/publications/780280/), swagger, */
func NewTrackerOrderPostgres(db *sqlx.DB) *TrackerOrderPostgres {
	return &TrackerOrderPostgres{db: db}
}

func (r *TrackerOrderPostgres) GetOrderBook(exchangeName, pair string) ([]*statTracker.OrderBook, error) {
	var orderBooks []*statTracker.OrderBook
	query := fmt.Sprintf("SELECT exchange, pair, asks, bids FROM %s WHERE exchange = $1 AND pair = $2", orderBookTable)
	rows, err := r.db.Query(query, exchangeName, pair)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var exchange string
		var pair string
		var asks json.RawMessage
		var bids json.RawMessage
		err := rows.Scan(&exchange, &pair, &asks, &bids)
		if err != nil {
			return nil, err
		}
		var askSlice []*statTracker.DepthOrder
		err = json.Unmarshal(asks, &askSlice)
		if err != nil {
			return nil, err
		}
		var bidSlice []*statTracker.DepthOrder
		err = json.Unmarshal(bids, &bidSlice)
		if err != nil {
			return nil, err
		}
		orderBook := &statTracker.OrderBook{
			Exchange: exchange,
			Pair:     pair,
			Asks:     askSlice,
			Bids:     bidSlice,
		}
		orderBooks = append(orderBooks, orderBook)
	}
	return orderBooks, nil
}

func (r *TrackerOrderPostgres) SaveOrderBook(exchangeName, pair string, orderBook []*statTracker.DepthOrder) error {
	asks := make([]*statTracker.DepthOrder, 0)
	bids := make([]*statTracker.DepthOrder, 0)

	for _, order := range orderBook {
		if order.BaseQty > 0 {
			asks = append(asks, order)
		} else {
			bids = append(bids, order)
		}
	}

	asksJson, err := json.Marshal(asks)
	if err != nil {
		return err
	}
	bidsJson, err := json.Marshal(bids)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (exchange, pair, asks, bids) VALUES ($1, $2, $3, $4)", orderBookTable)

	_, err = r.db.Exec(query, exchangeName, pair, asksJson, bidsJson)
	return err
}

func (r *TrackerOrderPostgres) GetOrderHistory(client *statTracker.Client) ([]*statTracker.HistoryOrder, error) {
	log.Println("GetOrderHistory called with client:", client.Client_name)
	queryGetOrders := `SELECT * FROM order_history WHERE client_name = $1`
	rows, err := r.db.Query(queryGetOrders, client.Client_name)
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	var orders []*statTracker.HistoryOrder
	for rows.Next() {
		order := &statTracker.HistoryOrder{}
		err := rows.Scan(
			&order.Client_name,
			&order.Exchange_name,
			&order.Label,
			&order.Pair,
			&order.Side,
			&order.Types,
			&order.Base_qty,
			&order.Price,
			&order.Algorithm_name_placed,
			&order.Lowest_sell_prc,
			&order.Highest_buy_prc,
			&order.Commission_quote_qty,
			&order.Time_placed,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *TrackerOrderPostgres) SaveOrder(client *statTracker.Client, order *statTracker.HistoryOrder) error {
	querySaveOrder := `INSERT INTO order_history (
        client_name, 
        exchange_name, 
        label, 
        pair, 
        side, 
        types, 
        base_qty, 
        price, 
        algorithm_name_placed, 
        lowest_sell_prc, 
        highest_buy_prc, 
        commission_quote_qty, 
        time_placed
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := r.db.Exec(querySaveOrder,
		client.Client_name,
		client.Exchange_name,
		client.Label,
		client.Pair,
		order.Side,
		order.Types,
		order.Base_qty,
		order.Price,
		order.Algorithm_name_placed,
		order.Lowest_sell_prc,
		order.Highest_buy_prc,
		order.Commission_quote_qty,
		order.Time_placed,
	)
	return err
}
