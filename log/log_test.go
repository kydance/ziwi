package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog_Init(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	opts := NewOptions()
	opts.Compress = false
	opts.Directory = "/Users/kyden/git-space/ziwi/logs"
	opts.DisableCaller = false
	opts.DisableSplitError = false
	opts.DisableStacktrace = false
	opts.Format = "console"
	opts.Level = "debug"
	opts.MaxBackups = 2
	opts.MaxSize = 1
	opts.Prefix = "Ziwi_Log_Test_"
	opts.TimeLayout = "2006-01-02 15:04:05.000"
	assert.NotNil(t)

	logger := NewLogger(opts)
	defer logger.Sync()

	for i := range 10_000 {
		logger.Debugw("test debug", "i", i)
		logger.Infow("test info", "i", i)
		logger.Warnw("test warn", "i", i)
		logger.Errorw("test error", "i", i)
	}
}
