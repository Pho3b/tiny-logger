package encoders

import (
	"bytes"
	"sync"

	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
)

type YAMLEncoder struct {
	baseEncoder
	DateTimePrinter services.DateTimePrinter
	yamlMarshaler   services.YamlMarshaler
}

// Log formats and prints a log message to the given output type.
// Internally used by all the encoder Log methods.
func (y *YAMLEncoder) Log(
	logger s.LoggerConfigsInterface,
	logLvlName ll.LogLvlName,
	outType s.OutputType,
	args ...any,
) {
	dEnabled, tEnabled := logger.GetDateTimeEnabled()
	msgBuffer := y.getBuffer()

	y.composeMsgInto(
		msgBuffer,
		y.yamlMarshaler,
		logLvlName,
		dEnabled,
		tEnabled,
		logger.GetShowLogLevel(),
		y.castToString(args[0]),
		logger.GetDateTimeFormat(),
		args[1:]...,
	)

	msgBuffer.WriteByte('\n')
	y.printLog(outType, msgBuffer, logger.GetLogFile())
	y.putBuffer(msgBuffer)
}

// Color formats and prints a colored Log message using the specified color.
func (y *YAMLEncoder) Color(logger s.LoggerConfigsInterface, color c.Color, args ...any) {
	if len(args) > 0 {
		msgBuffer := y.getBuffer()
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer.WriteString(color.String())

		y.composeMsgInto(
			msgBuffer,
			y.yamlMarshaler,
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			false,
			y.castToString(args[0]),
			logger.GetDateTimeFormat(),
			args[1:]...,
		)

		msgBuffer.WriteString(c.Reset.String())
		msgBuffer.WriteByte('\n')
		y.printLog(s.StdOutput, msgBuffer, logger.GetLogFile())
		y.putBuffer(msgBuffer)
	}
}

// composeMsgInto formats and writes the given 'msg' into the given buffer.
func (y *YAMLEncoder) composeMsgInto(
	buf *bytes.Buffer,
	yamlMarshaler services.YamlMarshaler,
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	showLogLevel bool,
	msg string,
	dateTimeFormat s.DateTimeFormat,
	extras ...any,
) {
	buf.Grow((averageWordLen * len(extras)) + len(msg) + 60)
	date, time, dateTime := y.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled, dateTimeFormat)

	if !showLogLevel {
		logLevel = ""
	}

	yamlMarshaler.MarshalInto(
		buf,
		services.YamlLogEntry{
			Level:    logLevel.String(),
			Date:     date,
			Time:     time,
			DateTime: dateTime,
			Message:  msg,
			Extras:   extras,
		},
	)
}

// NewYAMLEncoder initializes and returns a new YAMLEncoder instance.
func NewYAMLEncoder() *YAMLEncoder {
	encoder := &YAMLEncoder{DateTimePrinter: services.NewDateTimePrinter(), yamlMarshaler: services.NewYamlMarshaler()}
	encoder.encoderType = s.YamlEncoderType
	encoder.bufferSyncPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	return encoder
}
