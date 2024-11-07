package logger

import (
	"fmt"
)

// Config defines the configuration for the logger.
// It includes options for the log type, log level, output location, and format.
type Config struct {
	// Type specifies the type of logger to use (e.g., "zap", "default").
	Type LoggerType `yaml:"type"`

	// Level defines the log level (e.g., "info", "error", "debug").
	Level LogLevel `yaml:"level"`

	// Output specifies the log output location (e.g., "stdout", "file").
	Output LogOutput `yaml:"output"`

	// Format determines the log format (e.g., "json", "plain").
	Format LogFormat `yaml:"format"`
}

// warnForDefaultLogger outputs a warning if the log level or formatting options
// are provided for the default Go logger, as these options are ignored.
// This function helps inform users that certain configuration settings are not applicable.
func warnForDefaultLogger(conf Config) {
	if conf.Level != "" || conf.Format != "" {
		fmt.Println("Warning: Log level and formatting are ignored when using the default Go logger.")
	}
}

// warnConcurrency outputs a warning indicating that the default Go logger
// does not handle concurrency safely, and should be used cautiously in concurrent environments.
// This serves as a reminder to users about potential concurrency issues.
func warnConcurrency() {
	fmt.Println("Warning: The default Go logger does not handle concurrency safely.")
}
