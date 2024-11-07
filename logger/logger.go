package logger

// Logger defines the interface for a logging system that supports different log levels and formatting.
type Logger interface {
	// Debug logs a debug-level message.
	Debug(msg string)
	// Info logs an informational message.
	Info(msg string)
	// Warn logs a warning message.
	Warn(msg string)
	// Error logs an error message.
	Error(msg string)
	// Fatal logs a fatal-level message and exits the program.
	Fatal(msg string)
	// Panic logs a panic-level message and panics.
	Panic(msg string)

	// Debugf logs a formatted debug-level message.
	Debugf(format string, args ...interface{})
	// Infof logs a formatted informational message.
	Infof(format string, args ...interface{})
	// Warnf logs a formatted warning message.
	Warnf(format string, args ...interface{})
	// Errorf logs a formatted error message.
	Errorf(format string, args ...interface{})
	// Fatalf logs a formatted fatal-level message and exits the program.
	Fatalf(format string, args ...interface{})
	// Panicf logs a formatted panic-level message and panics.
	Panicf(format string, args ...interface{})
}

var globalLogger Logger // The global logger instance.

// SetGlobalLogger sets the given logger as the global logger.
func SetGlobalLogger(l Logger) {
	globalLogger = l
}

// GlobalLogger returns the globally set logger instance.
// If no global logger is set, it panics to inform the caller.
func GlobalLogger() Logger {
	if globalLogger == nil {
		panic("global logger is not set")
	}
	return globalLogger
}

// New initializes a new logger based on the provided configuration.
// It returns a Zap logger if specified; otherwise, it defaults to the standard Go logger.
func New(conf Config) Logger {
	switch conf.Type {
	case "zap":
		return newZapLogger(conf)
	default:
		warnForDefaultLogger(conf)
		warnConcurrency()
		return newDefaultLogger(conf)
	}
}
