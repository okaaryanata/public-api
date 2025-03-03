package domain

const (
	MainRoute = "/public-api"
)

type (
	PaginationArgs struct {
		Page int
		Size int
	}

	ErrorMessage struct {
		Errors interface{} `json:"errors"`
	}
)
