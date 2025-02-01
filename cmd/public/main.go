package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/okaaryanata/public-api/internal/api/health"
	"github.com/okaaryanata/public-api/internal/api/listing"
	"github.com/okaaryanata/public-api/internal/api/middleware"
	"github.com/okaaryanata/public-api/internal/api/user"
	"github.com/okaaryanata/public-api/internal/app"
	"github.com/okaaryanata/public-api/internal/domain"
	"github.com/okaaryanata/public-api/internal/service"
	listingClient "github.com/okaaryanata/public-api/pkg/listing"
	userClient "github.com/okaaryanata/public-api/pkg/user"
)

func main() {
	startService()
}

func startService() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &app.AppConfig{}
	app.InitService()

	// Packages
	userClient := userClient.NewUserClient(os.Getenv("URL_USER"))
	listingClient := listingClient.NewListingClient(os.Getenv("URL_LISTING"))

	// Services
	userSvc := service.NewUserService(userClient)
	listingSvc := service.NewListingService(userSvc, listingClient)

	// Controllers
	healthController := health.NewHealthController()
	userController := user.NewUserController(userSvc)
	listingController := listing.NewListingController(listingSvc)

	// Create main route
	router := gin.New()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: middleware.GetListSkipLogPath(),
	}))
	router.Use(gin.Recovery())
	router.Use(middleware.SetCORSMiddleware())

	// Register main route
	mainRoute := router.Group(domain.MainRoute)

	// Register routes
	healthController.RegisterRoutes(mainRoute)
	listingController.RegisterRoutes(mainRoute)
	userController.RegisterRoutes(mainRoute)

	host := fmt.Sprintf("%s:%s", app.Host, app.Port)
	router.Run(host)
}
