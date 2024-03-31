package storage

import (
	"context"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/pkg/entities"
)

type IListingStorage interface {
	GetFeed(context.Context, dto.FeedOptions) ([]*dto.FeedListingInfo, error)
	Create(context.Context, dto.NewListingInfo) (*entities.Listing, error)
}

type IAuthStorage interface {
	Auth(context.Context, dto.LoginInfo) (*entities.User, error)
	Register(context.Context, dto.SignupInfo) (*entities.User, error)
}
