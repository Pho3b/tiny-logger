package services

import (
	c "gitlab.com/docebo/libraries/go/tiny-logger/logs/colors"
)

type ColorsPrinter struct {
}

// PrintColors returns an array of colors to be used in log output based on color settings.
// If enableColors is true, it returns the provided color followed by a reset color;
// if false, it returns an array of empty strings.
func (d *ColorsPrinter) PrintColors(enableColors bool, color c.Color) []c.Color {
	var res = []c.Color{"", ""}

	if enableColors {
		res[0] = color
		res[1] = c.Reset
	}

	return res
}
