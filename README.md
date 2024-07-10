

**Order Book API**
================

### Endpoints

#### GET /orders/

Retrieve a list of orders for a specific client.

**Query Parameters**

* `client_name`: The name of the client to retrieve orders for.

**Example Request**

```
GET http://localhost:8000/orders/?client_name=John
```

#### POST /orders/

Create a new order.

**Request Body**

```json
{
  "Client": {
    "client_name": "John",
    "exchange_name": "Binance",
    "label": {
      "myLabel",
      "pair": "BTCUSDT"
    },
    "Order": {
      "side": "BUY",
      "types": {
        "LIMIT",
        "base_qty": 0.01,
        "price": 35000.0,
        "algorithm_name_placed": "My Algorithm",
        "lowest_sell_prc": 34900.0,
        "highest_buy_prc": 35100.0,
        "commission_quote_qty": 0.0001,
        "time_placed": "2023-02-20T14:30:00Z"
      }
    }
  }
}
```

#### GET /orderBook/

Retrieve the order book for a specific exchange and pair.

**Query Parameters**

* `exchangeName`: The name of the exchange to retrieve the order book for.
* `pair`: The pair to retrieve the order book for.

**Example Request**

```
GET http://localhost:8000/orderBook/?exchange=BTC&pair=USDT
```

#### POST /orderBook/

Create a new order book.

**Request Body**

```json
{
  "exchangeName": "BTC",
  "pair": "USDT",
  "orderBook": [
    {
      "price": 100.0,
      "baseQty": 10.0
    },
    {
      "price": 101.0,
      "baseQty": -20.0
    },
    {
      "price": 102.0,
      "baseQty": 30.0
    }
  ]
}
```

### Database Migrations

To bring up the database, run the following command:

```
make local-migration-up
```

This will apply the database migrations using Goose.

