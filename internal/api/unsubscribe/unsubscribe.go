package unsubscribe

import (
	"context"
	"errors"

	"http-server/internal/entities"
	desc "http-server/pkg/unsubscribe"
)

func (i *Implementation) Unsubscribe(ctx context.Context, req *desc.UnsubscribeRequest) (*desc.UnsubscribeResponse, error) {
	accountID := req.GetId()

	err := i.unsubscribeService.DeleteAccount(accountID)
	if err != nil {
		return nil, errors.New(entities.ErrFailedDelete)
	}

	return &desc.UnsubscribeResponse{
		Success: true,
		Message: "successful",
	}, nil
}
