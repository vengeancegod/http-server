package sqlite

import (
	"database/sql"

	"http-server/internal/entities"
)

type AccountEntity struct {
	DB *sql.DB
}

func (ae *AccountEntity) GetAll() ([]entities.Account, error) {
	statement := `SELECT id, access_token, refresh_token, expires FROM account ORDER BY id DESC`
	rows, err := ae.DB.Query(statement)
	if err != nil {
		return nil, err
	}

	accounts := []entities.Account{}
	for rows.Next() {
		a := entities.Account{}
		err := rows.Scan(&a.Id, &a.AccessToken, &a.RefreshToken, &a.Expires)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (ae *AccountEntity) CreateAccount(AccessToken, RefreshToken string, Expires int64) error {
	statement := `INSERT INTO account (access_token, refresh_token, expires) VALUES (?, ?, ?)`
	_, err := ae.DB.Exec(statement, AccessToken, RefreshToken, Expires)
	return err
}
