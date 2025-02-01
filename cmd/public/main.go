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
	"github.com/okaaryanata/public-api/internal/app"
	"github.com/okaaryanata/public-api/internal/domain"
	"github.com/okaaryanata/public-api/internal/service"
	listingClient "github.com/okaaryanata/public-api/pkg/listing"
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
	listingClient := listingClient.NewListingClient(os.Getenv("URL_LISTING"))

	// Services
	listingSvc := service.NewListingService(listingClient)

	// Controllers
	healthController := health.NewHealthController()
	listingController := listing.NewUserController(listingSvc)

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

	host := fmt.Sprintf("%s:%s", app.Host, app.Port)
	router.Run(host)
}
