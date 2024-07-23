-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reserve(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phoneNumber VARCHAR(30) NOT NULL,
    tableId VARCHAR(50) NOT NULL,
    data TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reserve
-- +goose StatementEnd
