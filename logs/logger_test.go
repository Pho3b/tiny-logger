package logs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"gitlab.com/docebo/libraries/go/tiny-logger/colors"
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
	assert.Equal(t, DebugLvl, logger.GetLogLvlIntValue())

	logger.SetLogLvl(WarnLvlName)
	assert.Equal(t, WarnLvlName, logger.GetLogLvlName())

	_ = os.Setenv(testLogsLvlVar1, InfoLvlName)
	logger.SetLogLvlEnvVariable(testLogsLvlVar1)
	assert.Equal(t, InfoLvlName, logger.GetLogLvlName())

	_ = os.Unsetenv(testLogsLvlVar1)
}

func TestLogger_SetLogLvl(t *testing.T) {
	logger := NewLogger()
	assert.Equal(t, DebugLvlName, logger.GetLogLvlName())

	logger.SetLogLvl(WarnLvlName)
	assert.Equal(t, WarnLvlName, logger.GetLogLvlName())

	logger.SetLogLvl(ErrorLvlName)
	assert.Equal(t, ErrorLvlName, logger.GetLogLvlName())

	logger.SetLogLvl("invalid log level string")
	assert.Equal(t, DebugLvlName, logger.GetLogLvlName())
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

func TestLogger_Log(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my GENERIC log test"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	logger := NewLogger()
	logger.Log("jldfald", testLog)
	logger.Log(colors.Black, testLog)
	logger.Log(colors.Cyan, testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.Contains(t, buf.String(), colors.White)
	assert.Contains(t, buf.String(), colors.Black)
	assert.Contains(t, buf.String(), colors.Cyan)
}

func TestLogger_BuildingMethods(t *testing.T) {
	logger := NewLogger()
	assert.IsType(t, &Logger{}, logger)
	assert.IsType(t, &Logger{}, logger.SetLogLvl(DebugLvlName))
	assert.IsType(t, &Logger{}, logger.SetLogLvlEnvVariable(InfoLvlName))
}
