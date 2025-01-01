package unisender_integration

import (
	"http-server/internal/repository"
)

type Service struct {
	unisenderIntegrationRepository repository.UnisenderIntegrationRepository
}

func NewService(unisenderIntegrationRepository repository.UnisenderIntegrationRepository) *Service {
	return &Service{
		unisenderIntegrationRepository: unisenderIntegrationRepository,
	}
}
