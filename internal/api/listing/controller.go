package listing

import (
	"github.com/gin-gonic/gin"

	"github.com/okaaryanata/public-api/internal/service"
)

type (
	Controller struct {
		listingSvc *service.ListingService
	}
)

func NewListingController(listingSvc *service.ListingService) *Controller {
	return &Controller{
		listingSvc: listingSvc,
	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	UserRouter := router.Group("/listings")
	{
		UserRouter.POST(``, c.CreateListing)
		UserRouter.GET(``, c.GetListings)
	}
}
