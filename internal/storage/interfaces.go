package storage

import (
	"context"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/pkg/entities"
)

type IListingStorage interface {
	GetListings(context.Context, dto.FeedOptions) ([]*entities.Listing, error)
	Create(context.Context, dto.NewListingInfo) (*entities.Listing, error)
}

type IAuthStorage interface {
	Auth(context.Context, dto.LoginInfo) (*entities.User, error)
	Register(context.Context, dto.SignupInfo) (*entities.User, error)
}
