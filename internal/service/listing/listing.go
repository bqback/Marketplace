package listing

import (
	"context"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/pkg/entities"
	"marketplace/internal/storage"
)

type ListingService struct {
	ls storage.IListingStorage
}

func NewListingService(listingStorage storage.IListingStorage) *ListingService {
	return &ListingService{
		ls: listingStorage,
	}
}

func (s *ListingService) Create(ctx context.Context, info dto.NewListingInfo) (*entities.Listing, error) {
	return s.ls.Create(ctx, info)
}

func (s *ListingService) GetListings(ctx context.Context, opts dto.FeedOptions) ([]*entities.Listing, error) {
	return s.ls.GetListings(ctx, opts)
}
