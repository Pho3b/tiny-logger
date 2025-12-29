package encoders

import (
	"bytes"
	"sync"
	"time"

	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
)

type DefaultEncoder struct {
	baseEncoder
	dateTimeFormat  s.DateTimeFormat
	PrinterService  services.Printer
	DateTimePrinter *services.DateTimePrinter
	LogsBuffer      *services.LogsBuffer
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
		args...,
	)

	msgBuffer.WriteByte('\n')

	if logger.GetBufferFlushInterval() == 0 {
		d.PrinterService.PrintLog(outType, msgBuffer, logger.GetLogFile())
	} else {
		d.LogsBuffer.AddLog(msgBuffer)
	}

	d.putBuffer(msgBuffer)
}

// Color formats and prints a colored Log message using the specified color.
func (d *DefaultEncoder) Color(logger s.LoggerConfigsInterface, color c.Color, args ...any) {
	if len(args) > 0 {
		msgBuffer := d.getBuffer()
		msgBuffer.WriteString(color.String())

		d.composeMsgInto(
			msgBuffer,
			ll.InfoLvlName,
			false,
			false,
			false,
			false,
			args...,
		)

		msgBuffer.WriteString(c.Reset.String())
		msgBuffer.WriteByte('\n')
		d.PrinterService.PrintLog(s.StdOutput, msgBuffer, logger.GetLogFile())
		d.putBuffer(msgBuffer)
	}
}

// SetDateTimeFormat updates the date and time format used by the encoder's DateTimePrinter.
// This method triggers an immediate update of the cached date and time strings to match the new format.
func (d *DefaultEncoder) SetDateTimeFormat(format s.DateTimeFormat) {
	d.DateTimePrinter.UpdateDateTimeFormat(format)
}

func (d *DefaultEncoder) SetBufferFlushInterval(interval time.Duration) {
	d.LogsBuffer.SetBufferFlushInterval(interval)
}

func (d *DefaultEncoder) FlushBuffer() {
	d.LogsBuffer.FlushLogs()
}

// composeMsgInto formats and writes the given 'msg' into the given buffer.
func (d *DefaultEncoder) composeMsgInto(
	buf *bytes.Buffer,
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	headerColorEnabled bool,
	showLogLevel bool,
	args ...any,
) {
	buf.Grow(len(args)*averageWordLen + defaultCharOverhead)

	isDateOrTimeEnabled := dateEnabled || timeEnabled
	colors := d.PrinterService.RetrieveColorsFromLogLevel(headerColorEnabled, ll.LogLvlNameToInt[logLevel])
	buf.WriteString(string(colors[0]))

	if showLogLevel {
		buf.WriteString(logLevel.String())

		if isDateOrTimeEnabled {
			buf.WriteByte(' ')
		}
	}

	if isDateOrTimeEnabled {
		dateStr, timeStr, dateTimeStr := d.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled)
		d.addFormattedDateTime(buf, dateStr, timeStr, dateTimeStr)
	}

	if showLogLevel || isDateOrTimeEnabled {
		buf.WriteByte(':')
		buf.WriteByte(' ')
	}

	buf.WriteString(string(colors[1]))
	d.castAndConcatenateInto(buf, args...)
}

// addFormattedDateTime correctly formats the dateTime string, adding and removing square brackets
// and white spaces as needed.
// While formatting, it adds the dateTime string to the given buffer.
func (d *DefaultEncoder) addFormattedDateTime(buf *bytes.Buffer, dateStr, timeStr, dateTimeStr string) {
	if dateStr == "" && timeStr == "" && dateTimeStr == "" {
		return
	}

	buf.Grow(averageWordLen)
	buf.WriteByte('[')

	if dateTimeStr != "" {
		buf.WriteString(dateTimeStr)
	} else {
		buf.WriteString(dateStr)

		if dateStr != "" && timeStr != "" {
			buf.WriteByte(' ')
		}

		buf.WriteString(timeStr)
	}

	buf.WriteByte(']')
}

// NewDefaultEncoder initializes and returns a new DefaultEncoder instance.
func NewDefaultEncoder() *DefaultEncoder {
	printer := services.NewPrinterService()
	encoder := &DefaultEncoder{
		PrinterService:  services.NewPrinterService(),
		DateTimePrinter: services.NewDateTimePrinter(),
		LogsBuffer:      services.NewLogsBuffer(0, nil, printer),
	}
	encoder.encoderType = s.DefaultEncoderType
	encoder.bufferSyncPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	return encoder
}
