package logger

import (
	"go.uber.org/zap"
)

// zapLogger implements the Logger interface using the Zap logging library.
type zapLogger struct {
	logger *zap.Logger // The Zap logger instance.
}

// newZapLogger initializes a Zap logger based on the provided configuration.
// It sets the log level, output location, and format, defaulting to JSON format
// and stdout output if none are provided.
func newZapLogger(conf Config) Logger {
	var zapCfg zap.Config
	if conf.Format == LogFormatJSON {
		zapCfg = zap.NewProductionConfig()
	} else {
		zapCfg = zap.NewDevelopmentConfig()
	}

	// Set log level based on configuration.
	switch conf.Level {
	case LogLevelDebug:
		zapCfg.Level.SetLevel(zap.DebugLevel)
	case LogLevelError:
		zapCfg.Level.SetLevel(zap.ErrorLevel)
	case LogLevelFatal:
		zapCfg.Level.SetLevel(zap.FatalLevel)
	case LogLevelPanic:
		zapCfg.Level.SetLevel(zap.PanicLevel)
	case LogLevelWarn:
		zapCfg.Level.SetLevel(zap.WarnLevel)
	default:
		zapCfg.Level.SetLevel(zap.InfoLevel)
	}

	// Set output location if specified.
	if conf.Output != LogOutputStdout && conf.Output != "" {
		zapCfg.OutputPaths = []string{string(conf.Output)}
	}

	logger, _ := zapCfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	return &zapLogger{logger: logger}
}

// Debug logs a debug-level message.
func (inst *zapLogger) Debug(msg string) {
	inst.logger.Debug(msg)
}

// Info logs an info-level message.
func (inst *zapLogger) Info(msg string) {
	inst.logger.Info(msg)
}

// Warn logs a warning-level message.
func (inst *zapLogger) Warn(msg string) {
	inst.logger.Warn(msg)
}

// Error logs an error-level message.
func (inst *zapLogger) Error(msg string) {
	inst.logger.Error(msg)
}

// Fatal logs a fatal-level message and exits the application.
func (inst *zapLogger) Fatal(msg string) {
	inst.logger.Fatal(msg)
}

// Panic logs a panic-level message and panics.
func (inst *zapLogger) Panic(msg string) {
	inst.logger.Panic(msg)
}

// Debugf logs a formatted debug-level message.
func (inst *zapLogger) Debugf(format string, args ...any) {
	inst.logger.Sugar().Debugf(format, args...)
}

// Infof logs a formatted info-level message.
func (inst *zapLogger) Infof(format string, args ...any) {
	inst.logger.Sugar().Infof(format, args...)
}

// Warnf logs a formatted warning-level message.
func (inst *zapLogger) Warnf(format string, args ...any) {
	inst.logger.Sugar().Warnf(format, args...)
}

// Errorf logs a formatted error-level message.
func (inst *zapLogger) Errorf(format string, args ...any) {
	inst.logger.Sugar().Errorf(format, args...)
}

// Fatalf logs a formatted fatal-level message and exits the application.
func (inst *zapLogger) Fatalf(format string, args ...any) {
	inst.logger.Sugar().Fatalf(format, args...)
}

// Panicf logs a formatted panic-level message and panics.
func (inst *zapLogger) Panicf(format string, args ...any) {
	inst.logger.Sugar().Panicf(format, args...)
}
