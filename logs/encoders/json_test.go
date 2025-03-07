package encoders

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

// decodeLogEntry decodes a JSON-encoded log entry into a jsonLogEntry struct.
func decodeLogEntry(t *testing.T, logOutput string) jsonLogEntry {
	var entry jsonLogEntry
	err := json.Unmarshal([]byte(logOutput), &entry)
	assert.NoError(t, err)
	return entry
}

func TestJSONEncoder_LogDebug(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true}

	output := captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)
}

func TestJSONEncoder_LogInfo(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true}

	output := captureOutput(func() {
		encoder.LogInfo(loggerConfig, "Test info message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "INFO", entry.Level)
	assert.Equal(t, "Test info message", entry.Message)
}

func TestJSONEncoder_LogInfoWithExtras(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true}

	output := captureOutput(func() {
		encoder.LogInfo(loggerConfig, "Test info message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "INFO", entry.Level)
	assert.Equal(t, "Test info message", entry.Message)
	assert.IsType(t, make(map[string]interface{}), entry.Extras)

	output = captureOutput(func() {
		encoder.LogInfo(loggerConfig, "Test info message with extras", "Location", "Italy", "Weather", "sunny", "Mood")
	})

	entry = decodeLogEntry(t, output)
	assert.Equal(t, "Test info message with extras", entry.Message)
	assert.IsType(t, make(map[string]interface{}), entry.Extras)
	assert.Equal(t, "Italy", entry.Extras["Location"])
	assert.Equal(t, "sunny", entry.Extras["Weather"])
	assert.Equal(t, nil, entry.Extras["Mood"])
}

func TestJSONEncoder_LogWarn(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true}

	output := captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test warning message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "WARN", entry.Level)
	assert.Equal(t, "Test warning message", entry.Message)
}

func TestJSONEncoder_LogError(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true}

	output := captureErrorOutput(func() {
		encoder.LogError(loggerConfig, "Test error message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "ERROR", entry.Level)
	assert.Equal(t, "Test error message", entry.Message)
}

func TestJSONEncoder_LogFatalError(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true}

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

func TestJSONEncoder_DateTime(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{TimeEnabled: true, ColorsEnabled: true}
	output := captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test msg")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Date)
	assert.Empty(t, entry.DateTime)
	assert.NotEmpty(t, entry.Time)

	loggerConfig = &LoggerConfigMock{DateEnabled: true, ColorsEnabled: true}
	output = captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test msg")
	})

	entry = decodeLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Time)
	assert.Empty(t, entry.DateTime)
	assert.NotEmpty(t, entry.Date)

	loggerConfig = &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true}
	output = captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test msg")
	})

	entry = decodeLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Time)
	assert.Empty(t, entry.Date)
	assert.NotEmpty(t, entry.DateTime)
}

func TestJSONEncoder_ExtraMessages(t *testing.T) {
	resMap := buildExtraMessages("user", "alice", "ip", "192.168.1.1")
	assert.NotNil(t, resMap)
	assert.NotNil(t, resMap["ip"])
	assert.Len(t, resMap, 2)

	resMap = buildExtraMessages("user", "alice", "ip")
	assert.Nil(t, resMap["ip"])
	assert.Len(t, resMap, 2)

	resMap = buildExtraMessages("user", "alice", "ip", "192.168.1.1", "city", "paris", "pass")
	assert.Len(t, resMap, 4)
	assert.Equal(t, "alice", resMap["user"])
	assert.Equal(t, "192.168.1.1", resMap["ip"])
	assert.Equal(t, "paris", resMap["city"])
	assert.Equal(t, nil, resMap["pass"])
}
