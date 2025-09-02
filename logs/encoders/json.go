package encoders

import (
	"bytes"
	"os"
	"sync"

	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
)

type JSONEncoder struct {
	baseEncoder
	DateTimePrinter services.DateTimePrinter
	jsonMarshaler   services.JsonMarshaler
}

// LogDebug formats and prints a debug-level log message in JSON format.
func (j *JSONEncoder) LogDebug(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		j.log(logger, ll.DebugLvlName, s.StdOutput, args...)
	}
}

// LogInfo formats and prints an info-level log message in JSON format.
func (j *JSONEncoder) LogInfo(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		j.log(logger, ll.InfoLvlName, s.StdOutput, args...)
	}
}

// LogWarn formats and prints a warning-level log message in JSON format.
func (j *JSONEncoder) LogWarn(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		j.log(logger, ll.WarnLvlName, s.StdOutput, args...)

	}
}

// LogError formats and prints an error-level log message in JSON format.
func (j *JSONEncoder) LogError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !j.areAllNil(args...) {
		j.log(logger, ll.ErrorLvlName, s.StdErrOutput, args...)
	}
}

// LogFatalError formats and prints a fatal error-level log message in JSON format and exits the program.
func (j *JSONEncoder) LogFatalError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !j.areAllNil(args...) {
		j.log(logger, ll.FatalErrorLvlName, s.StdErrOutput, args...)
		os.Exit(1)
	}
}

// Color formats and prints a colored log message using the specified color.
//
// Parameters:
//   - lConfig: the logger configuration.
//   - color: the color to apply to the log message.
//   - args: variadic arguments where the first is treated as the message and the rest are appended.
func (j *JSONEncoder) Color(lConfig s.LoggerConfigsInterface, color c.Color, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := lConfig.GetDateTimeEnabled()
		msgBuffer := j.getBuffer()
		msgBuffer.WriteString(color.String())

		j.composeMsgInto(
			msgBuffer,
			j.jsonMarshaler,
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			false,
			j.castToString(args[0]),
			args[1:]...,
		)

		msgBuffer.WriteString(c.Reset.String())
		j.printLog(s.StdOutput, msgBuffer, true)
		j.putBuffer(msgBuffer)
	}
}

// log formats and prints a log message to the given output type.
// Internally used by all the encoder Log methods.
func (j *JSONEncoder) log(
	logger s.LoggerConfigsInterface,
	logLvlName ll.LogLvlName,
	outType s.OutputType,
	args ...any,
) {
	dEnabled, tEnabled := logger.GetDateTimeEnabled()
	msgBuffer := j.getBuffer()

	j.composeMsgInto(
		msgBuffer,
		j.jsonMarshaler,
		logLvlName,
		dEnabled,
		tEnabled,
		logger.GetShowLogLevel(),
		j.castToString(args[0]),
		args[1:]...,
	)

	j.printLog(outType, msgBuffer, true)
	j.putBuffer(msgBuffer)
}

// composeMsgInto formats and writes the given 'msg' into the given buffer.
func (j *JSONEncoder) composeMsgInto(
	buf *bytes.Buffer,
	jsonMarshaler services.JsonMarshaler,
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	showLogLevel bool,
	msg string,
	extras ...any,
) {
	buf.Grow((averageWordLen * len(extras)) + len(msg) + 60)
	dateStr, timeStr, dateTimeStr := j.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled)

	if !showLogLevel {
		logLevel = ""
	}

	jsonMarshaler.MarshalInto(
		buf,
		services.JsonLogEntry{
			Level:    logLevel.String(),
			Date:     dateStr,
			DateTime: dateTimeStr,
			Time:     timeStr,
			Message:  msg,
			Extras:   extras,
		},
	)
}

// NewJSONEncoder initializes and returns a new JSONEncoder instance.
func NewJSONEncoder() *JSONEncoder {
	encoder := &JSONEncoder{DateTimePrinter: services.NewDateTimePrinter(), jsonMarshaler: services.JsonMarshaler{}}
	encoder.encoderType = s.JsonEncoderType
	encoder.bufferSyncPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	return encoder
}
