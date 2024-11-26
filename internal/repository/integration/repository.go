package integration

import (
	"errors"
	"http-server-fixed/internal/entities"
	rep "http-server-fixed/internal/repository"
	repoModel "http-server-fixed/internal/repository/integration/model"
	"strconv"
	"sync"
)

var _ rep.AccountIntegrationRepository = (*repository)(nil)

type repository struct {
	data map[string]*repoModel.AccountIntegration
	mu   sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]*repoModel.AccountIntegration),
	}
}

func (repo *repository) CreateIntegration(integration entities.AccountIntegration) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	integrationID := strconv.FormatInt(integration.ID, 10)

	if _, exists := repo.data[integrationID]; exists {
		return nil
	}

	repo.data[integrationID] = &repoModel.AccountIntegration{
		ID:                 integration.ID,
		AccountID:          integration.AccountID,
		SecretKey:          integration.SecretKey,
		ClientID:           integration.ClientID,
		RedirectURL:        integration.RedirectURL,
		AuthenticationCode: integration.AuthenticationCode,
		AuthorizationCode:  integration.AuthorizationCode,
	}

	return nil
}

func (repo *repository) GetAllIntegrations() ([]entities.AccountIntegration, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if len(repo.data) == 0 {
		return nil, errors.New(entities.ErrNotFoundInt)
	}

	var integrations []entities.AccountIntegration

	for _, integration := range repo.data {
		integrations = append(integrations, entities.AccountIntegration(*integration))
	}

	return integrations, nil
}

func (repo *repository) UpdateIntegration(id int64, integration entities.AccountIntegration) error {
	repo.mu.Lock()
	repo.mu.Unlock()

	integrationID := strconv.FormatInt(integration.ID, 10)
	if _, exists := repo.data[integrationID]; !exists {
		return errors.New(entities.ErrCreateInt)
	}

	repo.data[integrationID] = &repoModel.AccountIntegration{
		ID:                 integration.ID,
		AccountID:          integration.AccountID,
		SecretKey:          integration.SecretKey,
		ClientID:           integration.ClientID,
		RedirectURL:        integration.RedirectURL,
		AuthenticationCode: integration.AuthenticationCode,
		AuthorizationCode:  integration.AuthorizationCode,
	}

	return nil
}

func (repo *repository) DeleteIntegration(id int64) error {
	repo.mu.Lock()
	repo.mu.Unlock()

	integrationID := strconv.FormatInt(id, 10)

	if _, exists := repo.data[integrationID]; !exists {
		return errors.New(entities.ErrFailedDeleteI)
	}

	delete(repo.data, integrationID)

	return nil
}
