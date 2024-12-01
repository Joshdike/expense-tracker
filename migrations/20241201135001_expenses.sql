-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS expenses (
    transactionid SERIAL PRIMARY KEY,
    date DATE,
    amount NUMERIC(10,2),
    category VARCHAR(50),
    description TEXT,
    payment_method VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS expenses;
-- +goose StatementEnd
