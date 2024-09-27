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
	LoggerZap     LoggerType = "zap"     // Zap logger
	LoggerDefault LoggerType = "default" // Default Go logger

	// Log Levels
	LogLevelDebug LogLevel = "debug" // Debug log level
	LogLevelInfo  LogLevel = "info"  // Info log level
	LogLevelWarn  LogLevel = "warn"  // Warn log level
	LogLevelError LogLevel = "error" // Error log level
	LogLevelFatal LogLevel = "fatal" // Fatal log level
	LogLevelPanic LogLevel = "panic" // Panic log level

	// Log Outputs
	LogOutputStdout LogOutput = "stdout" // Log to stdout
	LogOutputFile   LogOutput = "file"   // Log to file

	// Log Formats
	LogFormatJSON  LogFormat = "json"  // JSON log format
	LogFormatPlain LogFormat = "plain" // Plain text log format
)
