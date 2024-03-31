package service

import (
	"context"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/pkg/entities"
)

type IListingService interface {
	Create(context.Context, dto.NewListingInfo) (*entities.Listing, error)
	GetFeed(context.Context, dto.FeedOptions) ([]*dto.FeedListingInfo, error)
}

type IAuthService interface {
	Authorize(context.Context, dto.LoginInfo) (*dto.JWT, error)
	Register(context.Context, dto.SignupInfo) (*dto.NewUserInfo, error)
}
