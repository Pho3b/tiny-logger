package colors

import (
	"runtime"
	"strings"
)

type Color string

var (
	Reset   Color = "\033[0m"
	Red     Color = "\033[31m"
	Magenta Color = "\033[91m"
	Yellow  Color = "\033[33m"
	Cyan    Color = "\033[36m"
	Gray    Color = "\033[37m"
	White   Color = "\033[97m"
	Black   Color = "\033[30m"
)

func init() {
	// Avoids printing unusable Color codes for Windows operating systems
	if strings.Contains(runtime.GOOS, "windows") {
		Reset = ""
		Red = ""
		Magenta = ""
		Yellow = ""
		Cyan = ""
		Gray = ""
		White = ""
		Black = ""
	}
}

// IsColorValid returns true if the given color is a valid and supported one, false otherwise.
func IsColorValid(color Color) bool {
	return color == Red || color == Magenta || color == Yellow ||
		color == Cyan || color == Gray || color == White || color == Black
}
