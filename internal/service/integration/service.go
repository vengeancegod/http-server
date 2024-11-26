package integration

import (
	"http-server/internal/repository"
)

type Service struct {
	accountIntegrationRepository repository.AccountIntegrationRepository
}

func NewService(accountIntegrationRepository repository.AccountIntegrationRepository) *Service {
	return &Service{
		accountIntegrationRepository: accountIntegrationRepository,
	}
}
