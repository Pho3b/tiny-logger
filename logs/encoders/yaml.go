package encoders

import (
	"fmt"
	"github.com/pho3b/tiny-logger/internal/services"
	"github.com/pho3b/tiny-logger/shared"
	"gopkg.in/yaml.v3"
	"os"
)

type YAMLEncoder struct {
	BaseEncoder
	DateTimePrinter *services.DateTimePrinter
}

// yamlLogEntry represents the structure of a YAML log entry.
type yamlLogEntry struct {
	Level    string `yaml:"level,omitempty"`
	Date     string `yaml:"date,omitempty"`
	Time     string `yaml:"time,omitempty"`
	DateTime string `json:"datetime,omitempty"`
	Message  string `yaml:"message"`
}

// LogDebug formats and prints a debug-level log message in YAML format.
func (y *YAMLEncoder) LogDebug(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		y.printYAMLLog("DEBUG", logger, shared.StdOutput, args...)
	}
}

// LogInfo formats and prints an info-level log message in YAML format.
func (y *YAMLEncoder) LogInfo(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		y.printYAMLLog("INFO", logger, shared.StdOutput, args...)
	}
}

// LogWarn formats and prints a warning-level log message in YAML format.
func (y *YAMLEncoder) LogWarn(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		y.printYAMLLog("WARN", logger, shared.StdOutput, args...)
	}
}

// LogError formats and prints an error-level log message in YAML format.
func (y *YAMLEncoder) LogError(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 && !y.areAllNil(args...) {
		y.printYAMLLog("ERROR", logger, shared.StdErrOutput, args...)
	}
}

// LogFatalError formats and prints a fatal error-level log message in YAML format and exits the program.
func (y *YAMLEncoder) LogFatalError(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 && !y.areAllNil(args...) {
		y.printYAMLLog("FATAL", logger, shared.StdErrOutput, args...)
		os.Exit(1)
	}
}

// printYAMLLog formats a log message as YAML and prints it to the appropriate output (stdout or stderr).
func (y *YAMLEncoder) printYAMLLog(
	level string,
	logger shared.LoggerConfigsInterface,
	outType shared.OutputType,
	args ...interface{},
) {
	dateStr, timeStr, dateTimeStr := y.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())

	if !logger.GetShowLogLevel() {
		level = ""
	}

	msgBytes, err := yaml.Marshal(
		yamlLogEntry{
			Level:    level,
			Date:     dateStr,
			Time:     timeStr,
			DateTime: dateTimeStr,
			Message:  y.buildMsg(args...),
		},
	)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "level: ERROR\nmessage: Failed to marshal YAML log: %s\n", err.Error())
		return
	}

	switch outType {
	case shared.StdOutput:
		_, _ = os.Stdout.Write(msgBytes)
	case shared.StdErrOutput:
		_, _ = os.Stderr.Write(msgBytes)
	}
}

// NewYAMLEncoder initializes and returns a new YAMLEncoder instance.
func NewYAMLEncoder() *YAMLEncoder {
	encoder := &YAMLEncoder{
		DateTimePrinter: services.NewDateTimePrinter(),
	}
	encoder.encoderType = shared.YamlEncoderType

	return encoder
}
