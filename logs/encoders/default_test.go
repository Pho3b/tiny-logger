package encoders

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/Pho3b/tiny-logger/internal/services"
	"github.com/Pho3b/tiny-logger/logs/colors"
	ll "github.com/Pho3b/tiny-logger/logs/log_level"
	s "github.com/Pho3b/tiny-logger/shared"
	"github.com/Pho3b/tiny-logger/test"
	"github.com/stretchr/testify/assert"
)

func TestLogDebug(t *testing.T) {
	encoder := NewDefaultEncoder(services.NewPrinter(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.DebugLvlName, s.StdOutput, "Test debug message")
	})

	assert.Contains(t, output, "DEBUG")
	assert.Contains(t, output, "Test debug message")
}

func TestLogInfo(t *testing.T) {
	encoder := NewDefaultEncoder(services.NewPrinter(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.InfoLvlName, s.StdOutput, "Test info message")
	})

	assert.Contains(t, output, "INFO")
	assert.Contains(t, output, "Test info message")
}

func TestLogWarn(t *testing.T) {
	encoder := NewDefaultEncoder(services.NewPrinter(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.WarnLvlName, s.StdOutput, "Test warning message")
	})

	assert.Contains(t, output, "WARN")
	assert.Contains(t, output, "Test warning message")
}

func TestLogError(t *testing.T) {
	encoder := NewDefaultEncoder(services.NewPrinter(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureErrorOutput(func() {
		encoder.Log(loggerConfig, ll.ErrorLvlName, s.StdErrOutput, "Test error message")
	})

	assert.Contains(t, output, "ERROR")
	assert.Contains(t, output, "Test error message")
}

func TestLogFatalError(t *testing.T) {
	encoder := NewDefaultEncoder(services.NewPrinter(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	if os.Getenv("BE_CRASHER") == "1" {
		encoder.Log(loggerConfig, ll.FatalErrorLvlName, s.StdOutput, "Test fatal error message")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestLogFatalError")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	exitError, ok := err.(*exec.ExitError)
	assert.False(t, ok && !exitError.Success())
}

func TestFormatDateTimeString(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	encoder := NewDefaultEncoder(services.NewPrinter(), services.GetDateTimePrinter())

	encoder.addFormattedDateTime(b, "dateTest", "timeTest", "")
	assert.Contains(t, b.String(), "[")
	assert.Contains(t, b.String(), "]")
	assert.Contains(t, b.String(), " ")

	b.Reset()
	encoder.addFormattedDateTime(b, "", "timeTest", "")
	assert.Contains(t, b.String(), "[")
	assert.Contains(t, b.String(), "]")

	b.Reset()
	encoder.addFormattedDateTime(b, "dateTest", "", "")
	assert.Contains(t, b.String(), "[")
	assert.Contains(t, b.String(), "]")

	b.Reset()
	encoder.addFormattedDateTime(b, "", "", "")
	assert.NotContains(t, b.String(), "[")
	assert.NotContains(t, b.String(), "]")
	assert.NotContains(t, b.String(), " ")
}

func TestShowLogLevel(t *testing.T) {
	encoder := NewDefaultEncoder(services.NewPrinter(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.DebugLvlName, s.StdOutput, "Test my-test message")
	})

	assert.Contains(t, output, "DEBUG")
	assert.Contains(t, output, "Test my-test message")

	loggerConfig = &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ShowLogLevel: false}

	output = test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.DebugLvlName, s.StdOutput, "Test my-test message")
	})

	assert.NotContains(t, output, "DEBUG:")
	assert.Contains(t, output, "Test my-test message")
}

func TestCheckColorsInTheOutput(t *testing.T) {
	encoder := NewDefaultEncoder(services.NewPrinter(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: false, TimeEnabled: false, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() { encoder.Log(loggerConfig, ll.DebugLvlName, s.StdOutput, "Test msg") })
	assert.Contains(t, output, colors.Gray.String())

	output = test.CaptureOutput(func() { encoder.Log(loggerConfig, ll.InfoLvlName, s.StdOutput, "Test my-test message") })
	assert.Contains(t, output, colors.Cyan.String())

	output = test.CaptureOutput(func() { encoder.Log(loggerConfig, ll.WarnLvlName, s.StdOutput, "Test my-test message") })
	assert.Contains(t, output, colors.Yellow.String())

	output = test.CaptureErrorOutput(func() { encoder.Log(loggerConfig, ll.ErrorLvlName, s.StdErrOutput, "Test my-test message") })
	assert.Contains(t, output, colors.Red.String())
}

func TestDefaultEncoder_Color(t *testing.T) {
	var output string
	testLog := "my testing Log"
	originalStdOut := os.Stdout
	encoder := NewDefaultEncoder(services.NewPrinter(), services.GetDateTimePrinter())
	lConfig := test.LoggerConfigMock{
		DateEnabled:   false,
		TimeEnabled:   false,
		ColorsEnabled: false,
		ShowLogLevel:  false,
	}

	output = test.CaptureOutput(func() { encoder.Color(&lConfig, colors.Magenta, testLog) })
	assert.Contains(t, output, colors.Magenta.String()+testLog)

	output = test.CaptureOutput(func() { encoder.Color(&lConfig, colors.Cyan, testLog) })
	assert.Contains(t, output, colors.Cyan.String()+testLog+colors.Reset.String())

	output = test.CaptureOutput(func() { encoder.Color(&lConfig, colors.Gray, testLog) })
	assert.Contains(t, output, colors.Gray.String()+testLog+colors.Reset.String())

	output = test.CaptureOutput(func() { encoder.Color(&lConfig, colors.Blue, testLog) })
	assert.Contains(t, output, colors.Blue.String()+testLog+colors.Reset.String())

	os.Stdout = originalStdOut
}
