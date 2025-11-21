-- +goose Up
CREATE TABLE user (
    id UUID 
    created_at TIMESTAMP
    updated_at TIMESTAMP
    name string
)

-- +goose Down
DROP TABLE users;
