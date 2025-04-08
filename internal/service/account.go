package service

import (
	"github.com/ftheodoro/go-gateway/internal/domain"
	"github.com/ftheodoro/go-gateway/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}
func (s *AccountService) CreateNewAccount(input dto.CreateAccount) (*dto.AccountReponse, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.ApiKey)

	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicateAPIKey
	}

	err = s.repository.Save(account)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}
func (s *AccountService) UpdateBalance(apikey string, amount float64) (*dto.AccountReponse, error) {
	account, err := s.repository.FindByAPIKey(apikey)
	if err != nil {
		return nil, err
	}
	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}
func (s *AccountService) FindByAPIKey(apikey string) (*dto.AccountReponse, error) {
	account, err := s.repository.FindByAPIKey(apikey)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}

func (s *AccountService) FindByID(id string) (*dto.AccountReponse, error) {
	account, err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}
