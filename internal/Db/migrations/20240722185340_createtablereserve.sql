-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reserve(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    tableId VARCHAR(100) NOT NULL FOREIGN KEY,
    userId VARCHAR(100) NOT NULL FOREIGN KEY,
    data TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reserve
-- +goose StatementEnd
