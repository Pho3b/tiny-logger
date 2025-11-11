package encoders

import (
	"bytes"
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

// Log formats and prints a log message to the given output type.
// Internally used by all the encoder Log methods.
func (j *JSONEncoder) Log(
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
		logger.GetDateTimeFormat(),
		args[1:]...,
	)

	msgBuffer.WriteByte('\n')
	j.printLog(outType, msgBuffer, logger.GetLogFile())
	j.putBuffer(msgBuffer)
}

// Color formats and prints a colored Log message using the specified color.
func (j *JSONEncoder) Color(logger s.LoggerConfigsInterface, color c.Color, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
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
			logger.GetDateTimeFormat(),
			args[1:]...,
		)

		msgBuffer.WriteString(c.Reset.String())
		msgBuffer.WriteByte('\n')
		j.printLog(s.StdOutput, msgBuffer, logger.GetLogFile())
		j.putBuffer(msgBuffer)
	}
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
	dateTimeFormat s.DateTimeFormat,
	extras ...any,
) {
	buf.Grow((averageWordLen * len(extras)) + len(msg) + 60)
	dateStr, timeStr, dateTimeStr := j.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled, dateTimeFormat)

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
