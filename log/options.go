package log

import "go.uber.org/zap/zapcore"

const (
	DefaultPrefix     = "ZIWI"
	DefaultDir        = "logs"
	DefaultLevel      = zapcore.InfoLevel
	DefaultTimeLayout = "2006-01-02 15:04:05.000"
	DefaultFormat     = "console" // console style
	DefaultMaxSize    = 100       // 100MB
	DefaultMaxBackups = 3         // Keep 3 old log files
	DefaultCompress   = false     // Not compress rotated log files
)

// Options for logger
type Options struct {
	Prefix    string // Log Prefix, e.g ZIWI
	Directory string // Log File Directory, e.g logs

	Level      string // Log Level, "debug", "info", "warn", "error", "dpanic", "panic", "fatal"
	TimeLayout string // Time Layout, e.g "2006-01-02 15:04:05.000"
	Format     string // Log Format, "console", "json"

	DisableCaller     bool // Disable caller information
	DisableStacktrace bool // Disable stack traces
	DisableSplitError bool // Disable separate error log files

	// Log rotation settings
	MaxSize    int  // Maximum size of log files in megabytes before rotation
	MaxBackups int  // Maximum number of old log files to retain
	Compress   bool // Whether to compress rotated log files
}

// NewOptions return the default Options.
//
// Default:
//
//	Prefix:    "ZIWI",
//	Directory: "logs",
//
//	Level:      "info",
//	TimeLayout: "2006-01-02 15:04:05.000",
//	Format:     "console",
//
//	DisableCaller:     false,
//	DisableStacktrace: false,
//	DisableSplitError: false,
//
//	// Default log rotation settings
//	MaxSize:    100, // 100MB
//	MaxBackups: 3,   // Keep 3 old log files
//	Compress:   false,
func NewOptions() *Options {
	return &Options{
		Prefix:    DefaultPrefix,
		Directory: DefaultDir,

		Level:      DefaultLevel.String(),
		TimeLayout: DefaultTimeLayout,
		Format:     DefaultFormat,

		DisableCaller:     false,
		DisableStacktrace: false,
		DisableSplitError: false,

		// Default log rotation settings
		MaxSize:    DefaultMaxSize,
		MaxBackups: DefaultMaxBackups,
		Compress:   DefaultCompress,
	}
}
