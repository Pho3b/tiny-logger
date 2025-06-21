package encoders

import (
	"bytes"
	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
	"os"
)

type JSONEncoder struct {
	BaseEncoder
	DateTimePrinter services.DateTimePrinter
	jsonMarshaler   services.JsonMarshaler
}

// LogDebug formats and prints a debug-level log message in JSON format.
func (j *JSONEncoder) LogDebug(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			j.jsonMarshaler,
			ll.DebugLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdOutput, msgBuffer, true)
	}
}

// LogInfo formats and prints an info-level log message in JSON format.
func (j *JSONEncoder) LogInfo(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			j.jsonMarshaler,
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdOutput, msgBuffer, true)
	}
}

// LogWarn formats and prints a warning-level log message in JSON format.
func (j *JSONEncoder) LogWarn(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			j.jsonMarshaler,
			ll.WarnLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdOutput, msgBuffer, true)
	}
}

// LogError formats and prints an error-level log message in JSON format.
func (j *JSONEncoder) LogError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !j.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			j.jsonMarshaler,
			ll.ErrorLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdErrOutput, msgBuffer, true)
	}
}

// LogFatalError formats and prints a fatal error-level log message in JSON format and exits the program.
func (j *JSONEncoder) LogFatalError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !j.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			j.jsonMarshaler,
			ll.FatalErrorLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdErrOutput, msgBuffer, true)
		os.Exit(1)
	}
}

// Color formats and prints a colored log message using the specified color.
//
// Parameters:
//   - color: the color to apply to the log message.
//   - args: variadic arguments where the first is treated as the message and the rest are appended.
func (j *JSONEncoder) Color(lConfig s.LoggerConfigsInterface, color c.Color, args ...any) {
	if len(args) > 0 {
		var b bytes.Buffer
		b.Grow((len(args) * averageWordLen) + averageWordLen)
		dEnabled, tEnabled := lConfig.GetDateTimeEnabled()

		msgBuffer := j.composeMsg(
			j.jsonMarshaler,
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			false,
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		b.WriteString(color.String())
		b.Write(msgBuffer.Bytes())
		b.WriteString(c.Reset.String())

		j.printLog(s.StdOutput, b, true)
	}
}

func (j *JSONEncoder) composeMsg(
	jsonMarshaler services.JsonMarshaler,
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	showLogLevel bool,
	msg string,
	extras ...any,
) bytes.Buffer {
	var b bytes.Buffer
	b.Grow((averageWordLen * len(extras)) + len(msg) + 60)
	dateStr, timeStr, dateTimeStr := j.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled)

	if !showLogLevel {
		logLevel = ""
	}

	b.Write(
		jsonMarshaler.Marshal(
			services.JsonLogEntry{
				Level:    logLevel.String(),
				Date:     dateStr,
				DateTime: dateTimeStr,
				Time:     timeStr,
				Message:  msg,
				Extras:   extras,
			},
		),
	)

	return b
}

// NewJSONEncoder initializes and returns a new JSONEncoder instance.
func NewJSONEncoder() *JSONEncoder {
	encoder := &JSONEncoder{DateTimePrinter: services.NewDateTimePrinter(), jsonMarshaler: services.JsonMarshaler{}}
	encoder.encoderType = s.JsonEncoderType

	return encoder
}
