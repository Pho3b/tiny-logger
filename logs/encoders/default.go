package encoders

import (
	"bytes"
	"github.com/pho3b/tiny-logger/internal/services"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	"github.com/pho3b/tiny-logger/shared"
	"os"
)

type DefaultEncoder struct {
	BaseEncoder
	ColorsPrinter   *services.ColorsPrinter
	DateTimePrinter *services.DateTimePrinter
}

// LogDebug formats and prints a debug-level log message to stdout.
// It includes date and/or time if enabled, with the text in gray if colors are enabled.
func (d *DefaultEncoder) LogDebug(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		d.printDefaultLog(ll.DebugLvlName, logger, shared.StdOutput, d.buildMsg(args...))
	}
}

// LogInfo formats and prints an info-level log message to stdout.
// It includes date and/or time if enabled, with the text in cyan if colors are enabled.
func (d *DefaultEncoder) LogInfo(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		d.printDefaultLog(ll.InfoLvlName, logger, shared.StdOutput, d.buildMsg(args...))
	}
}

// LogWarn formats and prints a warning-level log message to stdout.
// It includes date and/or time if enabled, with the text in yellow if colors are enabled.
func (d *DefaultEncoder) LogWarn(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		d.printDefaultLog(ll.WarnLvlName, logger, shared.StdOutput, d.buildMsg(args...))
	}
}

// LogError formats and prints an error-level log message to stderr.
// It includes date and/or time if enabled, with the text in red if colors are enabled.
func (d *DefaultEncoder) LogError(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 && !d.areAllNil(args...) {
		d.printDefaultLog(ll.ErrorLvlName, logger, shared.StdErrOutput, d.buildMsg(args...))
	}
}

// LogFatalError formats and prints a fatal error-level log message to stderr and terminates the program.
// It includes date and/or time if enabled, with the text in magenta if colors are enabled.
// NOTE: the LogFatalError also Terminates the program with a non-zero exit code.
func (d *DefaultEncoder) LogFatalError(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 && !d.areAllNil(args...) {
		d.printDefaultLog(ll.FatalErrorLvlName, logger, shared.StdErrOutput, d.buildMsg(args...))
		os.Exit(1)
	}
}

// printDefaultLog formats a default log message and prints it to the appropriate output (stdout or stderr).
func (d *DefaultEncoder) printDefaultLog(
	logLevelName ll.LogLvlName,
	logger shared.LoggerConfigsInterface,
	outType shared.OutputType,
	msg string,
) {
	dEnabled, tEnabled := logger.GetDateTimeEnabled()
	dateStr, timeStr, dateTimeStr := d.DateTimePrinter.RetrieveDateTime(dEnabled, tEnabled)
	colors := d.ColorsPrinter.RetrieveColorsFromLogLevel(logger.GetColorsEnabled(), ll.LogLvlNameToInt[logLevelName])

	// Composing the final log message
	var b bytes.Buffer
	b.Grow(len(msg) + 50)

	b.WriteString(string(colors[0]))

	if logger.GetShowLogLevel() {
		b.WriteString(logLevelName.String())
		b.WriteRune(':')
	}

	dtb := d.formatDateTimeString(dateStr, timeStr, dateTimeStr)
	b.Write(dtb.Bytes())
	b.WriteString(string(colors[1]))

	if logger.GetShowLogLevel() || dEnabled || tEnabled {
		b.WriteByte(' ')
	}

	b.WriteString(msg)
	b.WriteByte('\n')

	// Actual message print
	switch outType {
	case shared.StdOutput:
		_, _ = os.Stdout.Write(b.Bytes())
	case shared.StdErrOutput:
		_, _ = os.Stderr.Write(b.Bytes())
	}
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
	encoder := &DefaultEncoder{
		DateTimePrinter: services.NewDateTimePrinter(),
		ColorsPrinter:   &services.ColorsPrinter{},
	}
	encoder.encoderType = shared.DefaultEncoderType

	return encoder
}
