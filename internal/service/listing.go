package service

import (
	"context"
	"fmt"

	"github.com/okaaryanata/public-api/internal/domain"
	"github.com/okaaryanata/public-api/pkg/listing"
)

type (
	ListingService struct {
		userSvc *UserService

		listingClient *listing.ListingClient
	}
)

func NewListingService(
	userSvc *UserService,
	listingClient *listing.ListingClient,
) *ListingService {
	return &ListingService{
		userSvc:       userSvc,
		listingClient: listingClient,
	}
}

func (l *ListingService) GetListings(ctx context.Context, args *domain.GetListingsArgs) ([]domain.GetListingResponse, error) {
	listings, err := l.listingClient.GetListings(ctx, args)
	if err != nil {
		return nil, err
	}

	var resp []domain.GetListingResponse
	for idx := range listings {
		listing := listings[idx].MapperToGetListingResp()
		user, err := l.userSvc.GetUserByID(ctx, listings[idx].UserID)
		if err != nil {
			return nil, err
		}

		if user != nil {
			listing.User = user
		}

		resp = append(resp, *listing)
	}

	return resp, nil
}

func (l *ListingService) CreateListing(ctx context.Context, args *domain.CreateListingArgs) (*domain.ClientListingResponse, error) {
	user, err := l.userSvc.GetUserByID(ctx, uint(args.UserID))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user with id: %d didn't exist", args.UserID)
	}

	return l.listingClient.CreateListing(ctx, args)
}
