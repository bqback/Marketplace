package service

import (
	"context"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/pkg/entities"
)

type IListingService interface {
	Create(context.Context, dto.NewListingInfo) (*entities.Listing, error)
	GetListings(context.Context, dto.FeedOptions) ([]*entities.Listing, error)
}

type IAuthService interface {
	Authorize(context.Context, dto.LoginInfo) (*dto.JWT, error)
	Register(context.Context, dto.SignupInfo) (*dto.NewUserInfo, error)
}
