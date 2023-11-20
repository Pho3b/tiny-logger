package logs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestPackageLoggerLogLvl(t *testing.T) {
	assert.Equal(t, DebugLvl, packageLogLvl.lvl)
	assert.Equal(t, DebugLvlName, packageLogLvl.LvlName())

	testLogsLvlVar1 := "MY_INSTANCE_LOGS_LVL"
	testLogsLvlVar2 := "MY_INSTANCE_LOGS_LVL_2"

	_ = os.Setenv(testLogsLvlVar1, WarnLvlName)
	SetLogLvlEnvVariable(testLogsLvlVar1)
	assert.Equal(t, WarnLvl, LogLvlIntValue())
	assert.Equal(t, WarnLvlName, LogLvlName())

	_ = os.Setenv(testLogsLvlVar2, InfoLvlName)
	SetLogLvlEnvVariable(testLogsLvlVar2)
	assert.Equal(t, InfoLvl, LogLvlIntValue())
	assert.NotEqual(t, WarnLvl, LogLvlIntValue())

	SetLogLvl(DebugLvlName)
	_ = os.Unsetenv(testLogsLvlVar1)
	_ = os.Unsetenv(testLogsLvlVar2)
}

func TestPackageLoggerGetLogLvlIntValue(t *testing.T) {
	assert.Equal(t, DebugLvl, LogLvlIntValue())

	SetLogLvl(ErrorLvlName)
	assert.Equal(t, ErrorLvl, LogLvlIntValue())

	testLogsLvlVar1 := "MY_INSTANCE_LOGS_LVL"
	_ = os.Setenv(testLogsLvlVar1, InfoLvlName)
	SetLogLvlEnvVariable(testLogsLvlVar1)
	assert.Equal(t, InfoLvl, LogLvlIntValue())

	SetLogLvl(DebugLvlName)
	_ = os.Unsetenv(testLogsLvlVar1)
}

func TestPackageLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing DEBUG log"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	Info(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.Contains(t, buf.String(), testLog)
}

func TestPackageLogger_InfoNotLogging(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing DEBUG log"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	logger := NewLogger()
	logger.SetLogLvl(ErrorLvlName)
	logger.Info(testLog)
	logger.Warn(testLog)
	logger.Debug(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.NotContainsf(t, buf.String(), testLog, "error-msg")
}

func TestPackageLogger_Debug(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing INFO log"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	Debug(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.Contains(t, buf.String(), testLog)
}

func TestPackageLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing WARN log"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	Warn(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.Contains(t, buf.String(), testLog)
}

func TestPackageLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing ERROR log"
	originalStdErr := os.Stderr
	r, w, _ := os.Pipe()

	os.Stderr = w
	Error(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stderr = originalStdErr
	assert.Contains(t, buf.String(), testLog)
}
