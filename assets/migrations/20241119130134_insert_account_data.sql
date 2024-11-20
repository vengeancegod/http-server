-- +goose Up
-- +goose StatementBegin
INSERT INTO account(id, access_token, refresh_token, expires) VALUES
(1, 'access_token_1', 'refresh_token_1', '10'),
(2, 'access_token_2', 'refresh_token_2', '20'),
(3, 'access_token_3', 'refresh_token_3', '30'),
(4, 'access_token_4', 'refresh_token_4', '40'),
(5, 'access_token_5', 'refresh_token_5', '50');
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
