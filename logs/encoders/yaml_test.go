package encoders

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"testing"
)

func decodeYamlLogEntry(t *testing.T, logOutput string) jsonLogEntry {
	var entry jsonLogEntry
	err := yaml.Unmarshal([]byte(logOutput), &entry)
	assert.NoError(t, err)
	return entry
}

func TestYAMLEncoder_LogDebug(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)
}

func TestYAMLEncoder_LogInfo(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogInfo(loggerConfig, "Test info message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "INFO", entry.Level)
	assert.Equal(t, "Test info message", entry.Message)
}

func TestYAMLEncoder_LogWarn(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test warning message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "WARN", entry.Level)
	assert.Equal(t, "Test warning message", entry.Message)
}

func TestYAMLEncoder_LogError(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureErrorOutput(func() {
		encoder.LogError(loggerConfig, "Test error message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "ERROR", entry.Level)
	assert.Equal(t, "Test error message", entry.Message)
}

func TestYAMLEncoder_LogFatalError(t *testing.T) {
	encoder := NewYAMLEncoder()
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

func TestYAMLEncoder_ShowLogLevelLt(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)

	loggerConfig = &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: false}

	output = captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry = decodeYamlLogEntry(t, output)
	assert.Equal(t, "", entry.Level)
}
