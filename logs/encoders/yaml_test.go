package encoders

import (
	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/shared"
	"github.com/pho3b/tiny-logger/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"testing"
	"time"
)

func decodeYamlLogEntry(t *testing.T, logOutput string) shared.YamlLog {
	var yamlLog shared.YamlLog
	err := yaml.Unmarshal([]byte(logOutput), &yamlLog)
	assert.NoError(t, err)

	return yamlLog
}

func TestYAMLEncoder_LogDebug(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)
}

func TestYAMLEncoder_LogInfo(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogInfo(loggerConfig, "Test info message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "INFO", entry.Level)
	assert.Equal(t, "Test info message", entry.Message)
}

func TestYAMLEncoder_LogWarn(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogWarn(loggerConfig, "Test warning message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "WARN", entry.Level)
	assert.Equal(t, "Test warning message", entry.Message)
}

func TestYAMLEncoder_LogError(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureErrorOutput(func() {
		encoder.LogError(loggerConfig, "Test error message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "ERROR", entry.Level)
	assert.Equal(t, "Test error message", entry.Message)
}

func TestYAMLEncoder_LogFatalError(t *testing.T) {
	encoder := NewYAMLEncoder()
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

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
	loggerConfig := &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: true}

	output := captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry := decodeYamlLogEntry(t, output)
	assert.Equal(t, "DEBUG", entry.Level)
	assert.Equal(t, "Test debug message", entry.Message)

	loggerConfig = &test.LoggerConfigMock{DateEnabled: true, TimeEnabled: true, ColorsEnabled: true, ShowLogLevel: false}

	output = captureOutput(func() {
		encoder.LogDebug(loggerConfig, "Test debug message")
	})

	entry = decodeYamlLogEntry(t, output)
	assert.Equal(t, "", entry.Level)
}

func TestYAMLEncoder_Color(t *testing.T) {
	var output string
	testLog := "my testing log"
	originalStdOut := os.Stdout
	encoder := NewYAMLEncoder()
	lConfig := test.LoggerConfigMock{
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

func TestYAMLEncoder_ValidYAMLOutput(t *testing.T) {
	var yamlMsg string

	originalStdOut := os.Stdout
	testLog := "my testing log"
	yamlEncoder := NewYAMLEncoder()
	lConfig := &test.LoggerConfigMock{
		DateEnabled:   false,
		TimeEnabled:   false,
		ColorsEnabled: false,
		ShowLogLevel:  false,
	}

	yamlMsg = captureOutput(func() { yamlEncoder.LogInfo(lConfig, testLog, "id", 3) })
	assert.NoError(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	yamlMsg = captureOutput(func() { yamlEncoder.LogInfo(lConfig, testLog, "id", 3, 34, []string{"test", "test2"}) })
	assert.NoError(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	yamlMsg = captureOutput(func() {
		yamlEncoder.LogInfo(lConfig, testLog, "id", 3, 34, []string{"test", "test2"}, []string{"k", "k2"}, 2.3, 'f', 'A')
	})
	assert.NoError(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	yamlMsg = ":'test'}"
	assert.Error(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	yamlMsg = "level: DEBUG\ndatetime: 21/06/2025 11:34:56\nmsg: my testing log\nextras:\n  id: 3\n  34: [test test2]\n  [k k2]: 2.3\n  f: \"A\""
	assert.Error(t, yaml.Unmarshal([]byte(yamlMsg), &shared.YamlLog{}))

	os.Stdout = originalStdOut
}
