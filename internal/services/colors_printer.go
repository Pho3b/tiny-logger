package services

import (
	c "gitlab.com/docebo/libraries/go/tiny-logger/logs/colors"
)

type ColorsPrinter struct {
}

func (d *ColorsPrinter) PrintColors(enableColors bool, color c.Color) []c.Color {
	var res = []c.Color{"", ""}

	if enableColors {
		res[0] = color
		res[1] = c.Reset
	}

	return res
}
