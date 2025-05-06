package config

import (
	"io/ioutil"
	"os"      // Import os for setting environment variables
	"testing" // Import the testing package for writing unit tests

	"github.com/stretchr/testify/suite" // Import testify/suite for test suites
)

// ConfigTestSuite defines the test suite for the config package
type ConfigTestSuite struct {
	suite.Suite        // Embed the testify suite
	tempDir     string // Temporary directory for test files
}

// SetupTest sets up the test environment before each test
func (s *ConfigTestSuite) SetupTest() {
	// Create a temporary directory for test files
	tempDir, err := ioutil.TempDir("", "config-test")
	s.Require().NoError(err, "Failed to create temp directory")
	s.tempDir = tempDir

	// Change the working directory to the temp directory
	err = os.Chdir(s.tempDir)
	s.Require().NoError(err, "Failed to change working directory")

	// Clear environment variables
	os.Unsetenv("PORT")
	os.Unsetenv("PACK_SIZES")
}

// TearDownTest cleans up the test environment after each test
func (s *ConfigTestSuite) TearDownTest() {
	// Remove the temporary directory
	err := os.RemoveAll(s.tempDir)
	s.Require().NoError(err, "Failed to remove temp directory")
}

// TestConfigTestSuite runs the test suite
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

// TestDefaultValues tests loading with default values
func (s *ConfigTestSuite) TestDefaultValues() {
	// Load the configuration
	cfg, err := LoadConfig()
	s.Assert().NoError(err, "Expected no error")

	// Verify default values
	s.Assert().Equal(":3000", cfg.Port, "Port should match default")
	s.Assert().Equal([]int{250, 500, 1000, 2000, 5000}, cfg.PackSizes, "Pack sizes should match default")
}

// TestEnvironmentVariables tests loading from environment variables
func (s *ConfigTestSuite) TestEnvironmentVariables() {
	// Set environment variables
	os.Setenv("PORT", "4000")
	os.Setenv("PACK_SIZES", "100,200,300")

	// Load the configuration
	cfg, err := LoadConfig()
	s.Assert().NoError(err, "Expected no error")

	// Verify environment variable values
	s.Assert().Equal(":4000", cfg.Port, "Port should match environment variable")
	s.Assert().Equal([]int{100, 200, 300}, cfg.PackSizes, "Pack sizes should match environment variable")
}

// TestConfigFile tests loading from a config.yaml file
func (s *ConfigTestSuite) TestConfigFile() {
	// Create a temporary config.yaml file
	configContent := `
port: "5000"
pack_sizes: "50,100,150"
`
	err := ioutil.WriteFile("config.yaml", []byte(configContent), 0644)
	s.Require().NoError(err, "Failed to create config.yaml")

	// Load the configuration
	cfg, err := LoadConfig()
	s.Assert().NoError(err, "Expected no error")

	// Verify config file values
	s.Assert().Equal(":5000", cfg.Port, "Port should match config file")
	s.Assert().Equal([]int{50, 100, 150}, cfg.PackSizes, "Pack sizes should match config file")
}

// TestInvalidPackSizes tests handling of invalid pack sizes in config
func (s *ConfigTestSuite) TestInvalidPackSizes() {
	// Create a temporary config.yaml file with invalid pack sizes
	configContent := `
port: "6000"
pack_sizes: "invalid,100,200"
`
	err := ioutil.WriteFile("config.yaml", []byte(configContent), 0644)
	s.Require().NoError(err, "Failed to create config.yaml")

	// Load the configuration
	cfg, err := LoadConfig()
	s.Assert().NoError(err, "Expected no error")

	// Verify that invalid pack sizes are skipped
	s.Assert().Equal(":6000", cfg.Port, "Port should match config file")
	s.Assert().Equal([]int{100, 200}, cfg.PackSizes, "Invalid pack sizes should be skipped")
}
