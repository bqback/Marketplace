package service

import (
	"marketplace/internal/auth"
	authService "marketplace/internal/service/auth"
	listingService "marketplace/internal/service/listing"
	"marketplace/internal/storage"
)

type Services struct {
	Auth    IAuthService
	Listing IListingService
}

func NewServices(storages *storage.Storages, manager *auth.AuthManager) *Services {
	return &Services{
		Auth:    authService.NewAuthService(storages.Auth, manager),
		Listing: listingService.NewListingService(storages.Listing),
	}
}
