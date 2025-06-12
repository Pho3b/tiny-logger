package colors

import (
	"runtime"
	"strings"
)

type Color string

var (
	Reset        Color = "\033[0m"
	Red          Color = "\033[31m"
	Magenta      Color = "\033[91m"
	Yellow       Color = "\033[33m"
	Cyan         Color = "\033[36m"
	Gray         Color = "\033[37m"
	Black        Color = "\033[30m"
	Green        Color = "\033[32m"
	Blue         Color = "\033[34m"
	White        Color = "\033[97m"
	DarkGray     Color = "\033[90m"
	BrightGreen  Color = "\033[92m"
	BrightYellow Color = "\033[93m"
	BrightBlue   Color = "\033[94m"
	Pink         Color = "\033[95m"
	BrightCyan   Color = "\033[96m"
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
		Gray = ""
		Black = ""
		Green = ""
		Blue = ""
		White = ""
		DarkGray = ""
		BrightGreen = ""
		BrightYellow = ""
		BrightBlue = ""
		Pink = ""
		BrightCyan = ""
	}
}
