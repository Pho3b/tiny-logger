package logs

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/log_level"
	"github.com/pho3b/tiny-logger/shared"
	"github.com/stretchr/testify/assert"
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

func TestLogger_FatalError(t *testing.T) {
	var buf bytes.Buffer
	var testLog any
	originalStdErr := os.Stderr
	r, w, _ := os.Pipe()

	os.Stderr = w
	NewLogger().FatalError(testLog)

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stderr = originalStdErr
	assert.Equal(t, buf.String(), "")
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
	var output string
	testLog := "my testing DEBUG log"
	originalStdOut := os.Stdout
	logger := NewLogger()

	output = captureOutput(func() { logger.Color(colors.Magenta, testLog) })
	assert.Contains(t, output, colors.Magenta.String()+testLog)

	output = captureOutput(func() { logger.Color(colors.Cyan, testLog) })
	assert.Contains(t, output, colors.Cyan.String()+testLog+colors.Reset.String())

	output = captureOutput(func() { logger.Color(colors.Gray, testLog) })
	assert.Contains(t, output, colors.Gray.String()+testLog+colors.Reset.String())

	output = captureOutput(func() { logger.Color(colors.Blue, testLog) })
	assert.Contains(t, output, colors.Blue.String()+testLog+colors.Reset.String())

	os.Stdout = originalStdOut
}

func TestLogger_ShowLogLevel(t *testing.T) {
	logger := NewLogger().AddDateTime(false).
		ShowLogLevel(true).
		EnableColors(false)

	output := captureOutput(func() {
		logger.Info("my testing log")
	})
	assert.Contains(t, output, "INFO: my testing log")

	logger.ShowLogLevel(false)
	output = captureOutput(func() {
		logger.Info("my testing log")
	})
	assert.NotContains(t, output, "INFO: my testing log")
	assert.Contains(t, output, "my testing log")
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

func TestLogger_CorrectLogsFormattingDefaultEncoder(t *testing.T) {
	logger := NewLogger().SetEncoder(shared.DefaultEncoderType).AddDateTime(true).ShowLogLevel(true)
	re := regexp.MustCompile(`^([A-Z]+) \[(\d{2}\/\d{2}\/\d{4} \d{2}:\d{2}:\d{2})\]: (.+)\n?$`)

	outMsg := captureOutput(func() { logger.Debug("testing log") })

	matches := re.FindStringSubmatch(outMsg)
	assert.Equal(t, "DEBUG", matches[1])
	assert.NotNil(t, matches[2])
	assert.Equal(t, "testing log", matches[3])

	outMsg = captureOutput(func() {
		logger.Warn("testing log")
	})

	matches = re.FindStringSubmatch(outMsg)
	assert.Equal(t, "WARN", matches[1])
	assert.NotNil(t, matches[2])
	assert.Equal(t, "testing log", matches[3])

	outMsg = captureErrorOutput(func() {
		logger.Error("testing log")
	})

	matches = re.FindStringSubmatch(outMsg)
	assert.Equal(t, "ERROR", matches[1])
	assert.NotNil(t, matches[2])
	assert.Equal(t, "testing log", matches[3])

	outMsg = captureOutput(func() {
		logger.Info("testing log")
	})

	matches = re.FindStringSubmatch(outMsg)
	assert.Equal(t, "INFO", matches[1])
	assert.NotNil(t, matches[2])
	assert.Equal(t, "testing log", matches[3])
}

// captureOutput redirects os.Stdout to capture the output of the function f
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	defer r.Close()

	origStdout := os.Stdout
	os.Stdout = w

	f()
	w.Close()
	os.Stdout = origStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

// captureErrorOutput redirects os.Stderr to capture the output of the function f
func captureErrorOutput(f func()) string {
	r, w, _ := os.Pipe()
	defer r.Close()

	origStderr := os.Stderr
	os.Stderr = w

	f()
	w.Close()
	os.Stderr = origStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}
