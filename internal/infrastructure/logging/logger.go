package logging // Define the package name as "logging" for logging utilities

import (
	"log" // Import the log package for basic logging
	"os"  // Import the os package for file handling
)

// Logger wraps a standard logger with custom functionality
type Logger struct {
	logger *log.Logger // Underlying logger instance
}

// NewLogger creates a new Logger instance
func NewLogger() *Logger {
	// Create a logger that writes to stdout with a prefix and flags for date/time
	logger := log.New(os.Stdout, "order-packs-calculator: ", log.LstdFlags)
	return &Logger{logger: logger}
}

// Info logs an info-level message
func (l *Logger) Info(msg string) {
	l.logger.Printf("[INFO] %s", msg) // Log the message with an INFO level
}

// Error logs an error-level message
func (l *Logger) Error(msg string, err error) {
	l.logger.Printf("[ERROR] %s: %v", msg, err) // Log the message and error with an ERROR level
}
