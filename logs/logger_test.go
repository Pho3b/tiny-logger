package logs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/log_level"
	"github.com/pho3b/tiny-logger/shared"
	"github.com/pho3b/tiny-logger/test"
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

	output = test.CaptureOutput(func() { logger.Color(colors.Magenta, testLog) })
	assert.Contains(t, output, colors.Magenta.String()+testLog)

	output = test.CaptureOutput(func() { logger.Color(colors.Cyan, testLog) })
	assert.Contains(t, output, colors.Cyan.String()+testLog+colors.Reset.String())

	output = test.CaptureOutput(func() { logger.Color(colors.Gray, testLog) })
	assert.Contains(t, output, colors.Gray.String()+testLog+colors.Reset.String())

	output = test.CaptureOutput(func() { logger.Color(colors.Blue, testLog) })
	assert.Contains(t, output, colors.Blue.String()+testLog+colors.Reset.String())

	os.Stdout = originalStdOut
}

func TestLogger_ShowLogLevel(t *testing.T) {
	logger := NewLogger().AddDateTime(false).
		ShowLogLevel(true).
		EnableColors(false)

	output := test.CaptureOutput(func() {
		logger.Info("my testing log")
	})
	assert.Contains(t, output, "INFO: my testing log")

	logger.ShowLogLevel(false)
	output = test.CaptureOutput(func() {
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

	outMsg := test.CaptureOutput(func() { logger.Debug("testing log") })

	matches := re.FindStringSubmatch(outMsg)
	assert.Equal(t, "DEBUG", matches[1])
	assert.NotNil(t, matches[2])
	assert.Equal(t, "testing log", matches[3])

	outMsg = test.CaptureOutput(func() {
		logger.Warn("testing log")
	})

	matches = re.FindStringSubmatch(outMsg)
	assert.Equal(t, "WARN", matches[1])
	assert.NotNil(t, matches[2])
	assert.Equal(t, "testing log", matches[3])

	outMsg = test.CaptureErrorOutput(func() {
		logger.Error("testing log")
	})

	matches = re.FindStringSubmatch(outMsg)
	assert.Equal(t, "ERROR", matches[1])
	assert.NotNil(t, matches[2])
	assert.Equal(t, "testing log", matches[3])

	outMsg = test.CaptureOutput(func() {
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
	file := createMockOutFile(testFileName)

	// Clean up any existing test file
	defer os.Remove(testFileName)

	// Test setting a new log file
	result := logger.SetLogFile(file)
	assert.NotNil(t, result)
	assert.IsType(t, &Logger{}, result)
	assert.NotNil(t, logger.outFile)
	assert.NotNil(t, logger.GetLogFile())

	// Verify the file was created
	_, err := os.Stat(testFileName)
	assert.NoError(t, err)

	// Test that logging to a file works
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
	file := createMockOutFile(testFileName)

	// Create a file with initial content
	initialContent := "initial content\n"
	err := os.WriteFile(testFileName, []byte(initialContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(testFileName)

	// Set the log file (should append, not overwrite)
	logger.SetLogFile(file)
	logger.Info("appended message")
	logger.CloseLogFile()

	// Verify content was appended
	content, err := os.ReadFile(testFileName)
	assert.NoError(t, err)
	contentStr := string(content)
	assert.Contains(t, contentStr, "initial content")
	assert.Contains(t, contentStr, "appended message")
}

func TestLogger_SetLogFile_Nil(t *testing.T) {
	logger := NewLogger()
	warnOut := test.CaptureOutput(func() { logger.SetLogFile(nil) })
	assert.Equal(t, "WARN: the given log file is nil, skipping logs redirection\n", warnOut)
	assert.Nil(t, logger.outFile)
	assert.Nil(t, logger.GetLogFile())
}

func TestLogger_CloseLogFile(t *testing.T) {
	logger := NewLogger()
	testFileName := "close_test_log_file.txt"
	file := createMockOutFile(testFileName)

	// Clean up
	os.Remove(testFileName)
	defer os.Remove(testFileName)

	// Set a log file first
	logger.SetLogFile(file)
	assert.NotNil(t, logger.GetLogFile())

	// Close the log file
	logger.CloseLogFile()
	assert.Nil(t, logger.GetLogFile())
	assert.Nil(t, logger.outFile)
}

func TestLogger_CloseLogFile_NoFileSet(t *testing.T) {
	logger := NewLogger()

	// Test closing when no file is set - should log warning and not crash
	output := test.CaptureOutput(func() {
		logger.CloseLogFile()
	})

	assert.Contains(t, output, "no log file opened, skipping close")
	assert.Nil(t, logger.GetLogFile())
}

func TestLogger_GetLogFile(t *testing.T) {
	logger := NewLogger()

	// Initially should return nil
	assert.Nil(t, logger.GetLogFile())

	// After setting a file should return a file pointer
	testFileName := "get_log_file_test.txt"
	file := createMockOutFile(testFileName)
	defer os.Remove(testFileName)

	logger.SetLogFile(file)
	lFile := logger.GetLogFile()
	assert.NotNil(t, lFile)
	assert.IsType(t, &os.File{}, lFile)

	logger.CloseLogFile()
	assert.Nil(t, logger.GetLogFile())
}

func TestLogger_LogsRedirectedToFile(t *testing.T) {
	logger := NewLogger()
	testFileName := "redirect_test_log_file.txt"
	file := createMockOutFile(testFileName)
	defer os.Remove(testFileName)

	// Set the log file
	logger.SetLogFile(file)

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

func TestLogger_NestedStructParameterCorrectLogging(t *testing.T) {
	type Address struct {
		Street string
		City   string
		Zip    int
	}

	type Contact struct {
		Email string
		Phone string
		ad    Address
	}

	type User struct {
		ID      int
		Name    string
		Address Address // Named field
		Contact         // Embedded (Promoted) field
	}

	user := User{
		ID:   1,
		Name: "Alice",
		Address: Address{
			Street: "123 Go Lane",
			City:   "Tech City",
			Zip:    90210,
		},
		Contact: Contact{
			Email: "alice@example.com",
			Phone: "555-0199",
			ad: Address{
				Street: "123 Go Lane",
				City:   "Tech City",
				Zip:    90210,
			},
		},
	}

	logger := NewLogger().SetEncoder(shared.JsonEncoderType)
	outMsg := test.CaptureOutput(func() { logger.Debug(user) })
	assert.Contains(t,
		outMsg,
		"{1 Alice {123 Go Lane Tech City 90210} {alice@example.com 555-0199 {123 Go Lane Tech City 90210}}}",
	)
}

func createMockOutFile(fileName string) *os.File {
	file, err := os.OpenFile(fmt.Sprintf("./%s", fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		println("ERROR: cannot open out file", err)
		return nil
	}

	return file
}
