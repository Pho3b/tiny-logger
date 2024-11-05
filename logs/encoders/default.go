package encoders

import (
	"fmt"
	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
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
		d.printDefaultLog("DEBUG", logger, shared.StdOutput, c.Gray, d.buildMsg(args...))
	}
}

// LogInfo formats and prints an info-level log message to stdout.
// It includes date and/or time if enabled, with the text in cyan if colors are enabled.
func (d *DefaultEncoder) LogInfo(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		d.printDefaultLog("INFO", logger, shared.StdOutput, c.Cyan, d.buildMsg(args...))
	}
}

// LogWarn formats and prints a warning-level log message to stdout.
// It includes date and/or time if enabled, with the text in yellow if colors are enabled.
func (d *DefaultEncoder) LogWarn(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		d.printDefaultLog("WARN", logger, shared.StdOutput, c.Yellow, d.buildMsg(args...))
	}
}

// LogError formats and prints an error-level log message to stderr.
// It includes date and/or time if enabled, with the text in red if colors are enabled.
func (d *DefaultEncoder) LogError(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		d.printDefaultLog("ERROR", logger, shared.StdErrOutput, c.Red, d.buildMsg(args...))
	}
}

// LogFatalError formats and prints a fatal error-level log message to stderr and terminates the program.
// It includes date and/or time if enabled, with the text in magenta if colors are enabled.
// NOTE: the LogFatalError also Terminates the program with a non-zero exit code.
func (d *DefaultEncoder) LogFatalError(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		d.printDefaultLog("FATAL ERROR", logger, shared.StdErrOutput, c.Magenta, d.buildMsg(args...))
		os.Exit(1)
	}
}

// printDefaultLog formats a default log message and prints it to the appropriate output (stdout or stderr).
func (d *DefaultEncoder) printDefaultLog(
	level string,
	logger shared.LoggerConfigsInterface,
	outType shared.OutputType,
	color c.Color,
	args ...interface{},
) {
	var output *os.File
	dateStr, timeStr := d.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
	colors := d.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), color)
	whitespace := " "
	if dateStr == "" {
		whitespace = ""
	}

	switch outType {
	case shared.StdOutput:
		output = os.Stdout
	case shared.StdErrOutput:
		output = os.Stderr
	}

	_, _ = fmt.Fprintln(
		output,
		fmt.Sprintf(
			"%v%s[%s%s%s]:%v %s",
			colors[0],
			level,
			dateStr,
			whitespace,
			timeStr,
			colors[1],
			d.buildMsg(args...),
		),
	)
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
