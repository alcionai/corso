package logger_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/logger"
)

type LoggerUnitSuite struct {
	tester.Suite
}

func TestLoggerUnitSuite(t *testing.T) {
	suite.Run(t, &LoggerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *LoggerUnitSuite) TestAddLoggingFlags() {
	t := suite.T()

	logger.DebugAPIFV = false
	logger.ReadableLogsFV = false

	cmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			assert.True(t, logger.DebugAPIFV, logger.DebugAPIFN)
			assert.True(t, logger.ReadableLogsFV, logger.ReadableLogsFN)
			assert.Equal(t, string(logger.LLError), logger.LogLevelFV, logger.LogLevelFN)
			assert.Equal(t, string(logger.LFText), logger.LogFormatFV, logger.LogFormatFN)
			assert.True(t, logger.MaskSensitiveDataFV, logger.MaskSensitiveDataFN)
			// empty assertion here, instead of matching "log-file", because the LogFile
			// var isn't updated by running the command (this is expected and correct),
			// while the logFileFV remains unexported.
			assert.Empty(t, logger.ResolvedLogFile, logger.LogFileFN)
		},
	}

	logger.AddLoggingFlags(cmd)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + logger.DebugAPIFN,
		"--" + logger.LogFileFN, "log-file",
		"--" + logger.LogLevelFN, string(logger.LLError),
		"--" + logger.LogFormatFN, string(logger.LFText),
		"--" + logger.ReadableLogsFN,
		"--" + logger.MaskSensitiveDataFN,
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *LoggerUnitSuite) TestPreloadLoggingFlags() {
	t := suite.T()

	logger.DebugAPIFV = false
	logger.ReadableLogsFV = false

	args := []string{
		"--" + logger.DebugAPIFN,
		"--" + logger.LogFileFN, "log-file",
		"--" + logger.LogLevelFN, string(logger.LLError),
		"--" + logger.LogFormatFN, string(logger.LFText),
		"--" + logger.ReadableLogsFN,
		"--" + logger.MaskSensitiveDataFN,
	}

	settings := logger.PreloadLoggingFlags(args)

	assert.True(t, logger.DebugAPIFV, logger.DebugAPIFN)
	assert.True(t, logger.ReadableLogsFV, logger.ReadableLogsFN)
	assert.Equal(t, "log-file", settings.File, "settings.File")
	assert.Equal(t, logger.LLError, settings.Level, "settings.Level")
	assert.Equal(t, logger.LFText, settings.Format, "settings.Format")
	assert.Equal(t, logger.PIIHash, settings.PIIHandling, "settings.PIIHandling")
}

func (suite *LoggerUnitSuite) TestPreloadLoggingFlags_badArgsEnsureDefault() {
	t := suite.T()

	logger.DebugAPIFV = false
	logger.ReadableLogsFV = false

	args := []string{
		"--" + logger.DebugAPIFN,
		"--" + logger.LogFileFN, "log-file",
		"--" + logger.LogLevelFN, "not-a-level",
		"--" + logger.LogFormatFN, "not-a-format",
		"--" + logger.ReadableLogsFN,
		"--" + logger.MaskSensitiveDataFN,
	}

	settings := logger.PreloadLoggingFlags(args)
	settings = settings.EnsureDefaults()

	assert.Equal(t, logger.LLInfo, settings.Level, "settings.Level")
	assert.Equal(t, logger.LFText, settings.Format, "settings.Format")
}

func (suite *LoggerUnitSuite) TestSettings_ensureDefaults() {
	t := suite.T()

	s := logger.Settings{}
	require.Empty(t, s.File, "file")
	require.Empty(t, s.Level, "level")
	require.Empty(t, s.Format, "format")
	require.Empty(t, s.PIIHandling, "piialg")

	s = s.EnsureDefaults()
	require.NotEmpty(t, s.File, "file")
	require.NotEmpty(t, s.Level, "level")
	require.NotEmpty(t, s.Format, "format")
	require.NotEmpty(t, s.PIIHandling, "piialg")
}
