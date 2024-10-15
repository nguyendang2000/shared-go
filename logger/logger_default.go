package logger

import (
	"fmt"
	"log"
	"os"
)

// defaultLogger implements the Logger interface using the Go standard log package.
type defaultLogger struct {
	logger *log.Logger
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

	stdLogger := log.New(output, "", log.LstdFlags)
	return &defaultLogger{logger: stdLogger}
}

func (l *defaultLogger) Debug(msg string) {
	l.logger.Println("DEBUG: " + msg)
}

func (l *defaultLogger) Info(msg string) {
	l.logger.Println("INFO: " + msg)
}

func (l *defaultLogger) Warn(msg string) {
	l.logger.Println("WARN: " + msg)
}

func (l *defaultLogger) Error(msg string) {
	l.logger.Println("ERROR: " + msg)
}

func (l *defaultLogger) Fatal(msg string) {
	l.logger.Println("FATAL: " + msg)
	os.Exit(1)
}

func (l *defaultLogger) Panic(msg string) {
	l.logger.Println("PANIC: " + msg)
	panic(msg)
}

func (l *defaultLogger) Debugf(format string, args ...interface{}) {
	l.logger.Printf("DEBUG: "+format, args...)
}

func (l *defaultLogger) Infof(format string, args ...interface{}) {
	l.logger.Printf("INFO: "+format, args...)
}

func (l *defaultLogger) Warnf(format string, args ...interface{}) {
	l.logger.Printf("WARN: "+format, args...)
}

func (l *defaultLogger) Errorf(format string, args ...interface{}) {
	l.logger.Printf("ERROR: "+format, args...)
}

func (l *defaultLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Printf("FATAL: "+format, args...)
	os.Exit(1)
}

func (l *defaultLogger) Panicf(format string, args ...interface{}) {
	l.logger.Printf("PANIC: "+format, args...)
	panic(fmt.Sprintf(format, args...))
}
