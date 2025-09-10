package logs

import (
	"bytes"
	"io"
	"os"
	"os/exec"
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

func TestAreAllNil(t *testing.T) {
	logger := NewLogger()

	// Test with all nil arguments
	result := logger.areAllNil(nil, nil, nil)
	assert.True(t, result)

	// Test with no arguments
	result = logger.areAllNil()
	assert.True(t, result)

	// Test with some nil and some non-nil arguments
	result = logger.areAllNil(nil, "test", nil)
	assert.False(t, result)

	// Test with all non-nil arguments
	result = logger.areAllNil("test", 123, true)
	assert.False(t, result)
}

func TestLogger_SetLogFile(t *testing.T) {
	logger := NewLogger()
	testFileName := "test_log_file.txt"

	// Clean up any existing test file
	os.Remove(testFileName)
	defer os.Remove(testFileName)

	// Test setting a new log file
	result := logger.SetLogFile(testFileName)
	assert.NotNil(t, result)
	assert.IsType(t, &Logger{}, result)
	assert.NotNil(t, logger.outFile)
	assert.NotNil(t, logger.GetLogFile())

	// Verify file was created
	_, err := os.Stat(testFileName)
	assert.NoError(t, err)

	// Test that logging to file works
	logger.Info("test log message")

	// Read file content
	content, err := os.ReadFile(testFileName)
	assert.NoError(t, err)
	assert.Contains(t, string(content), "test log message")

	// Clean up
	logger.CloseLogFile()
}

func TestLogger_SetLogFile_ExistingFile(t *testing.T) {
	logger := NewLogger()
	testFileName := "existing_test_log_file.txt"

	// Create file with initial content
	initialContent := "initial content\n"
	err := os.WriteFile(testFileName, []byte(initialContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(testFileName)

	// Set log file (should append, not overwrite)
	logger.SetLogFile(testFileName)
	logger.Info("appended message")
	logger.CloseLogFile()

	// Verify content was appended
	content, err := os.ReadFile(testFileName)
	assert.NoError(t, err)
	contentStr := string(content)
	assert.Contains(t, contentStr, "initial content")
	assert.Contains(t, contentStr, "appended message")
}

func TestLogger_SetLogFile_InvalidPath(t *testing.T) {
	// This test verifies the logger handles fatal errors on invalid file paths
	if os.Getenv("BE_CRASHER") == "1" {
		logger := NewLogger()
		// Try to create file in non-existent directory
		logger.SetLogFile("/non/existent/directory/test.log")
		return
	}

	// Run the test in a subprocess to catch the fatal error
	cmd := exec.Command(os.Args[0], "-test.run=TestLogger_SetLogFile_InvalidPath")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	exitError, ok := err.(*exec.ExitError)
	assert.True(t, ok && !exitError.Success())
}

func TestLogger_CloseLogFile(t *testing.T) {
	logger := NewLogger()
	testFileName := "close_test_log_file.txt"

	// Clean up
	os.Remove(testFileName)
	defer os.Remove(testFileName)

	// Set a log file first
	logger.SetLogFile(testFileName)
	assert.NotNil(t, logger.GetLogFile())

	// Close the log file
	logger.CloseLogFile()
	assert.Nil(t, logger.GetLogFile())
	assert.Nil(t, logger.outFile)
}

func TestLogger_CloseLogFile_NoFileSet(t *testing.T) {
	logger := NewLogger()

	// Test closing when no file is set - should log warning and not crash
	output := captureOutput(func() {
		logger.CloseLogFile()
	})

	assert.Contains(t, output, "no log file opened, skipping close")
	assert.Nil(t, logger.GetLogFile())
}

func TestLogger_GetLogFile(t *testing.T) {
	logger := NewLogger()

	// Initially should return nil
	assert.Nil(t, logger.GetLogFile())

	// After setting file should return file pointer
	testFileName := "get_log_file_test.txt"
	os.Remove(testFileName)
	defer os.Remove(testFileName)

	logger.SetLogFile(testFileName)
	file := logger.GetLogFile()
	assert.NotNil(t, file)
	assert.IsType(t, &os.File{}, file)

	logger.CloseLogFile()
	assert.Nil(t, logger.GetLogFile())
}

func TestLogger_LogsRedirectedToFile(t *testing.T) {
	logger := NewLogger()
	testFileName := "redirect_test_log_file.txt"

	os.Remove(testFileName)
	defer os.Remove(testFileName)

	// Set log file
	logger.SetLogFile(testFileName)

	// Log various levels
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")

	logger.CloseLogFile()

	// Verify all messages were written to file
	content, err := os.ReadFile(testFileName)
	assert.NoError(t, err)
	contentStr := string(content)

	assert.Contains(t, contentStr, "debug message")
	assert.Contains(t, contentStr, "info message")
	assert.Contains(t, contentStr, "warn message")
	assert.Contains(t, contentStr, "error message")
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
