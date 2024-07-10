-- +goose Up
CREATE TABLE order_history (
                client_name                 VARCHAR(255),
                exchange_name               VARCHAR(255),
                label                       VARCHAR(255),
                pair                        VARCHAR(255),
                side                        VARCHAR(255),
                types                       VARCHAR(255),
                base_qty                    DECIMAL(18, 8),
                price                       DECIMAL(18, 8),
                algorithm_name_placed       VARCHAR(255),
                lowest_sell_prc             DECIMAL(18, 8),
                highest_buy_prc             DECIMAL(18, 8),
                commission_quote_qty        DECIMAL(18, 8),
                time_placed                 TIMESTAMP
);

-- +goose Down
DROP TABLE order_history;