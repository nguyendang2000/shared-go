package logger

// LoggerType defines a type for supported logger names.
type LoggerType string

// LogLevel defines a type for supported log levels.
type LogLevel string

// LogOutput defines a type for supported log output locations.
type LogOutput string

// LogFormat defines a type for supported log formats.
type LogFormat string

const (
	// Logger Types
	LoggerZap     LoggerType = "zap"     // Zap logger, a structured logger for high-performance logging.
	LoggerDefault LoggerType = "default" // Default Go logger, suitable for basic logging needs.

	// Log Levels
	LogLevelDebug LogLevel = "debug" // Debug log level, used for detailed system information.
	LogLevelInfo  LogLevel = "info"  // Info log level, used for general operational information.
	LogLevelWarn  LogLevel = "warn"  // Warn log level, used for non-critical issues that require attention.
	LogLevelError LogLevel = "error" // Error log level, used for serious issues that affect functionality.
	LogLevelFatal LogLevel = "fatal" // Fatal log level, used for severe issues causing a program to exit.
	LogLevelPanic LogLevel = "panic" // Panic log level, used for issues causing panic and stack trace output.

	// Log Outputs
	LogOutputStdout LogOutput = "stdout" // Log to standard output.
	LogOutputFile   LogOutput = "file"   // Log to a specified file.

	// Log Formats
	LogFormatJSON  LogFormat = "json"  // JSON log format, suitable for structured and machine-readable logs.
	LogFormatPlain LogFormat = "plain" // Plain text log format, suitable for human-readable logs.
)
