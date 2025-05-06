package repository

import (
	"testing" // Import the testing package for writing unit tests

	"github.com/stretchr/testify/suite" // Import testify/suite for test suites
)

// PackRepositoryTestSuite defines the test suite for the repository package
type PackRepositoryTestSuite struct {
	suite.Suite                         // Embed the testify suite
	repo        *InMemoryPackRepository // Repository under test
}

// SetupTest sets up the test environment before each test
func (s *PackRepositoryTestSuite) SetupTest() {
	// Initialize the repository with default pack sizes
	defaultSizes := []int{250, 500, 1000}
	s.repo = NewInMemoryPackRepository(defaultSizes)
}

// TestPackRepositoryTestSuite runs the test suite
func TestPackRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PackRepositoryTestSuite))
}

// TestGetPackSizes tests retrieving pack sizes
func (s *PackRepositoryTestSuite) TestGetPackSizes() {
	// Get the pack sizes
	sizes, err := s.repo.GetPackSizes()
	s.Assert().NoError(err, "Expected no error")
	s.Assert().Equal([]int{250, 500, 1000}, sizes, "Pack sizes should match initial value")
}

// TestUpdatePackSizes tests updating pack sizes
func (s *PackRepositoryTestSuite) TestUpdatePackSizes() {
	// Update the pack sizes
	newSizes := []int{100, 200, 300}
	err := s.repo.UpdatePackSizes(newSizes)
	s.Assert().NoError(err, "Expected no error")

	// Verify the updated pack sizes
	sizes, err := s.repo.GetPackSizes()
	s.Assert().NoError(err, "Expected no error")
	s.Assert().Equal(newSizes, sizes, "Pack sizes should match updated value")
}
