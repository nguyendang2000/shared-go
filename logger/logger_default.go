package logger

import (
	"fmt"
	"log"
	"os"
)

// defaultLogger implements the Logger interface using the Go standard log package.
type defaultLogger struct {
	logger *log.Logger // The Go standard logger instance.
}

// newDefaultLogger initializes the default Go logger based on the provided configuration.
// It logs to the specified output location, or defaults to stdout if no output is provided.
func newDefaultLogger(conf Config) Logger {
	output := os.Stdout
	if conf.Output != LogOutputStdout && conf.Output != "" {
		f, err := os.OpenFile(string(conf.Output), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed to open log file: %v", err)
		}
		output = f
	}

	// Initialize the Go standard logger with the output and default flags.
	stdLogger := log.New(output, "", log.LstdFlags)
	return &defaultLogger{logger: stdLogger}
}

// Debug logs a debug-level message.
func (inst *defaultLogger) Debug(msg string) {
	inst.logger.Println("DEBUG: " + msg)
}

// Info logs an info-level message.
func (inst *defaultLogger) Info(msg string) {
	inst.logger.Println("INFO: " + msg)
}

// Warn logs a warning-level message.
func (inst *defaultLogger) Warn(msg string) {
	inst.logger.Println("WARN: " + msg)
}

// Error logs an error-level message.
func (inst *defaultLogger) Error(msg string) {
	inst.logger.Println("ERROR: " + msg)
}

// Fatal logs a fatal-level message and exits the application.
func (inst *defaultLogger) Fatal(msg string) {
	inst.logger.Println("FATAL: " + msg)
	os.Exit(1)
}

// Panic logs a panic-level message and panics.
func (inst *defaultLogger) Panic(msg string) {
	inst.logger.Println("PANIC: " + msg)
	panic(msg)
}

// Debugf logs a formatted debug-level message.
func (inst *defaultLogger) Debugf(format string, args ...interface{}) {
	inst.logger.Printf("DEBUG: "+format, args...)
}

// Infof logs a formatted info-level message.
func (inst *defaultLogger) Infof(format string, args ...interface{}) {
	inst.logger.Printf("INFO: "+format, args...)
}

// Warnf logs a formatted warning-level message.
func (inst *defaultLogger) Warnf(format string, args ...interface{}) {
	inst.logger.Printf("WARN: "+format, args...)
}

// Errorf logs a formatted error-level message.
func (inst *defaultLogger) Errorf(format string, args ...interface{}) {
	inst.logger.Printf("ERROR: "+format, args...)
}

// Fatalf logs a formatted fatal-level message and exits the application.
func (inst *defaultLogger) Fatalf(format string, args ...interface{}) {
	inst.logger.Printf("FATAL: "+format, args...)
	os.Exit(1)
}

// Panicf logs a formatted panic-level message and panics.
func (inst *defaultLogger) Panicf(format string, args ...interface{}) {
	inst.logger.Printf("PANIC: "+format, args...)
	panic(fmt.Sprintf(format, args...))
}
