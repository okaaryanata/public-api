package listing

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/okaaryanata/public-api/internal/domain"
)

func (c *Controller) CreateListing(ctx *gin.Context) {
	var (
		args domain.CreateListingArgs
		err  error
	)

	defer func() {
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}()

	if err = ctx.ShouldBind(&args); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "errors": err.Error()})
		return
	}

	listing, err := c.listingSvc.CreateListing(ctx, &args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "listing": listing})
}

func (c *Controller) GetListings(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}()

	pageNum, err := strconv.Atoi(ctx.DefaultQuery("page_num", "1"))
	if err != nil || pageNum < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "error": "Invalid page_num"})
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "error": "Invalid page_size"})
		return
	}

	userID, err := strconv.Atoi(ctx.DefaultQuery("user_id", "0"))
	if err != nil || pageSize < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "error": "Invalid page_size"})
		return
	}

	users, err := c.listingSvc.GetListings(ctx, &domain.GetListingsArgs{
		UserID: userID,
		PaginationArgs: domain.PaginationArgs{
			Page: pageNum,
			Size: pageSize,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"result": false, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "users": users})
}
