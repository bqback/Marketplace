package service

import (
	"marketplace/internal/auth"
	authService "marketplace/internal/service/auth"
	"marketplace/internal/storage"
)

type Services struct {
	Auth    IAuthService
	Listing IListingService
}

func NewServices(storages *storage.Storages, manager *auth.AuthManager) *Services {
	return &Services{
		Auth:    authService.NewAuthService(storages.Auth, manager),
		Listing: authService.NewListingService(storages.Listing),
	}
}
