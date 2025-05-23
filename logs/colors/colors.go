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
)

func (c Color) String() string {
	return string(c)
}

func init() {
	// Avoids printing unusable Color codes for Windows operating systems
	if strings.Contains(runtime.GOOS, "windows") {
		Reset = ""
		Red = ""
		Magenta = ""
		Yellow = ""
		Cyan = ""
	}
}
