package sqlite

import (
	"database/sql"

	"http-server/internal/entities"
)

type AccountIntegrationEntity struct {
	DB *sql.DB
}

func (ae *AccountIntegrationEntity) GetAllIntegrations() ([]entities.AccountIntegration, error) {
	statement := `SELECT id, account_id, secret_key, client_id, redirect_url, authentication_code, authorization_code FROM account_integration ORDER BY id DESC`
	rows, err := ae.DB.Query(statement)
	if err != nil {
		return nil, err
	}

	accountIntegration := []entities.AccountIntegration{}
	for rows.Next() {
		ai := entities.AccountIntegration{}
		err := rows.Scan(&ai.Id, &ai.AccountId, &ai.SecretKey, &ai.ClientId, &ai.RedirectURL, &ai.AuthenticationCode, &ai.AuthorizationCode)
		if err != nil {
			return nil, err
		}

		accountIntegration = append(accountIntegration, ai)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accountIntegration, nil
}

func (ae *AccountIntegrationEntity) CreateIntegration(AccountId, SecretKey, ClientId, RedirectURL, AuthenticationCode, AuthorizationCode string) error {
	statement := `INSERT INTO account_integration (account_id, secret_key, client_id, redirect_url, authentication_code, authorization_code) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := ae.DB.Exec(statement, AccountId, SecretKey, ClientId, RedirectURL, AuthenticationCode, AuthorizationCode)
	return err
}
