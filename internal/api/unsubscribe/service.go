package unsubscribe

import (
	"http-server/internal/service"
	desc "http-server/pkg/unsubscribe"
)

type Implementation struct {
	desc.UnimplementedUnsubscribeServer
	unsubscribeService service.AccountService
}

func NewImplementation(unsubscribeService service.AccountService) *Implementation {
	return &Implementation{
		unsubscribeService: unsubscribeService,
	}
}
