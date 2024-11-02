package services

import (
	c "gitlab.com/docebo/libraries/go/tiny-logger/logs/colors"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

type ColorsPrinter struct {
}

func (d *ColorsPrinter) PrintColors(enableColors bool, logType log_level.LogLvlName) []c.Color {
	var res = []c.Color{"", ""}

	if enableColors {
		switch logType {
		case log_level.DebugLvlName:
			res[0] = c.Gray
			res[1] = c.Reset
		}
	}

	return res
}
