package domain

type (
	UserResponse struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		CreatedAt int64  `json:"created_at"`
		UpdatedAt int64  `json:"updated_at"`
	}

	CreateUserArgs struct {
		Name string `json:"name" binding:"required"`
	}
)
