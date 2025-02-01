package domain

type (
	ListingResponse struct {
		ID        uint          `json:"id"`
		Type      string        `json:"listing_type"`
		Price     int           `json:"price"`
		CreatedAt int64         `json:"created_at"`
		UpdatedAt int64         `json:"updated_at"`
		User      *UserResponse `json:"user"`
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
