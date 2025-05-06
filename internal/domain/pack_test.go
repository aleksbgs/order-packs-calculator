package domain

import (
	"reflect" // Import the reflect package to compare maps
	"testing" // Import the testing package for writing unit tests

	"github.com/stretchr/testify/suite" // Import testify/suite for test suites
)

// PackTestSuite defines the test suite for the domain layer
type PackTestSuite struct {
	suite.Suite // Embed the testify suite
}

// TestPackTestSuite runs the test suite
func TestPackTestSuite(t *testing.T) {
	suite.Run(t, new(PackTestSuite)) // Run the suite
}

// TestCalculatePacks tests the CalculatePacks function with various test cases
func (s *PackTestSuite) TestCalculatePacks() {
	tests := []struct { // Define a slice of test cases
		name          string      // Name of the test case for better reporting
		packSizes     []int       // Input pack sizes for the test (reverted from []Pack to []int)
		orderAmount   int         // Input order amount for the test
		expected      map[int]int // Expected result map (pack size -> quantity)
		expectedTotal int         // Expected total items fulfilled
		expectErr     error       // Expected error (if any)
	}{
		{ // Test case 1: Order 263 with default pack sizes
			name:          "Order 263 with default pack sizes", // Name of the test case
			packSizes:     []int{250, 500, 1000, 2000, 5000},   // Pack sizes as []int
			orderAmount:   263,                                 // Order amount to fulfill
			expected:      map[int]int{500: 1},                 // Expected result: 1 pack of 500
			expectedTotal: 500,                                 // Expected total items: 500
			expectErr:     nil,                                 // Expected error: none
		},
		{ // Test case 2: Edge case from the hints
			name:          "Edge case from hints",              // Name of the test case
			packSizes:     []int{23, 31, 53},                   // Pack sizes as []int
			orderAmount:   500000,                              // Order amount to fulfill
			expected:      map[int]int{53: 9429, 31: 7, 23: 2}, // Expected result from the problem
			expectedTotal: 500000,                              // Expected total items: 500,000
			expectErr:     nil,                                 // Expected error: none
		},
		{ // Test case 3: Negative order amount
			name:          "Negative order amount", // Name of the test case
			packSizes:     []int{250, 500},         // Pack sizes as []int
			orderAmount:   -1,                      // Order amount (invalid: negative)
			expected:      nil,                     // Expected result: nil (due to error)
			expectedTotal: 0,                       // Expected total items: 0 (due to error)
			expectErr:     ErrInvalidOrderAmount,   // Expected error: invalid order amount
		},
	}

	for _, tt := range tests { // Loop through each test case
		s.Run(tt.name, func() { // Use s.Run to run each test case as a subtest
			result, total, err := CalculatePacks(tt.packSizes, tt.orderAmount)                                                            // Call the function to test
			s.Assert().Equal(tt.expectErr, err, "Error should match expected")                                                            // Check if the error matches
			s.Assert().True(reflect.DeepEqual(result, tt.expected), "Result should match expected: got %v, want %v", result, tt.expected) // Check the result map
			s.Assert().Equal(tt.expectedTotal, total, "Total items should match expected")                                                // Check the total items
		})
	}
}
