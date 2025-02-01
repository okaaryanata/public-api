package service

import (
	"context"

	"github.com/okaaryanata/public-api/internal/domain"
	"github.com/okaaryanata/public-api/pkg/listing"
)

type (
	ListingService struct {
		listingClient *listing.ListingClient
	}
)

func NewListingService(listingClient *listing.ListingClient) *ListingService {
	return &ListingService{
		listingClient: listingClient,
	}
}

func (l *ListingService) GetListings(ctx context.Context, args *domain.GetListingsArgs) ([]domain.ListingResponse, error) {
	return l.listingClient.GetListings(ctx, args)
}

func (l *ListingService) CreateListing(ctx context.Context, args *domain.CreateListingArgs) (*domain.ListingResponse, error) {
	return l.listingClient.CreateListing(ctx, args)
}
