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

type DefaultEncoder struct {
	baseEncoder
	ColorsPrinter   services.ColorsPrinter
	DateTimePrinter services.DateTimePrinter
}

// LogDebug formats and prints a debug-level log message to stdout.
// It includes date and/or time if enabled, with the text in gray if colors are enabled.
func (d *DefaultEncoder) LogDebug(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		d.log(logger, ll.DebugLvlName, s.StdOutput, args...)
	}
}

// LogInfo formats and prints an info-level log message to stdout.
// It includes date and/or time if enabled, with the text in cyan if colors are enabled.
func (d *DefaultEncoder) LogInfo(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		d.log(logger, ll.InfoLvlName, s.StdOutput, args...)
	}
}

// LogWarn formats and prints a warning-level log message to stdout.
// It includes date and/or time if enabled, with the text in yellow if colors are enabled.
func (d *DefaultEncoder) LogWarn(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		d.log(logger, ll.WarnLvlName, s.StdOutput, args...)
	}
}

// LogError formats and prints an error-level log message to stderr.
// It includes date and/or time if enabled, with the text in red if colors are enabled.
func (d *DefaultEncoder) LogError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !d.areAllNil(args...) {
		d.log(logger, ll.ErrorLvlName, s.StdErrOutput, args...)
	}
}

// LogFatalError formats and prints a fatal error-level log message to stderr and terminates the program if any give
// args is not nil.
// It includes date and/or time if enabled, with the text in magenta if colors are enabled.
func (d *DefaultEncoder) LogFatalError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !d.areAllNil(args...) {
		d.log(logger, ll.FatalErrorLvlName, s.StdErrOutput, args...)
		os.Exit(1)
	}
}

// Color formats and prints a colored log message using the specified color.
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
		d.printLog(s.StdOutput, msgBuffer, true, logger.GetLogFile())
		d.putBuffer(msgBuffer)
	}
}

// log formats and prints a log message to the given output type.
// Internally used by all the encoder Log methods.
func (d *DefaultEncoder) log(
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

	d.printLog(outType, msgBuffer, true, logger.GetLogFile())
	d.putBuffer(msgBuffer)
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
	colors := d.ColorsPrinter.RetrieveColorsFromLogLevel(headerColorEnabled, ll.LogLvlNameToInt[logLevel])
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

	if showLogLevel || dateEnabled || timeEnabled {
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
	encoder := &DefaultEncoder{DateTimePrinter: services.NewDateTimePrinter(), ColorsPrinter: services.ColorsPrinter{}}
	encoder.encoderType = s.DefaultEncoderType
	encoder.bufferSyncPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	return encoder
}
