package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/okaaryanata/public-api/helper"
	"github.com/okaaryanata/public-api/internal/domain"
)

const (
	pathUsers = "/users"
)

type (
	UserClient struct {
		client *http.Client
		url    string
	}
)

func NewUserClient(baseURL string) *UserClient {
	return &UserClient{
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
		url: baseURL,
	}
}

func (c *UserClient) CreateUser(ctx context.Context, args *domain.CreateUserArgs) (*domain.UserResponse, error) {
	formData := url.Values{}
	formData.Set("name", args.Name)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.url, pathUsers), strings.NewReader(formData.Encode()))
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
		err = helper.ConverErrors(respErr.Errors, fmt.Sprintf("error occur when create user with status code: %d", resp.StatusCode))

		return nil, err
	}

	type CreateUserResp struct {
		User domain.UserResponse `json:"user"`
	}

	var resBody CreateUserResp
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return nil, errors.New("failed unmarshal response")
	}

	return &resBody.User, nil
}

func (c *UserClient) GetUserByID(ctx context.Context, userID uint) (*domain.UserResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s/%d",
		c.url,
		pathUsers,
		userID,
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
		err = helper.ConverErrors(respErr.Errors, fmt.Sprintf("error occur when get user with status code: %d", resp.StatusCode))

		return nil, err
	}

	type GetUserByIDResp struct {
		User *domain.UserResponse `json:"user"`
	}

	var resBody GetUserByIDResp
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return nil, errors.New("failed unmarshal response")
	}

	return resBody.User, nil
}
