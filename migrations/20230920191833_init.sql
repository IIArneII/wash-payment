-- +goose Up
-- +goose StatementBegin

ALTER DATABASE wash_payment SET default_transaction_isolation = 'serializable';

CREATE TYPE USER_ROLE_ENUM              AS ENUM ('system_manager', 'admin', 'no_access');
CREATE TYPE TRANSACTIONS_OPERATION_ENUM AS ENUM ('deposit', 'debit');

create table organizations (
    id           uuid                             PRIMARY KEY,
    name         TEXT    NOT NULL  DEFAULT '',
    display_name TEXT    NOT NULL                 UNIQUE,
    description  TEXT    NOT NULL  DEFAULT '',
    balance      BIGINT  NOT NULL  DEFAULT 0      CHECK (balance >= 0),
    deleted      BOOLEAN NOT NULL  DEFAULT false,
    version      BIGINT  NOT NULL  DEFAULT 1      CHECK (version >= 0)
);

create table groups (
    id              uuid                            PRIMARY KEY,
    organization_id uuid    NOT NULL                REFERENCES organizations(id) ON DELETE RESTRICT,
    name            TEXT    NOT NULL DEFAULT '',
    description     TEXT    NOT NULL DEFAULT '',
    deleted         BOOLEAN NOT NULL DEFAULT false,
    version         BIGINT  NOT NULL DEFAULT 1      CHECK (version >= 0)
);

CREATE TABLE users (
    id              TEXT                                 PRIMARY KEY,
    name            TEXT           NOT NULL  DEFAULT '',
    email           TEXT           NOT NULL  DEFAULT '',
    role            USER_ROLE_ENUM NOT NULL,
    organization_id uuid                                 REFERENCES organizations(id) ON DELETE RESTRICT,
    version         BIGINT         NOT NULL  DEFAULT 1   CHECK (version >= 0)
);

create table transactions (
    id              uuid                                     PRIMARY KEY,
    organization_id uuid                        NOT NULL     REFERENCES organizations(id) ON DELETE RESTRICT,
    amount          BIGINT                      NOT NULL     CHECK (amount > 0),
    operation       TRANSACTIONS_OPERATION_ENUM NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE    NOT NULL     DEFAULT NOW(),
    sevice          TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS organizations;
DROP TYPE  IF EXISTS TRANSACTIONS_OPERATION_ENUM;
DROP TYPE  IF EXISTS USER_ROLE_ENUM;
ALTER DATABASE wash_payment SET default_transaction_isolation = 'read committed';

-- +goose StatementEnd
