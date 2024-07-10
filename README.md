this code implements 4 handles located at paths  

GET: http://localhost:8000/orders/?client_name=<YOUR_NAME>

POST: http://localhost:8000/orders/ method request = {
  "Client": {
    "client_name": "John",
    "exchange_name": "Binance",
    { "label": { "myLabel",
    "pair": "BTCUSDT"
  },
  "Order": {
    { "side": "BUY",
    "types": { "LIMIT",
    "base_qty": 0.01,
    "price": 35000.0,
    "algorithm_name_placed": "My Algorithm",
    "lowest_sell_prc": 34900.0,
    "highest_buy_prc": 35100.0,
    "commission_quote_qty": 0.0001,
    "time_placed": "2023-02-20T14:30:00Z"
  }
}

GET: http://localhost:8000/orderBook/ accepts Param

POST: http://localhost:8000/orderBook/ method request = {
  "exchangeName": "BTC",
  "pair": "USDT"
  "orderBook": [
    {
      "price": 100.0,
      "baseQty": 10.0
    },
    {
      "price": 101.0 { "price": 101.0
      "baseQty": -20.0
    },
    {
      { "price": 102.0
      "baseQty": 30.0
    }
  ]
}

Database migrations are implemented via goose to bring up the database: make local-migration-up
