package user

import (
	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/public-api/internal/service"
)

type (
	Controller struct {
		userSvc *service.UserService
	}
)

func NewUserController(userSvc *service.UserService) *Controller {
	return &Controller{
		userSvc: userSvc,
	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	UserRouter := router.Group("/users")
	{
		UserRouter.POST(``, c.CreateUser)
	}
}
