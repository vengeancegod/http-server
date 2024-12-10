package unisender_integration

import (
	"errors"
	"http-server/internal/entities"
)

func (s *Service) GetUnisenderKey() ([]entities.UnisenderIntegration, error) {
	unisenderKey, err := s.unisenderIntegrationRepository.GetUnisenderKey()
	if err != nil {
		return nil, errors.New(entities.ErrGetUnisenderKey)
	}

	return unisenderKey, nil
}

func (s *Service) SaveUnisenderKey(unisenderKey *entities.UnisenderIntegration) error {
	err := s.unisenderIntegrationRepository.SaveUnisenderKey(unisenderKey)
	if err != nil {
		return errors.New(entities.ErrSaveUnisenderKey)
	}
	return nil
}
