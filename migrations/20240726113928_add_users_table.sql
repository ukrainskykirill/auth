-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM ('ADMIN', 'USER', 'UNKNOW');
CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(5000),
    email varchar(5000),
    password TEXT,
    role user_role,
    created_at timestamp,
    updated_at timestamp
);
ALTER TABLE users ADD CONSTRAINT unique_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TYPE user_role;
-- +goose StatementEnd
