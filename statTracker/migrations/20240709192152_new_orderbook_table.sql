-- +goose Up
CREATE TABLE order_book (
                            id SERIAL PRIMARY KEY,
                            exchange VARCHAR(255),
                            pair VARCHAR(255),
                            asks JSONB,
                            bids JSONB
);

-- +goose Down
DROP TABLE order_book;
