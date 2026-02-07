package encoders

import (
	"bytes"
	"sync"

	"github.com/Pho3b/tiny-logger/internal/services"
	c "github.com/Pho3b/tiny-logger/logs/colors"
	ll "github.com/Pho3b/tiny-logger/logs/log_level"
	s "github.com/Pho3b/tiny-logger/shared"
)

type DefaultEncoder struct {
	baseEncoder
	dateTimeFormat  s.DateTimeFormat
	printer         services.Printer
	DateTimePrinter *services.DateTimePrinter
}

// Log formats and prints a Log message to the given output type.
func (d *DefaultEncoder) Log(
	logger s.LoggerConfigsInterface,
	logLvlName ll.LogLvlName,
	outType s.OutputType,
	args ...any,
) {
	dEnabled, tEnabled := logger.GetDateTimeEnabled()
	msgBuffer := d.getBuffer()

	d.composeMsgInto(
		msgBuffer,
		logLvlName,
		dEnabled,
		tEnabled,
		logger.GetColorsEnabled(),
		logger.GetShowLogLevel(),
		logger.GetDateTimeFormat(),
		args...,
	)

	msgBuffer.WriteByte('\n')
	d.printer.PrintLog(outType, msgBuffer, logger.GetLogFile())
	d.putBuffer(msgBuffer)
}

// Color formats and prints a colored Log message using the specified color.
func (d *DefaultEncoder) Color(logger s.LoggerConfigsInterface, color c.Color, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := d.getBuffer()
		msgBuffer.WriteString(color.String())

		d.composeMsgInto(
			msgBuffer,
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			false,
			false,
			logger.GetDateTimeFormat(),
			args...,
		)

		msgBuffer.WriteString(c.Reset.String())
		msgBuffer.WriteByte('\n')
		d.printer.PrintLog(s.StdOutput, msgBuffer, logger.GetLogFile())
		d.putBuffer(msgBuffer)
	}
}

// composeMsgInto formats and writes the given 'msg' into the given buffer.
func (d *DefaultEncoder) composeMsgInto(
	buf *bytes.Buffer,
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	headerColorEnabled bool,
	showLogLevel bool,
	dateTimeFormat s.DateTimeFormat,
	args ...any,
) {
	buf.Grow(len(args)*averageWordLen + defaultCharOverhead)

	isDateOrTimeEnabled := dateEnabled || timeEnabled
	colors := d.printer.RetrieveColorsFromLogLevel(headerColorEnabled, ll.LogLvlNameToInt[logLevel])
	buf.WriteString(string(colors[0]))

	if showLogLevel {
		buf.WriteString(logLevel.String())

		if isDateOrTimeEnabled {
			buf.WriteByte(' ')
		}
	}

	if isDateOrTimeEnabled {
		dateStr, timeStr, unixTs := d.DateTimePrinter.RetrieveDateTime(dateTimeFormat, dateEnabled, timeEnabled)
		d.addFormattedDateTime(buf, dateStr, timeStr, unixTs)
	}

	if showLogLevel || isDateOrTimeEnabled {
		buf.WriteByte(':')
		buf.WriteByte(' ')
	}

	buf.WriteString(string(colors[1]))
	d.castAndConcatenateInto(buf, args...)
}

// addFormattedDateTime formats and adds the date and time strings enclosed in square brackets to the given buffer.
func (d *DefaultEncoder) addFormattedDateTime(buf *bytes.Buffer, dateStr, timeStr, unixTs string) {
	if unixTs != "" {
		buf.WriteByte('[')
		buf.WriteString(unixTs)
		buf.WriteByte(']')

		return
	}

	if dateStr == "" && timeStr == "" {
		return
	}

	buf.WriteByte('[')
	buf.WriteString(dateStr)

	if dateStr != "" && timeStr != "" {
		buf.WriteByte(' ')
	}

	buf.WriteString(timeStr)
	buf.WriteByte(']')
}

// NewDefaultEncoder initializes and returns a new DefaultEncoder instance.
func NewDefaultEncoder(
	printer services.Printer,
	dateTimePrinter *services.DateTimePrinter,
) *DefaultEncoder {
	encoder := &DefaultEncoder{DateTimePrinter: dateTimePrinter, printer: printer}
	encoder.encoderType = s.DefaultEncoderType
	encoder.bufferSyncPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	return encoder
}
