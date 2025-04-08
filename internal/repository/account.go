package repository

import (
	"database/sql"
	"time"

	"github.com/ftheodoro/go-gateway/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}
func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare("INSERT INTO accounts (id, name, email, api_key,balance,created_at,updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.ID, account.Name, account.Email, account.ApiKey, account.Balance, account.CreatedAt, account.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account

	row := r.db.QueryRow("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = $1", apiKey)

	err := row.Scan(&account.ID, &account.Name, &account.Email, &account.ApiKey, &account.Balance, &account.CreatedAt, &account.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	return &account, nil
}
func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	var account domain.Account

	row := r.db.QueryRow("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE id = $1", id)

	err := row.Scan(&account.ID, &account.Name, &account.Email, &account.ApiKey, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	return &account, nil
}
func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var currentBalance float64

	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", account.ID).Scan(&currentBalance)
	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
		UPADTE accounts
		SET balance = $1,updated_at = $2
		WHERE id = $3
	`, account.Balance, time.Now(), account.ID)

	if err != nil {
		return err
	}

	return nil
}
