-- +goose Up
-- +goose StatementBegin
CREATE TABLE tables (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    number int UNIQUE NOT NULL,
    capacity int NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tables;
-- +goose StatementEnd
