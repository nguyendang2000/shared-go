package logger

import (
	"go.uber.org/zap"
)

// zapLogger implements the Logger interface using the Zap logging library.
type zapLogger struct {
	logger *zap.Logger
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

	// Set log level
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

	// Set output location
	if conf.Output != LogOutputStdout && conf.Output != "" {
		zapCfg.OutputPaths = []string{string(conf.Output)}
	}

	logger, _ := zapCfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	return &zapLogger{logger: logger}
}

func (l *zapLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *zapLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *zapLogger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *zapLogger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *zapLogger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *zapLogger) Panic(msg string) {
	l.logger.Panic(msg)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.logger.Sugar().Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.logger.Sugar().Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.logger.Sugar().Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.logger.Sugar().Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Sugar().Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.logger.Sugar().Panicf(format, args...)
}
