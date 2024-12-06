package account

import (
	"database/sql"
	"errors"
	"http-server/internal/entities"
	"http-server/internal/infrastructure/database"
	rep "http-server/internal/repository"
)

var _ rep.AccountRepository = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository() (*repository, error) {
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}
	return &repository{
		db: db,
	}, nil
}

func (repo *repository) GetAllContacts() ([]entities.Contacts, error) {
	query := "SELECT name, email FROM contacts"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []entities.Contacts
	for rows.Next() {
		var contact entities.Contacts
		if err := rows.Scan(&contact.Name, &contact.Email); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(contacts) == 0 {
		return nil, errors.New(entities.ErrNotFoundContacts)
	}

	return contacts, nil
}
func (repo *repository) Authorization(request entities.AuthRequest) (entities.AuthResponse, error) {

	query := "SELECT access_token, refresh_token, expires FROM accounts WHERE code = ?"
	row := repo.db.QueryRow(query, request.Code)

	var accessToken, refreshToken string
	var expires int64

	if err := row.Scan(&accessToken, &refreshToken, &expires); err != nil {
		if err == sql.ErrNoRows {
			return entities.AuthResponse{}, errors.New(entities.ErrNotFoundAcc)
		}
		return entities.AuthResponse{}, err
	}

	return entities.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(expires),
	}, nil
}

func (repo *repository) CreateAccount(account entities.Account) error {

	query := "INSERT INTO accounts (access_token, refresh_token, expires) VALUES (?, ?, ?)"
	_, err := repo.db.Exec(query, account.AccessToken, account.RefreshToken, account.Expires)
	return err
}

func (repo *repository) GetAllAccounts() ([]entities.Account, error) {
	query := "SELECT * FROM accounts"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []entities.Account
	for rows.Next() {
		var account entities.Account
		if err := rows.Scan(&account.ID, &account.AccessToken, &account.RefreshToken, &account.Expires); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(accounts) == 0 {
		return nil, errors.New(entities.ErrNotFoundAcc)
	}

	return accounts, nil
}
