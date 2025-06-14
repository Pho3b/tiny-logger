package encoders

import (
	"bytes"
	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestLogDebug(t *testing.T) {
	encoder := NewDefaultEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	assert.Contains(t, output, "DEBUG")
	assert.Contains(t, output, "Test debug message")
}

func TestLogInfo(t *testing.T) {
	encoder := NewDefaultEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogInfo(loggerConfig, "Test info message")
	})

	assert.Contains(t, output, "INFO")
	assert.Contains(t, output, "Test info message")
}

func TestLogWarn(t *testing.T) {
	encoder := NewDefaultEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test warning message")
	})

	assert.Contains(t, output, "WARN")
	assert.Contains(t, output, "Test warning message")
}

func TestLogError(t *testing.T) {
	encoder := NewDefaultEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureErrorOutput(func() {
		encoder.LogError(loggerConfig, "Test error message")
	})

	assert.Contains(t, output, "ERROR")
	assert.Contains(t, output, "Test error message")
}

func TestLogFatalError(t *testing.T) {
	encoder := NewDefaultEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	if os.Getenv("BE_CRASHER") == "1" {
		encoder.LogFatalError(loggerConfig, "Test fatal error message")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestLogFatalError")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	exitError, ok := err.(*exec.ExitError)
	assert.True(t, ok && !exitError.Success())
}

func TestFormatDateTimeString(t *testing.T) {
	var b bytes.Buffer
	encoder := NewDefaultEncoder()

	b = encoder.formatDateTimeString("dateTest", "timeTest", "")
	assert.Contains(t, b.String(), "[")
	assert.Contains(t, b.String(), "]")
	assert.Contains(t, b.String(), " ")

	b = encoder.formatDateTimeString("", "timeTest", "")
	assert.Contains(t, b.String(), "[")
	assert.Contains(t, b.String(), "]")

	b = encoder.formatDateTimeString("dateTest", "", "")
	assert.Contains(t, b.String(), "[")
	assert.Contains(t, b.String(), "]")

	b = encoder.formatDateTimeString("", "", "")
	assert.NotContains(t, b.String(), "[")
	assert.NotContains(t, b.String(), "]")
	assert.NotContains(t, b.String(), " ")
}

func TestShowLogLevel(t *testing.T) {
	encoder := NewDefaultEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test my-test message")
	})

	assert.Contains(t, output, "DEBUG")
	assert.Contains(t, output, "Test my-test message")

	loggerConfig = &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ShowLogLevel: false}

	output = captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test my-test message")
	})

	assert.NotContains(t, output, "DEBUG:")
	assert.Contains(t, output, "Test my-test message")
}

func TestCheckColorsInTheOutput(t *testing.T) {
	encoder := NewDefaultEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: false, TimeEnabled: false, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() { encoder.LogDebug(loggerConfig, "Test msg") })
	assert.Contains(t, output, colors.Gray.String())

	output = captureOutput(func() { encoder.LogInfo(loggerConfig, "Test msg") })
	assert.Contains(t, output, colors.Cyan.String())

	output = captureOutput(func() { encoder.LogWarn(loggerConfig, "Test msg") })
	assert.Contains(t, output, colors.Yellow.String())

	output = captureErrorOutput(func() { encoder.LogError(loggerConfig, "Test msg") })
	assert.Contains(t, output, colors.Red.String())
}

func TestDefaultEncoder_Color(t *testing.T) {
	var output string
	testLog := "my testing log"
	originalStdOut := os.Stdout
	encoder := NewDefaultEncoder()
	lConfig := LoggerConfigMock{
		DateEnabled:   false,
		TimeEnabled:   false,
		ColorsEnabled: false,
		ShowLogLevel:  false,
	}

	output = captureOutput(func() { encoder.Color(&lConfig, colors.Magenta, testLog) })
	assert.Contains(t, output, colors.Magenta.String()+testLog)

	output = captureOutput(func() { encoder.Color(&lConfig, colors.Cyan, testLog) })
	assert.Contains(t, output, colors.Cyan.String()+testLog+colors.Reset.String())

	output = captureOutput(func() { encoder.Color(&lConfig, colors.Gray, testLog) })
	assert.Contains(t, output, colors.Gray.String()+testLog+colors.Reset.String())

	output = captureOutput(func() { encoder.Color(&lConfig, colors.Blue, testLog) })
	assert.Contains(t, output, colors.Blue.String()+testLog+colors.Reset.String())

	os.Stdout = originalStdOut
}
