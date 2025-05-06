package main // Define the package name as "main" for the application entry point

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log" // Import the log package for logging errors
	"order-packs-calculator/internal/presentation/http"

	"github.com/gofiber/fiber/v2"                               // Import the Fiber framework for the web server
	"order-packs-calculator/internal/infrastructure/config"     // Import the config package for loading configuration
	"order-packs-calculator/internal/infrastructure/logging"    // Import the logging package for logging
	"order-packs-calculator/internal/infrastructure/repository" // Import the repository package for data access
	"order-packs-calculator/internal/service"                   // Import the service package for business logic
)

// main is the entry point of the application
func main() {
	// Load the application configuration
	cfg, err := config.LoadConfig() // Call LoadConfig to get the configuration
	if err != nil {                 // Check if there was an error loading the configuration
		log.Fatalf("Failed to load configuration: %v", err) // Log the error and exit
	}

	// Initialize the logger
	logger := logging.NewLogger() // Create a new logger instance

	// Initialize the in-memory repository with the default pack sizes from the config
	repo := repository.NewInMemoryPackRepository(cfg.PackSizes)

	// Initialize the service with the repository
	calculatePacksService := service.NewCalculatePacksUseCase(repo)

	// Initialize the controller with the service and logger
	packController := http.NewPackController(calculatePacksService, logger)

	// Create a new Fiber application instance
	app := fiber.New()

	// Add CORS middleware to allow cross-origin requests
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:63342", // Allow requests from both origins
		AllowMethods:     "GET,POST,OPTIONS",                             // Include OPTIONS for preflight requests
		AllowHeaders:     "Content-Type",                                 // Allow Content-Type header
		AllowCredentials: false,                                          // Set to true if credentials (e.g., cookies) are needed
		MaxAge:           86400,                                          // Cache preflight response for 24 hours
	}))

	// Serve static files from the ./web directory (for the UI)
	app.Static("/", "./web")

	// Create a group for API routes under the /api prefix
	api := app.Group("/api")
	// Define the POST /api/calculate endpoint for calculating packs
	api.Post("/calculate", packController.CalculatePacks)
	// Define the POST /api/pack-sizes endpoint for updating pack sizes
	api.Post("/pack-sizes", packController.UpdatePackSizes)
	// Define the GET /api/pack-sizes endpoint for retrieving pack sizes
	api.Get("/pack-sizes", packController.GetPackSizes)

	// Start the Fiber server on the configured port
	if err := app.Listen(cfg.Port); err != nil { // Start the server and handle any errors
		log.Fatalf("Failed to start server: %v", err) // Log the error and exit
	}
}
