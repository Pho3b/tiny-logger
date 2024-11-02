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

func (d *DefaultEncoder) LogDebug(logger interfaces.LoggerInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.servicesWrapper.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.servicesWrapper.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Gray)

		_, _ = fmt.Fprintln(
			os.Stdout,
			fmt.Sprintf("%vDEBUG%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)
	}
}

func (d *DefaultEncoder) LogInfo(logger interfaces.LoggerInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.servicesWrapper.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.servicesWrapper.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Cyan)

		_, _ = fmt.Fprintln(
			os.Stdout,
			fmt.Sprintf("%vINFO%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)
	}
}

func (d *DefaultEncoder) LogWarn(logger interfaces.LoggerInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.servicesWrapper.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.servicesWrapper.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Yellow)

		_, _ = fmt.Fprintln(
			os.Stdout,
			fmt.Sprintf("%vWARN%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)
	}
}

func (d *DefaultEncoder) LogError(logger interfaces.LoggerInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.servicesWrapper.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.servicesWrapper.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Red)

		_, _ = fmt.Fprintln(
			os.Stderr,
			fmt.Sprintf("%vERROR%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)
	}
}

func (d *DefaultEncoder) LogFatalError(logger interfaces.LoggerInterface, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.servicesWrapper.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())
		colors := d.servicesWrapper.ColorsPrinter.PrintColors(logger.GetColorsEnabled(), c.Magenta)

		_, _ = fmt.Fprintln(
			os.Stderr,
			fmt.Sprintf("%vFATAL ERROR%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)

		os.Exit(1)
	}
}

func NewDefaultEncoder() *DefaultEncoder {
	return &DefaultEncoder{}
}
