package encoders

import (
	"bytes"
	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
	"os"
)

type YAMLEncoder struct {
	BaseEncoder
	DateTimePrinter services.DateTimePrinter
	yamlMarshaler   services.YamlMarshaler
}

// LogDebug formats and prints a debug-level log message in YAML format.
func (y *YAMLEncoder) LogDebug(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			y.yamlMarshaler,
			ll.DebugLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args[0]),
			args[1:]...,
		)

		y.printLog(s.StdOutput, msgBuffer, false)
	}
}

// LogInfo formats and prints an info-level log message in YAML format.
func (y *YAMLEncoder) LogInfo(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			y.yamlMarshaler,
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args[0]),
			args[1:]...,
		)

		y.printLog(s.StdOutput, msgBuffer, false)
	}
}

// LogWarn formats and prints a warning-level log message in YAML format.
func (y *YAMLEncoder) LogWarn(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			y.yamlMarshaler,
			ll.WarnLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args[0]),
			args[1:]...,
		)

		y.printLog(s.StdOutput, msgBuffer, false)
	}
}

// LogError formats and prints an error-level log message in YAML format.
func (y *YAMLEncoder) LogError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !y.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			y.yamlMarshaler,
			ll.ErrorLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args[0]),
			args[1:]...,
		)

		y.printLog(s.StdErrOutput, msgBuffer, false)
	}
}

// LogFatalError formats and prints a fatal error-level log message in YAML format and exits the program.
func (y *YAMLEncoder) LogFatalError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !y.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := y.composeMsg(
			y.yamlMarshaler,
			ll.FatalErrorLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			y.castAndConcatenate(args[0]),
			args[1:]...,
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
			y.yamlMarshaler,
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			false,
			y.castAndConcatenate(args[0]),
			args[1:]...,
		)

		b.WriteString(color.String())
		b.Write(msgBuffer.Bytes())
		b.WriteString(c.Reset.String())

		y.printLog(s.StdOutput, b, false)
	}
}

func (y *YAMLEncoder) composeMsg(
	yamlMarshaler services.YamlMarshaler,
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	showLogLevel bool,
	msg string,
	extras ...any,
) bytes.Buffer {
	var b bytes.Buffer
	b.Grow(len(msg) + 60)
	date, time, dateTime := y.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled)

	if !showLogLevel {
		logLevel = ""
	}

	b.Write(
		yamlMarshaler.Marshal(
			services.YamlLogEntry{
				Level:    logLevel.String(),
				Date:     date,
				Time:     time,
				DateTime: dateTime,
				Message:  msg,
				Extras:   extras,
			},
		),
	)

	return b
}

// NewYAMLEncoder initializes and returns a new YAMLEncoder instance.
func NewYAMLEncoder() *YAMLEncoder {
	encoder := &YAMLEncoder{DateTimePrinter: services.NewDateTimePrinter(), yamlMarshaler: services.NewYamlMarshaler()}
	encoder.encoderType = s.YamlEncoderType

	return encoder
}
