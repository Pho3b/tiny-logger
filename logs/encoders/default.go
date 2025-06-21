package encoders

import (
	"bytes"
	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
	"os"
)

type DefaultEncoder struct {
	BaseEncoder
	ColorsPrinter   services.ColorsPrinter
	DateTimePrinter services.DateTimePrinter
}

// LogDebug formats and prints a debug-level log message to stdout.
// It includes date and/or time if enabled, with the text in gray if colors are enabled.
func (d *DefaultEncoder) LogDebug(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := d.composeMsg(
			ll.DebugLvlName,
			dEnabled, tEnabled,
			logger.GetColorsEnabled(),
			logger.GetShowLogLevel(),
			d.castAndConcatenate(args...),
		)

		d.printLog(s.StdOutput, msgBuffer, true)
	}
}

// LogInfo formats and prints an info-level log message to stdout.
// It includes date and/or time if enabled, with the text in cyan if colors are enabled.
func (d *DefaultEncoder) LogInfo(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := d.composeMsg(
			ll.InfoLvlName,
			dEnabled, tEnabled,
			logger.GetColorsEnabled(),
			logger.GetShowLogLevel(),
			d.castAndConcatenate(args...),
		)

		d.printLog(s.StdOutput, msgBuffer, true)
	}
}

// LogWarn formats and prints a warning-level log message to stdout.
// It includes date and/or time if enabled, with the text in yellow if colors are enabled.
func (d *DefaultEncoder) LogWarn(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := d.composeMsg(
			ll.WarnLvlName,
			dEnabled, tEnabled,
			logger.GetColorsEnabled(),
			logger.GetShowLogLevel(),
			d.castAndConcatenate(args...),
		)

		d.printLog(s.StdOutput, msgBuffer, true)
	}
}

// LogError formats and prints an error-level log message to stderr.
// It includes date and/or time if enabled, with the text in red if colors are enabled.
func (d *DefaultEncoder) LogError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !d.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := d.composeMsg(
			ll.ErrorLvlName,
			dEnabled, tEnabled,
			logger.GetColorsEnabled(),
			logger.GetShowLogLevel(),
			d.castAndConcatenate(args...),
		)

		d.printLog(s.StdErrOutput, msgBuffer, true)
	}
}

// LogFatalError formats and prints a fatal error-level log message to stderr and terminates the program if any give
// args is not nil.
// It includes date and/or time if enabled, with the text in magenta if colors are enabled.
func (d *DefaultEncoder) LogFatalError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !d.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := d.composeMsg(
			ll.DebugLvlName,
			dEnabled, tEnabled,
			logger.GetColorsEnabled(),
			logger.GetShowLogLevel(),
			d.castAndConcatenate(args...),
		)

		d.printLog(s.StdErrOutput, msgBuffer, true)
		os.Exit(1)
	}
}

// Color formats and prints a colored log message using the specified color.
//
// Parameters:
//   - color: the color to apply to the log message.
//   - args: variadic msg arguments.
func (d *DefaultEncoder) Color(_ s.LoggerConfigsInterface, color c.Color, args ...any) {
	if len(args) > 0 {
		var b bytes.Buffer
		b.Grow((len(args) * averageWordLen) + averageWordLen)
		msgBuffer := d.composeMsg(
			ll.InfoLvlName,
			false,
			false,
			false,
			false,
			d.castAndConcatenate(args...),
		)

		b.WriteString(color.String())
		b.Write(msgBuffer.Bytes())
		b.WriteString(c.Reset.String())

		d.printLog(s.StdOutput, b, true)
	}
}

func (d *DefaultEncoder) composeMsg(
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	headerColorEnabled bool,
	showLogLevel bool,
	msg string,
) bytes.Buffer {
	var b bytes.Buffer
	b.Grow(len(msg) + 50)

	dateStr, timeStr, dateTimeStr := d.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled)
	colors := d.ColorsPrinter.RetrieveColorsFromLogLevel(headerColorEnabled, ll.LogLvlNameToInt[logLevel])

	b.WriteString(string(colors[0]))

	if showLogLevel {
		b.WriteString(logLevel.String())
	}

	dateTime := d.formatDateTimeString(dateStr, timeStr, dateTimeStr)
	b.Write(dateTime.Bytes())

	if showLogLevel || dateEnabled || timeEnabled {
		b.WriteRune(':')
		b.WriteByte(' ')
	}

	b.WriteString(string(colors[1]))
	b.WriteString(msg)

	return b
}

// formatDateTimeString correctly formats the dateTime string adding and removing square brackets
// and white spaces as needed.
func (d *DefaultEncoder) formatDateTimeString(dateStr, timeStr, dateTimeStr string) bytes.Buffer {
	var sb bytes.Buffer

	if dateStr == "" && timeStr == "" && dateTimeStr == "" {
		return sb
	}

	sb.Grow(averageWordLen)
	sb.WriteByte('[')

	if dateTimeStr != "" {
		sb.WriteString(dateTimeStr)
	} else {
		sb.WriteString(dateStr)

		if dateStr != "" && timeStr != "" {
			sb.WriteByte(' ')
		}

		sb.WriteString(timeStr)
	}

	sb.WriteByte(']')

	return sb
}

// NewDefaultEncoder initializes and returns a new DefaultEncoder instance.
func NewDefaultEncoder() *DefaultEncoder {
	encoder := &DefaultEncoder{DateTimePrinter: services.NewDateTimePrinter(), ColorsPrinter: services.ColorsPrinter{}}
	encoder.encoderType = s.DefaultEncoderType

	return encoder
}
