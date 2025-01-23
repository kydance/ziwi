// Package log is a log package.
package log

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = &ziwiLogger{}

var (
	mu sync.Mutex

	std = NewLogger(NewOptions())
)

// Logger defines the interface, includes the supported logging methods.
// For each log level, it exposes four methods:
//
//   - methods named after the log level for log.Print-style logging
//   - methods ending in "w" for loosely-typed structured logging (read as "info with")
//   - methods ending in "f" for log.Printf-style logging
//   - methods ending in "ln" for log.Println-style logging
type Logger interface {
	Sync()

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})

	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	Panicln(args ...interface{})
	Fatalln(args ...interface{})
}

// ziwiLogger is the implement of Logger interface.
// It wraps zap.Logger.
type ziwiLogger struct {
	z *zap.Logger
}

// Init Logger with the specified options.
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

func NewLogger(opts *Options) *ziwiLogger {
	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,

		// zap internal error output path
		ErrorOutputPaths: []string{"stderr"},
	}

	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &ziwiLogger{z: z}

	zap.RedirectStdLog(z)

	return logger
}

// Sync flushs any buffered log entries. Applications should take care to call Sync before exiting.
func Sync() { std.Sync() }

// Sync flushs any buffered log entries. Applications should take care to call Sync before exiting.
func (l *ziwiLogger) Sync() { _ = l.z.Sync() }

func Debug(args ...interface{})                 { std.z.Sugar().Debug(args...) }
func (l *ziwiLogger) Debug(args ...interface{}) { l.z.Sugar().Debug(args...) }

func Info(args ...interface{})                 { std.z.Sugar().Info(args...) }
func (l *ziwiLogger) Info(args ...interface{}) { l.z.Sugar().Info(args...) }

func Warn(args ...interface{})                 { std.z.Sugar().Warn(args...) }
func (l *ziwiLogger) Warn(args ...interface{}) { l.z.Sugar().Warn(args...) }

func Error(args ...interface{})                 { std.z.Sugar().Error(args...) }
func (l *ziwiLogger) Error(args ...interface{}) { l.z.Sugar().Error(args...) }

func Panic(args ...interface{})                 { std.z.Sugar().Panic(args...) }
func (l *ziwiLogger) Panic(args ...interface{}) { l.z.Sugar().Panic(args...) }

func Fatal(args ...interface{})                 { std.z.Sugar().Fatal(args...) }
func (l *ziwiLogger) Fatal(args ...interface{}) { l.z.Sugar().Fatal(args...) }

func Debugln(args ...interface{})                 { std.z.Sugar().Debugln(args...) }
func (l *ziwiLogger) Debugln(args ...interface{}) { l.z.Sugar().Debugln(args...) }

func Infoln(args ...interface{})                 { std.z.Sugar().Infoln(args...) }
func (l *ziwiLogger) Infoln(args ...interface{}) { l.z.Sugar().Infoln(args...) }

func Warnln(args ...interface{})                 { std.z.Sugar().Warnln(args...) }
func (l *ziwiLogger) Warnln(args ...interface{}) { l.z.Sugar().Warnln(args...) }

func Errorln(args ...interface{})                 { std.z.Sugar().Errorln(args...) }
func (l *ziwiLogger) Errorln(args ...interface{}) { l.z.Sugar().Errorln(args...) }

func Panicln(args ...interface{})                 { std.z.Sugar().Panicln(args...) }
func (l *ziwiLogger) Panicln(args ...interface{}) { l.z.Sugar().Panicln(args...) }

func Fatalln(args ...interface{})                 { std.z.Sugar().Fatalln(args...) }
func (l *ziwiLogger) Fatalln(args ...interface{}) { l.z.Sugar().Fatalln(args...) }

// Debugw logs a message with some additional context.
// The variadic key-value pairs are treated as they are in With.
func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

// Debugw logs a message with some additional context.
// The variadic key-value pairs are treated as they are in With.
func (l *ziwiLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context.
// The variadic key-value pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

// Infow logs a message with some additional context.
// The variadic key-value pairs are treated as they are in With.
func (l *ziwiLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context.
// The variadic key-value pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

// Warnw logs a message with some additional context.
// The variadic key-value pairs are treated as they are in With.
func (l *ziwiLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context.
// The variadic key-value pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context.
// The variadic key-value pairs are treated as they are in With.
func (l *ziwiLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Errorw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics.
// The variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics.
// The variadic key-value pairs are treated as they are in With.
func (l *ziwiLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit.
// The variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit.
// The variadic key-value pairs are treated as they are in With.
func (l *ziwiLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Fatalw(msg, keysAndValues...)
}

// ----------

// Debugf formats the message according to the format specifier and logs it.
func Debugf(template string, args ...interface{}) {
	std.z.Sugar().Debugf(template, args...)
}

// Debugf formats the message according to the format specifier and logs it.
func (l *ziwiLogger) Debugf(template string, args ...interface{}) {
	l.z.Sugar().Debugf(template, args...)
}

// Infof formats the message according to the format specifier and logs it.
func Infof(template string, args ...interface{}) {
	std.z.Sugar().Infof(template, args...)
}

// Infof formats the message according to the format specifier and logs it.
func (l *ziwiLogger) Infof(template string, args ...interface{}) {
	l.z.Sugar().Infof(template, args...)
}

// Warnf formats the message according to the format specifier and logs it.
func Warnf(template string, args ...interface{}) {
	std.z.Sugar().Warnf(template, args...)
}

// Warnf formats the message according to the format specifier and logs it.
func (l *ziwiLogger) Warnf(template string, args ...interface{}) {
	l.z.Sugar().Warnf(template, args...)
}

// Errorf formats the message according to the format specifier and logs it.
func Errorf(template string, args ...interface{}) {
	std.z.Sugar().Errorf(template, args...)
}

// Errorf formats the message according to the format specifier and logs it.
func (l *ziwiLogger) Errorf(template string, args ...interface{}) {
	l.z.Sugar().Errorf(template, args...)
}

// Panicf formats the message according to the format specifier and panics.
func Panicf(template string, args ...interface{}) {
	std.z.Sugar().Panicf(template, args...)
}

// Panicf formats the message according to the format specifier and panics.
func (l *ziwiLogger) Panicf(template string, args ...interface{}) {
	l.z.Sugar().Panicf(template, args...)
}

// Fatalf formats the message according to the format specifier and calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	std.z.Sugar().Fatalf(template, args...)
}

// Fatalf formats the message according to the format specifier and calls os.Exit.
func (l *ziwiLogger) Fatalf(template string, args ...interface{}) {
	l.z.Sugar().Fatalf(template, args...)
}
