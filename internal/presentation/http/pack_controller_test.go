package http

import (
	"bytes"             // Import bytes for creating request bodies
	"encoding/json"     // Import json for encoding/decoding
	"net/http/httptest" // Import httptest for HTTP testing
	"testing"           // Import the testing package for writing unit tests

	"github.com/gofiber/fiber/v2"                            // Import Fiber for creating a test app
	"github.com/golang/mock/gomock"                          // Import gomock for mocking
	"github.com/stretchr/testify/suite"                      // Import testify/suite for test suites
	"order-packs-calculator/internal/infrastructure/logging" // Import logging package
	"order-packs-calculator/internal/service/mocks"          // Import mocks for the service
)

// PackControllerTestSuite defines the test suite for the presentation layer
type PackControllerTestSuite struct {
	suite.Suite                                  // Embed the testify suite
	app         *fiber.App                       // Fiber app for testing
	controller  *PackController                  // Controller under test
	mockService *mocks.MockCalculatePacksService // Use gomock-generated mock type
	logger      *logging.Logger                  // Logger instance
	ctrl        *gomock.Controller               // Gomock controller for managing mocks
}

// SetupTest sets up the test environment before each test
func (s *PackControllerTestSuite) SetupTest() {
	// Create a gomock controller
	s.ctrl = gomock.NewController(s.T())

	// Create a mock logger
	s.logger = logging.NewLogger()

	// Create a mock service using gomock
	s.mockService = mocks.NewMockCalculatePacksService(s.ctrl)

	// Create a new PackController
	s.controller = NewPackController(s.mockService, s.logger)

	// Create a new Fiber app
	s.app = fiber.New()

	// Create a group for API routes under the /api prefix
	api := s.app.Group("/api")
	api.Post("/calculate", s.controller.CalculatePacks)
	api.Post("/pack-sizes", s.controller.UpdatePackSizes)
	api.Get("/pack-sizes", s.controller.GetPackSizes)
}

// TearDownTest cleans up the test environment after each test
func (s *PackControllerTestSuite) TearDownTest() {
	// Finish the gomock controller
	s.ctrl.Finish()
}

// TestPackControllerTestSuite runs the test suite
func TestPackControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PackControllerTestSuite))
}

// TestCalculatePacks_Success tests a successful CalculatePacks request
func (s *PackControllerTestSuite) TestCalculatePacks_Success() {
	// Set up the mock expectation using gomock API
	s.mockService.EXPECT().Execute(263).Return(map[int]int{500: 1}, 500, nil)

	// Create a request body
	reqBody := map[string]int{"orderAmount": 263}
	body, _ := json.Marshal(reqBody)

	// Create a new HTTP request
	req := httptest.NewRequest("POST", "/api/calculate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := s.app.Test(req)
	s.Assert().NoError(err, "Expected no error")

	// Check the response
	s.Assert().Equal(fiber.StatusOK, resp.StatusCode, "Expected status OK")

	// Decode the response body
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	s.Assert().NoError(err, "Expected no error decoding response")

	// Verify the response contents
	s.Assert().Equal(float64(500), response["totalItems"], "Total items should match") // JSON numbers are decoded as float64
	s.Assert().Equal(map[string]interface{}{"500": float64(1)}, response["packs"], "Packs should match")
}

// TestCalculatePacks_InvalidRequest tests an invalid request to CalculatePacks
func (s *PackControllerTestSuite) TestCalculatePacks_InvalidRequest() {
	// Create a request with an invalid body
	req := httptest.NewRequest("POST", "/api/calculate", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := s.app.Test(req)
	s.Assert().NoError(err, "Expected no error")

	// Check the response
	s.Assert().Equal(fiber.StatusBadRequest, resp.StatusCode, "Expected status BadRequest")

	// Decode the response body
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	s.Assert().NoError(err, "Expected no error decoding response")

	// Verify the error message
	s.Assert().Equal("Invalid request", response["error"], "Error message should match")
}

// TestGetPackSizes_Success tests a successful GetPackSizes request
func (s *PackControllerTestSuite) TestGetPackSizes_Success() {
	// Set up the mock expectation using gomock API
	s.mockService.EXPECT().GetPackSizes().Return([]int{10, 20, 50}, nil)

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/api/pack-sizes", nil)

	// Perform the request
	resp, err := s.app.Test(req)
	s.Assert().NoError(err, "Expected no error")

	// Check the response
	s.Assert().Equal(fiber.StatusOK, resp.StatusCode, "Expected status OK")

	// Decode the response body
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	s.Assert().NoError(err, "Expected no error decoding response")

	// Verify the response contents
	s.Assert().Equal([]interface{}{float64(10), float64(20), float64(50)}, response["packSizes"], "Pack sizes should match")
}

// TestUpdatePackSizes_Success tests a successful UpdatePackSizes request
func (s *PackControllerTestSuite) TestUpdatePackSizes_Success() {
	// Set up the mock expectation using gomock API
	s.mockService.EXPECT().UpdatePackSizes([]int{100, 200, 300}).Return(nil)

	// Create a request body
	reqBody := map[string][]int{"packSizes": {100, 200, 300}}
	body, _ := json.Marshal(reqBody)

	// Create a new HTTP request
	req := httptest.NewRequest("POST", "/api/pack-sizes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := s.app.Test(req)
	s.Assert().NoError(err, "Expected no error")

	// Check the response
	s.Assert().Equal(fiber.StatusOK, resp.StatusCode, "Expected status OK")

	// Decode the response body
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	s.Assert().NoError(err, "Expected no error decoding response")

	// Verify the response contents
	s.Assert().Equal("Pack sizes updated successfully", response["message"], "Message should match")
}
