package encoders

import (
	"bytes"
	"fmt"
	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
	"gopkg.in/yaml.v3"
	"os"
)

type YAMLEncoder struct {
	BaseEncoder
	DateTimePrinter services.DateTimePrinter
}

// yamlLogEntry represents the structure of a YAML log entry.
type yamlLogEntry struct {
	Level    string `yaml:"level,omitempty"`
	Date     string `yaml:"date,omitempty"`
	Time     string `yaml:"time,omitempty"`
	DateTime string `yaml:"datetime,omitempty"`
	Message  string `yaml:"message"`
}

// LogDebug formats and prints a debug-level log message in YAML format.
func (y *YAMLEncoder) LogDebug(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			ll.DebugLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args...),
		)

		y.printLog(s.StdOutput, msgBuffer, false)
	}
}

// LogInfo formats and prints an info-level log message in YAML format.
func (y *YAMLEncoder) LogInfo(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args...),
		)

		y.printLog(s.StdOutput, msgBuffer, false)
	}
}

// LogWarn formats and prints a warning-level log message in YAML format.
func (y *YAMLEncoder) LogWarn(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			ll.WarnLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args...),
		)

		y.printLog(s.StdOutput, msgBuffer, false)
	}
}

// LogError formats and prints an error-level log message in YAML format.
func (y *YAMLEncoder) LogError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !y.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			ll.ErrorLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args...),
		)

		y.printLog(s.StdErrOutput, msgBuffer, false)
	}
}

// LogFatalError formats and prints a fatal error-level log message in YAML format and exits the program.
func (y *YAMLEncoder) LogFatalError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !y.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			ll.FatalErrorLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args...),
		)

		y.printLog(s.StdErrOutput, msgBuffer, false)
		os.Exit(1)
	}
}

// Color formats and prints a colored log message using the specified color.
//
// Parameters:
//   - color: the color to apply to the log message.
//   - args: variadic msg arguments.
func (y *YAMLEncoder) Color(lConfig s.LoggerConfigsInterface, color c.Color, args ...any) {
	if len(args) > 0 {
		var b bytes.Buffer
		b.Grow((len(args) * averageWordLen) + averageWordLen)
		dEnabled, tEnabled := lConfig.GetDateTimeEnabled()

		msgBuffer := y.composeMsg(
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			false,
			y.castAndConcatenate(args...),
		)

		b.WriteString(color.String())
		b.Write(msgBuffer.Bytes())
		b.WriteString(c.Reset.String())

		y.printLog(s.StdOutput, b, false)
	}
}

func (y *YAMLEncoder) composeMsg(
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	showLogLevel bool,
	msg string,
) bytes.Buffer {
	var b bytes.Buffer
	b.Grow(len(msg) + 60)
	dateStr, timeStr, dateTimeStr := y.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled)

	if !showLogLevel {
		logLevel = ""
	}

	msgBytes, err := yaml.Marshal(
		yamlLogEntry{
			Level:    logLevel.String(),
			Date:     dateStr,
			Time:     timeStr,
			DateTime: dateTimeStr,
			Message:  msg,
		},
	)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "level: ERROR\nmessage: Failed to marshal YAML log: %s\n", err.Error())
		return bytes.Buffer{}
	}

	b.Write(msgBytes)
	return b
}

// NewYAMLEncoder initializes and returns a new YAMLEncoder instance.
func NewYAMLEncoder() *YAMLEncoder {
	encoder := &YAMLEncoder{
		DateTimePrinter: services.NewDateTimePrinter(),
	}
	encoder.encoderType = s.YamlEncoderType

	return encoder
}
