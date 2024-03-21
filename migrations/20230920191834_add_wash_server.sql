-- +goose Up
-- +goose StatementBegin

ALTER TABLE transactions ADD COLUMN wash_server_id uuid;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE transactions DROP COLUMN wash_server_id;

-- +goose StatementEnd
