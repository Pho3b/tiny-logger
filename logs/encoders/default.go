package encoders

import (
	"fmt"
	"gitlab.com/docebo/libraries/go/tiny-logger/interfaces"
	c "gitlab.com/docebo/libraries/go/tiny-logger/logs/colors"
	"os"
)

type DefaultEncoder struct {
	BaseEncoder
}

// LogDebug formats and prints a debug-level log message to stdout.
// It includes date and/or time if enabled, with the text in gray if colors are enabled.
func (d *DefaultEncoder) LogDebug(logger interfaces.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Gray)

		_, _ = fmt.Fprintln(
			os.Stdout,
			fmt.Sprintf("%vDEBUG%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)
	}
}

// LogInfo formats and prints an info-level log message to stdout.
// It includes date and/or time if enabled, with the text in cyan if colors are enabled.
func (d *DefaultEncoder) LogInfo(logger interfaces.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Cyan)

		_, _ = fmt.Fprintln(
			os.Stdout,
			fmt.Sprintf("%vINFO%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)
	}
}

// LogWarn formats and prints a warning-level log message to stdout.
// It includes date and/or time if enabled, with the text in yellow if colors are enabled.
func (d *DefaultEncoder) LogWarn(logger interfaces.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Yellow)

		_, _ = fmt.Fprintln(
			os.Stdout,
			fmt.Sprintf("%vWARN%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)
	}
}

// LogError formats and prints an error-level log message to stderr.
// It includes date and/or time if enabled, with the text in red if colors are enabled.
func (d *DefaultEncoder) LogError(logger interfaces.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Red)

		_, _ = fmt.Fprintln(
			os.Stderr,
			fmt.Sprintf("%vERROR%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)
	}
}

// LogFatalError formats and prints a fatal error-level log message to stderr and terminates the program.
// It includes date and/or time if enabled, with the text in magenta if colors are enabled.
// NOTE: the LogFatalError also Terminates the program with a non-zero exit code.
func (d *DefaultEncoder) LogFatalError(logger interfaces.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Magenta)

		_, _ = fmt.Fprintln(
			os.Stderr,
			fmt.Sprintf("%vFATAL ERROR%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)

		os.Exit(1)
	}
}

// NewDefaultEncoder initializes and returns a new DefaultEncoder instance.
func NewDefaultEncoder() *DefaultEncoder {
	return &DefaultEncoder{}
}
