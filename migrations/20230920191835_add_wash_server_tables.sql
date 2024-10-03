-- +goose Up
-- +goose StatementBegin

create table wash_servers (
    id              uuid                            PRIMARY KEY,
    title           TEXT    NOT NULL DEFAULT '',
    description     TEXT    NOT NULL DEFAULT '',
    group_id        uuid    NOT NULL                REFERENCES groups(id) ON DELETE RESTRICT,
    deleted         BOOLEAN NOT NULL DEFAULT false,
    version         BIGINT  NOT NULL DEFAULT 0      CHECK (version >= 0)
);

ALTER TABLE transactions ADD FOREIGN KEY (wash_server_id) REFERENCES wash_servers(id);
ALTER TABLE transactions DROP CONSTRAINT transactions_amount_check;
ALTER TABLE transactions ADD CONSTRAINT transactions_amount_check CHECK (amount >= 0);

ALTER TABLE organizations ALTER COLUMN version SET DEFAULT 0;
ALTER TABLE groups ALTER COLUMN version SET DEFAULT 0;
ALTER TABLE users ALTER COLUMN version SET DEFAULT 0;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE organizations ALTER COLUMN version SET DEFAULT 1;
ALTER TABLE groups ALTER COLUMN version SET DEFAULT 1;
ALTER TABLE users ALTER COLUMN version SET DEFAULT 1;
ALTER TABLE transactions DROP CONSTRAINT transactions_amount_check;
ALTER TABLE transactions ADD CONSTRAINT transactions_amount_check CHECK (amount > 0);
ALTER TABLE transactions DROP CONSTRAINT transactions_wash_servers_id_fk;
DROP TABLE IF EXISTS wash_servers;

-- +goose StatementEnd
