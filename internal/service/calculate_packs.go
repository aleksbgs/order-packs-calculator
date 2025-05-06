package service // Define the package name as "service" for the service layer (application logic)

import (
	"order-packs-calculator/internal/domain"                    // Changed from internal/entity to internal/domain
	"order-packs-calculator/internal/infrastructure/repository" // Changed from internal/repository to internal/infrastructure/repository
)

// CalculatePacksService defines the interface for the CalculatePacksUseCase
type CalculatePacksService interface {
	Execute(orderAmount int) (map[int]int, int, error)
	UpdatePackSizes(newSizes []int) error
	GetPackSizes() ([]int, error)
}

// CalculatePacksUseCase defines the service for calculating packs
type CalculatePacksUseCase struct {
	repo repository.PackRepository // Repository interface to fetch pack sizes
}

// Ensure CalculatePacksUseCase implements CalculatePacksService
var _ CalculatePacksService = (*CalculatePacksUseCase)(nil)

// NewCalculatePacksUseCase creates a new instance of CalculatePacksUseCase
func NewCalculatePacksUseCase(repo repository.PackRepository) *CalculatePacksUseCase {
	return &CalculatePacksUseCase{repo: repo} // Initialize the service with the provided repository
}

// Execute runs the service to calculate packs for an order
func (uc *CalculatePacksUseCase) Execute(orderAmount int) (map[int]int, int, error) {
	// Fetch pack sizes from the repository (could be a database in a real app)
	packSizes, err := uc.repo.GetPackSizes() // Call the repository to get the current pack sizes
	if err != nil {                          // Check if there was an error fetching pack sizes
		return nil, 0, err // Return the error if fetching failed
	}

	// Call the domain function to calculate packs using the fetched pack sizes
	return domain.CalculatePacks(packSizes, orderAmount) // Pass the []int directly to the domain layer
}

// UpdatePackSizes updates the pack sizes in the repository
func (uc *CalculatePacksUseCase) UpdatePackSizes(newSizes []int) error {
	return uc.repo.UpdatePackSizes(newSizes) // Call the repository to update pack sizes
}

// GetPackSizes retrieves the current pack sizes from the repository
func (uc *CalculatePacksUseCase) GetPackSizes() ([]int, error) {
	return uc.repo.GetPackSizes() // Delegate to the repository to fetch pack sizes
}
