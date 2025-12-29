package services

import (
	"bytes"
	"os"
	"testing"

	"github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
	"github.com/pho3b/tiny-logger/test"

	c "github.com/pho3b/tiny-logger/logs/colors"
	"github.com/stretchr/testify/assert"
)

func TestPrinter_PrintLog_Stdout(t *testing.T) {
	p := NewPrinter()
	buf := bytes.NewBufferString("hello stdout")

	output := test.CaptureOutput(func() {
		p.PrintLog(s.StdOutput, buf, nil)
	})

	assert.Equal(t, "hello stdout", output)
}

func TestPrinter_PrintLog_Stderr(t *testing.T) {
	p := NewPrinter()
	buf := bytes.NewBufferString("hello stderr")

	output := test.CaptureErrorOutput(func() {
		p.PrintLog(s.StdErrOutput, buf, nil)
	})

	assert.Equal(t, "hello stderr", output)
}

func TestPrinter_PrintLog_FileOutput_WritesToFile(t *testing.T) {
	p := NewPrinter()
	buf := bytes.NewBufferString("file log")

	tmpFile, err := os.CreateTemp("", "printer-log-*")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	p.PrintLog(s.FileOutput, buf, tmpFile)

	content, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, "file log", string(content))
}

func TestPrinter_PrintLog_FileOutput_NilFile_WritesErrorToStderr(t *testing.T) {
	p := NewPrinter()
	buf := bytes.NewBufferString("ignored")

	output := test.CaptureErrorOutput(func() {
		p.PrintLog(s.FileOutput, buf, nil)
	})

	assert.Contains(t, output, "tiny-logger-err: given out file is nil")
}

func TestPrinter_PrintLog_WriteError_LogsToStderr(t *testing.T) {
	p := NewPrinter()
	buf := bytes.NewBufferString("data")

	// Create an invalid *os.File to trigger a write error.
	// The fd value here is intentionally bogus.
	badFile := os.NewFile(^uintptr(0), "bad")

	output := test.CaptureErrorOutput(func() {
		p.PrintLog(s.FileOutput, buf, badFile)
	})

	assert.Contains(t, output, "tiny-logger-err:")
}

func TestPrintColors_EnableColorsTrue(t *testing.T) {
	printer := Printer{}

	result := printer.RetrieveColorsFromLogLevel(true, log_level.FatalErrorLvl)
	assert.Equal(t, c.Magenta, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")

	result = printer.RetrieveColorsFromLogLevel(true, log_level.ErrorLvl)
	assert.Equal(t, c.Red, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")

	result = printer.RetrieveColorsFromLogLevel(true, log_level.WarnLvl)
	assert.Equal(t, c.Yellow, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")

	result = printer.RetrieveColorsFromLogLevel(true, log_level.InfoLvl)
	assert.Equal(t, c.Cyan, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")

	result = printer.RetrieveColorsFromLogLevel(true, log_level.DebugLvl)
	assert.Equal(t, c.Gray, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")
}

func TestPrintColors_EnableColorsFalse(t *testing.T) {
	printer := Printer{}

	result := printer.RetrieveColorsFromLogLevel(false, log_level.DebugLvl)
	assert.Equal(t, c.Color(""), result[0], "Expected first element to be an empty string when colors are disabled")
	assert.Equal(t, c.Color(""), result[1], "Expected second element to be an empty string when colors are disabled")

	result = printer.RetrieveColorsFromLogLevel(false, log_level.InfoLvl)
	assert.Equal(t, c.Color(""), result[0], "Expected first element to be an empty string when colors are disabled")
	assert.Equal(t, c.Color(""), result[1], "Expected second element to be an empty string when colors are disabled")
}
