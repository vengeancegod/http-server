-- +goose Up
-- +goose StatementBegin
CREATE TABLE account_integration (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    secret_key TEXT NOT NULL,
    client_id TEXT NOT NULL,
    redirect_url TEXT NOT NULL,
    authentication_code TEXT NOT NULL,
    authorization_code TEXT NOT NULL,
    FOREIGN KEY(account_id) REFERENCES account(id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account_integration
-- +goose StatementEnd
