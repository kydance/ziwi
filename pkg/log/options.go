package log

import "go.uber.org/zap/zapcore"

// Options for logger
type Options struct {
	DisableCaller     bool
	DisableStacktrace bool
	Level             string // debug, info, warn, error, dpanic, panic,  fatal
	Format            string // console, json
	OutputPaths       []string
}

// LogOptions return the default Options.
//
//	DisableCaller: false,
//	DisableStacktrace: false,
//	Level: Info,
//	Format: console,
//	OutputPaths: stdout,
func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
