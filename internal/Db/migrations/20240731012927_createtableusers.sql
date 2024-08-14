-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    phoneNumber VARCHAR(30) NOT NULL,
    isAdmin BOOLEAN NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user
-- +goose StatementEnd
