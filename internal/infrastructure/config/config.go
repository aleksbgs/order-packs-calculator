package config // Define the package name as "config" for configuration management

import (
	"log"     // Import the log package for logging
	"strconv" // Import strconv to check if the port is a number
	"strings" // Import the strings package for string manipulation

	"github.com/spf13/viper" // Import the Viper library for configuration management
)

// Config holds the application configuration settings
type Config struct {
	Port      string // Port on which the server will listen (e.g., ":3000")
	PackSizes []int  // Default pack sizes for the application
}

// LoadConfig loads the configuration using Viper
func LoadConfig() (*Config, error) {
	// Initialize Viper
	v := viper.New() // Create a new Viper instance

	// Set configuration file name and paths
	v.SetConfigName("config") // Set the base name of the config file (e.g., config.yaml, config.json)
	v.AddConfigPath(".")      // Look for the config file in the current directory
	v.AddConfigPath("/root")  // Look in /root directory (matches WORKDIR in Dockerfile)

	// Set the file type (optional, Viper will try multiple formats like JSON, YAML, etc.)
	v.SetConfigType("yaml") // Specify that we prefer YAML format (you can change to "json" if needed)

	// Enable environment variable support
	v.AutomaticEnv() // Automatically read environment variables

	// Bind specific environment variables to Viper keys
	v.BindEnv("port", "PORT")             // Bind PORT environment variable to "port" key
	v.BindEnv("pack_sizes", "PACK_SIZES") // Bind PACK_SIZES environment variable to "pack_sizes" key

	// Set default values
	v.SetDefault("port", ":3000")                        // Default port if not specified
	v.SetDefault("pack_sizes", "250,500,1000,2000,5000") // Default pack sizes as a comma-separated string

	// Read the configuration file (if it exists)
	if err := v.ReadInConfig(); err != nil { // Attempt to read the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { // Check if the error is something other than file not found
			return nil, err // Return the error if it's not a "file not found" error
		}
		log.Println("Config file not found; falling back to environment variables or defaults")
	} else {
		log.Printf("Loaded configuration from file: %s", v.ConfigFileUsed()) // Log the config file path
	}

	// Create a new Config instance
	cfg := &Config{}

	// Load the port from Viper and ensure it's in the correct format
	port := v.GetString("port") // Get the port value as a string
	// If the port is a plain number (e.g., "3000"), prepend ":" to make it ":3000"
	if _, err := strconv.Atoi(port); err == nil { // Check if port is a number
		port = ":" + port // Prepend ":" to make it a valid address
	}
	// If the port already contains a colon (e.g., ":3000" or "localhost:3000"), use it as-is
	cfg.Port = port
	log.Printf("Using port: %s", cfg.Port) // Log the port value

	// Load the pack sizes from Viper
	packSizesStr := v.GetString("pack_sizes")                  // Get the pack sizes as a comma-separated string
	log.Printf("Raw pack sizes from config: %s", packSizesStr) // Log the raw pack sizes string
	sizes := strings.Split(packSizesStr, ",")                  // Split the string by commas
	cfg.PackSizes = make([]int, 0, len(sizes))                 // Initialize the PackSizes slice
	for _, size := range sizes {                               // Loop through each size string
		size = strings.TrimSpace(size) // Remove any whitespace
		if size == "" {                // Skip empty entries
			continue
		}
		num, err := strconv.Atoi(size) // Convert the string to an integer
		if err != nil || num <= 0 {    // If conversion fails or the number is invalid, skip it
			log.Printf("Skipping invalid pack size: %s", size) // Log invalid pack size
			continue
		}
		cfg.PackSizes = append(cfg.PackSizes, num) // Add the valid pack size to the slice
	}

	// Ensure there are pack sizes (fall back to defaults if none were parsed)
	if len(cfg.PackSizes) == 0 { // Check if the PackSizes slice is empty
		log.Println("No valid pack sizes found; using default pack sizes: 250,500,1000,2000,5000")
		cfg.PackSizes = []int{250, 500, 1000, 2000, 5000} // Set default pack sizes
	} else {
		log.Printf("Loaded pack sizes: %v", cfg.PackSizes) // Log the final pack sizes
	}

	return cfg, nil // Return the loaded configuration and nil error
}
