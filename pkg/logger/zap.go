package logger

import (
	"github.com/yagoluiz/user-api/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.Logger
}

type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Panic(args ...any)
	Panicf(format string, args ...any)
	WithField(key string, value any) Logger
}

func NewLogger(cfg *configs.Config) (Logger, error) {
	var c zap.Config
	if cfg.Debug {
		c = zap.NewDevelopmentConfig()
	} else {
		c = zap.NewProductionConfig()
	}
	c.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	c.Encoding = "json"

	logger, err := c.Build()
	if err != nil {
		return nil, err
	}

	return &zapLogger{logger: logger}, nil
}

func (l *zapLogger) Debug(args ...any) {
	l.logger.Sugar().Debug(args...)
}

func (l *zapLogger) Debugf(format string, args ...any) {
	l.logger.Sugar().Debugf(format, args...)
}

func (l *zapLogger) Info(args ...any) {
	l.logger.Sugar().Info(args...)
}

func (l *zapLogger) Infof(format string, args ...any) {
	l.logger.Sugar().Infof(format, args...)
}

func (l *zapLogger) Warn(args ...any) {
	l.logger.Sugar().Warn(args...)
}

func (l *zapLogger) Warnf(format string, args ...any) {
	l.logger.Sugar().Warnf(format, args...)
}

func (l *zapLogger) Error(args ...any) {
	l.logger.Sugar().Error(args...)
}

func (l *zapLogger) Errorf(format string, args ...any) {
	l.logger.Sugar().Errorf(format, args...)
}

func (l *zapLogger) Fatal(args ...any) {
	l.logger.Sugar().Fatal(args...)
}

func (l *zapLogger) Fatalf(format string, args ...any) {
	l.logger.Sugar().Fatalf(format, args...)
}

func (l *zapLogger) Panic(args ...any) {
	l.logger.Sugar().Panic(args...)
}

func (l *zapLogger) Panicf(format string, args ...any) {
	l.logger.Sugar().Panicf(format, args...)
}

func (l *zapLogger) WithField(key string, value any) Logger {
	return &zapLogger{logger: l.logger.With(zap.Any(key, value))}
}
