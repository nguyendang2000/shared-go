package logger

type Logger interface {
	Debug(msg string) // Logs debug information
	Info(msg string)  // Logs informational messages
	Warn(msg string)  // Logs warning messages
	Error(msg string) // Logs error messages
	Fatal(msg string) // Logs critical errors and exits the program
	Panic(msg string) // Logs critical errors and panics

	Debugf(format string, args ...interface{}) // Formatted debug logs
	Infof(format string, args ...interface{})  // Formatted info logs
	Warnf(format string, args ...interface{})  // Formatted warning logs
	Errorf(format string, args ...interface{}) // Formatted error logs
	Fatalf(format string, args ...interface{}) // Formatted fatal logs and exits the program
	Panicf(format string, args ...interface{}) // Formatted panic logs and panics
}

var globalLogger Logger

// SetGlobalLogger sets the given logger as the global logger.
func SetGlobalLogger(l Logger) {
	globalLogger = l
}

// GlobalLogger returns the globally set logger instance.
// If no global logger is set, it returns nil, so you should handle this in the calling code.
func GlobalLogger() Logger {
	if globalLogger == nil {
		panic("global logger is not set")
	}
	return globalLogger
}

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
