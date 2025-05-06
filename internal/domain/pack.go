package domain // Define the package name as "domain" for core business logic

import (
	"errors"
	"sort"
) // Import the sort package for sorting pack sizes

// Pack represents a pack size
type Pack struct {
	Size int // Size field to store the size of the pack (e.g., 250, 500)
}

// CalculatePacks calculates the minimum packs needed to fulfill an order
func CalculatePacks(packSizes []int, orderAmount int) (map[int]int, int, error) {
	if orderAmount < 0 { // Check if the order amount is negative
		return nil, 0, ErrInvalidOrderAmount // Return an error if the order amount is invalid
	}
	if len(packSizes) == 0 { // Check if the pack sizes slice is empty
		return nil, 0, ErrNoPackSizes // Return an error if no pack sizes are provided
	}

	// Sort pack sizes in descending order to start with the largest packs (greedy approach)
	sortedSizes := make([]int, len(packSizes))          // Create a new slice to hold sorted pack sizes
	copy(sortedSizes, packSizes)                        // Copy the input pack sizes to avoid modifying the original slice
	sort.Sort(sort.Reverse(sort.IntSlice(sortedSizes))) // Sort in descending order

	result := make(map[int]int) // Initialize a map to store pack size -> quantity needed
	remaining := orderAmount    // Track the remaining items to fulfill
	totalItems := 0             // Track the total items fulfilled by the selected packs

	// Greedy approach: start with largest packs to minimize overage
	for _, size := range sortedSizes { // Loop through each pack size in descending order
		if size <= 0 { // Skip invalid pack sizes (less than or equal to 0)
			continue // Move to the next pack size
		}
		// Calculate how many of this pack size can be used to fulfill the remaining amount
		count := remaining / size // Integer division to get the number of packs
		if count > 0 {            // If we can use at least one pack of this size
			result[size] = count  // Add the count to the result map
			items := count * size // Calculate the total items from these packs
			totalItems += items   // Add to the total items fulfilled
			remaining -= items    // Subtract from the remaining items needed
		}
	}

	// If we couldn't fulfill the order exactly, we need to overshoot minimally
	if remaining > 0 { // Check if there are still items remaining to fulfill
		for _, size := range sortedSizes { // Loop through pack sizes again
			if size >= remaining { // Find the smallest pack size that can cover the remaining items
				result[size]++     // Increment the count for this pack size
				totalItems += size // Add the pack size to the total items
				break              // Exit the loop since we've fulfilled the order
			}
		}
	}

	// If we still couldn't fulfill the order, the pack sizes are insufficient
	if totalItems < orderAmount { // Check if the total items are less than the order amount
		return nil, 0, ErrInsufficientPackSizes // Return an error if we can't fulfill the order
	}

	// Optimize to minimize the number of packs (call the optimization function)
	return optimizePacks(sortedSizes, result, orderAmount, totalItems) // Return the optimized result
}

// optimizePacks tries to reduce the number of packs while keeping total items minimal
func optimizePacks(sortedSizes []int, result map[int]int, orderAmount, totalItems int) (map[int]int, int, error) {
	bestResult := result                // Store the current result as the best result initially
	bestTotalItems := totalItems        // Store the current total items as the best total
	bestPackCount := countPacks(result) // Calculate the total number of packs in the current result

	// Try reducing larger packs and compensating with smaller ones to minimize pack count
	for i := 0; i < len(sortedSizes); i++ { // Loop through each pack size index
		size := sortedSizes[i] // Get the current pack size
		if result[size] == 0 { // Skip if we aren't using any packs of this size
			continue // Move to the next pack size
		}
		// Reduce one pack of this size to see if we can optimize
		tempResult := copyMap(result)                     // Create a copy of the result map to modify
		tempResult[size]--                                // Reduce the count of the current pack size by 1
		tempTotal := bestTotalItems - size                // Update the total items after removing one pack
		tempRemaining := orderAmount - (tempTotal - size) // Calculate remaining items to fulfill

		// Try to make up the difference with smaller packs
		for j := i + 1; j < len(sortedSizes); j++ { // Loop through smaller pack sizes
			smallSize := sortedSizes[j] // Get the smaller pack size
			if tempRemaining <= 0 {     // If we've fulfilled the remaining items
				break // Exit the inner loop
			}
			count := (tempRemaining + smallSize - 1) / smallSize // Ceiling division to get packs needed
			tempResult[smallSize] += count                       // Add the count of smaller packs to the result
			tempTotal += count * smallSize                       // Update the total items
			tempRemaining -= count * smallSize                   // Update the remaining items
		}

		if tempTotal >= orderAmount { // If the new combination fulfills the order
			tempPackCount := countPacks(tempResult) // Calculate the total number of packs
			// Check if this combination uses fewer packs or same packs with fewer total items
			if tempPackCount < bestPackCount || (tempPackCount == bestPackCount && tempTotal < bestTotalItems) {
				bestResult = tempResult       // Update the best result
				bestTotalItems = tempTotal    // Update the best total items
				bestPackCount = tempPackCount // Update the best pack count
			}
		}
	}

	return bestResult, bestTotalItems, nil // Return the optimized result, total items, and no error
}

// countPacks calculates the total number of packs in the result map
func countPacks(result map[int]int) int {
	total := 0                     // Initialize the total pack count to 0
	for _, count := range result { // Loop through each pack size's count in the result
		total += count // Add the count to the total
	}
	return total // Return the total number of packs
}

// copyMap creates a deep copy of the input map
func copyMap(m map[int]int) map[int]int {
	newMap := make(map[int]int) // Create a new empty map
	for k, v := range m {       // Loop through each key-value pair in the input map
		newMap[k] = v // Copy the key-value pair to the new map
	}
	return newMap // Return the copied map
}

// Define custom error types for different failure scenarios
var (
	ErrInvalidOrderAmount    = errors.New("order amount cannot be negative")          // Error for negative order amounts
	ErrNoPackSizes           = errors.New("no pack sizes provided")                   // Error for empty pack sizes
	ErrInsufficientPackSizes = errors.New("pack sizes insufficient to fulfill order") // Error for insufficient pack sizes
)
