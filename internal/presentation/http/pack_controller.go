package http // Define the package name as "presentation" for HTTP handlers

import (
	"github.com/gofiber/fiber/v2"                            // Import the Fiber framework for handling HTTP requests
	"order-packs-calculator/internal/infrastructure/logging" // Import the logging package for logging
	"order-packs-calculator/internal/service"                // Import the service package for business logic
)

// PackController handles HTTP requests for pack calculations
type PackController struct {
	calculatePacks service.CalculatePacksService // Changed to interface for testability
	logger         *logging.Logger               // Logger instance for logging requests and errors
}

// NewPackController creates a new instance of PackController
func NewPackController(calculatePacks service.CalculatePacksService, logger *logging.Logger) *PackController {
	return &PackController{
		calculatePacks: calculatePacks, // Initialize the service
		logger:         logger,         // Initialize the logger
	}
}

// CalculatePacks handles the POST /api/calculate endpoint to calculate packs
func (c *PackController) CalculatePacks(ctx *fiber.Ctx) error {
	c.logger.Info("Received request to calculate packs") // Log the incoming request

	var request struct { // Define a struct to parse the JSON request body
		OrderAmount int `json:"orderAmount"` // Field to hold the order amount from the request
	}
	if err := ctx.BodyParser(&request); err != nil { // Parse the request body into the struct
		c.logger.Error("Failed to parse request body", err) // Log the error
		// Return a 400 Bad Request response if parsing fails
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Call the service to calculate packs for the given order amount
	result, totalItems, err := c.calculatePacks.Execute(request.OrderAmount)
	if err != nil { // Check if there was an error during calculation
		c.logger.Error("Failed to calculate packs", err) // Log the error
		// Return a 500 Internal Server Error response if calculation fails
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.logger.Info("Successfully calculated packs") // Log the successful calculation
	// Return a 200 OK response with the calculation result and total items
	return ctx.JSON(fiber.Map{
		"packs":      result,     // Include the pack size -> quantity map
		"totalItems": totalItems, // Include the total items fulfilled
	})
}

// UpdatePackSizes handles the POST /api/pack-sizes endpoint to update pack sizes
func (c *PackController) UpdatePackSizes(ctx *fiber.Ctx) error {
	c.logger.Info("Received request to update pack sizes") // Log the incoming request

	var request struct { // Define a struct to parse the JSON request body
		PackSizes []int `json:"packSizes"` // Field to hold the new pack sizes from the request
	}
	if err := ctx.BodyParser(&request); err != nil { // Parse the request body into the struct
		c.logger.Error("Failed to parse request body", err) // Log the error
		// Return a 400 Bad Request response if parsing fails
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Call the service to update the pack sizes in the repository
	if err := c.calculatePacks.UpdatePackSizes(request.PackSizes); err != nil {
		c.logger.Error("Failed to update pack sizes", err) // Log the error
		// Return a 500 Internal Server Error response if updating fails
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.logger.Info("Successfully updated pack sizes") // Log the successful update
	// Return a 200 OK response with a success message
	return ctx.JSON(fiber.Map{"message": "Pack sizes updated successfully"})
}

// GetPackSizes handles the GET /api/pack-sizes endpoint to retrieve current pack sizes
func (c *PackController) GetPackSizes(ctx *fiber.Ctx) error {
	c.logger.Info("Received request to get pack sizes") // Log the incoming request

	// Fetch the current pack sizes from the service layer (which delegates to the repository)
	packSizes, err := c.calculatePacks.GetPackSizes() // Call the service method instead of accessing repo directly
	if err != nil {                                   // Check if there was an error fetching pack sizes
		c.logger.Error("Failed to get pack sizes", err) // Log the error
		// Return a 500 Internal Server Error response if fetching fails
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.logger.Info("Successfully retrieved pack sizes") // Log the successful retrieval
	// Return a 200 OK response with the current pack sizes
	return ctx.JSON(fiber.Map{"packSizes": packSizes})
}
