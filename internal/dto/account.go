package dto

import (
	"time"

	"github.com/ftheodoro/go-gateway/internal/domain"
)

type CreateAccount struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type AccountReponse struct {
	ID        string    `json:id`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAd time.Time `json:"updated_at"`
}

func ToAccount(account CreateAccount) *domain.Account {
	return domain.NewAccount(account.Name, account.Email)
}
func FromAccount(account *domain.Account) *AccountReponse {
	return &AccountReponse{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		APIKey:    account.ApiKey,
		CreatedAt: account.CreatedAt,
		UpdatedAd: account.UpdatedAt,
	}
}
