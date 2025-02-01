package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/public-api/internal/domain"
)

func (c *Controller) CreateUser(ctx *gin.Context) {
	var (
		args domain.CreateUserArgs
		err  error
	)

	defer func() {
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}()

	if err = ctx.ShouldBind(&args); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "errors": "User name is required"})
		return
	}

	user, err := c.userSvc.CreateUser(ctx, &args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "user": user})
}
