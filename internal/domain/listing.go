package domain

type (
	GetListingResponse struct {
		ID        uint          `json:"id"`
		Type      string        `json:"listing_type"`
		Price     int           `json:"price"`
		CreatedAt int64         `json:"created_at"`
		UpdatedAt int64         `json:"updated_at"`
		User      *UserResponse `json:"user"`
	}

	ClientListingResponse struct {
		ID        uint   `json:"id"`
		UserID    uint   `json:"user_id"`
		Type      string `json:"listing_type"`
		Price     int    `json:"price"`
		CreatedAt int64  `json:"created_at"`
		UpdatedAt int64  `json:"updated_at"`
	}

	CreateListingArgs struct {
		UserID int    `json:"user_id" binding:"required"`
		Type   string `json:"listing_type" binding:"required"`
		Price  int    `json:"price" binding:"required"`
	}

	GetListingsArgs struct {
		UserID int
		PaginationArgs
	}
)

func (cr *ClientListingResponse) MapperToGetListingResp() *GetListingResponse {
	return &GetListingResponse{
		ID:        cr.ID,
		Type:      cr.Type,
		Price:     cr.Price,
		CreatedAt: cr.CreatedAt,
		UpdatedAt: cr.UpdatedAt,
		User: &UserResponse{
			ID: cr.UserID,
		},
	}
}
