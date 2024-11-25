package account

import (
	"http-server-fixed/internal/repository"
)

type Service struct {
	accountRepository repository.AccountRepository
}

func NewService(accountRepository repository.AccountRepository) *Service {
	return &Service{
		accountRepository: accountRepository,
	}
}
