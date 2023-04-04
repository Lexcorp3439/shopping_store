-- +goose Up
-- +goose StatementBegin

CREATE
    EXTENSION IF NOT EXISTS pgcrypto;

CREATE
    OR REPLACE FUNCTION set_updated_at_column() RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at
        = now() at time zone 'utc';
    RETURN NEW;
END;
$$
    language 'plpgsql';

CREATE TABLE shopping_store
(
    id         serial                                 not null,
    user_id    bigint                                 not null,
    staff_id   bigint                                 not null,
    count      int                                    not null,
    created_at timestamp with time zone DEFAULT now() not null,
    updated_at timestamp with time zone DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE shopping_store;
-- +goose StatementEnd