package service

import (
	"github.com/stretchr/testify/assert"
	"order-packs-calculator/internal/infrastructure/repository/mocks" // Import the mocks package
	"testing"

	"github.com/golang/mock/gomock"     // Import gomock for mocking
	"github.com/stretchr/testify/suite" // Import testify/suite for test suites
)

// CalculatePacksUseCaseTestSuite defines the test suite for the service layer
type CalculatePacksUseCaseTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockPackRepository // Use gomock-generated mock type
	uc       *CalculatePacksUseCase    // Use case under test
	ctrl     *gomock.Controller        // Gomock controller for managing mocks
}

// SetupTest sets up the test environment before each test
func (s *CalculatePacksUseCaseTestSuite) SetupTest() {
	// Create a gomock controller
	s.ctrl = gomock.NewController(s.T())

	// Create a mock repository using gomock
	s.mockRepo = mocks.NewMockPackRepository(s.ctrl)

	// Create a new use case instance
	s.uc = NewCalculatePacksUseCase(s.mockRepo)
}

// TearDownTest cleans up the test environment after each test
func (s *CalculatePacksUseCaseTestSuite) TearDownTest() {
	// Finish the gomock controller
	s.ctrl.Finish()
}

// TestCalculatePacksUseCaseTestSuite runs the test suite
func TestCalculatePacksUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePacksUseCaseTestSuite)) // Run the suite
}

// TestExecute tests the Execute method of CalculatePacksUseCase
func (s *CalculatePacksUseCaseTestSuite) TestExecute() {
	s.Run("Success", func() {
		// Set up the mock expectation using gomock API
		s.mockRepo.EXPECT().GetPackSizes().Return([]int{250, 500, 1000, 2000, 5000}, nil)

		// Call the Execute method
		result, total, err := s.uc.Execute(263)
		s.Assert().NoError(err, "Expected no error")
		s.Assert().Equal(map[int]int{500: 1}, result, "Result should match expected")
		s.Assert().Equal(500, total, "Total items should match expected")
	})

	s.Run("RepositoryError", func() {
		// Set up the mock expectation using gomock API
		s.mockRepo.EXPECT().GetPackSizes().Return([]int{}, assert.AnError)

		// Call the Execute method
		result, total, err := s.uc.Execute(263)
		s.Assert().Error(err, "Expected an error")
		s.Assert().Nil(result, "Result should be nil on error")
		s.Assert().Equal(0, total, "Total should be 0 on error")
	})
}
