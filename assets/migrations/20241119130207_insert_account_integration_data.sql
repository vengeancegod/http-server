-- +goose Up
-- +goose StatementBegin
INSERT INTO account_integration (account_id, secret_key, client_id, redirect_url, authentication_code, authorization_code) VALUES
(1, 'secret_key_1', 'client_id_1', 'https://example.com/redirect1', 'auth_code_1', 'auth_code_1'),
(2, 'secret_key_2', 'client_id_2', 'https://example.com/redirect2', 'auth_code_2', 'auth_code_2'),
(3, 'secret_key_3', 'client_id_3', 'https://example.com/redirect3', 'auth_code_3', 'auth_code_3'),
(4, 'secret_key_4', 'client_id_4', 'https://example.com/redirect4', 'auth_code_4', 'auth_code_4'),
(5, 'secret_key_5', 'client_id_5', 'https://example.com/redirect5', 'auth_code_5', 'auth_code_5');
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
