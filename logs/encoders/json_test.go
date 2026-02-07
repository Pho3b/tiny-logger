package encoders

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/Pho3b/tiny-logger/internal/services"
	"github.com/Pho3b/tiny-logger/logs/colors"
	ll "github.com/Pho3b/tiny-logger/logs/log_level"
	"github.com/Pho3b/tiny-logger/shared"
	"github.com/Pho3b/tiny-logger/test"
	"github.com/stretchr/testify/assert"
)

// decodeLogEntry decodes a JSON-encoded Log entry into a JsonLogEntry struct.
func decodeLogEntry(t *testing.T, logOutput string) shared.JsonLog {
	var entry shared.JsonLog
	err := json.Unmarshal([]byte(logOutput), &entry)
	assert.NoError(t, err)

	return entry
}

func TestJSONEncoder_LogDebug(t *testing.T) {
	encoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.DebugLvlName, shared.StdOutput, "Test debug message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)
}

func TestJSONEncoder_LogInfo(t *testing.T) {
	encoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.InfoLvlName, shared.StdOutput, "Test info message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "INFO", entry.Level)
	assert.Equal(t, "Test info message", entry.Message)
}

func TestJSONEncoder_LogInfoWithExtras(t *testing.T) {
	encoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.InfoLvlName, shared.StdOutput, "Test info message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "INFO", entry.Level)
	assert.Equal(t, "Test info message", entry.Message)
	assert.IsType(t, make(map[string]any), entry.Extras)

	output = test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.InfoLvlName, shared.StdOutput, "Test info message with extras", "Location", "Italy", "Weather", "sunny", "Mood")
	})

	entry = decodeLogEntry(t, output)
	assert.Equal(t, "Test info message with extras", entry.Message)
	assert.IsType(t, make(map[string]any), entry.Extras)
	assert.Equal(t, "Italy", entry.Extras["Location"])
	assert.Equal(t, "sunny", entry.Extras["Weather"])
	assert.Equal(t, nil, entry.Extras["Mood"])
}

func TestJSONEncoder_LogWarn(t *testing.T) {
	encoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.WarnLvlName, shared.StdOutput, "Test warning message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "WARN", entry.Level)
	assert.Equal(t, "Test warning message", entry.Message)
}

func TestJSONEncoder_LogError(t *testing.T) {
	encoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureErrorOutput(func() {
		encoder.Log(loggerConfig, ll.ErrorLvlName, shared.StdErrOutput, "Test error message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "ERROR", entry.Level)
	assert.Equal(t, "Test error message", entry.Message)
}

func TestJSONEncoder_LogFatalError(t *testing.T) {
	encoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	if os.Getenv("BE_CRASHER") == "1" {
		encoder.Log(loggerConfig, ll.FatalErrorLvlName, shared.StdErrOutput, "Test fatal error message")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestLogFatalError")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	exitError, ok := err.(*exec.ExitError)
	assert.False(t, ok && !exitError.Success())
}

func TestJSONEncoder_DateTime(t *testing.T) {
	encoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}
	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.WarnLvlName, shared.StdOutput, "Test msg")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Date)
	assert.Empty(t, entry.DateTime)
	assert.NotEmpty(t, entry.Time)

	loggerConfig = &test.LoggerConfigMock{DateEnabled: true, ColorsEnabled: true, ShowLogLevel: true}
	output = test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.WarnLvlName, shared.StdOutput, "Test msg")
	})

	entry = decodeLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Time)
	assert.Empty(t, entry.DateTime)
	assert.NotEmpty(t, entry.Date)

	loggerConfig = &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}
	output = test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.WarnLvlName, shared.StdOutput, "Test msg")
	})

	entry = decodeLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Time)
	assert.Empty(t, entry.Date)
	assert.NotEmpty(t, entry.DateTime)
}

func TestJSONEncoder_ExtraMessages(t *testing.T) {
	jsonEncoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	lConfig := &test.LoggerConfigMock{DateEnabled: false, TimeEnabled: false, ColorsEnabled: false, ShowLogLevel: false}

	output := test.CaptureOutput(func() {
		jsonEncoder.Log(lConfig, ll.InfoLvlName, shared.StdOutput, "test", "user", "alice", "ip", "192.168.1.1")
	})
	entry := decodeLogEntry(t, output)
	assert.NotNil(t, entry)
	assert.NotNil(t, entry.Extras["ip"])
	assert.Len(t, entry.Extras, 2)

	output = test.CaptureOutput(func() {
		jsonEncoder.Log(lConfig, ll.InfoLvlName, shared.StdOutput, "test", "user", "alice", "ip")
	})
	entry = decodeLogEntry(t, output)
	assert.Nil(t, entry.Extras["ip"])
	assert.Len(t, entry.Extras, 2)

	output = test.CaptureOutput(func() {
		jsonEncoder.Log(lConfig, ll.InfoLvlName, shared.StdOutput, "test", "user", "alice", "ip", "192.168.1.1", "city", "paris", "pass")
	})
	entry = decodeLogEntry(t, output)
	assert.Len(t, entry.Extras, 4)
	assert.Equal(t, "alice", entry.Extras["user"])
	assert.Equal(t, "192.168.1.1", entry.Extras["ip"])
	assert.Equal(t, "paris", entry.Extras["city"])
	assert.Equal(t, nil, entry.Extras["pass"])
}

func TestJSONEncoder_ShowLogLevelLt(t *testing.T) {
	encoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.DebugLvlName, shared.StdOutput, "Test debug message")
	})

	entry := decodeLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)

	loggerConfig = &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: false}

	output = test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.DebugLvlName, shared.StdOutput, "Test debug message")
	})

	entry = decodeLogEntry(t, output)
	assert.Equal(t, "", entry.Level)
}

func TestJSONEncoder_Color(t *testing.T) {
	var output string

	testLog := "my testing Log"
	originalStdOut := os.Stdout
	encoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	lConfig := test.LoggerConfigMock{
		DateEnabled:   false,
		TimeEnabled:   false,
		ColorsEnabled: false,
		ShowLogLevel:  false,
	}

	output = test.CaptureOutput(func() { encoder.Color(&lConfig, colors.Magenta, testLog) })
	assert.Contains(t, output, colors.Magenta.String())
	assert.Contains(t, output, testLog)
	assert.NotContains(t, output, time.Now().Format("02/01/2006"))
	assert.Contains(t, output, colors.Reset.String())

	lConfig.DateEnabled = true
	output = test.CaptureOutput(func() { encoder.Color(&lConfig, colors.Cyan, testLog) })
	assert.Contains(t, output, colors.Cyan.String())
	assert.Contains(t, output, time.Now().Format("02/01/2006"))
	assert.Contains(t, output, testLog)
	assert.Contains(t, output, colors.Reset.String())

	output = test.CaptureOutput(func() { encoder.Color(&lConfig, colors.Gray, testLog) })
	assert.Contains(t, output, colors.Gray.String())
	assert.Contains(t, output, testLog)
	assert.Contains(t, output, colors.Reset.String())

	output = test.CaptureOutput(func() { encoder.Color(&lConfig, colors.Blue, testLog) })
	assert.Contains(t, output, colors.Blue.String())
	assert.Contains(t, output, testLog)
	assert.Contains(t, output, colors.Reset.String())

	os.Stdout = originalStdOut
}

func TestJSONEncoder_ValidJSONOutput(t *testing.T) {
	var jsonMsg string

	originalStdOut := os.Stdout
	testLog := "my testing Log"
	jsonEncoder := NewJSONEncoder(services.NewPrinter(), services.NewJsonMarshaler(), services.GetDateTimePrinter())
	lConfig := &test.LoggerConfigMock{
		DateEnabled:   false,
		TimeEnabled:   false,
		ColorsEnabled: false,
		ShowLogLevel:  false,
	}

	jsonMsg = test.CaptureOutput(
		func() {
			jsonEncoder.Log(lConfig, ll.DebugLvlName, shared.StdOutput, testLog, "id", 3)
		},
	)
	assert.NoError(t, json.Unmarshal([]byte(jsonMsg), &shared.JsonLog{}))

	jsonMsg = test.CaptureOutput(
		func() {
			jsonEncoder.Log(lConfig, ll.DebugLvlName, shared.StdOutput, testLog, "id", 3, 34, []string{"test", "test2"})
		},
	)
	assert.NoError(t, json.Unmarshal([]byte(jsonMsg), &shared.JsonLog{}))

	jsonMsg = test.CaptureOutput(func() {
		jsonEncoder.Log(lConfig, ll.DebugLvlName, shared.StdOutput, testLog, "id", 3, 34, []string{"test", "test2"}, []string{"k", "k2"}, 2.3, 'f', 'A')
	})
	assert.NoError(t, json.Unmarshal([]byte(jsonMsg), &shared.JsonLog{}))

	jsonMsg = "{{'test'}"
	assert.Error(t, json.Unmarshal([]byte(jsonMsg), &shared.JsonLog{}))

	jsonMsg = "{\"msg\"\"This is my Warn Log\",\"extras\":{\"Test arg\":[\"efsdaf\",\"dfas\"],\"[k3 k2]\":3}}"
	assert.Error(t, json.Unmarshal([]byte(jsonMsg), &shared.JsonLog{}))

	os.Stdout = originalStdOut
}
