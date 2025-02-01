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

	"github.com/okaaryanata/public-api/helper"
	"github.com/okaaryanata/public-api/internal/domain"
)

const (
	pathListing = "/listings"
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

func (c *ListingClient) GetListings(ctx context.Context, args *domain.GetListingsArgs) ([]domain.ClientListingResponse, error) {
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
		pathListing,
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error occur when read resp body (detail: %s)", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		var respErr domain.ErrorMessage
		json.Unmarshal(body, &respErr)
		err = helper.ConverErrors(respErr.Errors, fmt.Sprintf("error occur when get listing with status code: %d", resp.StatusCode))

		return nil, err
	}

	type GetListingResp struct {
		Listings []domain.ClientListingResponse `json:"listings"`
	}

	var resBody GetListingResp
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return nil, errors.New("failed unmarshal response")
	}

	return resBody.Listings, nil
}

func (c *ListingClient) CreateListing(ctx context.Context, args *domain.CreateListingArgs) (*domain.ClientListingResponse, error) {
	formData := url.Values{}
	formData.Set("user_id", strconv.Itoa(args.UserID))
	formData.Set("listing_type", args.Type)
	formData.Set("price", strconv.Itoa(int(args.Price)))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.url, pathListing), strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error occur when read resp body (detail: %s)", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		var respErr domain.ErrorMessage
		json.Unmarshal(body, &respErr)
		err = helper.ConverErrors(respErr.Errors, fmt.Sprintf("error occur when create listing with status code: %d", resp.StatusCode))

		return nil, err
	}

	type CreateListingResp struct {
		Listing *domain.ClientListingResponse `json:"listing"`
	}

	var resBody CreateListingResp
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return nil, errors.New("failed unmarshal response")
	}

	return resBody.Listing, nil
}
