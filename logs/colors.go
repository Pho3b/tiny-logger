package logs

import (
	"runtime"
	"strings"
)

var (
	reset        = "\033[0m"
	red          = "\033[31m"
	lightMagenta = "\033[91m"
	yellow       = "\033[33m"
	cyan         = "\033[36m"
	gray         = "\033[37m"
)

func init() {
	// Avoids printing unusable Color codes for Windows operating systems
	if strings.Contains(runtime.GOOS, "windows") {
		reset = ""
		lightMagenta = ""
		yellow = ""
		cyan = ""
		gray = ""
	}
}
