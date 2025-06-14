package logs

import (
	"bytes"
	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/log_level"
	"github.com/pho3b/tiny-logger/shared"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"sync"
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
	testLog := "my testing INFO log"
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
	testLog := "my testing DEBUG log"
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

func TestLogger_AddDateTime(t *testing.T) {
	logger := NewLogger()
	logger.AddDate(true)
	assert.True(t, logger.dateEnabled)

	logger.AddDate(false)
	assert.False(t, logger.dateEnabled)

	logger.AddTime(true)
	assert.True(t, logger.timeEnabled)

	logger.AddTime(false)
	assert.False(t, logger.timeEnabled)

	logger.AddDateTime(true)
	assert.True(t, logger.dateEnabled)
	assert.True(t, logger.timeEnabled)

	logger.AddDateTime(false)
	assert.False(t, logger.dateEnabled)
	assert.False(t, logger.timeEnabled)
}

func TestLogger_EnableColors(t *testing.T) {
	logger := NewLogger()
	logger.EnableColors(true)
	assert.True(t, logger.GetColorsEnabled())

	logger.EnableColors(false)
	assert.False(t, logger.GetColorsEnabled())
}

func TestLoggerConcurrency(t *testing.T) {
	var wg sync.WaitGroup

	logger := NewLogger()
	numGoroutines := 100
	numMessages := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go writeDummyLogsWithNewLogger(numMessages, &wg, logger)
	}

	wg.Wait()
}

func writeDummyLogsWithNewLogger(logsNumber int, wg *sync.WaitGroup, logger *Logger) {
	defer wg.Done()

	for i := 0; i < logsNumber; i++ {
		logger.Debug("This is a test message", ";", 2)
	}
}

func TestLogger_Color(t *testing.T) {
	var buf bytes.Buffer
	testLog := "my testing DEBUG log"
	originalStdOut := os.Stdout
	logger := NewLogger()

	r, w, _ := os.Pipe()
	os.Stdout = w
	logger.Color(colors.Magenta, testLog)
	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	assert.Contains(t, buf.String(), colors.Magenta.String()+testLog)

	buf.Reset()
	r, w, _ = os.Pipe()
	os.Stdout = w
	logger.Color(colors.Cyan, testLog)
	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	assert.Contains(t, buf.String(), colors.Cyan.String()+testLog+colors.Reset.String())

	buf.Reset()
	r, w, _ = os.Pipe()
	os.Stdout = w
	logger.Color(colors.Cyan, testLog)
	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	assert.Contains(t, buf.String(), colors.Cyan.String()+testLog+colors.Reset.String())

	buf.Reset()
	r, w, _ = os.Pipe()
	os.Stdout = w
	logger.Color(colors.Blue, testLog)
	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	assert.Contains(t, buf.String(), colors.Blue.String()+testLog+colors.Reset.String())

	os.Stdout = originalStdOut
}

func TestLogger_ShowLogLevel(t *testing.T) {
	var buf bytes.Buffer
	originalStdOut := os.Stdout
	logger := NewLogger().AddDateTime(false).
		ShowLogLevel(true).
		EnableColors(false)

	r, w, _ := os.Pipe()
	os.Stdout = w
	logger.Info("my testing log")
	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	assert.Contains(t, buf.String(), "INFO: my testing log")

	logger.ShowLogLevel(false)

	buf.Reset()
	r, w, _ = os.Pipe()
	os.Stdout = w
	logger.Color(colors.Cyan, "my testing log")
	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	assert.NotContains(t, buf.String(), "INFO: my testing log")
	assert.Contains(t, buf.String(), "my testing log")

	os.Stdout = originalStdOut
}

func TestLogger_SetEncoder(t *testing.T) {
	l := NewLogger().SetEncoder(shared.DefaultEncoderType)
	assert.Equal(t, shared.DefaultEncoderType, l.encoder.GetType())
	assert.Equal(t, shared.DefaultEncoderType, l.GetEncoderType())

	l.SetEncoder(shared.JsonEncoderType)
	assert.Equal(t, shared.JsonEncoderType, l.encoder.GetType())
	assert.Equal(t, shared.JsonEncoderType, l.GetEncoderType())

	l.SetEncoder(shared.YamlEncoderType)
	assert.Equal(t, shared.YamlEncoderType, l.GetEncoderType())
	assert.Equal(t, shared.YamlEncoderType, l.encoder.GetType())
}
