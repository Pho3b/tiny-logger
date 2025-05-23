package services

import (
	c "github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/log_level"
)

type ColorsPrinter struct {
}

// RetrieveColorsFromLogLevel returns an array of colors as strings to be used in log output based on given log level.
// if enableColors is false, it returns an array of empty strings.
func (d *ColorsPrinter) RetrieveColorsFromLogLevel(enableColors bool, logLevelInt int8) []c.Color {
	var res = []c.Color{"", ""}

	if enableColors {
		switch logLevelInt {
		case log_level.FatalErrorLvl:
			res[0] = c.Magenta
		case log_level.ErrorLvl:
			res[0] = c.Red
		case log_level.WarnLvl:
			res[0] = c.Yellow
		case log_level.InfoLvl:
			res[0] = c.Cyan
		case log_level.DebugLvl:
			res[0] = c.Gray
		}

		res[1] = c.Reset
	}

	return res
}
