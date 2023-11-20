package logs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestLoggerLogLvl(t *testing.T) {
	testLogsLvlVar1 := "MY_INSTANCE_LOGS_LVL"
	testLogsLvlVar2 := "MY_INSTANCE_LOGS_LVL_2"
	logger := NewLogger()
	assert.Equal(t, DebugLvl, logger.logLvl.lvl)

	_ = os.Setenv(testLogsLvlVar1, WarnLvlName)
	logger = NewLogger()
	logger.SetLogLvlEnvVariable(testLogsLvlVar1)
	assert.Equal(t, WarnLvl, logger.logLvl.lvl)

	_ = os.Setenv(testLogsLvlVar2, InfoLvlName)
	logger = NewLogger()
	logger.SetLogLvlEnvVariable(testLogsLvlVar2)
	assert.NotEqual(t, WarnLvl, logger.logLvl)
	assert.Equal(t, InfoLvl, logger.logLvl.lvl)

	_ = os.Unsetenv(testLogsLvlVar1)
	_ = os.Unsetenv(testLogsLvlVar2)
}

func TestLogger_GetLogLvlIntValue(t *testing.T) {
	testLogsLvlVar1 := "MY_INSTANCE_LOGS_LVL"
	logger := NewLogger()
	assert.Equal(t, DebugLvl, logger.LogLvlIntValue())

	logger.SetLogLvl(WarnLvlName)
	assert.Equal(t, WarnLvlName, logger.LogLvlName())

	_ = os.Setenv(testLogsLvlVar1, InfoLvlName)
	logger.SetLogLvlEnvVariable(testLogsLvlVar1)
	assert.Equal(t, InfoLvlName, logger.LogLvlName())

	_ = os.Unsetenv(testLogsLvlVar1)
}

func TestLogger_SetLogLvl(t *testing.T) {
	logger := NewLogger()
	assert.Equal(t, DebugLvlName, logger.LogLvlName())

	logger.SetLogLvl(WarnLvlName)
	assert.Equal(t, WarnLvlName, logger.LogLvlName())

	logger.SetLogLvl(ErrorLvlName)
	assert.Equal(t, ErrorLvlName, logger.LogLvlName())

	logger.SetLogLvl("invalid log level string")
	assert.Equal(t, DebugLvlName, logger.LogLvlName())
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
	logger := NewLogger()
	logger.SetLogLvl(ErrorLvlName)
	logger.Info(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.NotContainsf(t, buf.String(), testLog, "error-msg")
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
