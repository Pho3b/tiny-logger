package encoders

import (
	"encoding/json"
	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
	"time"
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
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)
}

func TestJSONEncoder_LogInfo(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogInfo(loggerConfig, "Test info message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "INFO", entry.Level)
	assert.Equal(t, "Test info message", entry.Message)
}

func TestJSONEncoder_LogInfoWithExtras(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

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
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test warning message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "WARN", entry.Level)
	assert.Equal(t, "Test warning message", entry.Message)
}

func TestJSONEncoder_LogError(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureErrorOutput(func() {
		encoder.LogError(loggerConfig, "Test error message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "ERROR", entry.Level)
	assert.Equal(t, "Test error message", entry.Message)
}

func TestJSONEncoder_LogFatalError(t *testing.T) {
	encoder := NewJSONEncoder()
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

func TestJSONEncoder_DateTime(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}
	output := captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test msg")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Date)
	assert.Empty(t, entry.DateTime)
	assert.NotEmpty(t, entry.Time)

	loggerConfig = &LoggerConfigMock{DateEnabled: true, ColorsEnabled: true, ShowLogLevel: true}
	output = captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test msg")
	})

	entry = decodeLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Time)
	assert.Empty(t, entry.DateTime)
	assert.NotEmpty(t, entry.Date)

	loggerConfig = &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}
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
	jsonEncoder := NewJSONEncoder()

	resMap := jsonEncoder.buildExtraMessages("user", "alice", "ip", "192.168.1.1")
	assert.NotNil(t, resMap)
	assert.NotNil(t, resMap["ip"])
	assert.Len(t, resMap, 2)

	resMap = jsonEncoder.buildExtraMessages("user", "alice", "ip")
	assert.Nil(t, resMap["ip"])
	assert.Len(t, resMap, 2)

	resMap = jsonEncoder.buildExtraMessages("user", "alice", "ip", "192.168.1.1", "city", "paris", "pass")
	assert.Len(t, resMap, 4)
	assert.Equal(t, "alice", resMap["user"])
	assert.Equal(t, "192.168.1.1", resMap["ip"])
	assert.Equal(t, "paris", resMap["city"])
	assert.Equal(t, nil, resMap["pass"])
}

func TestJSONEncoder_ShowLogLevelLt(t *testing.T) {
	encoder := NewJSONEncoder()
	loggerConfig := &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)

	loggerConfig = &LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: false}

	output = captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry = decodeLogEntry(t, output)
	assert.Equal(t, "", entry.Level)
}

func TestJSONEncoder_Color(t *testing.T) {
	var output string

	testLog := "my testing log"
	originalStdOut := os.Stdout
	encoder := NewJSONEncoder()
	lConfig := LoggerConfigMock{
		DateEnabled:   false,
		TimeEnabled:   false,
		ColorsEnabled: false,
		ShowLogLevel:  false,
	}

	output = captureOutput(func() { encoder.Color(&lConfig, colors.Magenta, testLog) })
	assert.Contains(t, output, colors.Magenta.String())
	assert.Contains(t, output, testLog)
	assert.NotContains(t, output, time.Now().Format("02/01/2006"))
	assert.Contains(t, output, colors.Reset.String())

	lConfig.DateEnabled = true
	output = captureOutput(func() { encoder.Color(&lConfig, colors.Cyan, testLog) })
	assert.Contains(t, output, colors.Cyan.String())
	assert.Contains(t, output, time.Now().Format("02/01/2006"))
	assert.Contains(t, output, testLog)
	assert.Contains(t, output, colors.Reset.String())

	output = captureOutput(func() { encoder.Color(&lConfig, colors.Gray, testLog) })
	assert.Contains(t, output, colors.Gray.String())
	assert.Contains(t, output, testLog)
	assert.Contains(t, output, colors.Reset.String())

	output = captureOutput(func() { encoder.Color(&lConfig, colors.Blue, testLog) })
	assert.Contains(t, output, colors.Blue.String())
	assert.Contains(t, output, testLog)
	assert.Contains(t, output, colors.Reset.String())

	os.Stdout = originalStdOut
}

func TestJSONEncoder_ValidJSONOutput(t *testing.T) {
	var jsonMsg string

	originalStdOut := os.Stdout
	testLog := "my testing log"
	jsonEncoder := NewJSONEncoder()
	lConfig := &LoggerConfigMock{
		DateEnabled:   false,
		TimeEnabled:   false,
		ColorsEnabled: false,
		ShowLogLevel:  false,
	}

	jsonMsg = captureOutput(func() { jsonEncoder.LogInfo(lConfig, testLog, "id", 3) })
	assert.NoError(t, json.Unmarshal([]byte(jsonMsg), &jsonLogEntry{}))

	jsonMsg = captureOutput(func() { jsonEncoder.LogInfo(lConfig, testLog, "id", 3, 34, []string{"test", "test2"}) })
	assert.NoError(t, json.Unmarshal([]byte(jsonMsg), &jsonLogEntry{}))

	jsonMsg = captureOutput(func() {
		jsonEncoder.LogInfo(lConfig, testLog, "id", 3, 34, []string{"test", "test2"}, []string{"k", "k2"}, 2.3, 'f', 'A')
	})
	assert.NoError(t, json.Unmarshal([]byte(jsonMsg), &jsonLogEntry{}))

	jsonMsg = "{{'test'}"
	assert.Error(t, json.Unmarshal([]byte(jsonMsg), &jsonLogEntry{}))

	jsonMsg = "{\"msg\"\"This is my Warn log\",\"extras\":{\"Test arg\":[\"efsdaf\",\"dfas\"],\"[k3 k2]\":3}}"
	assert.Error(t, json.Unmarshal([]byte(jsonMsg), &jsonLogEntry{}))

	os.Stdout = originalStdOut
}
