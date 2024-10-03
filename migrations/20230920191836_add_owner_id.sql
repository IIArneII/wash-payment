-- +goose Up
-- +goose StatementBegin

ALTER TABLE wash_servers ADD COLUMN owner_id TEXT REFERENCES users(id);
ALTER TABLE wash_servers RENAME COLUMN title TO name;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE wash_servers DROP COLUMN owner_id;
ALTER TABLE wash_servers RENAME COLUMN name TO title;

-- +goose StatementEnd
