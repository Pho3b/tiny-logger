package encoders

import (
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/pho3b/tiny-logger/internal/services"
	"github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	"github.com/pho3b/tiny-logger/shared"
	"github.com/pho3b/tiny-logger/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func decodeYamlLogEntry(t *testing.T, logOutput string) shared.YamlLog {
	var yamlLog shared.YamlLog
	err := yaml.Unmarshal([]byte(logOutput), &yamlLog)
	assert.NoError(t, err)

	return yamlLog
}

func TestYAMLEncoder_LogDebug(t *testing.T) {
	encoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.DebugLvlName, shared.StdOutput, "Test debug message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)
}

func TestYAMLEncoder_LogInfo(t *testing.T) {
	encoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.InfoLvlName, shared.StdOutput, "Test info message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "INFO", entry.Level)
	assert.Equal(t, "Test info message", entry.Message)
}

func TestYAMLEncoder_LogInfoWithExtras(t *testing.T) {
	encoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.InfoLvlName, shared.StdOutput, "Test info message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "INFO", entry.Level)
	assert.Equal(t, "Test info message", entry.Message)
	assert.IsType(t, make(map[string]any), entry.Extras)

	output = test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.InfoLvlName, shared.StdOutput, "Test info message with extras", "Location", "Italy", "Weather", "sunny", "Mood")
	})

	entry = decodeYamlLogEntry(t, output)
	assert.Equal(t, "Test info message with extras", entry.Message)
	assert.IsType(t, make(map[string]any), entry.Extras)
	assert.Equal(t, "Italy", entry.Extras["Location"])
	assert.Equal(t, "sunny", entry.Extras["Weather"])
	assert.Equal(t, nil, entry.Extras["Mood"])
}

func TestYAMLEncoder_LogWarn(t *testing.T) {
	encoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.WarnLvlName, shared.StdOutput, "Test warning message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "WARN", entry.Level)
	assert.Equal(t, "Test warning message", entry.Message)
}

func TestYAMLEncoder_LogError(t *testing.T) {
	encoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureErrorOutput(func() {
		encoder.Log(loggerConfig, ll.ErrorLvlName, shared.StdErrOutput, "Test error message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "ERROR", entry.Level)
	assert.Equal(t, "Test error message", entry.Message)
}

func TestYAMLEncoder_LogFatalError(t *testing.T) {
	encoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
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

func TestYAMLEncoder_DateTime(t *testing.T) {
	encoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}
	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.WarnLvlName, shared.StdOutput, "Test msg")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Date)
	assert.Empty(t, entry.DateTime)
	assert.NotEmpty(t, entry.Time)

	loggerConfig = &test.LoggerConfigMock{DateEnabled: true, ColorsEnabled: true, ShowLogLevel: true}
	output = test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.WarnLvlName, shared.StdOutput, "Test msg")
	})

	entry = decodeYamlLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Time)
	assert.Empty(t, entry.DateTime)
	assert.NotEmpty(t, entry.Date)

	loggerConfig = &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}
	output = test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.WarnLvlName, shared.StdOutput, "Test msg")
	})

	entry = decodeYamlLogEntry(t, output)
	assert.Equal(t, "Test msg", entry.Message)
	assert.Empty(t, entry.Time)
	assert.Empty(t, entry.Date)
	assert.NotEmpty(t, entry.DateTime)
}

func TestYAMLEncoder_ExtraMessages(t *testing.T) {
	yamlEncoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
	lConfig := &test.LoggerConfigMock{DateEnabled: false, TimeEnabled: false, ColorsEnabled: false, ShowLogLevel: false}

	output := test.CaptureOutput(func() {
		yamlEncoder.Log(lConfig, ll.InfoLvlName, shared.StdOutput, "test", "user", "alice", "ip", "192.168.1.1")
	})
	entry := decodeYamlLogEntry(t, output)
	assert.NotNil(t, entry)
	assert.NotNil(t, entry.Extras["ip"])
	assert.Len(t, entry.Extras, 2)

	output = test.CaptureOutput(func() {
		yamlEncoder.Log(lConfig, ll.InfoLvlName, shared.StdOutput, "test", "user", "alice", "ip")
	})
	entry = decodeYamlLogEntry(t, output)
	assert.Nil(t, entry.Extras["ip"])
	assert.Len(t, entry.Extras, 2)

	output = test.CaptureOutput(func() {
		yamlEncoder.Log(lConfig, ll.InfoLvlName, shared.StdOutput, "test", "user", "alice", "ip", "192.168.1.1", "city", "paris", "pass")
	})
	entry = decodeYamlLogEntry(t, output)
	assert.Len(t, entry.Extras, 4)
	assert.Equal(t, "alice", entry.Extras["user"])
	assert.Equal(t, "192.168.1.1", entry.Extras["ip"])
	assert.Equal(t, "paris", entry.Extras["city"])
	assert.Equal(t, nil, entry.Extras["pass"])
}

func TestYAMLEncoder_ShowLogLevelLt(t *testing.T) {
	encoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.DebugLvlName, shared.StdOutput, "Test debug message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)

	loggerConfig = &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: false}

	output = test.CaptureOutput(func() {
		encoder.Log(loggerConfig, ll.DebugLvlName, shared.StdOutput, "Test debug message")
	})

	entry = decodeYamlLogEntry(t, output)
	assert.Equal(t, "", entry.Level)
}

func TestYAMLEncoder_Color(t *testing.T) {
	var output string
	testLog := "my testing Log"
	originalStdOut := os.Stdout
	encoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
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

func TestYAMLEncoder_ValidYAMLOutput(t *testing.T) {
	var yamlMsg string

	originalStdOut := os.Stdout
	testLog := "my testing Log"
	yamlEncoder := NewYAMLEncoder(services.NewPrinter(), services.NewYamlMarshaler(), services.NewDateTimePrinter())
	lConfig := &test.LoggerConfigMock{
		DateEnabled:   false,
		TimeEnabled:   false,
		ColorsEnabled: false,
		ShowLogLevel:  false,
	}

	yamlMsg = test.CaptureOutput(
		func() {
			yamlEncoder.Log(lConfig, ll.DebugLvlName, shared.StdOutput, testLog, "id", 3)
		},
	)
	assert.NoError(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	yamlMsg = test.CaptureOutput(
		func() {
			yamlEncoder.Log(lConfig, ll.DebugLvlName, shared.StdOutput, testLog, "id", 3, 34, []string{"test", "test2"})
		},
	)
	assert.NoError(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	yamlMsg = test.CaptureOutput(func() {
		yamlEncoder.Log(lConfig, ll.DebugLvlName, shared.StdOutput, testLog, "id", 3, 34, []string{"test", "test2"}, []string{"k", "k2"}, 2.3, 'f', 'A')
	})
	assert.NoError(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	yamlMsg = ":'test'}"
	assert.Error(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	yamlMsg = "level: DEBUG\ndatetime: 21/06/2025 11:34:56\nmsg: my testing Log\nextras:\n  id: 3\n  34: [test test2]\n  [k k2]: 2.3\n  f: \"A\""
	assert.Error(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	os.Stdout = originalStdOut
}
