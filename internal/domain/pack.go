package domain

import (
	"errors"
	"sort"
)

// Pack represents a pack size
type Pack struct {
	Size int
}

// CalculatePacks calculates the minimum packs needed to fulfill an order
func CalculatePacks(packSizes []int, orderAmount int) (map[int]int, int, error) {
	if orderAmount < 0 {
		return nil, 0, ErrInvalidOrderAmount
	}
	if len(packSizes) == 0 {
		return nil, 0, ErrNoPackSizes
	}

	// Sort pack sizes in descending order
	sortedSizes := make([]int, len(packSizes))
	copy(sortedSizes, packSizes)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedSizes)))

	// Remove invalid pack sizes (<= 0)
	validSizes := []int{}
	for _, size := range sortedSizes {
		if size > 0 {
			validSizes = append(validSizes, size)
		}
	}
	if len(validSizes) == 0 {
		return nil, 0, ErrNoPackSizes
	}
	sortedSizes = validSizes

	// Dynamic programming to find the combination that exactly matches or minimally overshoots
	// dp[i] = minimum number of packs to reach amount i
	maxAmount := orderAmount + sortedSizes[0] // Allow slight overshoot
	dp := make([]int, maxAmount+1)
	prev := make([]int, maxAmount+1) // Track the last pack size used
	for i := range dp {
		dp[i] = -1 // -1 means unreachable
	}
	dp[0] = 0 // Base case: 0 packs needed for amount 0

	// First pass: Fill dp to find minimum packs
	for i := 1; i <= maxAmount; i++ {
		for _, size := range sortedSizes {
			if i < size {
				continue
			}
			if dp[i-size] != -1 {
				packs := dp[i-size] + 1
				if dp[i] == -1 || packs < dp[i] {
					dp[i] = packs
					prev[i] = size
				}
			}
		}
	}

	// Find the smallest amount >= orderAmount that can be fulfilled
	bestTotalItems := -1
	bestPackCount := -1
	bestResult := make(map[int]int)

	for amount := orderAmount; amount <= maxAmount; amount++ {
		if dp[amount] == -1 {
			continue
		}
		// Reconstruct the solution for this amount
		result := make(map[int]int)
		tempAmount := amount
		for tempAmount > 0 {
			size := prev[tempAmount]
			result[size]++
			tempAmount -= size
		}
		packCount := countPacks(result)
		// Prioritize exact match, then minimize pack count, then minimize overage
		if bestTotalItems == -1 ||
			(amount == orderAmount && bestTotalItems != orderAmount) || // Exact match takes priority
			(amount == bestTotalItems && packCount < bestPackCount) ||
			(amount < bestTotalItems) {
			bestTotalItems = amount
			bestPackCount = packCount
			bestResult = result
		}
	}

	if bestTotalItems == -1 {
		return nil, 0, ErrInsufficientPackSizes
	}

	return bestResult, bestTotalItems, nil
}

func countPacks(result map[int]int) int {
	total := 0
	for _, count := range result {
		total += count
	}
	return total
}

var (
	ErrInvalidOrderAmount    = errors.New("order amount cannot be negative")
	ErrNoPackSizes           = errors.New("no pack sizes provided")
	ErrInsufficientPackSizes = errors.New("pack sizes insufficient to fulfill order")
)
