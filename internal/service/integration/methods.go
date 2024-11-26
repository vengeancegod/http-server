package integration

import (
	"errors"
	"http-server/internal/entities"
)

func (s *Service) CreateIntegration(integration entities.AccountIntegration) error {
	err := s.accountIntegrationRepository.CreateIntegration(integration)
	if err != nil {
		return errors.New(entities.ErrCreateInt)
	}

	return nil
}

func (s *Service) GetAllIntegrations() ([]entities.AccountIntegration, error) {
	integration, err := s.accountIntegrationRepository.GetAllIntegrations()
	if err != nil {
		return nil, errors.New(entities.ErrNotFoundInt)
	}

	return integration, nil
}

func (s *Service) UpdateIntegration(id int64, integration entities.AccountIntegration) error {
	err := s.accountIntegrationRepository.UpdateIntegration(id, integration)

	if err != nil {
		return errors.New(entities.ErrCreateInt)
	}

	return nil
}

func (s *Service) DeleteIntegration(id int64) error {
	err := s.accountIntegrationRepository.DeleteIntegration(id)

	if err != nil {
		return errors.New(entities.ErrFailedDeleteI)
	}

	return nil
}
