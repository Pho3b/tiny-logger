package logs

import (
	"fmt"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/colors"
	"os"
	"strings"
)

// log prints the given objects to the 'standard output' coloring the messages with the given Color.
// If the given color is not valid, the message is printed in WHITE by default.
//
// HINT: Colors can be retrieved from the 'colors' package
func log(color colors.Color, args ...interface{}) {
	if len(args) > 0 {
		if !colors.IsColorValid(color) {
			color = colors.White
		}

		_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%v%s%v ", color, buildMsg(args...), colors.Reset))
	}
}

// logInfo prints the given objects as strings to the 'standard output' and colors the prefix in CYAN in
// supported by the operating system.
func logInfo(args ...interface{}) {
	if len(args) > 0 {
		_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%vINFO:%v %s", colors.Cyan, colors.Reset, buildMsg(args...)))
	}
}

// logWarn prints the given objects as strings to the 'standard output' and colors the prefix in YELLOW if
// supported by the operating system.
func logWarn(args ...interface{}) {
	if len(args) > 0 {
		_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%vWARNING:%v %s", colors.Yellow, colors.Reset, buildMsg(args...)))
	}
}

// logError prints the given objects as strings to the 'standard logError' and colors the prefix in RED if
// supported by the operating system.
// It does not print anything if all the given args result to be nil.
func logError(args ...interface{}) {
	if len(args) > 0 && !areAllNil(args...) {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf("%vERROR:%v %s", colors.Red, colors.Reset, buildMsg(args...)))
	}
}

// logFatalError prints the given objects as strings to the 'standard logError' and colors the prefix in MAGENTA if
// supported by the operating system, it also exits the current process.
// logFatalError does not print anything and does not exit the current process if all the given args result to be nil.
func logFatalError(args ...interface{}) {
	if len(args) > 0 && !areAllNil(args...) {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf("%vFATAL ERROR:%v %s", colors.Magenta, colors.Reset, buildMsg(args...)))
		os.Exit(1)
	}
}

// buildMsg returns a string containing all the given arguments cast to strings concatenated with a white space.
func buildMsg(args ...interface{}) string {
	res := strings.Builder{}

	for _, arg := range args {
		res.WriteString(fmt.Sprintf("%v ", arg))
	}

	return strings.TrimSuffix(res.String(), " ")
}

// areAllNil returns true if all the given args are 'nil', false otherwise.
func areAllNil(args ...interface{}) bool {
	for _, arg := range args {
		if arg != nil {
			return false
		}
	}

	return true
}