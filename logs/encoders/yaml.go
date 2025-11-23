package encoders

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
)

type YAMLEncoder struct {
	baseEncoder
	DateTimePrinter *services.DateTimePrinter
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
		logger.GetDateTimeFormat(),
		y.castToString(args[0]),
		args[1:]...,
	)

	msgBuffer.WriteByte('\n')
	if err := y.printLog(outType, msgBuffer, logger.GetLogFile()); err != nil {
		fmt.Fprintf(os.Stderr, "tiny-logger: write error: %v\n", err)
	}
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
			logger.GetDateTimeFormat(),
			y.castToString(args[0]),
			args[1:]...,
		)

		msgBuffer.WriteString(c.Reset.String())
		msgBuffer.WriteByte('\n')
		if err := y.printLog(s.StdOutput, msgBuffer, logger.GetLogFile()); err != nil {
			fmt.Fprintf(os.Stderr, "tiny-logger: write error: %v\n", err)
		}
		y.putBuffer(msgBuffer)
	}
}

// SetDateTimeFormat updates the date and time format used by the encoder's DateTimePrinter.
// This method triggers an immediate update of the cached date and time strings to match the new format.
func (y *YAMLEncoder) SetDateTimeFormat(format s.DateTimeFormat) {
	y.DateTimePrinter.UpdateDateTimeFormat(format)
}

// composeMsgInto formats and writes the given 'msg' into the given buffer.
func (y *YAMLEncoder) composeMsgInto(
	buf *bytes.Buffer,
	yamlMarshaler services.YamlMarshaler,
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	showLogLevel bool,
	dateTimeFormat s.DateTimeFormat,
	msg string,
	extras ...any,
) {
	buf.Grow((averageWordLen * len(extras)) + len(msg) + 60)
	date, time, dateTime := y.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled)

	if !showLogLevel {
		logLevel = ""
	}

	yamlMarshaler.MarshalInto(
		buf,
		services.YamlLogEntry{
			Level:          logLevel.String(),
			Date:           date,
			Time:           time,
			DateTime:       dateTime,
			DateTimeFormat: dateTimeFormat,
			Message:        msg,
			Extras:         extras,
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
