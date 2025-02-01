package listing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/okaaryanata/public-api/internal/domain"
)

const (
	listingSvcGetListings   = "/listings"
	listingSvcCreateListing = "/listings"
)

type (
	ListingClient struct {
		client *http.Client
		url    string
	}
)

func NewListingClient(baseURL string) *ListingClient {
	return &ListingClient{
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
		url: baseURL,
	}
}

func (c *ListingClient) GetListings(ctx context.Context, args *domain.GetListingsArgs) ([]domain.ListingResponse, error) {
	queryUserID := ""
	if args.UserID > 0 {
		queryUserID = fmt.Sprintf("&user_id=%d", args.UserID)
	}
	if args.Page <= 0 {
		args.Page = 1
	}
	if args.Size <= 0 {
		args.Size = 10
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?page_num=%d&page_size=%d%s",
		c.url,
		listingSvcGetListings,
		args.Page,
		args.Size,
		queryUserID,
	), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occur when get listing with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return nil, fmt.Errorf("error occur when read resp body (detail: %s)", err.Error())
	}

	type GetListingResp struct {
		Listings []domain.ListingResponse `json:"listings"`
	}

	var resBody GetListingResp
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return nil, errors.New("failed unmarshal response")
	}

	return resBody.Listings, nil
}

func (c *ListingClient) CreateListing(ctx context.Context, args *domain.CreateListingArgs) (*domain.ListingResponse, error) {
	formData := url.Values{}
	formData.Set("user_id", strconv.Itoa(args.UserID))
	formData.Set("listing_type", args.Type)
	formData.Set("price", strconv.Itoa(int(args.Price)))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.url, listingSvcGetListings), strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error occur while create listing with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return nil, fmt.Errorf("error occur when read resp body (detail: %s)", err.Error())
	}

	type CreateListingResp struct {
		Listing domain.ListingResponse `json:"listing"`
	}

	var resBody CreateListingResp
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return nil, errors.New("failed unmarshal response")
	}

	return &resBody.Listing, nil
}
