package logger

import (
	"fmt"
)

// Config defines the configuration for the logger.
// It includes options for the log type, log level, output location, and format.
type Config struct {
	Type   LoggerType `yaml:"type"`   // Type of logger to use (e.g., "zap", "default")
	Level  LogLevel   `yaml:"level"`  // Log level (e.g., "info", "error", "debug")
	Output LogOutput  `yaml:"output"` // Log output location (e.g., "stdout", "file")
	Format LogFormat  `yaml:"format"` // Log format (e.g., "json", "plain")
}

// warnForDefaultLogger outputs a warning if the log level or formatting options
// are provided for the default Go logger, as these options are ignored.
func warnForDefaultLogger(cfg Config) {
	if cfg.Level != "" || cfg.Format != "" {
		fmt.Println("Warning: Log level and formatting are ignored when using the default Go logger.")
	}
}

// warnConcurrency outputs a warning indicating that the default Go logger
// does not handle concurrency safely, and should be used cautiously in concurrent environments.
func warnConcurrency() {
	fmt.Println("Warning: The default Go logger does not handle concurrency safely.")
}
