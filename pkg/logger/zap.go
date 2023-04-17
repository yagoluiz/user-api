package logger

import (
	"github.com/yagoluiz/user-api/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
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

	return &ZapLogger{logger: logger}, nil
}

func (l *ZapLogger) Debug(args ...any) {
	l.logger.Sugar().Debug(args...)
}

func (l *ZapLogger) Debugf(format string, args ...any) {
	l.logger.Sugar().Debugf(format, args...)
}

func (l *ZapLogger) Info(args ...any) {
	l.logger.Sugar().Info(args...)
}

func (l *ZapLogger) Infof(format string, args ...any) {
	l.logger.Sugar().Infof(format, args...)
}

func (l *ZapLogger) Warn(args ...any) {
	l.logger.Sugar().Warn(args...)
}

func (l *ZapLogger) Warnf(format string, args ...any) {
	l.logger.Sugar().Warnf(format, args...)
}

func (l *ZapLogger) Error(args ...any) {
	l.logger.Sugar().Error(args...)
}

func (l *ZapLogger) Errorf(format string, args ...any) {
	l.logger.Sugar().Errorf(format, args...)
}

func (l *ZapLogger) Fatal(args ...any) {
	l.logger.Sugar().Fatal(args...)
}

func (l *ZapLogger) Fatalf(format string, args ...any) {
	l.logger.Sugar().Fatalf(format, args...)
}

func (l *ZapLogger) Panic(args ...any) {
	l.logger.Sugar().Panic(args...)
}

func (l *ZapLogger) Panicf(format string, args ...any) {
	l.logger.Sugar().Panicf(format, args...)
}

func (l *ZapLogger) WithField(key string, value any) Logger {
	return &ZapLogger{logger: l.logger.With(zap.Any(key, value))}
}
