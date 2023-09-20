-- +goose Up
-- +goose StatementBegin

CREATE TYPE USER_ROLE_ENUM AS ENUM ('system_manager', 'admin');

create table organizations
(
    id          uuid    PRIMARY KEY,
    name        TEXT    NOT NULL,
    description TEXT,
    deleted     BOOLEAN NOT NULL
);

CREATE TABLE users (
    id              TEXT           PRIMARY KEY,
    name            TEXT           NOT NULL,
    email           TEXT           NOT NULL,
    role            USER_ROLE_ENUM NOT NULL,
    organization_id uuid           REFERENCES organizations(id),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS organizations;
DROP TYPE USER_ROLE_ENUM;
-- +goose StatementEnd
