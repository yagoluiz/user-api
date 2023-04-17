package logger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestZapLogger(t *testing.T) {
	tests := []struct {
		name    string
		act     func(*testing.T, Logger, string, []any)
		level   zapcore.Level
		format  string
		args    []any
		message string
	}{
		{
			name: "Debug",
			act: func(_t *testing.T, l Logger, s string, _ []any) {
				l.Debug(s)
			},
			level:   zapcore.DebugLevel,
			format:  "anything",
			args:    nil,
			message: "anything",
		},
		{
			name: "Debugf",
			act: func(_t *testing.T, l Logger, s string, a []any) {
				l.Debugf(s, a...)
			},
			level:   zapcore.DebugLevel,
			format:  "1 + 1 = %d",
			args:    []any{2},
			message: "1 + 1 = 2",
		},
		{
			name: "Info",
			act: func(_t *testing.T, l Logger, s string, _ []any) {
				l.Info(s)
			},
			level:   zapcore.InfoLevel,
			format:  "anything",
			args:    nil,
			message: "anything",
		},
		{
			name: "Infof",
			act: func(_t *testing.T, l Logger, s string, a []any) {
				l.Infof(s, a...)
			},
			level:   zapcore.InfoLevel,
			format:  "1 + 1 = %d",
			args:    []any{2},
			message: "1 + 1 = 2",
		},
		{
			name: "Warn",
			act: func(_t *testing.T, l Logger, s string, _ []any) {
				l.Warn(s)
			},
			level:   zapcore.WarnLevel,
			format:  "anything",
			args:    nil,
			message: "anything",
		},
		{
			name: "Warnf",
			act: func(_t *testing.T, l Logger, s string, a []any) {
				l.Warnf(s, a...)
			},
			level:   zapcore.WarnLevel,
			format:  "1 + 1 = %d",
			args:    []any{2},
			message: "1 + 1 = 2",
		},
		{
			name: "Error",
			act: func(_t *testing.T, l Logger, s string, _ []any) {
				l.Error(s)
			},
			level:   zapcore.ErrorLevel,
			format:  "anything",
			args:    nil,
			message: "anything",
		},
		{
			name: "Errorf",
			act: func(_t *testing.T, l Logger, s string, a []any) {
				l.Errorf(s, a...)
			},
			level:   zapcore.ErrorLevel,
			format:  "1 + 1 = %d",
			args:    []any{2},
			message: "1 + 1 = 2",
		},
		{
			name: "Panic",
			act: func(t *testing.T, l Logger, s string, _ []any) {
				t.Helper()

				defer func() {
					if err := recover(); err == nil {
						t.Errorf("was expecting a panic")
					}
				}()

				l.Panic(s)
			},
			level:   zapcore.PanicLevel,
			format:  "anything",
			args:    nil,
			message: "anything",
		},
		{
			name: "Panicf",
			act: func(t *testing.T, l Logger, s string, a []any) {
				t.Helper()

				defer func() {
					if err := recover(); err == nil {
						t.Errorf("was expecting a panic")
					}
				}()

				l.Panicf(s, a...)
			},
			level:   zapcore.PanicLevel,
			format:  "1 + 1 = %d",
			args:    []any{2},
			message: "1 + 1 = 2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			z, logs := newZapLoggerWithLogs(zapcore.DebugLevel)

			test.act(t, z, test.format, test.args)

			if logs.Len() != 1 {
				t.Errorf("was expecting only one log")
			}

			log := logs.All()[0]
			if log.Level != test.level {
				t.Errorf("want: %v, got: %v", test.level, log.Level)
			}
			if log.Message != test.message {
				t.Errorf("want: %v, got: %v", test.message, log.Message)
			}
		})
	}
}

func TestZapLoggerWithField(t *testing.T) {
	z, logs := newZapLoggerWithLogs(zapcore.DebugLevel)
	key := "key"
	value := "value"

	z = z.WithField(key, value)

	z.Info("message")

	if logs.Len() != 1 {
		t.Errorf("was expecting only one log")
	}

	f := logs.All()[0].Context[0]

	if f.Key != key {
		t.Errorf("want: %v, got: %v", key, f.Key)
	}
	if f.String != value {
		t.Errorf("want: %v, got: %v", value, f.String)
	}
}

func newZapLoggerWithLogs(l zapcore.Level) (Logger, *observer.ObservedLogs) {
	core, logs := observer.New(l)

	return &ZapLogger{logger: zap.New(core)}, logs
}
