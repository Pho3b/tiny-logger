package logs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
	"io"
	"os"
	"testing"
)

func TestLoggerLogLvl(t *testing.T) {
	testLogsLvlVar1 := "MY_INSTANCE_LOGS_LVL"
	testLogsLvlVar2 := "MY_INSTANCE_LOGS_LVL_2"
	logger := NewLogger()
	assert.Equal(t, log_level.DebugLvl, logger.logLvl.Lvl)

	_ = os.Setenv(testLogsLvlVar1, string(log_level.WarnLvlName))
	logger = NewLogger()
	logger.SetLogLvlEnvVariable(testLogsLvlVar1)
	assert.Equal(t, log_level.WarnLvl, logger.logLvl.Lvl)

	_ = os.Setenv(testLogsLvlVar2, string(log_level.InfoLvlName))
	logger = NewLogger()
	logger.SetLogLvlEnvVariable(testLogsLvlVar2)
	assert.NotEqual(t, log_level.WarnLvl, logger.logLvl.Lvl)
	assert.Equal(t, log_level.InfoLvl, logger.logLvl.Lvl)

	_ = os.Unsetenv(testLogsLvlVar1)
	_ = os.Unsetenv(testLogsLvlVar2)
}

func TestLogger_GetLogLvlIntValue(t *testing.T) {
	testLogsLvlVar1 := "MY_INSTANCE_LOGS_LVL"
	logger := NewLogger()
	assert.Equal(t, log_level.DebugLvl, logger.GetLogLvlIntValue())

	logger.SetLogLvl(log_level.WarnLvlName)
	assert.Equal(t, log_level.WarnLvlName, logger.GetLogLvlName())

	_ = os.Setenv(testLogsLvlVar1, string(log_level.InfoLvlName))
	logger.SetLogLvlEnvVariable(testLogsLvlVar1)
	assert.Equal(t, log_level.InfoLvlName, logger.GetLogLvlName())

	_ = os.Unsetenv(testLogsLvlVar1)
}

func TestLogger_SetLogLvl(t *testing.T) {
	logger := NewLogger()
	assert.Equal(t, log_level.DebugLvlName, logger.GetLogLvlName())

	logger.SetLogLvl(log_level.WarnLvlName)
	assert.Equal(t, log_level.WarnLvlName, logger.GetLogLvlName())

	logger.SetLogLvl(log_level.ErrorLvlName)
	assert.Equal(t, log_level.ErrorLvlName, logger.GetLogLvlName())

	logger.SetLogLvl("invalid log level string")
	assert.Equal(t, log_level.DebugLvlName, logger.GetLogLvlName())
}

func TestLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing DEBUG log"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	NewLogger().Info(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.Contains(t, buf.String(), testLog)
}

func TestLogger_InfoNotLogging(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing DEBUG log"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	logger := NewLogger().SetLogLvl(log_level.ErrorLvlName)
	logger.Info(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.NotContainsf(t, buf.String(), testLog, "logError-msg")
}

func TestLogger_Debug(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing INFO log"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	NewLogger().Debug(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.Contains(t, buf.String(), testLog)
}

func TestLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing WARN log"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	NewLogger().Warn(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.Contains(t, buf.String(), testLog)
}

func TestLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing ERROR log"
	originalStdErr := os.Stderr
	r, w, _ := os.Pipe()

	os.Stderr = w
	NewLogger().Error(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stderr = originalStdErr
	assert.Contains(t, buf.String(), testLog)
}

func TestLogger_BuildingMethods(t *testing.T) {
	logger := NewLogger()
	assert.IsType(t, &Logger{}, logger)
	assert.IsType(t, &Logger{}, logger.SetLogLvl(log_level.DebugLvlName))
	assert.IsType(t, &Logger{}, logger.SetLogLvlEnvVariable("test-env-var"))
}
